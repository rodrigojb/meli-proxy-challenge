package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rodrigojb/meli-proxy/meli-proxy/internal"
)

func Handler(u *url.URL, rdb *redis.Client, criteria []internal.Criterion) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.RequestURI, "apidoc") {
			requestDump, err := httputil.DumpRequest(r, true)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(requestDump))
		}

		if strings.Compare("/metrics", r.URL.Path) == 0 {
			promhttp.Handler().ServeHTTP(rw, r)
			return
		}

		request := internal.Request{
			Host: r.Host,
			Path: r.URL.Path,
		}

		if strings.Contains(r.RequestURI, "apidoc") {
			return
		}

		err := internal.LimitRequest(r.Context(), rdb, criteria, request)
		if err != nil && errors.Is(err, internal.TooManyRequestErr) {
			rw.WriteHeader(http.StatusTooManyRequests)
			return
		}

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rp := httputil.NewSingleHostReverseProxy(u)
		r.Host = u.Host
		rp.ServeHTTP(rw, r)
	}
}
