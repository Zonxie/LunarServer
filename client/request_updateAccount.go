package client

import (
	"LunarAssignment/server/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (client *Client) UpdateAccount(account *model.Account) (string, error) {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(account)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:6000/update", buf)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer SecretInternalKey"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("bad status %s", resp.Status)
		return "", err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(responseData), nil
}
