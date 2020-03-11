package service

import (
	"LunarAssignment/LunarServer/model"
)

func (s *BridgeService) GetAccount(id int) (*model.Account, error) {

	account, err := s.client.GetAccount(id)
	if err != nil {
		return nil, err
	}

	return account, nil

}
