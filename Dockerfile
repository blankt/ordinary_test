# 编译：docker build --no-cache -t tpler .
# 运行：docker run --rm -it -p 8080:8080 tpler
# 导出：docker save tpler > tpler.tar
# 导入：docker load < tpler.tar

FROM golang:1.17-buster AS builder
ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOFLAGS -mod=vendor
COPY . /go/src
WORKDIR /go/src
RUN go build -v -o ordinarytest -ldflags "-s -w" .

# ===========================

FROM debian:buster
LABEL maintainer="tlf <tlf@qq.com>"
COPY --from=builder /go/src/ordinarytest /
WORKDIR /
ENV TZ Asia/Shanghai
ENV LOG_LEVEL 0
EXPOSE 8080
ENV db postgres://postgres:postgres@43.143.109.42:5432/drone?sslmode=disable
ENTRYPOINT /ordinarytest