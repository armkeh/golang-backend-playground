package main

import (
	"github.com/armkeh/golang-backend-playground/internal/gin/routes"
)

func main() {
	r := routes.Setup()

	r.Run(":8080")
}
