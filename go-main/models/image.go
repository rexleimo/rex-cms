package models

import "time"

type Image struct {
	Id        int       `json:"id" bson:"_id,omitempty gorm:"primaryKey"`
	URL       string    `json:"url" bson:"url,omitempty"`
	Width     int       `json:"width" bson:"width,omitempty"`
	Height    int       `json:"height" bson:"height,omitempty"`
	Size      int       `json:"size" bson:"size,omitempty"`
	MimeType  string    `json:"mime_type" bson:"mime_type,omitempty"`
	CreateAt  time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
