package linebot

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// define reply const content
const (
	ConfirmMessageYes = "好的!"
	ConfirmMessageNo  = "不要!"
)

// TODO: 允許複數個 text
func (app *YamchaLineBot) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		log.Println("reply text err:", err)
		return err
	}
	return nil
}

// TODO: 自訂 button 內容和功用
func (app *YamchaLineBot) quickReplyButton(replyToken, text string) error {
	_, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text).
			WithQuickReplies(linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton(
					"option1",
					linebot.NewMessageAction("account link", "link")),
				linebot.NewQuickReplyButton(
					"option2",
					linebot.NewMessageAction("list user", "list")),
			)),
	).Do()
	if err != nil {
		log.Println("reply quickReplyButton err:", err)
		return err
	}

	return nil
}

// TODO: buttom
func (app *YamchaLineBot) replyImageButton(replyToken, text string) error {
	imageURL := "https://www.edureka.co/blog/wp-content/uploads/2018/09/Golang-Logo-Golang-Tutorial-Edureka.jpg"
	template := linebot.NewButtonsTemplate(
		imageURL, "My button sample", "Hello, my button",
		linebot.NewURIAction("Go to line.me", "https://line.me"),
		linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
		linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
		linebot.NewMessageAction("Say message", "Rice=米"),
	)
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTemplateMessage("Buttons alt text", template),
	).Do(); err != nil {
		log.Println("reply buttons err:", err)
		return err
	}

	return nil
}

func (app *YamchaLineBot) confirm(replyToken, text string) error {
	template := linebot.NewConfirmTemplate(
		text,
		linebot.NewMessageAction(ConfirmMessageYes, ConfirmMessageYes),
		linebot.NewMessageAction(ConfirmMessageNo, ConfirmMessageNo),
	)
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTemplateMessage("Confirm alter text", template),
	).Do(); err != nil {
		log.Println("reply confirm err:", err)
		return err
	}
	return nil
}
