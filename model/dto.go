package model

type NotificationRequest struct {
	UserId     uint
	NotifTitle string
	NotifType  string
	Seen       bool
}

type NotificationSettingsRequest struct {
	UserId              uint
	ResRequest          bool
	ResCancel           bool
	HostReview          bool
	AccommodationReview bool
	ReservationResponse bool
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
