package service

import "LunarAssignment/server/client"

type BridgeService struct {
	client *client.Client
}

func NewBridgeService(client *client.Client) *BridgeService {
	return &BridgeService{client: client}
}
