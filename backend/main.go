package main

import (
	"log"
	"net/http"
	"os" // Make sure os is imported for os.Getenv and os.Stat

	"social-sync-backend/lib"
	"social-sync-backend/routes"
	"social-sync-backend/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	_ "github.com/lib/pq"
)

// CORSMiddleware sets CORS headers.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	// --- MODIFICATION START ---

	// Load environment variables conditionally
	// Check if APP_ENV is NOT "production". You will set APP_ENV=production on Render.
	if os.Getenv("APP_ENV") != "production" {
		// Check if .env file actually exists before trying to load it
		if _, err := os.Stat(".env"); err == nil {
			if err := godotenv.Load(); err != nil {
				// Log a warning if .env fails to load in non-production, but don't fatal
				// This allows the app to proceed if variables are set through other means (e.g., CI/CD)
				log.Printf("⚠️ Warning: Error loading .env file (continuing assuming variables are set externally): %v", err)
			} else {
				log.Println("✅ .env file loaded successfully.")
			}
		} else {
			log.Println("ℹ️ .env file not found locally, continuing with environment variables.")
		}
	} else {
		log.Println("✅ Running in production environment, loading variables from OS environment.")
	}

	// --- MODIFICATION END ---

	// Initialize database
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

	// Initialize Cloudinary
	if err := lib.InitCloudinary(); err != nil {
		// This will now check variables loaded from Render's environment
		log.Fatalf("❌ Failed to initialize Cloudinary: %v", err)
	}
	log.Println("✅ Cloudinary initialized!")

	// Setup cron job for social account sync
	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger),
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))
	if _, err := c.AddFunc("@every 24h", func() { // Note: Your log says "every 12h" but code is "every 24h"
		log.Println("🔁 Running scheduled social account sync...")
		utils.SyncAllSocialAccountsTask(lib.DB)
	}); err != nil {
		log.Fatalf("❌ Failed to schedule cron: %v", err)
	}
	c.Start()
	defer c.Stop()
	log.Println("✅ Cron job started (every 24h).") // Corrected log message for clarity

	// Setup routes and middleware
	r := routes.InitRoutes()
	handler := CORSMiddleware(r)

	// Start server
	port := os.Getenv("PORT") // This will now correctly get PORT from Render's environment
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server running at: http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}