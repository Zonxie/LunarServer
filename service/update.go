package service

import "LunarAssignment/server/model"

func (s *BridgeService) Update(account *model.Account) (string, error) {

	response, err := s.client.UpdateAccount(account)
	if err != nil {
		return "", err
	}

	return response, nil
}
