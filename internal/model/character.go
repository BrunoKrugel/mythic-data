package model

import "time"

type Characters struct {
	Name              string    `json:"name"`
	Race              string    `json:"race"`
	Class             string    `json:"class"`
	ActiveSpecName    string    `json:"active_spec_name"`
	ActiveSpecRole    string    `json:"active_spec_role"`
	Gender            string    `json:"gender"`
	Faction           string    `json:"faction"`
	AchievementPoints int       `json:"achievement_points"`
	HonorableKills    int       `json:"honorable_kills"`
	ThumbnailURL      string    `json:"thumbnail_url"`
	Region            string    `json:"region"`
	Realm             string    `json:"realm"`
	LastCrawledAt     time.Time `json:"last_crawled_at"`
	ProfileURL        string    `json:"profile_url"`
	ProfileBanner     string    `json:"profile_banner"`
}
