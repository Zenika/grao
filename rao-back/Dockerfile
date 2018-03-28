FROM golang:1.8 as builder

ADD . /go/src/github.com/Zenika/rao/rao-back
WORKDIR /go/src/github.com/Zenika/rao/rao-back
RUN CGO_ENABLED=0 GOOS=linux go build -o ./rao

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/github.com/Zenika/rao/rao-back
COPY --from=builder /go/src/github.com/Zenika/rao/rao-back/rao .

ENV GRAO_APP_PORT=8080

EXPOSE 8080
CMD ["/go/src/github.com/Zenika/rao/rao-back/rao"]
