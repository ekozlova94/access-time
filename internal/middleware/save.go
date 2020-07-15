package middleware

import (
	"net/http"
	"time"

	"github.com/ekozlova94/internal/ctxutils"
	"github.com/ekozlova94/internal/model"
	"github.com/ekozlova94/internal/storage"
)

func Save(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := ctxutils.IpFromContext(r.Context())
		var data = model.AccessTime{
			Ip:   clientIP,
			Time: time.Now(),
		}
		if err := storage.Save(data, r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
