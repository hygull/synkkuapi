//	https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac (Read it)
//	http://dghubble.com/blog/posts/json-web-tokens-and-go/
package main

import "fmt"
import jwt "github.com/dgrijalva/jwt-go"
import "time"

func main() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2016, 11, 4, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString(hmacSampleSecret)

	tokenString, err := token.SignedString([]byte("rishi"))
	fmt.Printf("%T %v %v", tokenString, tokenString, err)

	/* Parse */
	// sample token string taken from the New example

	//tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
	tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDY0NjU2MDB9.-DZIwM-3wYxBoTiRlQvVLcfd9noc7GpW1FPhcZgji1M"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return hmacSampleSecret, nil
		return []byte("rishi"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Details : ", claims["foo"], claims["nbf"])
	} else {
		fmt.Println("Error : ", err)
	}
}
