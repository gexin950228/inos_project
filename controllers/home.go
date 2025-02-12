package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	pageId, err := h.GetInt("pageId")
	if err != nil {
		pageId = 0
	}
	h.Data["pageId"] = pageId
	h.TplName = "index.html"
}

func (h *HomeController) Welcome() {
	h.TplName = "welcome.html"
}
