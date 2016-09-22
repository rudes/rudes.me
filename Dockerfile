FROM golang

ADD . /go/src/github.com/rudes/rudes.me
RUN go install github.com/rudes/rudes.me
ENTRYPOINT /go/bin/rudes.me

EXPOSE 8080
