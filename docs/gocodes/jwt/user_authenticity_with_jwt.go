package main

import jwt "github.com/dgrijalva/jwt-go"
import "fmt"
import "net/http"
import "time"
import "strings"

func getToken(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	password := req.Form["password"][0]
	fmt.Println("Your entered password : ", password)
	fmt.Printf("Length : %d\n\n", len(password))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"name": "Rishikesh",
		"nbf":  time.Date(2015, 11, 4, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":  time.Now().Add(1 * time.Minute).Unix(),
	})

	//Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString(hmacSampleSecret)

	tokenString, err := token.SignedString([]byte(password))
	fmt.Println(tokenString, err)
	rw.Write([]byte(tokenString))
}

func authenticate(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	tokenString := strings.TrimSpace(req.Form["token"][0])
	password := strings.TrimSpace(req.Form["password"][0])

	//fmt.Println("Header : ", req.Header.Get("Authorization"))
	fmt.Println(tokenString, " | ", password)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return hmacSampleSecret, nil
		return []byte(password), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Details : ", claims["foo"], claims["nbf"], claims["name"])
	} else {
		fmt.Println("Input combination is not correct ... ", err)

		fmt.Println("If you forgot your token then please re-request it again...")
		rw.Write([]byte("Invalid signature...Or the token has been expired...1 minute"))
	}
}

func main() {
	fmt.Println("This is authentication using jwt...")

	http.HandleFunc("/getToken", getToken)
	http.HandleFunc("/authenticate", authenticate)

	fmt.Printf("%s %d", "Server is listening on ", 8000)
	http.ListenAndServe(":8000", nil)
}
