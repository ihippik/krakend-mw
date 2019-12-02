package relyingparty

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSelfHostMiddleware(t *testing.T) {

	sf := &RelyingParty{
		//logger: liblog.New("testApp", liblog.Config{Level: "debug"}),
		cfg: &rpConfig{
			AuthServerURI:         "http://test.com",
			AuthorizationEndpoint: "/auth",
			TokenEndpoint:         "/token",
		},
	}

	tests := []struct {
		name     string
		path     string
		wantCode int
		wantBody string
	}{
		{
			name:     "well_known",
			path:     "/api/.well-known",
			wantCode: http.StatusOK,
			wantBody: `{"authServerURI":"http://test.com","authorizationEndpoint":"/auth","tokenEndpoint":"/token"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router := gin.Default()
			router.GET(tt.path, newEndpointRelyingPartyMw(sf)(nil))
			w := performRequest(router, "GET", tt.path)

			assert.EqualValues(t, tt.wantCode, w.Code)
			assert.EqualValues(t, tt.wantBody, w.Body.String())
		})
	}
}
