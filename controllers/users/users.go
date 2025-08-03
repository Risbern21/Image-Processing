package users

import (
	users "image/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	m := users.New()

	if err := c.BodyParser(&m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := m.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(m)
}

func Get(c *fiber.Ctx) error {
	m := users.New()

	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}

	m.ID = userId

	if err := m.Get(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(m)
}

func Update(c *fiber.Ctx) error {
	m := users.New()

	if err := c.BodyParser(&m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("inavlid user id")
	}

	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("inavlid user id")
	}

	m.ID = userId
	if err := m.Update(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func Delete(c *fiber.Ctx) error {
	m := users.New()

	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid user id")
	}

	m.ID = userId

	if err := m.Delete(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("user not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
