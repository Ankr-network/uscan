package apis

import (
	"context"
	"fmt"
	"log"

	"github.com/Ankr-network/uscan/share"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Apis(ctx context.Context) error {
	svc := fiber.New(fiber.Config{
		Prefork:               false,
		ServerHeader:          "uscan team",
		DisableStartupMessage: true,
	})

	g := svc.Group("/uscan/v1")

	g.Get("/hello/:name", func(c *fiber.Ctx) error {
		c.JSON(fmt.Sprintf("Hi %s", c.Params("name")))
		return nil
	})

	addr := fmt.Sprintf("%s:%s", viper.GetString(share.HttpAddr), viper.GetString(share.HttpPort))
	log.Printf("service boot with: %s \n", addr)
	svc.Listen(addr)

	return nil
}
