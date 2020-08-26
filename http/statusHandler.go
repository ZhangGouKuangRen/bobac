package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (sh *statusHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	requestMethod := r.Method
	if requestMethod != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	statJson, err := json.Marshal(sh.GetStat())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(statJson)
}
