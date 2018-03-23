FROM golang:1.8

RUN apt-get update
RUN apt-get -y install tidy
RUN apt-get -y install wv
RUN apt-get -y install poppler-utils
RUN apt-get -y install unrtf

RUN export GOPATH=$(go env GOPATH)
RUN export PATH=$PATH:$(go env GOPATH)/bin

RUN go get golang.org/x/oauth2
RUN go get github.com/stacktic/dropbox
RUN go get code.sajari.com/docconv/...
RUN go get github.com/JalfResi/justext
RUN go get github.com/algolia/algoliasearch-client-go/algoliasearch
RUN go get github.com/robfig/cron
RUN go get gopkg.in/square/go-jose.v2
RUN go get github.com/auth0-community/go-auth0
RUN go get github.com/rs/cors
RUN go get -u github.com/gorilla/mux

WORKDIR $GOPATH/src/github.com/Zenika/RAO
COPY ./rao-back .
RUN go build -o bin/rao

ENV GRAO_APP_PORT 8080
EXPOSE 8080

CMD ["bin/rao"]
