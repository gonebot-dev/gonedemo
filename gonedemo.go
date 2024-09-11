package main

import (
	"github.com/gonebot-dev/gonebot"
	"github.com/gonebot-dev/grepo"
)

func main() {
	grepo.Require("echo", "latest")
	grepo.Require("status", "latest")

	gonebot.StartBackend("onebot11")
}
