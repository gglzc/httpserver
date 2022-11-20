package main

import (
	"github.com/gglzc/httpServer/controller"
	"github.com/gofiber/fiber/v2"
)


 func main(){
	
	router := fiber.New()
	router.Get("/Win",controller.CalWin)
	router.Listen(":8080")
}


