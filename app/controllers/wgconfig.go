package controllers

import (
	// "benzj/wgconfig/app/models"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"os"
	"os/exec"
)

type WgConfig struct {
	*revel.Controller
}

var path_wgconf string = "/etc/wireguard/wg0.conf"

// var path_wgconf string = "/etc/wireguard/wg0.conf"

func (c WgConfig) Index() revel.Result {
	data, err := ioutil.ReadFile(path_wgconf)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	wg0conf := string(data)
	return c.Render(wg0conf)
}

func (c WgConfig) SaveAndReastart(wgconftex string) revel.Result {
	cmd := exec.Command("wg-quick", "down", "wg0")
	cmd.Run()
	f, err := os.Create(path_wgconf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = f.WriteString(wgconftex)
	if err != nil {
		fmt.Println(err)
		f.Close()
		panic(err)
	}
	fmt.Println("写入成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	cmd = exec.Command("wg-quick", "up", "wg0")
	cmd.Run()
	c.Flash.Success("重启Wireguard完成")
	return c.Redirect(WgConfig.Index)
}

func (c WgConfig) checkUser() revel.Result {
	if user := c.Session["user"]; user == nil {
		c.Flash.Error("请先登陆")
		return c.Redirect(App.Index)
	}
	return nil
}
