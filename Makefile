cli:
	go run cmd/cli.go

web:
	pnpm --prefix ./web i
	pnpm --prefix ./web run build
	go run cmd/server.go
