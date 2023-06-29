package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/windbnb/notification-service/model"
	"github.com/windbnb/notification-service/repository"
	"github.com/windbnb/notification-service/tracer"
)

var ratingServiceUrl = os.Getenv("ratingServiceUrl") + "/api/notifications/"

type RatingService struct {
	Repo repository.IRepository
}

func (s *RatingService) UpdateUserNotificationSettings(NotificationSettingsRequest *model.NotificationSettingsRequest, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "saveHostRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Checks

	var notificationSettingsData = model.NotificationSettings{
		UserId:              NotificationSettingsRequest.UserId,
		ResRequest:          NotificationSettingsRequest.ResRequest,
		ResCancel:           NotificationSettingsRequest.ResCancel,
		HostReview:          NotificationSettingsRequest.HostReview,
		AccommodationReview: NotificationSettingsRequest.AccommodationReview,
		ReservationResponse: NotificationSettingsRequest.ReservationResponse}

	s.Repo.UpdateNotificationSettings(&notificationSettingsData, ctx)

	return &notificationSettingsData, nil
}

func (s *RatingService) GetUserNotificationSettings(userId uint, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserNotificationSettingsService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	notificationSettings, err := s.Repo.FindUserNotificationSettings(userId, ctx)
	if err != nil {
		return nil, err
	}

	return notificationSettings, nil
}

func (s *RatingService) SaveNotification(notification *model.NotificationRequest, ctx context.Context) (*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "saveNotificationService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Checks
	notificationSettings, err := s.Repo.FindUserNotificationSettings(notification.UserId, ctx)
	if err != nil {
		return nil, err
	}

	fmt.Print(notification.NotifType)

	switch notification.NotifType {
	case "ResRequest":
		if !notificationSettings.ResRequest {
			return nil, nil
		}
	case "ResCancel":
		if !notificationSettings.ResCancel {
			return nil, nil
		}

	case "HostReview":
		if !notificationSettings.HostReview {
			return nil, nil
		}

	case "AccommodationReview":
		if !notificationSettings.AccommodationReview {
			return nil, nil
		}

	case "ReservationResponse":
		if !notificationSettings.ReservationResponse {
			return nil, nil
		}
	}

	notificationData := model.Notification{
		UserId:     notification.UserId,
		Timestamp:  time.Now(),
		NotifTitle: notification.NotifTitle,
		NotifType:  notification.NotifType,
		Seen:       false,
	}

	notificationResult, err := s.Repo.AddNotification(&notificationData, ctx)
	if err != nil {
		return nil, err
	}

	return notificationResult, nil
}

func (s *RatingService) GetUserRecentNotification(userId uint, ctx context.Context) (*[]model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserNotificationSettingsService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	notifications, err := s.Repo.GetUserRecentNotification(userId, 20, ctx)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

// func (s *RatingService) GetAverageAccomodationRating(accomodationId uint, ctx context.Context) (*model.AccomodationAvgRating, error) {
// 	span := tracer.StartSpanFromContext(ctx, "getAverageAccomodationRatingService")
// 	defer span.Finish()

// 	ctx = tracer.ContextWithSpan(context.Background(), span)

// 	ratings, err := s.Repo.FindAllAccomodationRatings(uint(accomodationId), ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var avgRating float32 = 0
// 	for _, rating := range *ratings {
// 		avgRating += float32(rating.Raiting)

// 	}

// 	if len(*ratings) > 0 {
// 		avgRating /= float32(len(*ratings))
// 	}

// 	var result = model.AccomodationAvgRating{
// 		AccomodationId: uint(accomodationId),
// 		AverageRaiting: avgRating,
// 	}

// 	return &result, nil
// }
