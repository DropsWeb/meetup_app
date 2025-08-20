FROM golang:1.24.5-alpine

COPY ./ /builds

WORKDIR /builds

RUN CGO_ENABLED=0 GOOS=linux go build  -o /goapi

CMD ["/goapi"]