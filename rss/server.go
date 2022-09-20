// Copyright (c) 2022 coqchain team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rss

import (
	"fmt"

	"github.com/Ankr-network/coqchain/log"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/urfave/cli.v1"
)

func (r *Rebuild) Start(ctx *cli.Context) {
	svc := fiber.New(fiber.Config{
		ServerHeader:          "coqchain team",
		Prefork:               false,
		DisableStartupMessage: true,
	})

	svc.Post("/block", Block)
	svc.Post("/tx", Tx)

	addr := fmt.Sprintf("%s:%s", r.host, r.port)
	log.Info("rebuid", "addr", addr)
	svc.Listen(addr)
}
