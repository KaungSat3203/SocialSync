package utils

import(
	"os"
)

func GetCallbackURL(provider string) string {
	env := os.Getenv("APP_ENV")

	switch provider {
	case "facebook":
		if env == "production" {
			return os.Getenv("FACEBOOK_CALLBACK_PROD")
		}
		return os.Getenv("FACEBOOK_CALLBACK_LOCAL")

	case "google":
		if env == "production" {
			return os.Getenv("GOOGLE_CALLBACK_PROD")
		}
		return os.Getenv("GOOGLE_CALLBACK_LOCAL")

	case "twitter":
		if env == "production" {
			return os.Getenv("TWITTER_CALLBACK_PROD")
		}
		return os.Getenv("TWITTER_CALLBACK_LOCAL")

	case "youtube":
		if env == "production" {
			return os.Getenv("YOUTUBE_CALLBACK_PROD")
		}
		return os.Getenv("YOUTUBE_CALLBACK_LOCAL")

	case "instagram":
		if env == "production" {
			return os.Getenv("INSTAGRAM_CALLBACK_PROD")
		}
		return os.Getenv("INSTAGRAM_CALLBACK_LOCAL")

	case "mastodon":
		if env == "production" {
			return os.Getenv("MASTODON_CALLBACK_PROD")
		}
		return os.Getenv("MASTODON_CALLBACK_LOCAL")
	}

	return ""
}


func GetFrontendURL() string {
	if os.Getenv("APP_ENV") == "production" {
		return os.Getenv("FRONTEND_URL_PROD")
	}
	return os.Getenv("FRONTEND_URL_LOCAL")
}