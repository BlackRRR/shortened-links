run:
	docker-compose up -d --force-recreate --build

tests:
	go test ./test/links_service_test.go
