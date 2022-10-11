FROM golang:1.19.2 AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux GOARCH=amd64 go build -v -o curlrequester 

FROM alpine:3.16.1
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
EXPOSE 8080
ENV VERSION 1.0.0
CMD ["curlrequester"]
COPY --from=build /src/curlrequester /usr/local/bin/curlrequester
RUN chmod +x /usr/local/bin/curlrequester
