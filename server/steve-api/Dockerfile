# https://oneuptime.com/blog/post/2026-01-07-go-docker-multi-stage/view
# https://docs.docker.com/guides/go-prometheus-monitoring/
FROM golang:alpine3.24 AS builder

# static building of the binary to avoid dynamic dependencies
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app 

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go build -o steve-api .

FROM alpine:3.20 AS runner

WORKDIR /app

COPY --from=builder /app/steve-api .

EXPOSE 8089

CMD ["./steve-api"]