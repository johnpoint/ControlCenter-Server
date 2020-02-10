#!/bin/bash

Green_font_prefix="\033[32m" && Red_font_prefix="\033[31m" && Green_background_prefix="\033[42;37m" && Red_background_prefix="\033[41;37m" && Font_color_suffix="\033[0m"
Info="${Green_font_prefix}[信息]${Font_color_suffix}"
Error="${Red_font_prefix}[错误]${Font_color_suffix}"
Tip="${Green_font_prefix}[注意]${Font_color_suffix}"

[[ $EUID != 0 ]] && echo -e "${Error} 当前账号非ROOT(或没有ROOT权限)，无法继续操作，请使用${Green_background_prefix} sudo su ${Font_color_suffix}来获取临时ROOT权限（执行后会提示输入当前账号的密码）。" && exit 1

if [ -f /etc/redhat-release ]; then
    release="centos"
    PM='yum'
    echo -e "${Error}不兼容~"
    exit 0
elif cat /etc/issue | grep -Eqi "debian"; then
    release="debian"
    PM='apt'
elif cat /etc/issue | grep -Eqi "ubuntu"; then
    release="ubuntu"
    PM='apt'
elif cat /proc/version | grep -Eqi "debian"; then
    release="debian"
    PM='apt'
elif cat /proc/version | grep -Eqi "ubuntu"; then
    release="ubuntu"
    PM='apt'
else
    echo -e "${Error}无法识别~"
    exit 0
fi

$PM update -y
$PM upgrade
$PM install curl unzip python -y
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh --mirror Aliyun
systemctl enable docker
systemctl start docker
curl -L https://github.com/docker/compose/releases/download/1.25.4/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
mkdir /web/vhost /web/conf /web/ssl/auto /web/mysql/data -p
wget https://raw.githubusercontent.com/johnpoint/DNMP-lvcshu/master/docker-compose.yml
docker-compose up -d
