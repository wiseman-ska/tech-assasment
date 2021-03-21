package commons

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	privKeyPath = "user-manager-api/auth-keys/app.rsa"
	pubKeyPath  = "user-manager-api/auth-keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("[initKeys]: %s\n", err)
		panic(err)
	}
}

func GenerateToken(name, role string) (string, error) {
	Claims := jwt.MapClaims{}
	Claims["iss"] = "admin"
	Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), Claims)
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(signKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token",
					500,
				)
				return
			}
		default:
			DisplayAppError(w,
				err,
				"Error while parsing the Access Token",
				500,
			)
			return
		}
	}
	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
