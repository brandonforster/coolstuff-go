run:
	go run main.go

test:
	go test

docker_build:
	docker build --tag coolstuff:1.0 .

docker_run:
	docker run -p 5000:5000 --name cool coolstuff:1.0

docker_clean:
	docker stop cool
	docker rm cool
