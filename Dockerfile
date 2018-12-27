FROM golang:alpine as builder
RUN apk add git
RUN go get -u -v github.com/kramer0815/rss-reader
RUN mkdir /build 
ADD . /build
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
ENTRYPOINT [ "/app/main" ]
