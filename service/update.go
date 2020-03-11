package service

import "LunarAssignment/LunarServer/model"

func (s *BridgeService) Update(account *model.Account) (string, error) {

	response, err := s.client.UpdateAccount(account)
	if err != nil {
		return "", err
	}

	return response, nil
}
