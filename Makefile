# Variables
service_name = example-service
version = 1.0
org = assetsatlas

# Supported platforms
platforms = linux/amd64,linux/arm64

# Targets
.PHONY: go_test go_build docker_build tag push deploy jp_push clean

# Run Go tests
go_test:
	go test ./src/ -v

# Build the binary locally (for testing purposes)
go_build:
	GOOS=linux GOARCH=amd64 go build -o $(service_name) ./src

# Build and push the Docker image for multiple architectures
docker_build:
	docker buildx build \
		--platform $(platforms) \
		--provenance=true --sbom=true \
		-t $(org)/$(service_name):$(version) \
		--push .

# Tag the Docker image with "latest"
tag:
	docker tag $(org)/$(service_name):$(version) $(org)/$(service_name):$(version)

# Push the Docker image to the registry
push:
	docker push $(org)/$(service_name):$(version)

# Push to a Jumppad environment (example deployment step)
jp_push:
	jumppad push $(org)/$(service_name):$(version) resource.nomad_cluster.dev

# Build and deploy the service
deploy: docker_build

# Clean up local build artifacts
clean:
	rm -f $(service_name)