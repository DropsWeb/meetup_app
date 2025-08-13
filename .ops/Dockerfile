FROM golang:1.23.0

COPY . /builds

WORKDIR /builds/cmd/app

RUN CGO_ENABLED=0 GOOS=linux go build  -o /goapi

CMD ["/goapi"]