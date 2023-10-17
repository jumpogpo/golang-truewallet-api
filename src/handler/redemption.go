package handler

import "github.com/gofiber/fiber/v2"

type RedemptionHandler interface {
	redeemCode(c *fiber.Ctx, code string, phoneNumber string) error
}