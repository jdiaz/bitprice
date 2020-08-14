package main

import "bitprice/server"

const port = 8080

func main() {
	server.StartPriceServer(port)
}
