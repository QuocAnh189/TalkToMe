docker-compose up:
	docker-compose -f deployments/docker/docker-compose.yml up --build -d 

docker-compose down:
	docker-compose -f deployments/docker/docker-compose.yml down