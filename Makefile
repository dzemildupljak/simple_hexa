
build-dev-docker-container:
	docker build -t my_hexapp-dev .

run-dev-docker-container:
	docker run -p 8080:8080 my_hexapp-dev

http-get-user-list:
	curl http://localhost:8080/api/v1users
http-get-user-by-id:
	curl http://localhost:8080/api/v1users/999
http-create-user:
	curl -X POST -d '{"username": "dummy-username","email":"dummy@mail.com"}' http://localhost:8080/api/v1users

swagger_generate:
	swag init -g cmd/main.go --ot go,json