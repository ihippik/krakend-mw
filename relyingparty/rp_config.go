package relyingparty

import (
	"fmt"

	"github.com/devopsfaith/krakend/config"
)

// sfConfig is the custom config struct containing the params for the Auth Checker.
type sfConfig struct {
	// AuthServerURI is URI of admin API of OIDC.
	AuthServerURI string `json:"authServerURI"`
	// AuthorizationEndpoint is a relative path to authorization endpoint.
	AuthorizationEndpoint string `json:"authorizationEndpoint"`
	// TokenEndpoint is a relative path to authorization endpoint.
	TokenEndpoint string `json:"tokenEndpoint"`
}

// sfZeroCfg is the zero value for the sfConfig struct.
var sfZeroCfg = sfConfig{}

// sfNamespace is the key for getting config from extraConfig global section.
// Use underscores instead of dots.
const sfNamespace = "git_omprussia_ru/auth/krakendself"

// getSFConfig parses the extra config for the Auth Checker.
func getSFConfig(e config.ExtraConfig) *sfConfig {
	v, ok := e[sfNamespace]
	if !ok {
		return &sfZeroCfg
	}
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return &sfZeroCfg
	}

	cfg := &sfConfig{}
	if v, ok := tmp["auth_server_uri"]; ok {
		cfg.AuthServerURI = fmt.Sprintf("%v", v)
	}
	if v, ok := tmp["authorization_endpoint"]; ok {
		cfg.AuthorizationEndpoint = fmt.Sprintf("%v", v)
	}
	if v, ok := tmp["token_endpoint"]; ok {
		cfg.TokenEndpoint = fmt.Sprintf("%v", v)
	}

	return cfg
}
