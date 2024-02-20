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

play:
	echo $(ARG)
