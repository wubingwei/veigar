veigar:
	go build -o veigar main.go

dev: veigar
	DEV_MODE=true ./veigar

syntax:
	golangci-lint run --timeout=60m ./...

clean:
	rm -f veigar
