.PHONY: build

build:
	rm -rf build
	mkdir build

	# compile the app
	GOOS=windows GOARCH=amd64 go build -o ./build/app.exe main.go 
	GOOS=linux GOARCH=amd64 go build -o ./build/app main.go
	
