package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserId     uint               `bson:"userId"`
	Timestamp  time.Time          `bson:"timestamp"`
	NotifTitle string             `bson:"notifTitle"`
	NotifType  string             `bson:"notifType"`
	Seen       bool               `bson:"seen"`
}

type NotificationSettings struct {
	ID                  primitive.ObjectID `bson:"_id"`
	UserId              uint               `bson:"userId"`
	ResRequest          bool               `bson:"resRequest"`
	ResCancel           bool               `bson:"resCancel"`
	HostReview          bool               `bson:"HostReview"`
	AccommodationReview bool               `bson:"accommodationReview"`
	ReservationResponse bool               `bson:"rservationResponse"`
}
