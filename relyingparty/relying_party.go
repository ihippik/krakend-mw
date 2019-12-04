package relyingparty

import (
	"github.com/devopsfaith/krakend/config"
	"github.com/sirupsen/logrus"
)

// RelyingParty checks access to endpoint by access_token.
type RelyingParty struct {
	cfg *rpConfig
}

// New creates a new RelyingParty.
func New(e config.ExtraConfig) (*RelyingParty, error) {
	rpCfg, err := getRpConfig(e)
	if err != nil {
		logrus.WithError(err).Error("getRpConfig error")
	}
	return &RelyingParty{
		cfg: rpCfg,
	}, nil
}
