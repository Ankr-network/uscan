FROM golang:1.19.2-alpine3.16 AS builder
LABEL stage=builder
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./ ./
ARG GoVersion="" 
ARG Branch=""
ARG Commit=""
ARG Date=""
ARG Author=""
ARG Email=""
ARG GoVersion=""
RUN GOOS=linux go build -o executor -a -installsuffix cgo -ldflags \
	"-X 'github.com/Ankr-network/uscan/cmd.Branch=${Branch}' \
	-X 'github.com/Ankr-network/uscan/cmd.Commit=${Commit}' \
	-X 'github.com/Ankr-network/uscan/cmd.Date=${Date}' \
	-X 'github.com/Ankr-network/uscan/cmd.Author=${Author}' \
	-X 'github.com/Ankr-network/uscan/cmd.Email=${Email}' \
	-X 'github.com/Ankr-network/uscan/cmd.GoVersion=${GoVersion}'" .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/executor .
COPY --from=builder /go/src/app/.uscan.yaml .
CMD ["/root/executor"]
