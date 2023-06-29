package repository

import (
	"context"

	"github.com/windbnb/notification-service/model"
)

type IRepository interface {
	AddNotification(notification *model.Notification, ctx context.Context) (*model.Notification, error)
	GetUserRecentNotification(userId uint, findLimit uint, ctx context.Context) (*[]model.Notification, error)
	// GetAllUserNotification(userId uint, ctx context.Context) (*[]model.Notification, error)
	// UpdateNotification(notification model.Notification, ctx context.Context) (*model.Notification, error)

	FindUserNotificationSettings(userId uint, ctx context.Context) (*model.NotificationSettings, error)
	AddNotificationSettings(notificationSettings *model.NotificationSettings, ctx context.Context) (*model.NotificationSettings, error)
	DeleteNotificationSettings(userId uint, ctx context.Context) (int64, error)
	UpdateNotificationSettings(notificationSettings *model.NotificationSettings, ctx context.Context) (*model.NotificationSettings, error)
}
