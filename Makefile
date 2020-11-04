.PHONY: build
.PHONY: start
.PHONY: list
.PHONY: kill
.PHONY: logs
#.PHONY: migrate
#.PHONY: migrate-undo
#.PHONY: migration-create

## build: Build and run server
build:
	go build -o ./build ./cmd/api
	./build/api # run App

## start: Start deamonized
start:
	./cmd/pmgo start ./cmd/api api-server

## list: List deamon process
list:
	./cmd/pmgo list

## logs: Show info log stream
logs:
	tail -f ${HOME}/.pmgo/api-server/api-server.err

## kill: Kill deamon
kill:
	./cmd/pmgo kill

### migrate: Migrate to database schema all unapplied migrations
#migrate:
#	migrate -path migrations -database postgres://localhost/quiz_development?sslmode=disable up
#
### migrate-undo: Undoing all migrations
#migrate-undo:
#	migrate -path migrations -database postgres://localhost/quiz_development?sslmode=disable down
#
### migration-create: Create new migrations file, please specify filename with "name" argument
#migration-create:
#	migrate create -ext sql -dir migrations $(name)

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

.DEFAULT_GOAL := build