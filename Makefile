build-dev-docker-container:
	docker build -t my_hexapp-dev .

run-dev-docker-container:
	docker run -p 8080:8080 my_hexapp-dev
