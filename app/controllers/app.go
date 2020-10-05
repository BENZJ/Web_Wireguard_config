package controllers

import (
	"bufio"
	"fmt"
	"github.com/revel/revel"
	"os"
)

type App struct {
	*revel.Controller
}

var path_userconf string = "/etc/wireguard/user.conf"

func (c App) Index() revel.Result {

	return c.Render()
}

func (c App) Login(myName string, myPassword string) revel.Result {
	file, err := os.Open(path_userconf)
	if err != nil {
		fmt.Println("读取失败")
		panic(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	a, _, _ := br.ReadLine()
	username := string(a)
	a, _, _ = br.ReadLine()
	password := string(a)

	if myName == username && myPassword == password {
		c.Session["user"] = myName
		c.Session.SetDefaultExpiration()
		c.Flash.Success("Welcome, " + myName)
		return c.Redirect(WgConfig.Index)

	} else {
		c.Flash.Error("Login failed")
		return c.Redirect(App.Index)
	}
}

func (c App) Loginout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}
