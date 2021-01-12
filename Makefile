.PHONY: clean

K8S_ROOT_DIR:=$(PWD)/k8s

GOOS := drawin
GOARCH := amd64
GO111MODULE := auto

build: build-scrapper build-notify

build-scrapper:
	go build -ldflags="-s -w" -o $(PWD)/scrapper scrapper.go

build-notify:
	go build -ldflags="-s -w" -o $(PWD)/main main.go

watch-scrapper:
	docker-compose run scrapper reflex -r '\.go' -s -- sh -c "go run ./scrapper.go"

watch-notify:
	docker-compose run scrapper reflex -r '\.go' -s -- sh -c "go run ./main.go"

clean:
	kubectl delete namespace/covid19-app-namespace

# Commons
up:
	docker-compose up -d --build

db:
	docker-compose up -d --build postgres

down:
	docker-compose down

logs:
	docker-compose logs -f

# Related to Docker
docker-build:
	docker build . -t m0ai/covid19-bot

docker-push: docker-build
	docker push m0ai/covid19-bot:latest

deploy-dev: docker-build
	kustomize build k8s/dev | kubectl apply -f -
