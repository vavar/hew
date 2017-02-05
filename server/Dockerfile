FROM golang:wheezy

RUN go get github.com/codegangsta/gin \
    && mkdir -p /go/src/app

WORKDIR /go/src/app

ADD docker-go.sh /root/
RUN chmod +x /root/docker-go.sh

ENTRYPOINT ["/root/docker-go.sh"]