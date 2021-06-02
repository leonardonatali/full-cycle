FROM golang:1.16-alpine as build

WORKDIR /app

COPY main.go .

RUN go mod init build-test

RUN go build -ldflags="-s -w" -o ./build

FROM scratch

WORKDIR /app

COPY --from=build /app/build /app/build

ENTRYPOINT [ "/app/build" ]