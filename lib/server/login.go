package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/horitaka/oidc-cli/lib/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	authClient := utils.LoadEnv()

	u, err := url.Parse(constants.AUTH_URL)
	if err != nil {
		wrappedError := errors.Wrap(err, "Failed to parse URL.")
		fmt.Printf("%+v\n", wrappedError)
	}

	q := u.Query()
	clienId := authClient.ClientId
	q.Set("client_id", clienId)
	q.Set("scope", "openid email profile")
	q.Set("response_type", "code")
	q.Set("redirect_uri", constants.CALLBACK_URL)
	q.Set("state", "state") // TODO
	q.Set("nonce", "nonce") // TODO

	u.RawQuery = q.Encode()

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", u.String())
	w.WriteHeader(http.StatusFound)
}
