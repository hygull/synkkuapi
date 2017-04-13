package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"strings"
	//"time"
)
import (
	"synkkuapi/conf"
	"synkkuapi/controllers/middlewares"
	"synkkuapi/views"
)

func GetLikes(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	//AuthenticationKey := conf.AuthenticationKey
	ctx.SetContentType("application/json; charset=utf8")

	if !ctx.IsPost() { //Checking whether the data is being sent using POST method or not(If already implemented in Android then not required)
		views.ShowSuccessOrErrorAsJSON(ctx, "PostMethodNotFoundError", "The data is not being sent using POST method", 107, 7)
		fmt.Println("The data is not being sent using POST method")
		return
	}
	l := len(ctx.PostBody())
	if l == 0 { //If there's no post data then inform the user...len([]byte )==0
		views.ShowSuccessOrErrorAsJSON(ctx, "NoPostDataError", "You haven't sent the POST data", 104, 4)
		fmt.Println("You haven't send the POST data(as it is empty)")
		return
	}
	fmt.Println("Length of PostBody() : ", l)

	uId := strings.TrimSpace(string(ctx.FormValue("uid")))
	uat := strings.TrimSpace(string(ctx.FormValue("uat")))
	ty := strings.TrimSpace(string(ctx.FormValue("ty")))
	pId := strings.TrimSpace(string(ctx.FormValue("pid")))

	/****************************************************************************************/
	fmt.Println("Extracted URL Parameters : ", uId, uat, ty, pId)
	if !middlewares.AreAllFieldsAreInValidForm(uId, uat, ty, pId) {
		views.ShowSuccessOrErrorAsJSON(ctx, "EmptyFieldOrKeyNameError", "You have to provide all fields (uid,uat,ty,pid) except uid or check if  you have provided a wrong key uid", 103, 3)
		return
	}

	fmt.Println("Connecting to DB...")
	db, err := sql.Open("postgres", "postgres://"+conf.DBUserName+":"+conf.DBPassword+"@"+conf.DBHost+":"+conf.DBPort+"/"+conf.DBName+"?sslmode=disable")

	if err != nil {
		fmt.Println("Error in connection...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "DBConnectionError", "Error in connection with database", 104, 4)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in connection test...")
		fmt.Println(err.Error())
		views.ShowSuccessOrErrorAsJSON(ctx, "DBConnectionTestError", "Error in connection with database", 112, 12)
		return
	}

	fmt.Println("Connected...")

	type LikesData struct {
		Likeid     int    `json:"likid"`
		Isliked    int    `json:"islkd"`
		Created_at string `json:"c_at"`
		Updated_at string `json:"u_at"`
		Uid        int    `json:"uid"`
		Uname      string `json:"uname"`
		P_pic      string `json:"p_pic"`
	}

	var Likeid int
	var Isliked int
	var Updated_at string
	var Created_at string
	var Uid int
	var Uname string
	var P_pic string

	LikesArr := []LikesData{}

	compOp := ""

	if ty == "1" {
		compOp = ">"
	} else {
		compOp = "<"
	}
	/*@******** here we are using fname as the username******/
	// query := "select likes.like_id, likes.is_liked,likes.updated_on,users.user_id,users.first_name,"+
	// " users.profilepic_url,posts.created_on from users inner join posts  on users.user_id=posts.user_id "+
	// " inner join likes  on posts.post_id=likes.post_id and likes.updated_on"+compOp+"'"+uAt+"';"

	query := "select likes.id, likes.is_liked,likes.updated_on,auth_users.id,auth_users.first_name," +
		" auth_users.profile_pic,posts.created_on from auth_users inner join posts  on auth_users.id=posts.id " +
		" and auth_users.id=" + uId + " inner join likes  on posts.id=likes.post_id and likes.updated_on" + compOp + "'" + uat + "' " +
		" and posts.id=" + pId + ";"
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error in query execution")
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in execution of query", 127, 27)
		return
	}

	/*		We are not checking the existence of data,
	because it is resposibilty of android developer to check whether likes array is empty or not
	*/

	for rows.Next() {
		rows.Scan(&Likeid, &Isliked, &Updated_at, &Uid, &Uname, &P_pic, &Created_at)
		LikesArr = append(LikesArr, LikesData{Likeid, Isliked, Created_at, Updated_at, Uid, Uname, P_pic})
	}

	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "l_arr"}, 1, "Likes Data", LikesArr))
	fmt.Println("Likes Data have been successfully displayed...")
}
