include .env

dev:
	air

test:
	go test -v ./...

dbmodel:
	go run -mod=mod entgo.io/ent/cmd/ent new $(ARG)

dbgen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent/generated ./ent/schema 

dbdesc:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

dbseed:
	go run ./cmd/seed/seed.go

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./tmp/frameapi ./server/server.go

copytoserver:
	echo "Copying frameapi to $(IP_ADDR)"
	scp -i $(KEY_PATH) ./tmp/frameapi $(USER)@$(IP_ADDR):$(REMOTE_PATH)/frameapi

deployframeapi: build copytoserver
	echo "Deploying frameapi to $(IP_ADDR)"
	ssh -i $(KEY_PATH) $(USER)@$(IP_ADDR) "cd $(REMOTE_PATH) && docker compose up -d --force-recreate --build"

play:
	echo $(ARG)
