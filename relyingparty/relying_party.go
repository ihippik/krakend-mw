package relyingparty

import (
	"github.com/devopsfaith/krakend/config"
)

// RelyingParty checks access to endpoint by access_token.
type RelyingParty struct {
	cfg *sfConfig
}

// New creates a new RelyingParty.
func New(e config.ExtraConfig) (*RelyingParty, error) {

	return &RelyingParty{
		cfg: getSFConfig(e),
	}, nil
}
