package main

import (
	"anylogibtc/api/handler"
)

func main() {
	var s handler.Server = handler.NewEchoServer(3000)

	s.Run()
}
