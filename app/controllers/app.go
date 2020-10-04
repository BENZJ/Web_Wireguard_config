package controllers

import (
	"benzj/wgconfig/app/models"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"os"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	file, err := os.Open("/Users/benz/go/src/benzj/wgconfig/wg0.conf")
	if err != nil {
		fmt.Println("读取失败")
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
	return c.Render()
}

func (c App) Login(myName string, myPassword string) revel.Result {
	xxuser := new(models.User)
	xxuser.Username = "hello"

	if myName == "benz" && myPassword == "jbcbenz" {
		c.Session["user"] = myName
		c.Session.SetDefaultExpiration()
		c.Flash.Success("Welcome, " + myName)
		return c.Redirect(WgConfig.Index)

	} else {
		c.Flash.Error("Login failed")
		return c.Redirect(App.Index)
	}

	fmt.Println(myName)
	fmt.Println(myPassword)
	fmt.Println(xxuser.Username)
	return c.Render()
}

func (c App) Loginout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}
