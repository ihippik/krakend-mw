package relyingparty

import (
	"github.com/devopsfaith/krakend/config"
)

// RelyingParty checks access to endpoint by access_token.
type RelyingParty struct {
	cfg *rpConfig
}

// New creates a new RelyingParty.
func New(e config.ExtraConfig) (*RelyingParty, error) {
	// TODO initialize MW.

	// common config
	return &RelyingParty{
		cfg: getRpConfig(e),
	}, nil
}
