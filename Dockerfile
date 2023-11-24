FROM golang:1.21-alpine AS builder

# Install prerequisites
RUN apk add --no-cache ca-certificates git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY config/*.json ./config/

COPY . .

RUN go build --installsuffix -buildvcs=false -o main ./cmd

EXPOSE 80

CMD ["./main"]
