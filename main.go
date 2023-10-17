package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func redeemCode(c *fiber.Ctx, code string, phoneNumber string) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	transport := &http.Transport{
		TLSClientConfig:   tlsConfig,
		DisableKeepAlives: false,
	}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://gift.truemoney.com/campaign/vouchers/%s/redeem", code),
		strings.NewReader(fmt.Sprintf(`{"mobile": "%s"}`, phoneNumber)),
	)

	userAgent := "MyApp/" + uuid.NewString()
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	res, err := client.Do(req)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return
	}

	c.Status(res.StatusCode).Send(body)
}

func main() {
	port := "1500"
	err := godotenv.Load()

	if err == nil {
		port = os.Getenv("PORT")
	}

	router := fiber.New()

    router.Post("/redeem/:code", func(c *fiber.Ctx) error {
        code := c.Params("code")

        var requestBody struct {
            MobilePhone string `json:"mobile"`
        }

        if err := c.BodyParser(&requestBody); err != nil {
            c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        redeemCode(c, code, requestBody.MobilePhone)
		return nil
    })

	router.Listen(fmt.Sprintf(":%s", port))
}