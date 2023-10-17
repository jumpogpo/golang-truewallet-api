package handler

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RedemptionService struct{}

func (rs *RedemptionService) RedeemCode(c *fiber.Ctx, code string, phoneNumber string) error {
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
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	c.Status(res.StatusCode).Send(body)
	return nil
}