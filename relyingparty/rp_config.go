package relyingparty

import (
	"fmt"

	"github.com/devopsfaith/krakend/config"
)

// rpConfig is the custom config struct containing the params for the Auth Checker.
type rpConfig struct {
	TokenSecret string `json:"token_secret"`
}

// rpNamespace is the key for getting config from extraConfig global section.
// Use underscores instead of dots.
const rpNamespace = "github_com/ihippik/krakend-mw/relyingparty"

// rpZeroCfg is the zero value for the rpConfig struct.
var rpZeroCfg = rpConfig{}

// getRpConfig parses the extra config for the RP.
func getRpConfig(e config.ExtraConfig) *rpConfig {
	v, ok := e[rpNamespace]
	if !ok {
		return &rpZeroCfg
	}
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return &rpZeroCfg
	}

	cfg := &rpConfig{}
	if v, ok := tmp["token_secret"]; ok {
		cfg.TokenSecret = fmt.Sprintf("%v", v)
	}
	return cfg
}
