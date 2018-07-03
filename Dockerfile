FROM golang:latest

RUN go get github.com/gin-gonic/gin

RUN mkdir /api
WORKDIR /api
ADD . /api

RUN go build -o main .
CMD ["/api/main"]