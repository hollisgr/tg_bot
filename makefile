APP := cmd/main/main.go
TGBOT := github.com/go-telegram/bot
OUTPUT := app

all: clean build run

build:
	go build -o $(OUTPUT) $(APP)

run:
	./$(OUTPUT)

mod:
	go mod init tg_bot

get:
	go get $(TGBOT)

clean:
	rm -rf $(OUTPUT)