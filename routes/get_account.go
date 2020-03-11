package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (rs *RestServer) GetAccount(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token != "Bearer SuperSecretAuth" {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
	} else {

		vars := mux.Vars(r)
		id, found := vars["id"]
		if !found {
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		account, err := rs.service.GetAccount(idInt)
		if err != nil {
			if account != nil {
				w.WriteHeader(500)
				w.Write([]byte("Internal Error"))
				return
			}
			w.WriteHeader(400)
			w.Write([]byte("Bad Request: No such account"))
			return
		}

		body, err := json.Marshal(account)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		w.WriteHeader(200)
		w.Write(body)

	}

}
