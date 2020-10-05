# Web 控制修改 Wirguard配置文件

使用Revel(go web 框架)，来搭建了一个管理wireguard conf的小程序

# 文件结构
```bash 
.
├── README.md
├── app
│   ├── controllers
│   │   ├── app.go
│   │   ├── init.go
│   │   └── wgconfig.go
│   ├── init.go
│   ├── models
│   ├── routes
│   ├── tmp
│   └── views
├── conf
│   ├── app.conf
│   └── routes
├── go.mod
├── go.sum
├── messages
│   └── sample.en
├── none
├── public
│   ├── css
│   │   └── bootstrap-3.3.6.min.css
│   ├── fonts
│   ├── img
│   │   └── favicon.png
│   └── js
├── target
├── tests
```
# 运行说明


