/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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

	svc.Static("/", "web")
	svc.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("web/index.html")
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
