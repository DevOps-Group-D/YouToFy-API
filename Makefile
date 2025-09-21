run:
	go run .

args?=./...
test:
	go test $(args)

build:
	go build -o youtofy-youtube .
