test: *.go
	go test ./...

out/example: cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example