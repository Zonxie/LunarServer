package main

import (
	"LunarAssignment/LunarServer/client"
	"LunarAssignment/LunarServer/routes"
	"LunarAssignment/LunarServer/service"
)

func main() {

	c := client.NewClient()
	s := service.NewBridgeService(c)
	routes.StartRestServer(s)
}
