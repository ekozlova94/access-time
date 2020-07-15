package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ekozlova94/internal/ctxutils"
	"github.com/ekozlova94/internal/model"
	"github.com/ekozlova94/internal/storage"
	"github.com/ekozlova94/pkg/forms"
)

func SetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "request not post", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetFirstTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "request not get", http.StatusMethodNotAllowed)
		return
	}
	result, err := storage.Get(ctxutils.IpFromContext(r.Context()), true, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result == nil {
		result.Time = time.Now()
	}
	createResponse(forms.NewAccessTime(&result.Time, ""), w)
}

func GetLastTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "request not post", http.StatusMethodNotAllowed)
		return
	}
	var m model.AccessTime
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := storage.Get(m.Ip, false, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var form *forms.AccessTime
	if result != nil {
		form = forms.NewAccessTime(&result.Time, "")
	} else {
		form = forms.NewAccessTime(nil, fmt.Errorf("no info about this %s", m.Ip).Error())
	}
	createResponse(form, w)
}

func createResponse(result *forms.AccessTime, w http.ResponseWriter) {
	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
