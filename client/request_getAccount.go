package client

import (
	"LunarAssignment/LunarServer/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *Client) GetAccount(id int) (*model.Account, error) {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:6000/account/%v", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer SecretInternalKey"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status %s", resp.Status)
		return nil, err
	}

	account := &model.Account{}
	json.NewDecoder(resp.Body).Decode(account)

	return account, nil
}
