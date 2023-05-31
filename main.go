package main

import (
	"log"

	"github.com/ArykaAnisaP/arykaanisap/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/ArykaAnisaP/arykaanisap/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/ArykaAnisaP/arykaanisap/docs"
)

// @title TES SWAG
// @version 1.0
// @description This is a sample swagger server

//@contact.name API Support
// @contact.name https://github.com/ArykaAnisaP
// @contact.email arykaanisap22@gmail.com

// @host aryka.herokuapp.com/
// @BasePath /
//@schemes https http

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
