package common

import (
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"

	"fmt"
	"github.com/dgrijalva/jwt-go/request"
	"encoding/json"
	"crypto/rsa"
	"context"
	"github.com/wjurmy/havaard/models"
)

// using asymmetric crypto/RSA keys
const (
	// Openssl genrsa -out app.rsa 1024
	privateKeyPath = "keys/app.rsa"

	// Openssl rsa -in app.rsa -pubout > app.rsa.pub
	publicKeyPath = "keys/app.rsa.pub"
)

type Response struct {
	Data string `json:"data"`
}

type Token struct {

Token string `json:"token"`
}

// Private key for signing and public key for verification

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {
	// TODO Add Error handling
	var err error
	signBytes, err := ioutil.ReadFile(privateKeyPath)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)


	verifyBytes, err := ioutil.ReadFile(publicKeyPath)


	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		log.Print(err)
	}

}

/*
	The private key is used for signing the JWT; the public
	key verifies the JWT in HTTP request to access the resources
	of the RESTful API.
*/


// Generation JWT token
func GenerateJWT(user models.User, role string )(string, error){
	// create a signer for RSA256


	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	t := jwt.New(jwt.SigningMethodRS256)

	// set claims for JWT token
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["email"] = user.Email
	claims["id"]  =	user.Id.Hex()

	claims["role"] = struct {
		Name string
		Role string
	}{user.Email, "admin"}


	// set the expire time for JWT token
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	t.Claims = claims
	tokenString, err := t.SignedString(signKey)
		if err != nil {
			return "", err
		}
	return tokenString, nil

}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// log.Print("Authorize", r.GetBody);
	// validate the token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err!=nil{
			log.Println(err)
		}


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("\nmerchant is set...  \n %#v\n", claims["role"])
		log.Printf("\nmerchant is set...  \n %#v\n", claims["iss"])

		ctx := context.WithValue(r.Context(),"iss", claims["iss"])
		ctx = context.WithValue(ctx,"email", claims["email"])
		ctx = context.WithValue(ctx,"id", claims["id"])
		r = r.WithContext(ctx)
	} else {
		log.Println("not ok: ", err)

	}

	log.Print("Authorize", token)

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}