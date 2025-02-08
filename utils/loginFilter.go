package utils

import (
	"github.com/astaxie/beego/context"
	"net/http"
)

func LoginFilter(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(http.StatusFound, "/")
	}
}
