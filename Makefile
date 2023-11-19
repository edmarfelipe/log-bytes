BINARY_NAME=log-bytes

build:
	go build -o bin/${BINARY_NAME} -ldflags="-s -w" .
	sudo mv bin/${BINARY_NAME} /usr/bin/${BINARY_NAME}
	sudo chmod +x /usr/bin/${BINARY_NAME}