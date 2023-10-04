FROM golang:1.21.1-alpine3.18 as build

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main *.go

FROM alpine:3.18
COPY --from=build /go/src/app .

CMD ["/main"]
