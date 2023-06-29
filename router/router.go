package router

import (
	"github.com/gorilla/mux"
	"github.com/windbnb/notification-service/handler"
	"github.com/windbnb/notification-service/metrics"
)

func ConfigureRouter(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/notifications", metrics.MetricProxy(handler.SaveNotification)).Methods("POST")
	router.HandleFunc("/api/notifications/user/{id}", metrics.MetricProxy(handler.GetUserRecentNotification)).Methods("GET")

	router.HandleFunc("/api/notifications/settings/userId/{id}", metrics.MetricProxy(handler.GetNotificationSettings)).Methods("GET")
	router.HandleFunc("/api/notifications/settings", metrics.MetricProxy(handler.PutNotificationSettings)).Methods("PUT")

	router.Path("/metrics").Handler(metrics.MetricsHandler())

	return router
}
