package main

//https://www.reddit.com/r/golang/comments/2lu55c/restful_session_token/

import "log"
import "time"
import "errors"
import "crypto/hmac"
import "crypto/sha256"
import "encoding/json"
import "encoding/base64"
import "fmt"

type AuthToken struct {
	Info string // Contains TokenInfo in base 64 encoded json
	HMAC string // Base 64 encoded hmac
}

type TokenInfo struct {
	UserID         string
	ExpirationDate time.Time
}

func NewAuthToken(uid string, expirationDate time.Time, secret string) *AuthToken {
	at := &AuthToken{
		Info: NewTokenInfo(uid, expirationDate).ToBase64(),
	}
	fmt.Println(*at)
	at.HMAC = ComputeHmac256(at.Info, secret)
	return at
}

func NewTokenInfo(uid string, expirationDate time.Time) *TokenInfo {
	return &TokenInfo{
		UserID:         uid,
		ExpirationDate: expirationDate,
	}
}

func (at *AuthToken) verify(secret string) bool {
	if ComputeHmac256(at.Info, secret) == at.HMAC {
		return true
	} else {
		return false
	}
}

func (at *AuthToken) GetTokenInfo(secret string) (*TokenInfo, error) {
	/* If the token is not valid, stop now. */
	if !at.verify(secret) {
		return nil, errors.New("The token is not valid.")
	}

	/* Convert from base64. */
	jsonString, err := base64.StdEncoding.DecodeString(at.Info)
	if err != nil {
		log.Fatal("Failed to decode base64 string: ", err)
	}
	/* Unmarshal json object. */
	var ti TokenInfo
	err = json.Unmarshal(jsonString, &ti)
	if err != nil {
		log.Fatal("Failed to decode TokenInfo: ", err)
	}

	/* Check if the token is expired. */
	if time.Now().Unix() > ti.ExpirationDate.Unix() {
		return nil, errors.New("The token is expired.")
	} else {
		fmt.Println("Not expired...")
		return &ti, nil
	}
}

func (ti *TokenInfo) ToBase64() string {
	bytes, err := json.Marshal(ti)
	if err != nil {
		log.Panic("Failed to marshal TokenInfo.")
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func ComputeHmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func main() {
	fmt.Println("Hello...")
	fmt.Println("Creating auth token...")
	//https://www.socketloop.com/references/golang-time-time-adddate-function-example
	authToken := NewAuthToken("23", time.Now().AddDate(2015, 10, 1), "^Go~lang@1729@$")
	fmt.Printf("%T, %v", authToken, authToken)
	fmt.Println(authToken.GetTokenInfo("^Go~lang@1729@$"))
}
