package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var (
	signingKey = "HELLO GO-MIAMI"
	privateKey []byte
	publicKey  []byte
)

func init() {
	var err error
	privateKey, err = ioutil.ReadFile("rsa_key.rsa")
	if err != nil {
		fmt.Println("Can't read private key...")
		return
	} else {
		fmt.Println("Succesfully read the private key from pwd...")
	}

	publicKey, err = ioutil.ReadFile("rsa_key.rsa.pub")
	if err != nil {
		fmt.Println("Can't read public key...")
		return
	} else {
		fmt.Println("Succefully read public key from pwd...")
	}
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "GO HOME SON")
	}
}

func APIHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "WIN")
}

func main() {
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token := jwt.New(jwt.GetSigningMethod("RS256"))
		tokenString, _ := token.SignedString(privateKey)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
	})

	router.Handle("/api", negroni.New(negroni.HandlerFunc(AuthMiddleware), negroni.HandlerFunc(APIHandler)))

	n.UseHandler(router)

	port := "9000"
	fmt.Println("Server is listening on the port : ", port)
	http.ListenAndServe(":"+port, n)
}
