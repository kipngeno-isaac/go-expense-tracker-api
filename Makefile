test: docker_build_test
	docker-compose up -d
	docker-compose exec -T http go test ./...
	docker-compose down

unit_test:
	go test `go list ./... | grep -v e2e_test

docker_build:
	docker build . -t service

docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	docker run --publish 8080:8080 service