MAIN_FILE="./cmd/mockPizza/main.go"
OUT_FILE="mockPizza"
DB_VOL = "/postgres-volume:/var/lib/postgresql/data"
DB_CONTAINER="postgres-pizza"
NETWORK="postgresnet"

postgres:
	docker run --rm --name ${DB_CONTAINER} \
		-p 5432:5432 \
		-e POSTGRES_USER=user \
		-e POSTGRES_PASSWORD=1234 \
		-v ${DB_VOL} \
		--network postgresnet \
		-d postgres
	docker run --rm --name adminer-pizza \
		-p 5050:8080 \
		--network ${NETWORK} \
		-d adminer

createdb:
	docker exec -it ${DB_CONTAINER} createdb --username=user --owner=user mockpizzadb

build:
	go build -o ${OUT_FILE} ${MAIN_FILE}

run:
	go run ${MAIN_FILE}

deps:
	go mod tidy
	go mod download
