package linebot

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func (app *YamchaLineBot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	// wake up yamcha!
	case "yamcha", "Yamcha", "飲茶":
	case "profile":
		if source.UserID != "" {
			profile, err := app.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return app.replyText(replyToken, err.Error())
			}
			if _, err := app.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return app.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case "buttons":
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
			return err
		}
	case "confirm":
		if err := app.confirm(replyToken, "確定嗎?"); err != nil {
			return err
		}
	case "quick":
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage("Select your favorite food category or send me your location!").
				WithQuickReplies(linebot.NewQuickReplyItems(
					linebot.NewQuickReplyButton(
						"https://www.edureka.co/blog/wp-content/uploads/2018/09/Golang-Logo-Golang-Tutorial-Edureka.jpg",
						linebot.NewMessageAction("Sushi", "Sushi")),
					linebot.NewQuickReplyButton(
						"https://pbs.twimg.com/profile_images/1142154201444823041/O6AczwfV.png",
						linebot.NewMessageAction("Tempura", "Tempura")),
					linebot.NewQuickReplyButton(
						"",
						linebot.NewLocationAction("Send location")),
				)),
		).Do(); err != nil {
			return err
		}
	case "bye":
		switch source.Type {
		case linebot.EventSourceTypeUser:
			return app.replyText(replyToken, "Bot can't leave from 1:1 chat")
		case linebot.EventSourceTypeGroup:
			if err := app.replyText(replyToken, "Leaving group"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveGroup(source.GroupID).Do(); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		case linebot.EventSourceTypeRoom:
			if err := app.replyText(replyToken, "Leaving room"); err != nil {
				return err
			}
			if _, err := app.bot.LeaveRoom(source.RoomID).Do(); err != nil {
				return app.replyText(replyToken, err.Error())
			}
		}
	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage("echo: "+message.Text),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (app *YamchaLineBot) handleSticker(message *linebot.StickerMessage, replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewStickerMessage(message.PackageID, message.StickerID),
	).Do(); err != nil {
		return err
	}
	return nil
}
