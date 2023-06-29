package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/windbnb/notification-service/model"
	"github.com/windbnb/notification-service/service"
	"github.com/windbnb/notification-service/tracer"
)

type Handler struct {
	Service *service.RatingService
	Tracer  opentracing.Tracer
	Closer  io.Closer
}

func (handler *Handler) PutNotificationSettings(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("putNotificationHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling rating host at %s\n", r.URL.Path)),
	)

	var notificationSettingsRequest model.NotificationSettingsRequest
	json.NewDecoder(r.Body).Decode(&notificationSettingsRequest)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	notificationSettings, err := handler.Service.UpdateUserNotificationSettings(&notificationSettingsRequest, ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(notificationSettings)
}

func (handler *Handler) GetNotificationSettings(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("getNotificationSettingsHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling rating host at %s\n", r.URL.Path)),
	)

	params := mux.Vars(r)
	hostId, _ := strconv.ParseUint(params["id"], 10, 32)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	hostRating, err := handler.Service.GetUserNotificationSettings(uint(hostId), ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(hostRating)
}

func (handler *Handler) SaveNotification(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("saveNotificationHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling saving notification at %s\n", r.URL.Path)),
	)

	var notificationRequest model.NotificationRequest
	json.NewDecoder(r.Body).Decode(&notificationRequest)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	notification, err := handler.Service.SaveNotification(&notificationRequest, ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(notification)
}

func (handler *Handler) GetUserRecentNotification(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("getUserRecentNotificationHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling reading user notification at %s\n", r.URL.Path)),
	)

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	notifications, err := handler.Service.GetUserRecentNotification(uint(userId), ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(notifications)
}

// func (handler *Handler) RateAccomodation(w http.ResponseWriter, r *http.Request) {
// 	span := tracer.StartSpanFromRequest("accomodationRatingHandler", handler.Tracer, r)
// 	defer span.Finish()
// 	span.LogFields(
// 		tracer.LogString("handler", fmt.Sprintf("handling rating accomodation at %s\n", r.URL.Path)),
// 	)

// 	var accomodationRatingRequest model.AccomodationRatingRequest
// 	json.NewDecoder(r.Body).Decode(&accomodationRatingRequest)

// 	ctx := tracer.ContextWithSpan(context.Background(), span)
// 	accomodationRating, err := handler.Service.SaveAccomodationRating(&accomodationRatingRequest, ctx)

// 	w.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		tracer.LogError(span, err)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
// 		return
// 	}

// 	json.NewEncoder(w).Encode(accomodationRating)
// }

// func (handler *Handler) GetAccomodationAverageReview(w http.ResponseWriter, r *http.Request) {
// 	span := tracer.StartSpanFromRequest("accomodationAverageReviewHandler", handler.Tracer, r)
// 	defer span.Finish()
// 	span.LogFields(
// 		tracer.LogString("handler", fmt.Sprintf("handling average accomodation at %s\n", r.URL.Path)),
// 	)

// 	params := mux.Vars(r)
// 	accomodationId, _ := strconv.ParseUint(params["id"], 10, 32)

// 	ctx := tracer.ContextWithSpan(context.Background(), span)
// 	accomodationRating, err := handler.Service.GetAverageAccomodationRating(uint(accomodationId), ctx)

// 	w.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		tracer.LogError(span, err)
// 		w.WriteHeader(http.StatusUnauthorized)
// 		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
// 		return
// 	}

// 	json.NewEncoder(w).Encode(accomodationRating)
// }
