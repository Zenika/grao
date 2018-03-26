FROM golang:1.8

ADD . /go/src/github.com/Zenika/rao/rao-back
WORKDIR /go/src/github.com/Zenika/rao/rao-back

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get github.com/pilu/fresh

CMD ["/bin/bash", "run.sh"]
