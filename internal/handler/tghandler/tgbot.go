package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sweetbot/conf/config"
	"sweetbot/internal/handler/openai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandler struct {
	Bot *tgbotapi.BotAPI
}

func NewTelegramHandler(webhookURL string) (*TelegramHandler, error) {
	bot, err := tgbotapi.NewBotAPI(config.Conf.TGBotToken)
	if err != nil {
		return nil, err
	}

	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		return nil, err
	}
	_, err = bot.Request(webhookConfig)
	if err != nil {
		return nil, err
	}

	return &TelegramHandler{Bot: bot}, nil
}

func (th *TelegramHandler) HandleUpdates(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Error decoding update: %v", err)
		return
	}

	th.handleUpdate(update)
}

func (th *TelegramHandler) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	answer, err := openai.AskGPT(update.Message.Text)
	if err != nil {
		log.Print(err)
		return
	}

	if answer == "" {
		answer = "抱歉，我無法回答"
	}

	if answer == "你應該去找其他人" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
		th.Bot.Send(msg)
		return
	}

	imgResultChan := make(chan string)
	go func() {
		imgURL, err := openai.GenerateImage(update.Message.Text)
		if err != nil {
			imgResultChan <- ""
		} else {
			imgResultChan <- imgURL
		}
		close(imgResultChan)
	}()

	imgURL := <-imgResultChan
	if imgURL != "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
		th.Bot.Send(msg)
		photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(imgURL))
		th.Bot.Send(photoMsg)
		txtMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "這是一張可能的成品圖片")
		th.Bot.Send(txtMsg)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
		th.Bot.Send(msg)
	}
}
