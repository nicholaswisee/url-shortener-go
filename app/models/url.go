package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ShortCode   string         `gorm:"uniqueIndex;size:20;not null" json:"short_code"`
	OriginalURL string         `gorm:"type:text;not null" json:"original_url"`
	Clicks      int64          `gorm:"default:0" json:"clicks"`
	ExpiresAt   *time.Time     `gorm:"index" json:"expires_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for the URL model
func (URL) TableName() string {
	return "urls"
}

// IsExpired checks if the URL has expired
func (u *URL) IsExpired() bool {
	if u.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*u.ExpiresAt)
}

// IncrementClicks increments the click counter
func (u *URL) IncrementClicks(db *gorm.DB) error {
	return db.Model(u).Update("clicks", gorm.Expr("clicks + 1")).Error
}
