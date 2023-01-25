package go_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	TELEGRAM_URL          = "https://api.telegram.org/bot"
	TELEGRAM_SEND_MESSAGE = "/sendMessage"
)

func (gl *GoLogger) sendTelegramMessage(message string, trace []string) {
	chatId := gl.Telegram.ChatId
	url := TELEGRAM_URL + gl.Telegram.BotKey + TELEGRAM_SEND_MESSAGE

	content := fmt.Sprintf("Приложение: <b>%s</b>\nОшибка: %s\nТрассировка:", os.Getenv("APP_NAME"), message)
	for _, v := range trace {
		content += "<code>" + v + "</code>"
	}

	payload, _ := json.Marshal(map[string]string{
		"chat_id":    chatId,
		"text":       content,
		"parse_mode": "html",
	})

	read := bytes.NewBuffer(payload)
	res, err := http.Post(url, "application/json", read)
	if err != nil {
		fmt.Println(err)
	}

	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
