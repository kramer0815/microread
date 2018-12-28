FROM golang:alpine as builder
RUN apk add git
RUN go get -u -v github.com/kramer0815/microread
RUN mkdir /app 
ADD . /app
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
