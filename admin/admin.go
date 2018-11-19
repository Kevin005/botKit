package admin

import (
	"github.com/ifchange/botKit/config"
	"github.com/ifchange/botKit/signature"
	"io"
	"net/http"
)

var (
	cfg *config.URIConfig
)

func init() {
	cfg = config.GetConfig().URI
	if cfg == nil || cfg.Admin == "" {
		panic("botKit-admin products config is nil")
	}
}

func getURI() string {
	return cfg.Admin
}

func AdminPOST(subURI string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", getURI()+subURI, body)
	if err != nil {
		return nil, err
	}
	err = signature.AddSignature(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
