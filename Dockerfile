FROM golang:1.21-alpine AS golang

ADD . /go/src/foo

WORKDIR /go/src/foo

RUN go mod download

RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main

FROM gcr.io/distroless/static-debian12

COPY --from=golang /go/bin/main /dummy-api

EXPOSE 80

CMD [ "/dummy-api" ]