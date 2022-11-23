FROM golang
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./ ./

RUN apt-get update && apt-get install -y libzstd-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o executor main.go
RUN chmod -R 777 ./pkg/files/
ENV http_addr $http_addr
ENV http_port $http_port
ENV rpc_urls $rpc_urls
ENV db_path $db_path

EXPOSE ${http_addr}
CMD ["./executor", "$http_addr", "$http_port", "$rpc_urls", "$db_path"]