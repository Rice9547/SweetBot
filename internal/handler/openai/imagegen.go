package openai

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
	"sweetbot/conf/config"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateImage(question string) (string, error) {
	client := openai.NewClient(config.Conf.OpenAIKey)

	resp, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         fmt.Sprintf("%s的成品", question),
			Model:          openai.CreateImageModelDallE3,
			Quality:        openai.CreateImageQualityStandard,
			Size:           openai.CreateImageSize1792x1024,
			ResponseFormat: openai.CreateImageResponseFormatB64JSON,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", nil
	}
	data, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("base64 decode error: %v", err)
		return "", nil
	}
	img, format, err := image.Decode(strings.NewReader(string(data)))
	if err != nil {
		fmt.Printf("image decode error: %v, format: %v\n", err, format)
		return "", nil
	}
	return getResizeImage(img, format, 256)
}

func getResizeImage(img image.Image, format string, width int) (string, error) {
	resizedImg := resize.Resize(uint(width), 0, img, resize.Lanczos3)
	filename := fmt.Sprintf("images/%s.%s", uuid.New().String(), format)
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("file creation error: %v", err)
	}
	fmt.Printf("The image was saved as %s\n", file.Name())
	defer file.Close()

	err = jpeg.Encode(file, resizedImg, nil)
	if err != nil {
		return "", fmt.Errorf("file write error: %v", err)
	}
	return filename, nil
}
