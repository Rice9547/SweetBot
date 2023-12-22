package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sweetbot/conf/config"
	"sweetbot/internal/handler/openai"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotHandler handles incoming requests from LINE platform
func LineBotHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		config.Conf.LineBotChannelSecret,
		config.Conf.LineBotChannelToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				answer, err := openai.AskGPT(message.Text)
				if err != nil {
					log.Print(err)
					continue
				}

				answerMessage := linebot.NewTextMessage(answer)

				if answer == "" {
					answer = "抱歉，我無法回答"
				}

				if answer == "你應該去找其他人。" {
					if _, err = bot.ReplyMessage(event.ReplyToken, answerMessage).Do(); err != nil {
						log.Print(err)
					}
					continue
				}

				imgResultChan := make(chan string)
				go func() {
					imgURL, err := openai.GenerateImage(message.Text)
					if err != nil {
						imgResultChan <- ""
					} else {
						imgResultChan <- imgURL
					}
					close(imgResultChan)
				}()

				imgURL := <-imgResultChan
				if imgURL != "" {
					defer func() {
						time.Sleep(time.Second * 20)
						os.Remove(imgURL)
					}()
					fullImageURL := fmt.Sprintf("%s/%s", config.Conf.BaseURL, imgURL)
					imgMessage := linebot.NewImageMessage(fullImageURL, fullImageURL)
					txtMessage := linebot.NewTextMessage("這是一張可能的成品圖片")
					if _, err = bot.ReplyMessage(event.ReplyToken, answerMessage, imgMessage, txtMessage).Do(); err != nil {
						log.Printf("send image failed with err: %v", err)
					}
				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, answerMessage).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}

}
