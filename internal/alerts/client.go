package alerts

import (
	"crypto/tls"
	"net/http"

	"github.com/natanel-ziv/oref-alerts-go/pkg/config"
)

type Client struct {
    http  *http.Client
    cfg   *config.Config
}

func NewClient(cfg *config.Config) *Client {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify},
    }
    return &Client{
        http: &http.Client{Transport: tr, Timeout: cfg.Timeout},
        cfg:  cfg,
    }
}

