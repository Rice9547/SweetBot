package openai

import (
	"context"
	"fmt"
	"sweetbot/conf/config"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateImage(question string) (string, error) {
	client := openai.NewClient(config.Conf.OpenAIKey)

	respUrl, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         fmt.Sprintf("%s的成品", question),
			Model:          openai.CreateImageModelDallE3,
			Size:           openai.CreateImageSize1792x1024,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", nil
	}
	return respUrl.Data[0].URL, nil
}
