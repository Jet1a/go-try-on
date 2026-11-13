package main

import (
	"log"

	"github.com/Jet1a/go-try-on/api/internal/config"
	"github.com/Jet1a/go-try-on/api/internal/db"
	"github.com/Jet1a/go-try-on/api/internal/server"
)

func main()  {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	gormDB := db.Connect(&cfg)

	app := server.New(gormDB, &cfg)

	log.Fatal(app.Listen(":" + cfg.DBPort))

	// fmt.Println("Application is running...")

	// http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "applications/json")
	// 	fmt.Fprintf(w, `{"message": Hello from Go Backend}`)
	// })

	// fmt.Println("Server is running on port 8080")
	// http.ListenAndServe(":8080", nil)
}
