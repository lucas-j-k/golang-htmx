package main

import (
	"fmt"
	"log"

	"example.com/go-links-htmx/auth"
	sqlcservice "example.com/go-links-htmx/database"
	"example.com/go-links-htmx/link"
	"example.com/go-links-htmx/user"
	"example.com/go-links-htmx/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT")
	redisPass := viper.Get("REDIS_PASSWORD")

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: fmt.Sprintf("%v", redisPass),
		DB:       0,
	})

	// Connect to MySQL
	if err := sqlcservice.Connect(); err != nil {
		log.Panic("Unable to connect to MYSQL")
	}

	// Register the templates directory with Go html template engine
	engine := html.New("./views", ".html")

	// Initialize SQLC repository
	service := sqlcservice.New(sqlcservice.DB)

	// Initialize the session Manager
	sessionManager := auth.RedisSessionManager{
		Client: redisClient,
	}

	// Set up fiber app with the template engine and our custom error handler
	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		ErrorHandler:      utils.ErrorView,
	})

	// set up controllers
	userController := user.New(service, sessionManager)
	linkController := link.New(service)

	// Hook up middlewares
	app.Use(logger.New())
	app.Use(auth.WithUser(sessionManager))
	app.Use(utils.WithFlash())

	// Serve static assets
	app.Static("/static", "./public")

	// Redirect root url to user links page
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/admin/links", fiber.StatusFound)
	})

	// Public user routes
	app.Post("/register", userController.RegisterPostHandler)
	app.Get("/register", userController.RegisterView)
	app.Get("/login", userController.LoginView)
	app.Post("/login", userController.LoginPostHandler)
	app.Post("/logout", userController.LogoutPostHandler)
	app.Get("/links/:username", userController.UserPublicProfileView)

	// Authenticated routes
	app.Use(auth.AuthGuard)

	app.Get("/profile", userController.UserProfileView)
	app.Get("/admin/links/new", linkController.CreateLinkView)
	app.Post("/admin/links/new", linkController.CreateLinkHandler)
	app.Get("/admin/links", linkController.ManageLinksView)
	app.Get("/admin/links/:id/update", linkController.UpdateLinkView)
	app.Post("/admin/links/:id/update", linkController.UpdateLinkHandler)
	app.Post("/admin/links/:id/delete", linkController.DeleteLinkHandler)

	// Start the server
	fmt.Printf("Server running on port [%v]\n\n", port)
	log.Fatal(app.Listen(fmt.Sprintf("localhost:%v", port)))
}
