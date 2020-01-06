package linebot

import (
	"github.com/line/line-bot-sdk-go/linebot"
	log "github.com/sirupsen/logrus"
)

func (app *YamchaLineBot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	// wake up yamcha!
	case "yamcha", "Yamcha", "飲茶":
		log.Info("reply token:", replyToken)
		_ = app.wakeUp(message, replyToken, source)
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
	case "Yamcha刪除訂單":
		_ = app.deleteConfirm(replyToken, source)
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

func (app *YamchaLineBot) deleteConfirm(replyToken string, source *linebot.EventSource) error {
	if err := app.replyDeleteConfirm(replyToken, source.GroupID); err != nil {
		return err
	}
	return nil
}

func (app *YamchaLineBot) wakeUp(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	log.Printf("Yamcha wakeup in group: %s", source.GroupID)

	// value := app.orderSvc.GetInt(source.GroupID + "Status")
	// if value == StatusYamchaSeelp {
	// err := app.orderSvc.Set(source.GroupID+"Status", 1)
	// if err != nil {
	// 	return err
	// }

	if err := app.replyFlex(replyToken, source.GroupID); err != nil {
		return err
	}

	// } else if value == StatusYamchaWakeUp {
	// 	err := app.Storage.Set(source.GroupID+"Status", 1)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if err := app.replyText(replyToken, "目前點餐的訂單有... (TODO)"); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	return ErrUnknow
	// }

	return nil
}
