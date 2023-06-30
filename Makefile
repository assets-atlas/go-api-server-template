service_name = "example-service"
version = "0.1.0"
org = "assets-atlas"

go_test:
	go test ./src -v
go_build:
	go build ./src -o $(service_name)
docker_build:
	docker build -t $(org)/$(service_name):$(version) .
tag:
	docker tag  $(org)/$(service_name):$(version) $(org)/$(service_name):$(version)
push:
	docker push  $(org)/$(service_name):$(version)
deploy: build tag push