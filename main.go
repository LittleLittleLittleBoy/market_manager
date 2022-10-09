package main

import (
	"log"

	"github.com/littlelittlelittleboy/market_manager/pkg/ocr"
	"github.com/littlelittlelittleboy/market_manager/pkg/settings"
)

func main() {
	settings.Setup()

	log.Println(ocr.GetImageContent("https://imgse.com/i/xJjwm8"))

	// r := routers.InitRouter()

	// r.Run()
}
