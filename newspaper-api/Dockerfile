FROM golang:1.23-alpine AS base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal

FROM base AS builder

WORKDIR cmd/newspaper-api
RUN go build -o /build/newspaper-api .

FROM base AS debug
WORKDIR cmd/newspaper-api
RUN go build -gcflags="all=-N -l" -o /build/newspaper-api .

ENV DEBUG_PORT=2345
EXPOSE $DEBUG_PORT
ENTRYPOINT /go/bin/dlv --listen=:$DEBUG_PORT --headless=true --api-version=2 --accept-multiclient exec /build/newspaper-api

FROM alpine AS run
COPY --from=builder /build/newspaper-api /
USER 1000
CMD ["/newspaper-api"]

