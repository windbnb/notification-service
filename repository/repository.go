package repository

import (
	"context"
	"time"

	"github.com/windbnb/notification-service/model"
	"github.com/windbnb/notification-service/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *Repository) FindUserNotificationSettings(userId uint, ctx context.Context) (*[]model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "findUserNotificationSettingsRepository")
	defer span.Finish()

	notificationSettingsList := []model.NotificationSettings{}
	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"userId", userId},
	}

	cursor, err := r.Db.Collection("notification-settings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	for cursor.Next(dbCtx) {
		var notificationSettings model.NotificationSettings
		err := cursor.Decode(&notificationSettings)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		notificationSettingsList = append(notificationSettingsList, notificationSettings)
	}

	return &notificationSettingsList, nil
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
