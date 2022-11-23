ROM golang:latest
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./ ./

RUN apt-get update && apt-get install -y libzstd-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o executor main.go
RUN chmod -R 777 ./pkg/files/

EXPOSE 4322
ENTRYPOINT ["./executor"]