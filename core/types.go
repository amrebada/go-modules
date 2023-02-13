package core

import "github.com/gofiber/fiber/v2"

type FiberHandler = func(*fiber.Ctx) error
