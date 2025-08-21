package main

import (
	"log"
	"net/http"
	"os"

	"social-sync-backend/lib"
	"social-sync-backend/routes"
	"social-sync-backend/utils"
	// "social-sync-backend/workers"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	_ "github.com/lib/pq"
)

// CORSMiddleware sets CORS headers dynamically based on environment
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appEnv := os.Getenv("APP_ENV")
		var allowedOrigin string

		if appEnv == "production" {
			allowedOrigin = os.Getenv("FRONTEND_URL_PROD")
		} else {
			allowedOrigin = os.Getenv("FRONTEND_URL_LOCAL")
		}

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load .env only in development
	if os.Getenv("APP_ENV") != "production" {
		if _, err := os.Stat(".env"); err == nil {
			if err := godotenv.Load(); err != nil {
				log.Printf("⚠️ Warning: Error loading .env file: %v", err)
			} else {
				log.Println("✅ .env file loaded successfully.")
			}
		} else {
			log.Println("ℹ️ .env file not found locally, continuing with environment variables.")
		}
	} else {
		log.Println("✅ Running in production environment.")
	}

	// Connect to DB
	lib.ConnectDB()
	defer func() {
		if lib.DB != nil {
			if err := lib.DB.Close(); err != nil {
				log.Printf("❌ Error closing database: %v", err)
			} else {
				log.Println("✅ Database connection closed.")
			}
		}
	}()
	log.Println("✅ Connected to PostgreSQL DB!")

	// Cloudinary setup
	if err := lib.InitCloudinary(); err != nil {
		log.Fatalf("❌ Failed to initialize Cloudinary: %v", err)
	}
	log.Println("✅ Cloudinary initialized!")

	// CRON Jobs
	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger),
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))

	// Social account sync every 24h
	if _, err := c.AddFunc("@every 24h", func() {
		log.Println("🔁 Running scheduled social account sync...")
		utils.SyncAllSocialAccountsTask(lib.DB)
	}); err != nil {
		log.Fatalf("❌ Failed to schedule social account sync: %v", err)
	}

	// Post analytics sync every 6h
	// if _, err := c.AddFunc("@every 1m", func() {
	// 	log.Println("📊 Running scheduled Facebook analytics sync...")
	// 	utils.SyncFacebookAnalytics(lib.DB)
	// }); err != nil {
	// 	log.Fatalf("❌ Failed to schedule Facebook analytics sync: %v", err)
	// }

	c.Start()
	defer c.Stop()
	log.Println("✅ All cron jobs started.")

	// HTTP Server
	r := routes.InitRoutes()
	handler := CORSMiddleware(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server running at: http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
