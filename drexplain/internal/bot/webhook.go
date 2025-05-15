package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitWebhook(bot *tgbotapi.BotAPI, url, port string) error {
	wh, err := tgbotapi.NewWebhook(url + port + "/" + bot.Token)
	if err != nil {
		return fmt.Errorf("failed to set webhook: %v", err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		return fmt.Errorf("failed to request webhook: %v", err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return fmt.Errorf("failed to get webhook info: %v", err)
	}

	if info.LastErrorDate != 0 {
		return fmt.Errorf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	return nil
}
