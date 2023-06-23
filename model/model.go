package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserId     uint               `bson:"userId"`
	NotifTitle string             `bson:"notifTitle"`
	NotifType  string             `bson:"notifType"`
	Seen       bool               `bson:"seen"`
}

type NotificationSettings struct {
	ID                  primitive.ObjectID `bson:"_id"`
	UserId              uint               `bson:"userId"`
	ResRequest          bool               `bson:"resRequest"`
	ResCancel           bool               `bson:"resCancel"`
	HostReview          bool               `bson:"hostReview"`
	AccommodationReview bool               `bson:"accommodationReview"`
	ReservationResponse bool               `bson:"reservationResponse"`
}
