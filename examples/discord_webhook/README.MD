# Discord Webhook Example

To get your Discord Webhook ID and token, see:https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks  
Once you get the webhook URL, you can get ID and token in following format:  
https://discord.com/api/webhooks/THIS_IS_YOUR_WEBHOOK_ID/THIS_IS_YOUR_WEBHOOK_TOKEN

then run Program

``go run ./examples/discord_webhook/cmd/main.go -id THIS_IS_YOUR_WEBHOOK_ID -token THIS_IS_YOUR_WEBHOOK_TOKEN -host localhost -port 8080 ``

## On Your Game Server
There is no authentication.  
You need to configure event logger on get5: ``get5_remote_log_url "http://localhost:8080/get5/event"``.  
