package main

import jwt "github.com/dgrijalva/jwt-go"
import "fmt"
import "net/http"
import "time"
import "strings"
import "gopkg.in/redis.v3"

func getToken(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.Form["email"][0]
	appKey := req.Form["appKey"][0]
	fmt.Println("Your entered email: ", email, " and appKey:", appKey)
	fmt.Printf("Length : %d\n\n", len(email))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"name": "Rishikesh",
		"nbf":  time.Date(2015, 11, 4, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":  time.Now().Add(1 * time.Minute).Unix(),
	})

	//Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString(hmacSampleSecret)
	tokenString, err := token.SignedString([]byte(email))
	fmt.Println(tokenString, err)

	/*........... REDIS  ........*/
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})
	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Could not connect to REDIS Server")
	}
	fmt.Println(ping, err)

	err = client.Set(tokenString, email, 0).Err()
	if err != nil {
		fmt.Println("Unable to set the key ", tokenString)
	}
	fmt.Println("{", tokenString, ":", email, "}")
	rw.Write([]byte(tokenString))
}

func authenticate(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	tokenString := strings.TrimSpace(req.Form["token"][0])

	//fmt.Println("Header : ", req.Header.Get("Authorization"))
	fmt.Println("JWT Token :", tokenString)

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})
	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Could not connect to REDIS Server")
	}
	fmt.Println(ping, err)

	email, err := client.Get(tokenString).Result()
	if err == redis.Nil {
		fmt.Println("Key : ", tokenString, " does not exist")
		rw.Write([]byte("Invalid token"))
	} else if err != nil {
		fmt.Println("Error while getting the key", err)
		rw.Write([]byte("Error while getting the key"))
	} else {
		fmt.Println("Key : ", tokenString, " found  with value : ", email)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return hmacSampleSecret, nil
		return []byte(email), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Details : ", claims["foo"], claims["nbf"], claims["name"])
	} else {
		fmt.Println("Input combination is not correct ... ", err)
		fmt.Println("If you forgot your token then please re-request it again...")
		rw.Write([]byte("Invalid signature...Or the token has been expired (1 minute is expiry time)"))
	}
}

func main() {
	fmt.Println("This is authentication using jwt...")

	http.HandleFunc("/getToken", getToken)
	http.HandleFunc("/authenticate", authenticate)

	fmt.Printf("%s %d", "Server is listening on ", 8000)
	http.ListenAndServe(":8000", nil)
}
