docker-compose up:
	docker-compose --env-file .env -f docker\docker-compose.yml up --build -d 

docker-compose down:
	docker-compose --env-file .env -f docker\docker-compose.yml down