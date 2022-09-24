ARG ALPINE=3.14
ARG GOLANG=1.17.5

FROM golang:${GOLANG}-alpine${ALPINE} as build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /twitter-translate cmd/twitter-translate/main.go

FROM alpine:${ALPINE}
RUN apk add ca-certificates
COPY --from=build /twitter-translate /app/twitter-translate
WORKDIR /app

ENTRYPOINT ["./twitter-translate"]
