package main

import "fmt"
import "net/http"
import "time"
import "os"
import "encoding/json"

func getCookie(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "text/html;charset=utf-8")
	cookieName, err := req.Cookie("goris") //goris
	fmt.Println("Details : ", cookieName, err)
	fmt.Println("More cookies\n\n", req.Cookies(), req.Header.Get("goris"))
	if err != nil {
		fmt.Println("This cookie does not exist", err.Error())
	} else {
		fmt.Println("The cookie exists")
		fmt.Println("Cookie's Name field value id : ", cookieName.Name)

		// rw.Write([]byte())
	}
}

func setCookie(rw http.ResponseWriter, req *http.Request) {
	//fmt.Println(time.Minute * 2)
	rw.Header().Set("Content-type", "text/html")
	cookie := http.Cookie{
		Name:    "goris",
		Value:   "^goris1729$",
		Expires: time.Now().Add(2 * time.Minute), //Cookie will expire after 1 minute
	}
	cookieBytes, _ := json.MarshalIndent(cookie, "", "\t")
	fmt.Printf("Type : %T\nCookie : %v", cookie, cookie)
	http.SetCookie(rw, &cookie)
	req.AddCookie(&cookie)
	fmt.Println("Cookie successfully set for ", 2*time.Minute, " minutes")
	fmt.Println("Cookie: \n", string(cookieBytes))
	rw.Write(cookieBytes)
}

func deleteCookie(rw http.ResponseWriter, req *http.Request) {

}

func home(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Println("Rendering html....")
	rw.Write([]byte("<center><h1 style='color:green;'>Cute Golangers</h1>" +
		"<a href='/v1/set/cookie'>Set Cookie</a><br>" +
		"<a href='/v1/get/cookie'>Get Cookie</a><br>" +
		"<a href='/v1/delete/cookie'>Delete Cookie</a><br>" +
		"<a href='/v1/check/cookie/status'>Check cookie's status</a><br>" +
		"</center>"))
	fmt.Println("\nContent-type:>> ", req.Header.Get("Content-Type"))
}

func checkCookieStatus(rw http.ResponseWriter, req *http.Request) {

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[goris : error] Please provide a listen port for the server")
		return
	}

	port := os.Args[1]
	http.HandleFunc("/", home)
	http.HandleFunc("/v1/get/cookie", getCookie)
	http.HandleFunc("/v1/set/cookie", setCookie)
	http.HandleFunc("/v1/delete/cookie", deleteCookie)
	http.HandleFunc("/v1/check/cookie/status", checkCookieStatus)
	fmt.Println("HTTP Server is listening on port : ", port)
	http.ListenAndServe(":"+port, nil)
}
