.PHONY: build tidy watch cert

build:
	go build .

tidy: build
	go fmt ./...
	go mod tidy

watch:
	modd

cert:
	openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out certificate.pem

docker_build: tidy
	docker build -t bouncer:latest .
	docker tag bouncer:latest andrebq/bouncer:latest 
	docker push andrebq/bouncer:latest

