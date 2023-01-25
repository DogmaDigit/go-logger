package go_logger

import "errors"

type GoLogger struct {
	Telegram   TelegramNotification
	FolderPath string
}

type TelegramNotification struct {
	Notification bool
	ChatId       string
	BotKey       string
}

var config GoLogger

// Установить обязательные параметры для логгера
func Config(args *GoLogger) error {
	if args == nil {
		return nil
	}
	config.FolderPath = args.FolderPath

	if args.Telegram.Notification && (len(args.Telegram.ChatId) == 0 || len(args.Telegram.BotKey) == 0) {
		return errors.New("required field - Telegram.ChatId Telegram.BotKey")
	}
	config.Telegram.ChatId = args.Telegram.ChatId
	config.Telegram.BotKey = args.Telegram.BotKey
	config.Telegram.Notification = args.Telegram.Notification
	return nil
}

func New() *GoLogger {
	return &config
}
