package conf

import "os"

func GetPort() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
