package controllers

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)
import (
	"synkkuapi/conf"
	"synkkuapi/views"
)

/***************** For now the following code is just Default, we have to change it *********************/

func Home(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	ctx.SetContentType("application/json; charset=utf8")
	fmt.Fprintf(ctx, views.UniversalMessageCreatorAsJSON([]string{"success", "message", "visit"}, 1, "Server is listening",
		"http://"+conf.HostServerIP+":"+conf.Port+"/v1/account/login"))
}
