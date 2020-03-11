package main

import (
	"LunarAssignment/server/client"
	"LunarAssignment/server/routes"
	"LunarAssignment/server/service"
)

func main() {

	c := client.NewClient()
	s := service.NewBridgeService(c)
	routes.StartRestServer(s)
}
