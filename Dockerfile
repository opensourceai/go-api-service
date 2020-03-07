FROM golang:latest

ENV GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE=on
WORKDIR $GOPATH/src/github.com/opensourceai/go-api-service
COPY . $GOPATH/src/github.com/opensourceai/go-api-service
RUN go get -u github.com/swaggo/swag/cmd/swag \
    && swag init
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-api-service"]
