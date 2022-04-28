BINARY_NAME=practice-sales-backend

build: setup
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME} cmd/*
# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go

run:
	bin/${BINARY_NAME}

setup:
	mkdir -p bin

clean:
	go clean
	rm -rf build