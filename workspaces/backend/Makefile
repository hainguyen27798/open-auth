run:
	MODE=dev air

build:
	go build -o backend ./cmd/server

wire:
	cd internal/wires && wire

sqlc:
	sqlc generate