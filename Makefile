.PHONY: build

K8S_ROOT_DIR:=$(PWD)/k8s

build:
	@cd src \
	&& env GO111MODULE=auto GOOS=linux go build -ldflags="-s -w" -o $(PWD)/main main.go

clean:
	kubectl delete namespace/covid19-app-namespace

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
	docker build . -t m0ai/covid19-bot

docker-push: docker-build
	docker push m0ai/covid19-bot:latest

deploy-dev:
	kustomize build k8s/dev | kubectl apply -f -
