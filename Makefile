.PHONY: build

K8S_ROOT_DIR:=$(PWD)/k8s

build:
	@cd src \
	&& env GO111MODULE=auto GOOS=linux go build -ldflags="-s -w" -o $(PWD)/main main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

watch:
	@cd src \
	&& reflex -r '\.go' -s -- sh -c "go run ./main.go"

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

docker-build:
	docker build . -t test

docker-push: docker-build
	docker

deploy-dev:
	@kubectl apply -k ${K8S_ROOT_DIR}/dev
