package controller_todoSM

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	model_todoSM "github.com/rrraf1/to-do_social_media/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func newPostResponse(post model_todoSM.Post) model_todoSM.PostResponse {
	return model_todoSM.PostResponse{
		Id:       post.Id,
		Title:    post.Title,
		Brand:    post.Brand,
		Platform: post.Platform,
		DueDate:  post.DueDate.Format("2006-01-02"),
	}
}

// GetPosts godoc
//
//	@Summary		Retrieve all posts
//	@Description	Get all social media posts with a response wrapper containing message and data
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model_todoSM.PostsResponse			"Posts retrieved successfully"
//	@Failure		500	{object}	model_todoSM.ServerErrorResponse	"Internal server error"
//	@Router			/posts [get]
func (r *Repository) GetPosts(context *fiber.Ctx) error {
	var posts []model_todoSM.Post

	if err := r.DB.Find(&posts).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "error getting data"})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Posts retrieved successfully", "data": posts})
}

// GetPostsByRange godoc
//
//	@Summary		Retrieve posts by due date range
//	@Description	Retrieve posts within the specified date range. Dates must be in YYYY-MM-DD format.
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			date1	query		string								true	"Start date in YYYY-MM-DD format"
//	@Param			date2	query		string								true	"End date in YYYY-MM-DD format"
//	@Success		200		{object}	model_todoSM.PostsResponse			"Posts found within range"
//	@Failure		400		{object}	model_todoSM.ErrorResponse			"Invalid date format or missing parameter"
//	@Failure		500		{object}	model_todoSM.ServerErrorResponse	"Database error"
//	@Router			/posts/due-date [get]
func (r *Repository) GetPostsByRange(context *fiber.Ctx) error {
	var posts []model_todoSM.Post
	date1Str := context.Query("date1")
	date2Str := context.Query("date2")

	if date1Str == "" || date2Str == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Both start and end dates are required"})
	}

	layout := "2006-01-02"
	date1, err := time.Parse(layout, date1Str)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Invalid start date format. Use YYYY-MM-DD"})
	}

	date2, err := time.Parse(layout, date2Str)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Invalid end date format. Use YYYY-MM-DD"})
	}

	if date2.Before(date1) {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "End date cannot be earlier than start date"})
	}

	result := r.DB.Where("due_date BETWEEN ? AND ?", date1, date2).Find(&posts)
	if result.Error != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "Database error", "details": result.Error.Error()})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Posts found!", "data": posts})
}

// GetClosestPost godoc
//
//	@Summary		Retrieve posts within a day range from now
//	@Description	Get posts with due dates between now and now+targetDate days.
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			targetDate	path		int									true	"Number of days from now to filter posts"
//	@Success		200			{object}	model_todoSM.PostsResponse			"Posts found in the specified range"
//	@Failure		400			{object}	model_todoSM.ErrorResponse			"Invalid targetDate parameter"
//	@Failure		500			{object}	model_todoSM.ServerErrorResponse	"Internal server error"
//	@Router			/posts/{targetDate} [get]
func (r *Repository) GetClosestPost(context *fiber.Ctx) error {
	var posts []model_todoSM.Post

	selectRangeStr := context.Params("targetDate")
	selectRange, err := strconv.Atoi(selectRangeStr)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Invalid targetDate range parameter"})
	}

	nowDate := time.Now()
	targetDate := time.Now().AddDate(0, 0, selectRange)

	if err := r.DB.Where("due_date BETWEEN ? AND ?", nowDate, targetDate).Find(&posts).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "Failed to get posts"})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Posts found", "data": posts})
}

// CreatePost godoc
//
//	@Summary		Create a new post
//	@Description	Create a new social media post with the required fields
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			post	body		model_todoSM.PostInput				true	"Post data"
//	@Success		201		{object}	model_todoSM.SinglePostResponse		"Post created successfully"
//	@Failure		400		{object}	model_todoSM.ErrorResponse			"Invalid request body or missing fields"
//	@Failure		500		{object}	model_todoSM.ServerErrorResponse	"Failed to create post due to server error"
//	@Router			/posts [post]
func (r *Repository) CreatePost(context *fiber.Ctx) error {
	var input model_todoSM.PostInput

	if err := context.BodyParser(&input); err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if input.DueDate == nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "due_date is required"})
	}

	dueDate, err := time.Parse("2006-01-02", *input.DueDate)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid due_date format. Use YYYY-MM-DD"})
	}

	post := model_todoSM.Post{
		Title:    *input.Title,
		Brand:    *input.Brand,
		Platform: *input.Platform,
		DueDate:  dueDate,
	}

	if err := r.DB.Create(&post).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post", "details": err.Error()})
	}

	return context.Status(http.StatusCreated).JSON(fiber.Map{"message": "Post created successfully", "data": newPostResponse(post)})
}



// UpdatePost godoc
//
//	@Summary		Update an existing post
//	@Description	Update an existing social media post by ID. Only non-empty fields will be updated.
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int									true	"Post ID"
//	@Param			post	body		model_todoSM.PostInput				true	"Post data"
//	@Success		200		{object}	model_todoSM.SinglePostResponse		"Post updated successfully"
//	@Failure		400		{object}	model_todoSM.ErrorResponse			"Invalid request body or missing ID"
//	@Failure		404		{object}	model_todoSM.NotFoundResponse		"Post not found"
//	@Failure		500		{object}	model_todoSM.ServerErrorResponse	"Failed to update post due to server error"
//	@Router			/posts/{id} [put]
func (r *Repository) UpdatePost(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Post ID is required"})
	}

	var post model_todoSM.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Database error", "details": err.Error()})
	}

	var input model_todoSM.PostInput
	if err := context.BodyParser(&input); err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if input.Title != nil {
		post.Title = *input.Title
	}
	if input.Brand != nil {
		post.Brand = *input.Brand
	}
	if input.Platform != nil {
		post.Platform = *input.Platform
	}
	if input.DueDate != nil {
		dueDate, err := time.Parse("2006-01-02", *input.DueDate)
		if err != nil {
			return context.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid due_date format. Use YYYY-MM-DD"})
		}
		post.DueDate = dueDate
	}

	if err := r.DB.Save(&post).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post", "details": err.Error()})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Post updated successfully", "data": newPostResponse(post)})
}

// DeletePost godoc
//
//	@Summary		Delete a post
//	@Description	Delete a social media post by its ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int									true	"Post ID"
//	@Success		200	{object}	model_todoSM.StandardResponse		"Post deleted successfully"
//	@Failure		400	{object}	model_todoSM.ErrorResponse			"Post ID is missing"
//	@Failure		404	{object}	model_todoSM.NotFoundResponse		"Post not found"
//	@Failure		500	{object}	model_todoSM.ServerErrorResponse	"Failed to delete post due to server error"
//	@Router			/posts/{id} [delete]
func (r *Repository) DeletePost(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "Post ID is required"})
	}

	var post model_todoSM.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(&fiber.Map{"error": "Post not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "Database error", "details": err.Error()})
	}

	if err := r.DB.Delete(&post).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"error": "Failed to delete post", "details": err.Error()})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Post deleted successfully"})
}
