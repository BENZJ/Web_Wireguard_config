#/bin/sh
# 安装go
wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
mkdir /root/go

# 配置环境变量
export GOPATH=/root/go
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin

# 配置系统环境变量重启之后有效
echo "export GOPATH=/root/go" >> /root/.bashrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> /root/.bashrc
echo "export PATH=\$PATH:\$GOPATH/bin" >> /root/.bashrc
source /root/.bashrc



# 安装revel
go get github.com/revel/revel
go get github.com/revel/cmd/revel


git clone https://github.com/BENZJ/Web_Wireguard_config

revel package Web_Wireguard_config

mkdir wgconfig

tar -zxvf ./Web_Wireguard_config/Web_Wireguard_config.tar.gz  -C ./wgconfig


