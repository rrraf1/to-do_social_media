package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	connection_todoSM "github.com/rrraf1/to-do_social_media/connection"
	routes_todoSM "github.com/rrraf1/to-do_social_media/routers"
	_ "github.com/rrraf1/to-do_social_media/docs"

)

// @title			TO-DO Api
// @version		1.0
// @description	API to manage social media post
// @host			localhost:3000
// @BasePath		/
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env")
	}

	config := &connection_todoSM.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := connection_todoSM.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := connection_todoSM.Migrate(db); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	r := routes_todoSM.NewRepository(db)
	r.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
