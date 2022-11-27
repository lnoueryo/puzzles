init: ## setup docker build, network, and databases
	docker network create --subnet=192.168.0.1/24 puzzles || true
	docker compose up -d

start:
	docker-compose  down
	docker-compose  up -d

stop:
	docker-compose down

restart:
	docker-compose  down
	docker-compose  build --no-cache
	docker-compose  up -d

delete:
	docker-compose --rmi all --volumes --remove-orphans

db:
	docker exec -it puzzles-mysql mysql -uroot -ppassword

