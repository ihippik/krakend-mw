package relyingparty

import (
	"github.com/devopsfaith/krakend/config"
	"github.com/mitchellh/mapstructure"
)

// rpConfig is the custom config struct containing the params for the Auth Checker.
type rpConfig struct {
	TokenSecret string `mapstructure:"token_secret"`
}

type epConfig struct {
	Roles []string `json:"roles"`
}

// rpNamespace is the key for getting config from extraConfig global section.
// Use underscores instead of dots.
const rpNamespace = "github_com/ihippik/krakend-mw/relyingparty"

// rpZeroCfg is the zero value for the rpConfig struct.
var rpZeroCfg = rpConfig{}

// getRpConfig parses the extra config for the RP.
func getRpConfig(e config.ExtraConfig) (*rpConfig, error) {
	cfg := new(rpConfig)
	input, ok := e[rpNamespace]
	if !ok {
		return &rpZeroCfg, nil
	}
	err := mapstructure.Decode(input, cfg)
	if err != nil {
		return &rpZeroCfg, err
	}
	return cfg, nil
}

// getRpConfig parses the extra config for the Endpoint.
func getEpConfig(extra interface{}) (*epConfig, error) {
	cfg := new(epConfig)
	err := mapstructure.Decode(extra, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
