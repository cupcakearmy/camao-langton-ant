cli:
	go run cmd/cli.go

web:
	pnpm --prefix ./web i
	pnpm --prefix ./web run build
	go run cmd/server.go

docker:
	docker build -t camao-langton-ant .
	docker run --rm -p 8000:8000 camao-langton-ant
