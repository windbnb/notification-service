package service

import (
	"context"
	"os"

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

// func (s *RatingService) SaveAccomodationRating(accomodationRatingRequest *model.AccomodationRatingRequest, ctx context.Context) (*model.AccomodationRating, error) {
// 	span := tracer.StartSpanFromContext(ctx, "saveAccomodationRatingService")
// 	defer span.Finish()

// 	ctx = tracer.ContextWithSpan(context.Background(), span)

// 	_, err := s.Repo.DeleteGuestAccomodationRatings(accomodationRatingRequest.GuestId, accomodationRatingRequest.AccomodationId, ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Checks
// 	if accomodationRatingRequest.Raiting < 1 || accomodationRatingRequest.Raiting > 5 {
// 		return nil, errors.New("rating must be between 1 and 5")
// 	}

// 	// Get if had accomodation

// 	var accomodationRatingRequestData = model.AccomodationRating{
// 		GuestId:        accomodationRatingRequest.GuestId,
// 		AccomodationId: accomodationRatingRequest.AccomodationId,
// 		Raiting:        accomodationRatingRequest.Raiting}

// 	s.Repo.RateAccomodation(&accomodationRatingRequestData, ctx)

// 	return &accomodationRatingRequestData, nil
// }

// func (s *RatingService) GetAverageHostRating(hostId uint, ctx context.Context) (*model.HostAvgRating, error) {
// 	span := tracer.StartSpanFromContext(ctx, "getAverageHostRatingService")
// 	defer span.Finish()

// 	ctx = tracer.ContextWithSpan(context.Background(), span)

// 	ratings, err := s.Repo.FindAllHostRatings(uint(hostId), ctx)
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

// 	var result = model.HostAvgRating{
// 		HostId:         uint(hostId),
// 		AverageRaiting: avgRating,
// 	}

// 	return &result, nil
// }

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
