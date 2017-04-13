/*
	Created on : 29/10/2016.
	Links	   : https://github.com/dgrijalva/jwt-go/blob/master/example_test.go
				 https://godoc.org/github.com/dgrijalva/jwt-go#MapClaims
*/
package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	_ "time"
)

func createJWT() {

	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			ExpiresAt: 120, //15000,

			Issuer: "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("Token : %v \nError : %v\n\n", ss, err)
}

func checkJWT() {
	/*****************/
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"

	// type MyCustomClaims struct {
	// 	Foo string `json:"foo"`
	// 	jwt.StandardClaims
	// }

	// sample token is expired.  override time so it parses as valid
	// at(time.Unix(0, 0), func() {
	// 	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("AllYourBase"), nil
	// 	})

	// 	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
	// 		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// })

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet

			fmt.Println("Great...The token is either active or expired...choose a different way for these...to know")

			fmt.Println("Timing is everything")

		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}

func main() {
	fmt.Printf("%s\n", "*********************")
	createJWT()
	checkJWT()

	fmt.Println("\n*********************")

	fmt.Println("*********************")

}
