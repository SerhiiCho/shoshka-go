package telegram

import (
	"os"
	"strconv"

	"github.com/SerhiiCho/shoshka_go/utils"
)

func getChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("BOT_CHAT_ID"), 10, 64)
	utils.HandleError(err, "Cannot convert chat id to integer")
	return chatID
}
