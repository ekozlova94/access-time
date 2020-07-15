package middleware

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ekozlova94/internal/ctxutils"
)

func Ip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, err := FromRequest(r)
		if err != nil {
			log.Fatal(err)
		}
		withContext := r.WithContext(ctxutils.NewIpContext(r.Context(), ip))
		next.ServeHTTP(w, withContext)
	})
}

func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
	}
	return userIP, nil
}
