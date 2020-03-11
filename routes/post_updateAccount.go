package routes

import (
	"LunarAssignment/server/model"
	"encoding/json"
	"log"
	"net/http"
)

func (rs *RestServer) UpdateAccount(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token != "Bearer SuperSecretAuth" {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	} else {

		account := &model.Account{}
		json.NewDecoder(r.Body).Decode(account)

		increasedVal, err := rs.service.Update(account)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(400)
			w.Write([]byte("Internal error"))
			return
		}

		body, err := json.Marshal(increasedVal)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(400)
			w.Write([]byte("Internal error"))
			return
		}

		w.WriteHeader(201)
		w.Write(body)
	}

}
