# 甜點問答機器人

甜點問答機器人是一個利用 Go 語言與 OpenAI 來回答各式甜點問題的 Line 與 Telegram Bot。

## 準備事項

在開始前，請確保你已經準備好以下事項：

- 安裝 Go 語言環境。
- 安裝 Git（用於 clone 專案）。
- 擁有 Line 與 Telegram 帳號。
- 準備一個支援 HTTPS 的伺服器，供 Line Bot 的 callback 使用。

## 設定 HTTPS 伺服器

由於 Line Bot 需要 HTTPS 的回呼 callback，您需要設定一個 HTTPS 伺服器。您可以透過以下方式設定：

- 使用 [Let's Encrypt](https://letsencrypt.org/) 等服務取得免費的 SSL/TLS 憑證。
- 開發階段可使用 [ngrok](https://ngrok.com/) 等工具提供安全連線至本地伺服器。

## 設定專案

專案的設定可透過 `config/env.yaml` 檔案或環境變數完成。

### 取得必要參數

#### OpenAI API 金鑰

1. 訪問 [OpenAI 官網](https://openai.com/) 並註冊或登入。
2. 在個人控制台中申請 API 金鑰。
3. 將此金鑰用於專案設定。

#### Line Bot 參數

1. 訪問 [Line Developers](https://developers.line.biz/) 並登入。
2. 建立一個新的 Messaging API 專案。
3. 在專案設定中找到 Channel Secret 和 Channel Access Token。
4. 使用這些參數進行專案設定。

#### Telegram Bot 參數

1. 在 Telegram 中尋找 @BotFather。
2. 依指示建立一個新的 Bot。
3. 記下 Bot 提供的 Token。

### 使用 `config/env.yaml`

1. 在專案根目錄下建立 `config` 資料夾（若尚未存在）。
2. 在 `config` 資料夾中建立 `env.yaml` 檔案。
3. 編輯 `env.yaml` 檔案，加入以下內容：

   ```yaml
   base_url: "https://你的伺服器位址"
   openai_key: "你的 OpenAI API 金鑰"
   line_bot_channel_secret: "你的 Line Bot Channel Secret"
   line_bot_channel_token: "你的 Line Bot Channel Token"
   tg_bot_token: "你的 Telegram Bot Token"
   ```
將括號內的內容替換為實際的值。

使用環境變數
或者，透過設定環境變數進行設定：
```bash
export BASE_URL="https://你的伺服器位址"
export OPENAI_API_KEY="你的 OpenAI API 金鑰"
export LINE_BOT_CHANNEL_SECRET="你的 Line Bot Channel Secret"
export LINE_BOT_CHANNEL_TOKEN="你的 Line Bot Channel Token"
export TG_BOT_TOKEN="你的 Telegram Bot Token"
```

### 啟動伺服器

1. clone 程式碼：
    ```bash
    git clone https://github.com/Rice9547/SweetBot
    ```
2. 進入專案目錄：
    ```bash
    cd SweetBot
    ```
3. 編譯並啟動伺服器
    ```bash
    go build -o server ./cmd/server/
    ./server
    ```