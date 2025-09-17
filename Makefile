run:
	docker compose up --build

args?=./...
test:
	go test $(args)
