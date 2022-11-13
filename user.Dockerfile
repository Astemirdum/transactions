FROM golang:1.19-alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../../.. .
RUN go build -ldflags="-s -w" -o /app/user ./cmd/user/main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/user /app/user
COPY --from=builder /build/.env /app

RUN mkdir configs
COPY --from=builder /build/configs/user.yml /app/configs
CMD ["./user", "-c", "configs/user.yml"]