build:
	go build -o ./build/erestourant cmd/api/main.go

run: build
	./build/erestourant

clear:
	rm -rf build