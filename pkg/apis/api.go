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

	"github.com/Ankr-network/uscan/pkg/log"
	"github.com/Ankr-network/uscan/share"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rakyll/statik/fs"

	_ "github.com/Ankr-network/uscan/statik"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Apis(ctx context.Context) error {
	svc := fiber.New(fiber.Config{
		Prefork:               false,
		ServerHeader:          "uscan team",
		DisableStartupMessage: true,
	})

	statikFs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	svc.Use("/", filesystem.New(filesystem.Config{
		Root: statikFs,
	}))

	g := svc.Group("/uscan/v1")
	g.Use(recover.New())
	g.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Headers",
		MaxAge:           86400,
	}))
	SetupRouter(g)

	addr := fmt.Sprintf("%s:%s", viper.GetString(share.HttpAddr), viper.GetString(share.HttpPort))
	log.Infof("service boot with: %s \n", addr)
	if err := svc.Listen(addr); err != nil {
		log.Fatalf("service boot with error: %s", err)
	}

	return nil
}
