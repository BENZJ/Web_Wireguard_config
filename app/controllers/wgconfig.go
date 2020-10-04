package controllers

import (
	// "benzj/wgconfig/app/models"
	"github.com/revel/revel"
)

type WgConfig struct {
	*revel.Controller
}

func (c WgConfig) Index() revel.Result {
	return c.Render()
}

func (c WgConfig) checkUser() revel.Result {
	if user := c.Session["user"]; user == nil {
		c.Flash.Error("请先登陆")
		return c.Redirect(App.Index)
	}
	return nil
}
