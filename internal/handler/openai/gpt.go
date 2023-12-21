package openai

import (
	"context"
	"fmt"
	"sweetbot/conf/config"

	openai "github.com/sashabaranov/go-openai"
)

type GPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func AskGPT(question string) (string, error) {
	client := openai.NewClient(config.Conf.OpenAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("當被問到關於一種甜點的問題，如單一甜點名稱，比如 '巧克力'，請提供該甜點的一個簡單食譜或製作方法。如果問題是關於一種具體甜點的詳細製作方法，比如 '巧克力蛋糕的製作方法'，請給出詳細的步驟和所需材料。如果問題與甜點無關，請回答 '你應該去找其他人'。%s", question),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
