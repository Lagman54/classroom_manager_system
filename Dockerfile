FROM golang:1.22.2 as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY internal/classroom-app/migration ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app ./cmd/classroom-app

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN ls -la

COPY --from=builder /app/demo-app .
COPY --from=builder /app/internal/classroom-app/migration ./migration

CMD ["./demo-app"]