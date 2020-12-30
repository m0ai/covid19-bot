.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/scrapper scrapper/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

watch:
	reflex -r '\.go' -s -- sh -c "go run ./main.go"

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f
