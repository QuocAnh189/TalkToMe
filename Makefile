run:
	go run cmd\app\main.go 

docker-compose up:
	docker-compose --env-file .env up --build -d 

docker-compose down:
	docker-compose --env-file .env down