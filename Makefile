run:
	go run main.go

test:
	gofmt -l .
	[ "`gofmt -l .`" = "" ]
	go test ./...

docker_build:
	docker build --tag coolstuff-go:1.0 .

docker_run:
	docker run -p 5000:5000 --name coolgo coolstuff-go:1.0

docker_clean:
	docker stop coolgo
	docker rm coolgo
