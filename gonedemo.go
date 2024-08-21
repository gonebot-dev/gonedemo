package main

import (
	"github.com/gonebot-dev/gonebot"
	"github.com/gonebot-dev/gonebot/plugins/builtinplugins"
)

func main() {
	gonebot.LoadPlugin(builtinplugins.Echo)

	gonebot.StartBackend("onebot11")
}
