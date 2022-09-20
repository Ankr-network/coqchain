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
	"github.com/Ankr-network/coqchain/core"
	"gopkg.in/urfave/cli.v1"
)

type Rebuild struct {
	host string
	port string
	bc   *core.BlockChain
}

var svc *Rebuild

func NewRebuild(bc *core.BlockChain, host, port string) *Rebuild {
	svc = &Rebuild{bc: bc, host: host, port: port}
	return svc
}

func Start(ctx *cli.Context) {
	go svc.Start(ctx)
}
