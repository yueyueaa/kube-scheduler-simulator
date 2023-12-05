kube-scheduler-simulator
基于Go语言进行开发的kube-scheduler调度器调度过程模拟器。

组件
Simulator的服务主要由以下3部分组成：
simulator-etcd：simulator使用etcd存储数据；
simulator-frontend：通过3000端口提供Web UI页面，是模拟器的客户端之一,可以在Web UI进行创建/编辑/删除Nodes、Pods、Persistent Volumes、Persistent Volume Claims、Storage Classes、Priority Classes这些资源来模拟集群；
simulator-server：simulator服务主体部分，暴露1212/tcp和3131/tcp两个端口

```bash
$ git clone https://github.com/kubernetes-sigs/kube-scheduler-simulator
$ cd kube-scheduler-simulator
$ vim simulator/Dockerfile
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn  #添加该行,修改go模块代理为国内源
修改apk源为国内：
$ vim web/Dockerfile
FROM node:16-alpine AS deps
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories &&  \
apk update && apk upgrade && \
apk add --no-cache make gcc g++ py-pip
```

build image
```bash
make docker_build_and_up
```

start simulator
```bash
docker-compose up
```