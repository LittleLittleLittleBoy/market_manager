package main

import (
	"github.com/littlelittlelittleboy/market_manager/pkg/settings"
	"github.com/littlelittlelittleboy/market_manager/routers"
)

func main() {
	settings.Setup()

	r := routers.InitRouter()

	r.Run()
}
