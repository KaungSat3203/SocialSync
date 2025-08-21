package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"userId"`
	Platform       string    `json:"platform"`
	PlatformPostID string    `json:"platformPostId"`
	Message        string    `json:"message"`
	MediaURLs      []string  `json:"mediaUrls,omitempty"` // multiple media URLs
	PostedAt       time.Time `json:"postedAt"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func SavePost(db *sql.DB, post Post) error {
	mediaURLsJSON, err := json.Marshal(post.MediaURLs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO posts (
			id, user_id, platform, platform_post_id, message,
			media_urls, posted_at, status, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`

	_, err = db.Exec(
		query,
		post.ID,
		post.UserID,
		post.Platform,
		post.PlatformPostID,
		post.Message,
		mediaURLsJSON,
		post.PostedAt,
		post.Status,
		post.CreatedAt,
		post.UpdatedAt,
	)

	return err
}
