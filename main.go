package main

import (
	"bitprice/web"
	"log"
)

const port = 8080

func main() {
	log.Println("Booting web server")
	web.StartPriceServer(port)
}
