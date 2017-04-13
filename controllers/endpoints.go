/*
	Created on : 18/10/2016.
	Aim        : To display JSON output on the browser...Either it can be an error message or success message
				 whatever.
*/

package controllers

import (
	"encoding/json"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)
import (
	"synkkuapi/conf"
	"synkkuapi/views"
)

func Endpoints(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	//rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.SetContentType("application/json; charset=utf8")

	type Link struct {
		Link string `json:"url"`
		Id   int    `json:"id"`
	}

	type Links struct {
		UrlDetails Link `json:"url details"`
	}

	var jsonStrData []byte
	linksCollection := map[string]Links{}

	linksCollection["HomePageLink"] = Links{Link{"http://" + conf.HostServerIP + ":" + conf.Port +
		"/", 0}}
	linksCollection["EndpointsLink"] = Links{Link{"http://" + conf.HostServerIP + ":" + conf.Port +
		"/v1/endpoints", 1}}
	linksCollection["AccountCreationLink"] = Links{Link{"http://" + conf.HostServerIP + ":" + conf.Port +
		"/v1/account/new", 2}}
	linksCollection["AccountLoginLink"] = Links{Link{"http://" + conf.HostServerIP + ":" + conf.Port +
		"/v1/account/login", 3}}

	jsonStrData, _ = json.Marshal(linksCollection)
	views.ShowDataAsJSON(ctx, string(jsonStrData))
}
