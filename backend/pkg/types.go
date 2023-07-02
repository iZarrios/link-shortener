package types

type Link struct {
    ID uint64 `json:"id" gorm:"primary_key"`
    URL string `json:"url" gorm:"not null unique"`
    Redirect string `json:"short_url" gorm:"not null unique"`
    Count uint64 `json:"count" gorm:"not null"`
    CreatedAt int64 `json:"created_at" gorm:"not null" gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt int64 `json:"updated_at" gorm:"not null" gorm:"default:CURRENT_TIMESTAMP"`
}
