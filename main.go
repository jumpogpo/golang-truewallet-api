package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func redeemCode(c *gin.Context, code string, phoneNumber string) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(res.StatusCode, "application/json", body)
}

func main() {
	port := "1500"
	err := godotenv.Load()

	if err == nil {
		port = os.Getenv("PORT")
	}

	router := gin.Default()

	router.POST("/redeem/:code", func(c *gin.Context) {
		code := c.Param("code")

        var requestBody struct {
            MobilePhone string `json:"mobile"`
        }

		if err := c.BindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

		redeemCode(c, code, requestBody.MobilePhone)
	})

	router.Run(":" + port)
}