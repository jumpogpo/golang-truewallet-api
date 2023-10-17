package config

import "os"

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "1500"
	}

	return port
}