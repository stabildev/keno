package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/stabildev/keno/models"
)

type createUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Website string `json:"website" binding:"required"`
}

type updateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Website string `json:"website" binding:"required"`
}

func CreateUser(c *fiber.Ctx) error {
	req := &createUserRequest{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{
		Name:    req.Name,
		Email:   req.Email,
		Website: req.Website,
	}

	models.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"user": user,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	models.DB.Find(&users)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"users": users,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User

	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"user": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// check if user exists
	user := models.User{}
	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	// parse response body
	req := &updateUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	// update user
	updateUser := models.User{
		Name:    req.Name,
		Email:   req.Email,
		Website: req.Website,
	}
	models.DB.Model(&user).Updates(&updateUser)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"user": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	// check if user exists
	user := models.User{}
	if err := models.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "Record not found!",
		})
	}

	// delete user
	models.DB.Delete(&user)

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User deleted!",
	})
}
