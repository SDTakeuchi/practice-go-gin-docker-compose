FROM golang:1.17.1-buster

WORKDIR /go/src/app

RUN export GO111MODULE=auto \
&& apt-get update

COPY ./src/app/main.go .

RUN go mod init main \
&& go get . \
&& go build

EXPOSE 8080

CMD ["go", "run", "main.go", "0.0.0.0:8080"]