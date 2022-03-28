## Build

FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd

RUN go build -o /meli-proxy

## Deploy

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /meli-proxy /meli-proxy

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/meli-proxy"]