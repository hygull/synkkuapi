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

func GetComments(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
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

	type CommentData struct {
		Cid    int    `json:"cid"`
		Cmtmsg string `json:"cmtmsg"`
		Cmtimg string `json:"cmtimg"`
		Sts    int    `json:"sts"`
		C_at   string `json:"c_at"`
		U_at   string `json:"u_at"`
		Uid    int    `json:"uid"`
		Uname  string `json:"uname"`
		Ppic   string `json:"p_pic"`
	}

	CommentsArr := []CommentData{}

	var cmt_id int
	var cmt_msg string
	var cmt_img string
	var status int
	var created_on string
	var updated_on string
	var uid int
	var uname string
	var ppic string

	comOp := ""
	if ty == "1" {
		comOp = ">"
	} else {
		comOp = "<"
	}
	queryString := "select comments.comment_id,comments.comment_message,comments.image_url,comments.status, " +
		"comments.created_on,comments.updated_on,auth_users.id,auth_users.user_name,auth_users.profile_pic from comments " +
		"inner join auth_users on auth_users.id=" + uId + " AND comments.user_id=" + uId + " AND auth_users.login_type=1 and comments.updated_on " + comOp + " '" + uat + "';"
	fmt.Println(queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error in query execution")
		views.ShowSuccessOrErrorAsJSON(ctx, "QueryExecutionError", "Error in execution of query", 127, 27)
		return
	}
	found := false
	for rows.Next() {
		rows.Scan(&cmt_id, &cmt_msg, &cmt_img, &status, &created_on, &updated_on, &uid, &uname, &ppic)
		CommentsArr = append(CommentsArr, CommentData{cmt_id, cmt_msg, cmt_img, status, created_on, updated_on, uid, uname, ppic})
		found = true
	}

	if found == false {
		views.ShowSuccessOrErrorAsJSON(ctx, "CommentsNotFound", "As of now, no one has commented here", 130, 30)
		fmt.Println("No comments found in this community")
		return
	}
	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "c_arr"}, 1, "Comments Data are", CommentsArr))
	fmt.Println("Comments Data have been successfully displayed...")
}
