# syntax=docker/dockerfile:1

#使用golang:1.16-alpine作为基本镜像
FROM golang:1.16-alpine

#为了在运行其余命令时更轻松，让我们在正在构建的镜像内创建一个目录。
#这也指示 Docker 将此目录用作所有后续命令的默认目标。
#这样，我们不必键入完整的文件路径，而是可以使用基于此目录的相对路径
WORKDIR /app

#把所有的文件复制到镜像目录中
COPY . /app

#为映像设置代理，方便依赖包的下载
ENV GOPROXY=https://goproxy.cn,direct

#下载依赖包
RUN go mod download

#编译我们的应用程序
RUN go build -o /CourseSelectionSystem

#这里要与你的项目名字一致
CMD ["/CourseSelectionSystem"]