package repository

import (
	"context"
	"time"

	"github.com/windbnb/notification-service/model"
	"github.com/windbnb/notification-service/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Db *mongo.Database
}

func (r *Repository) AddNotification(notification *model.Notification, ctx context.Context) (*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "addNotification")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	notification.ID = primitive.NewObjectID()
	_, err := r.Db.Collection("notifications").InsertOne(dbCtx, &notification)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return notification, nil
}

func (r *Repository) GetUserRecentNotification(userId uint, findLimit uint, ctx context.Context) (*[]model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserRecentNotificationRepository")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"userId", userId},
	}
	options := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(20) // Zamijenite "timestamp" sa poljem koje označava vreme dodavanja

	cursor, err := r.Db.Collection("notifications").Find(dbCtx, filter, options)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	var notifications = []model.Notification{}

	for cursor.Next(dbCtx) {
		var notification model.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		notifications = append(notifications, notification)
	}

	return &notifications, nil
}

func (r *Repository) FindUserNotificationSettings(userId uint, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "findUserNotificationSettingsRepository")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"userId", userId},
	}

	cursor := r.Db.Collection("notification-settings").FindOne(dbCtx, filter)

	var notificationSettings model.NotificationSettings
	err := cursor.Decode(&notificationSettings)
	if err != nil {
		emptySettings := model.NotificationSettings{
			UserId:              userId,
			ResRequest:          true,
			ResCancel:           true,
			HostReview:          true,
			AccommodationReview: true,
			ReservationResponse: true,
		}

		return &emptySettings, nil
	}

	return &notificationSettings, nil
}

func (r *Repository) DeleteNotificationSettings(userId uint, ctx context.Context) (int64, error) {
	span := tracer.StartSpanFromContext(ctx, "deleteNotificationSettingsRepository")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"userId", userId},
	}

	data, err := r.Db.Collection("notification-settings").DeleteMany(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return 0, err
	}

	return data.DeletedCount, nil
}

func (r *Repository) AddNotificationSettings(notificationSettings *model.NotificationSettings, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "addNotificationSettings")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	notificationSettings.ID = primitive.NewObjectID()
	_, err := r.Db.Collection("notification-settings").InsertOne(dbCtx, &notificationSettings)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return notificationSettings, nil
}

func (r *Repository) UpdateNotificationSettings(notificationSettings *model.NotificationSettings, ctx context.Context) (*model.NotificationSettings, error) {
	_, err := r.DeleteNotificationSettings(notificationSettings.UserId, ctx)
	if err != nil {
		return nil, err
	}

	return r.AddNotificationSettings(notificationSettings, ctx)
}
