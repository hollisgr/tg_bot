APP := cmd/main/main.go
OUTPUT := app

TGBOT := github.com/go-telegram/bot
GODOTENV := github.com/joho/godotenv

all: clean build run

build:
	go build -o $(OUTPUT) $(APP)

run:
	./$(OUTPUT)

mod:
	go mod init tg_bot

get:
	go get $(TGBOT)
	go get $(GODOTENV)

clean:
	rm -rf $(OUTPUT)