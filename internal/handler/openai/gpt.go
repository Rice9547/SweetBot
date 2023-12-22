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

var example1 = "巧克力：以下是製作巧克力的簡單食譜：\n\n所需材料：\n- 200克巧克力（可以選擇黑巧克力或牛奶巧克力，根據個人喜好）\n- 30克奶油\n- 可選配料（例如果仁、乾果、蔓越莓等）\n\n製作步驟：\n1. 將巧克力切成小塊，放入碗中。\n2. 在熱水中放一個鍋子，將碗放在鍋子上方，以水蒸氣加熱巧克力，讓巧克力慢慢融化。\n3. 一邊融化巧克力，一邊加入奶油，不斷攪拌混合，直到巧克力完全融化且與奶油充分結合。\n4. 如果你喜歡加入其他配料，可以現在加入，並輕輕攪拌混合。\n5. 將巧克力混合物倒入模具中，讓它冷卻至室溫，然後放入冰箱冷藏至固化。\n6. 等待巧克力固化後，即可取出並享用。\n\n希望這個簡單的巧克力食譜能幫助到你！如果你有其他疑問，歡迎繼續問我。"
var example2 = "果凍：以下是製作果凍的簡單食譜：\n\n材料：\n- 500克水果（可選擇你喜歡的水果，如草莓、藍莓、葡萄等）\n- 150克砂糖（可根據個人口味調整）\n- 10克明膠粉\n- 100毫升水\n\n步驟：\n1. 將水果洗淨，去皮去籽，切成小塊狀。\n2. 將切好的水果放入攪拌機中，搗碎成泥狀。\n3. 在一個小碗中，將明膠粉和水混合，攪拌均勻，待明膠充分融化。\n4. 在一個鍋子中，加入砂糖和水果泥，以中小火加熱，攪拌至糖完全融化。\n5. 將明膠溶液加入鍋中，繼續攪拌均勻，使所有成分充分混合。\n6. 關火後，將果醬倒入製模冷卻，在室溫下放置約30分鐘，然後放入冰箱中冷藏至果醬完全凝固。\n7. 完全凝固後，用刀子輕輕切成你喜歡的形狀，即可享用。\n\n提示：你還可以在水果中加入一些柠檬汁或橙汁，以增加果醬的酸味。並且可以根據個人喜好，在果醬中添加一些碎果肉或果仁，以增加口感。\n\n希望這個簡單的果凍食譜能幫助到你！如果你有其他疑問，歡迎繼續問我。"

func AskGPT(question string) (string, error) {
	client := openai.NewClient(config.Conf.OpenAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4,
			MaxTokens:   1024,
			Temperature: 0.75,
			TopP:        0.8,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "當被問到關於一種甜點的問題，請根據以下範例提供該甜點的一個簡單食譜或製作方法。如果問題是關於一種具體甜點的詳細製作方法，比如 '巧克力蛋糕的製作方法'，請給出詳細的步驟和所需材料。如果問題與甜點無關，請回答 '你應該去找其他人'。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("%s", question),
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: fmt.Sprintf("例子 1：%s\n\n例子 2：%s\n\n現在根據用戶的問題，請提供相關的甜點食譜。", example1, example2),
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
