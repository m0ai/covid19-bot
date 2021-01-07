.PHONY: build

K8S_ROOT_DIR:=$(PWD)/k8s

build: build-scrapper build-notify

build-scrapper:
	@cd src \
	&& env GO111MODULE=auto go build -ldflags="-s -w" -o $(PWD)/scrapper scrapper.go

build-notify:
	@cd src \
	&& env GO111MODULE=auto go build -ldflags="-s -w" -o $(PWD)/main main.go

clean:
	kubectl delete namespace/covid19-app-namespace

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
	docker build . -t m0ai/covid19-bot

docker-push: docker-build
	docker push m0ai/covid19-bot:latest

deploy-dev: docker-build
	kustomize build k8s/dev | kubectl apply -f -
