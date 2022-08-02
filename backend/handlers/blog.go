package handlers

import (
	"github.com/AlfarrelIndrawan/cmlabs-onboarding/backend/config"
	"github.com/AlfarrelIndrawan/cmlabs-onboarding/backend/entities"

	"github.com/gofiber/fiber/v2"
)

func GetBlogs(c *fiber.Ctx) error {
	var blogs []entities.Blog

	config.Database.Find(&blogs)
	return c.Status(200).JSON(blogs)
}

func GetBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog entities.Blog

	result := config.Database.Find(&blog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&blog)
}

func CreateBlog(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	blog := new(entities.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&blog)
	return c.Status(201).JSON(blog)
}

func UpdateBlog(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	blog := new(entities.Blog)
	id := c.Params("id")

	if err := c.BodyParser(blog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&blog)
	return c.Status(200).JSON(blog)
}

func RemoveBlog(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	id := c.Params("id")
	var blog entities.Blog

	result := config.Database.Delete(&blog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
