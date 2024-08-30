run:
	MODE=dev go run ./cmd/server/

build:
	go build -o backend ./cmd/server

wire:
	cd internal/wires && wire

sqlc:
	sqlc generate