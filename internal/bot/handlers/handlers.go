package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tg_bot/internal/user"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Register(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get_user ", bot.MatchTypePrefix, GetUserById)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get_userlist", bot.MatchTypeExact, GetUserList)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/create_user ", bot.MatchTypePrefix, CreateUser)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/delete_user ", bot.MatchTypePrefix, DeleteUserById)
}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	func1 := "/get_userlist to get list of all users\n"
	func2 := "/get_user {id} to get user by id\n"
	func3 := "/create_user {username} {password} {email} to create user\n"
	func4 := "/delete_user {id} to delete user\n"

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   func1 + func2 + func3 + func4,
	})
}

func GetUserList(ctx context.Context, b *bot.Bot, update *models.Update) {

	resp, err := http.Get("http://127.0.0.1:8080/users")

	if err != nil {
		fmt.Println(err)
		return
	}

	usersByte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(usersByte),
	})
}

func GetUserById(ctx context.Context, b *bot.Bot, update *models.Update) {

	body := update.Message.Text
	id := 0
	fmt.Sscanf(body, "/get_user %d", &id)

	getUrl := fmt.Sprintf("http://127.0.0.1:8080/users/%d", id)

	resp, err := http.Get(getUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	userByte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(userByte),
	})
}

func CreateUser(ctx context.Context, b *bot.Bot, update *models.Update) {
	body := update.Message.Text
	username := ""
	pwd := ""
	email := ""
	fmt.Sscanf(body, "/create_user %s %s %s", &username, &pwd, &email)
	newUser := user.User{
		Username: username,
		Password: pwd,
		Email:    email,
	}

	reqBody, _ := json.Marshal(newUser)

	resp, err := http.Post("http://127.0.0.1:8080/users", "json", bytes.NewReader(reqBody))

	if err != nil {
		fmt.Println(err)
		return
	}

	userByte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(userByte),
	})
}

func DeleteUserById(ctx context.Context, b *bot.Bot, update *models.Update) {

	body := update.Message.Text
	id := 0
	fmt.Sscanf(body, "/delete_user %d", &id)

	delUrl := fmt.Sprintf("http://127.0.0.1:8080/users/%d", id)

	req, err := http.NewRequest(http.MethodDelete, delUrl, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	deleteByte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(deleteByte),
	})
}
