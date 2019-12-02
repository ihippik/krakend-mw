package relyingparty

import (
	"errors"
	"fmt"
	"strings"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	namespace           = "github.com/ihippik/krakend-mw/relyingparty"
	HeaderAuthorization = "Authorization"
	HeaderUserID        = "User-Id"
)

type EndpointMw func(gin.HandlerFunc) gin.HandlerFunc

var unauthorizedErr = errors.New("unauthorized")

// NewHandlerFactory builds a oauth2 wrapper over the received handler factory.
// Run for each endpoints.
func NewHandlerFactory(next krakendgin.HandlerFactory, rp *RelyingParty) krakendgin.HandlerFactory {
	return func(remote *config.EndpointConfig, p proxy.Proxy) gin.HandlerFunc {
		handlerFunc := next(remote, p)
		_, ok := remote.ExtraConfig[namespace]
		if !ok {
			return handlerFunc
		}

		return newEndpointRelyingPartyMw(rp)(handlerFunc)
	}
}

// newEndpointRelyingPartyMw is the handler middlware that represents endpoints of
// gateway itself.
func newEndpointRelyingPartyMw(rp *RelyingParty) EndpointMw {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			userToken := c.GetHeader(HeaderAuthorization)

			if len(userToken) == 0 {
				logrus.Warnln("empty user token")
				_ = c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
				return
			}

			items := strings.Split(userToken, " ")
			if len(items) != 2 {
				logrus.Warnln("invalid token")
				_ = c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
				return
			}

			token, err := jwt.Parse(items[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(rp.cfg.TokenSecret), nil
			})
			if err != nil {
				logrus.WithError(err).Warnln("parse token err")
				_ = c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID, ok := claims["user_id"].(string)
				if !ok {
					logrus.Warnln("invalid user id")
					_ = c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
					return
				}
				c.Request.Header.Set(HeaderUserID, userID)
			} else {
				logrus.WithError(err).Warnln("claims err")
				_ = c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
				return
			}

			next(c)
		}
	}
}
