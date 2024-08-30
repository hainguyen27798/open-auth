FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/server

FROM scratch

COPY --from=builder /build/backend /
COPY --from=builder /build/configs /configs
COPY --from=builder /build/database/migrations /database/migrations

ENTRYPOINT ["./backend"]
