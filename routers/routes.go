package routes_todoSM

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	controller_todoSM "github.com/rrraf1/to-do_social_media/controller"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) SetupRoutes(app *fiber.App) {

	apiLimiter := limiter.New(limiter.Config{
		Max:        15,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			token := c.Get("Authorization")
			if token != "" {
				return token
			}
			return c.IP()
		},
	})

	postRepo := controller_todoSM.Repository{DB: r.DB}
	posts := app.Group("/posts", apiLimiter)

	posts.Get("/due-date", postRepo.GetPostsByRange)

	posts.Get("/:targetDate", postRepo.GetClosestPost)

	posts.Get("/", postRepo.GetPosts)

	posts.Post("/", postRepo.CreatePost)

	posts.Put("/:id", postRepo.UpdatePost)

	posts.Delete("/:id", postRepo.DeletePost)

}
