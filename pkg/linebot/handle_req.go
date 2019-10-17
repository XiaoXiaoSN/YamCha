package linebot

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func (app *YamchaLineBot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	// wake up yamcha!
	case "yamcha", "Yamcha", "飲茶":
		log.Println("reply token:", replyToken)
		_ = app.wekeUp(message, replyToken, source)
		// if _, err := app.bot.ReplyMessage(
		// 	replyToken,
		// 	linebot.NewTextMessage("echo: "+message.Text),
		// ).Do(); err != nil {
		// 	return err
		// }

		// default:
		// 	log.Printf("Echo message to %s: %s", replyToken, message.Text)
		// 	if _, err := app.bot.ReplyMessage(
		// 		replyToken,
		// 		linebot.NewTextMessage("echo: "+message.Text),
		// 	).Do(); err != nil {
		// 		return err
		// 	}

	case "成功了！":
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage("哇嗚你好厲害"),
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

func (app *YamchaLineBot) wekeUp(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	log.Printf("Yamcha wakeup in group: %s", source.GroupID)

	value := app.Storage.GetInt(source.GroupID + "Status")
	if value == StatusYamchaSeelp {
		err := app.Storage.Set(source.GroupID+"Status", 1)
		if err != nil {
			return err
		}
		// beta: return menu
		if err := app.replyFlex(replyToken, "嘿！ 今天想喝點什麼?"); err != nil {
			return err
		}
		// end beta
		// if err := app.replyText(replyToken, "嘿！ 今天想喝點什麼?"); err != nil {
		// 	return err
		// }
	} else if value == StatusYamchaWakeUp {
		err := app.Storage.Set(source.GroupID+"Status", 1)
		if err != nil {
			return err
		}

		if err := app.replyText(replyToken, "目前點餐的訂單有... (TODO)"); err != nil {
			return err
		}
	} else {
		return ErrUnknow
	}

	return nil
}
