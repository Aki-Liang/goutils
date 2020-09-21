package goutils

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

var (
	JwkUrl = "https://appleid.apple.com/auth/keys"
)

func getKey(token *jwt.Token) (interface{}, error) {
	// we want to verify a JWT
	set, err := jwk.FetchHTTP(JwkUrl)
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if key := set.LookupKeyID(keyID); len(key) == 1 {
		// SHA-256	== RS256
		var p interface{}
		err := key[0].Raw(&p)
		return p, err
	}

	return nil, fmt.Errorf("unable to find key %q", keyID)
}
