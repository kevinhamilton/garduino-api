default:
	@echo "============= Building Production API ============="
	docker build -f Dockerfile.prod -t garduino .

up: default
	@echo "=============starting docker locally============="
	docker-compose up -d

logs:
	docker logs garduino-api_web_1 --follow

down:
	docker-compose down

test:
	go test

kill:
	docker-compose kill