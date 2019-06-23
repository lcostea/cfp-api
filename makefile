IMG=lcostea/go-api-demo:0.1.0

all: build docker-package docker-push


build: dep
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/cfp-api


dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v


docker-package:
	docker build . -t ${IMG}

docker-push:
	docker login -u ${PERSONAL_DOCKER_USER} -p ${PERSONAL_DOCKER_PASS}
	docker push ${IMG}
