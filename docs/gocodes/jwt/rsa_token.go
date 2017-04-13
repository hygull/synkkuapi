package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)

var (
	publicKey []byte
)

func init() {
	publicKey, _ = ioutil.ReadFile("./rsa_key.rsa.pub") //("Location of your demo.rsa.pub")
}

func authRequest(rw http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) ([]byte, error) {
		return publicKey, nil
	})

	if token.Valid {
		//YAY!
		fmt.Println("Wow!")
		fmt.Fprintf(rw, "Wow!")
	} else {
		//Someone is being funny
		fmt.Println("Wrong!")
		fmt.Fprintf(rw, "Wrong!")
	}
}
func main() {
	//init()
	http.HandleFunc("/authorize/me", authRequest)
	http.ListenAndServe(":8000", nil)
}
