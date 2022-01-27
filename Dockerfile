FROM archlinux:latest

MAINTAINER johnpoint
ENV TZ=Asia/Shanghai

RUN echo 'Server = https://mirrors.aliyun.com/archlinux/$repo/os/$arch' > /etc/pacman.d/mirrorlist
RUN pacman -Sy --noconfirm
RUN pacman -Sy gnu-netcat --noconfirm
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR /usr/src
COPY ./ControlCenter /usr/src
COPY ./config_dev.yaml /usr/src
COPY ./start.sh /usr/src