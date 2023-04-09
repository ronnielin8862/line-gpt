# line-gpt
Gpt in Line


## How to use
1. tc => translate to Chinese , "tc how are you", will reply "你好嗎"
2. tj => translate to Japanese, "tj how are you?", will reply "お元気ですか？"
3. te => translate to English
4. tt => translate to Thai
5. tk => translate to Khmer
6. ai => talk with chatGPT, "ai 請用一個六歲小孩的日文，表達他很喜歡他母親", will reply "わたしのお母さん、だいすき！"
## How to build
1. apply Line developer account and get channel secret and channel token
2. apply GPT api token
2. Setting config in project root path named `config.toml`
    ````
    [line_server]
    channel_secret = "abc"
    channel_token = "abd"
    
    [gpt]
    auth_token = "abc"
    
    [server]
    host = ""
    port = 9487
    ````
2. Run `go run ./cmd/line-gpt/line-gpt.go`