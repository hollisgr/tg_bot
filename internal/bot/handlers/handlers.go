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

var API_URL string

func Register(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get_user ", bot.MatchTypePrefix, GetUserById)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get_info", bot.MatchTypeExact, GetUserInfo)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/get_userlist", bot.MatchTypeExact, GetUserList)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/create_user ", bot.MatchTypePrefix, CreateUser)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/delete_user ", bot.MatchTypePrefix, DeleteUserById)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/register_tg_user", bot.MatchTypeExact, RegisterTgUser)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/set_admin ", bot.MatchTypePrefix, SetAdminRole)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/delete_user ", bot.MatchTypePrefix, DeleteUserById)

}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	func1 := "/get_userlist to get list of all users\n"
	func2 := "/get_user {id} to get user by id\n"
	func3 := "/create_user {username} {password} {email} to create user\n"
	func4 := "/delete_user {id} to delete user\n"
	func5 := "/get_info to get user info\n"
	func6 := "/register_tg_user to register tg_user info\n"
	func7 := "/set_admin {password} to set admin role  info\n"

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   func1 + func2 + func3 + func4 + func5 + func6 + func7,
	})
}

func RegisterTgUser(ctx context.Context, b *bot.Bot, update *models.Update) {
	username := update.Message.Chat.Username
	tg_id := update.Message.Chat.ID

	newTgUser := user.TgUser{
		TG_ID:       tg_id,
		TG_USERNAME: username,
	}

	reqBody, _ := json.Marshal(newTgUser)

	RegisterUrl := fmt.Sprintf("%s/tg_users/", API_URL)

	resp, err := http.Post(RegisterUrl, "json", bytes.NewReader(reqBody))

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

func SetAdminRole(ctx context.Context, b *bot.Bot, update *models.Update) {

	body := update.Message.Text
	password := ""
	fmt.Sscanf(body, "/set_admin %s", &password)

	newAdmin := user.TgAdmin{
		TG_ID:     update.Message.Chat.ID,
		ADMIN_PWD: password,
	}

	reqBody, _ := json.Marshal(newAdmin)

	AdminUrl := fmt.Sprintf("%s/admin", API_URL)

	resp, err := http.Post(AdminUrl, "json", bytes.NewReader(reqBody))

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

func GetUserInfo(ctx context.Context, b *bot.Bot, update *models.Update) {

	return_string := fmt.Sprintf("tg_id: %d, tg_username: %s, tg_firstname: %s, lastname: %s", update.Message.Chat.ID, update.Message.Chat.Username, update.Message.Chat.FirstName, update.Message.Chat.LastName)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   return_string,
	})
}

func GetUserList(ctx context.Context, b *bot.Bot, update *models.Update) {

	respURL := API_URL + "/users"
	resp, err := http.Get(respURL)

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

	getUrl := fmt.Sprintf("%s/users/%d", API_URL, id)

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
		TG_ID:    update.Message.Chat.ID,
	}

	reqBody, _ := json.Marshal(newUser)
	respURL := API_URL + "/users"
	resp, err := http.Post(respURL, "json", bytes.NewReader(reqBody))

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

	newUser := user.User{
		TG_ID: update.Message.Chat.ID,
	}

	reqBody, _ := json.Marshal(newUser)

	delUrl := fmt.Sprintf("%s/users/%d", API_URL, id)

	req, err := http.NewRequest(http.MethodDelete, delUrl, bytes.NewReader(reqBody))

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
