package relyingparty

import (
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// EndpointMw is a function that decorates the received handlerFunc with
// gateway's own logic.
type EndpointMw func(gin.HandlerFunc) gin.HandlerFunc

const (
	namespace = "gitlab.com/r-stat/krakend-mw/relyingparty"
)

// NewHandlerFactory builds a oauth2 wrapper over the received handler factory.
func NewHandlerFactory(next krakendgin.HandlerFactory, sf *RelyingParty) krakendgin.HandlerFactory {
	return func(remote *config.EndpointConfig, p proxy.Proxy) gin.HandlerFunc {
		handlerFunc := next(remote, p)

		_, ok := remote.ExtraConfig[namespace]
		if !ok {
			return handlerFunc
		}

		return newEndpointSelfHostMw(sf)(handlerFunc)
	}
}

// newEndpointSelfHostMw is the handler middlware that represents endpoints of
// gateway itself.
func newEndpointSelfHostMw(sf *RelyingParty) EndpointMw {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			// Return config in response.
			c.AbortWithStatusJSON(http.StatusOK, sf.cfg)
			logrus.Info("SF: config provided.")
		}
	}
}
