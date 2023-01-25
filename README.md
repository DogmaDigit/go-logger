# Go-Logger

Пакет для пользовательского предствления вывода информации.

Ps. Использует пакет `github.com/pkg/errors`

## Установка

`go get gitlab.tap2v.com/dogma/utils/go-logger`

## Настройка

Перед использованием, необходимо обязательно настроить логгер.

Описание конфигурации:
```
type GoLogger struct {
	Telegram   TelegramNotification
	FolderPath string // путь к папке, где будут храниться логи, будут разбиваться на типы логов (error.log и т.д.) 
}

type TelegramNotification struct {
	Notification bool // флаг, сигнализирующи об необходимости отправки уведомлений в телеграм
	ChatId       string // ID чата в телеграме, обязательный параметры, если telegramNotification = true
	BotKey       string // Ключ бота в телеграме, обязательный параметры, если telegramNotification = true
}
```

Использование:
```
package app

import (
	"errors"

	"gitlab.tap2v.com/dogma/utils/go-logger"
)

func main() {
	// настройка логгера
	go_logger.Config(go_logger.GoLogger{
		FolderPath: ".logs",
		go_logger.TelegramNotification{
			Notification: true,
			ChatId: "-123456789",
			BotKey: "1022473515:AAFjlxWqSOXauX-ydyiUhEbnQrBGd_PEg",
		},
	})
}

func someHandler() {
	// получение экземпляра логгера
	logger := go_logger.New()

	// использование логгера
	logger.Info("info log")
	logger.Warn("warning log")
	logger.Error("error log", errors.New("some err"), false)
}
```
