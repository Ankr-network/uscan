FROM golang:1.19-alpine3.16 AS builder
LABEL stage=builder
WORKDIR /uscan
COPY ./ ./

RUN apk --no-cache add build-base linux-headers git bash ca-certificates libstdc++
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 make statik
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o executor main.go


FROM alpine:3.16
WORKDIR /app
RUN apk add --no-cache ca-certificates  libstdc++ tzdata
COPY --from=builder /uscan/executor /app/
EXPOSE 4322
ENTRYPOINT ["./executor"]
