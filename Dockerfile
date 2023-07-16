FROM golang:alpine As builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH="amd64" \
    GOPROXY="https://goproxy.cn"

WORKDIR /build

COPY . .

RUN go build -ldflags  "-X cmd.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X cmd.buildVersion=1.0.0 -X cmd.gitCommitID=429cf4dba0f3c5691b9e1ebf362b7c4950d3d886" -o blog-service .

FROM golang:alpine 

WORKDIR /blog

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone 

COPY --from=builder /build/blog-service /blog/blog-service 

COPY --from=builder /build/configs/config.yaml /blog/configs/config.yaml 

ENTRYPOINT ["/blog/blog-service"]

CMD ["--config", "/blog/configs"]
