FROM golang:1.8

RUN apt-get update && \
        apt-get install --yes tidy && \
        apt-get install --yes wv && \
        apt-get install --yes poppler-utils && \
        apt-get install --yes unrtf

RUN go get github.com/JalfResi/justext
RUN go get code.sajari.com/docconv/...

CMD "/go/bin/docd"
