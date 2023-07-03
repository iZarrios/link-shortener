package types

import "time"

type Link struct {
    ID uint64 `json:"id" gorm:"primary_key"`
    Url string `json:"url" gorm:"not null unique"`
    Redirect string `json:"short_url" gorm:"not null unique"`
    Count uint64 `json:"count" gorm:"not null"`
    CreatedAt time.Time `json:"created_at" gorm:"not null default CURRENT_TIMESTAMP"`
    UpdatedAt time.Time `json:"updated_at" gorm:"not null default CURRENT_TIMESTAMP "`
}
