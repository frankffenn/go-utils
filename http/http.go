package http

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	// Client return a http client
	Client = resty.New()
)

func init() {
	Client.SetTimeout(time.Second * 30)
	Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
