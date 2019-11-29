package linebot

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"

	log "github.com/sirupsen/logrus"
)

// define reply const content
const (
	ConfirmMessageYes = "好的!"
	ConfirmMessageNo  = "不要!"
)

// // Hero type structure
// type Hero struct {
// 	Type        string `json:"type"`
// 	URL         string `json:"url"`
// 	Size        string `json:"size"`
// 	AspectRatio string `json:"aspectRatio"`
// 	AspectMode  string `json:"aspectMode"`
// }

// type Button struct {
// 	Type   string `json:"type"`
// 	Style  string `json:"style"`
// 	Height string `json:"height"`
// 	Action
// }

// type Action struct {
// 	Type  string `json:"type"`
// 	Label string `json:"label"`
// 	URI   string `json:"uri"`
// }

// type FlexMessage struct {
//   Type string `json:"type"`
//   Hero
//   Body struct {
//     Type
//     Layout
//     Contents Button[]
//   } `json:"body"`
// }
var initJSONData = `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "http://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "飲茶 ver.0.0.1 beta",
        "weight": "bold",
        "size": "xl"
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "查詢訂單",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "刪除訂單",
          "uri": "https://www.google.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "新增品項",
          "uri": "line://app/1653300700-EjDoldvQ?order=<orderID>"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "查詢菜單",
          "uri": "line://app/1653300700-ydEGLgZR?order=<orderID>"
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      }
    ],
    "flex": 0
  }
}`

var progressMenuString = `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "http://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "飲茶 ver.0.0.1 beta",
        "weight": "bold",
        "size": "xl"
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "新增訂單",
          "uri": "line://app/1653300700-odMBaL1P?group=<groupID>"
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      }
    ],
    "flex": 0
  }
}`

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

func (app *YamchaLineBot) replyFlex(replyToken string, groupID string) error {
	// log.Println("reply token in replyFlex:", replyToken)
	// log.Println("groupID:", groupID)

	// check if order exists from db
	// log.Println("service", app.orderSvc)

	//  check if order exists
	if orderData, errMsg := app.orderSvc.GetGroupOrder(groupID); errMsg != nil {

		log.Println("err:", errMsg)

		// return create order menu if not
		progressMenuJSON := strings.Replace(progressMenuString, "<groupID>", groupID, -1)

		if container, err := linebot.UnmarshalFlexMessageJSON([]byte(progressMenuJSON)); err != nil {
			log.Println("err:", err)
			return err
		} else if _, errorMsg := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("alt message", container),
		).Do(); errorMsg != nil {
			log.Println("reply token:", replyToken)
			log.Println("reply text err:", errorMsg)
			return errorMsg
		}
		// return err
	} else {
		// return selection menu

		initMenuJSON := strings.Replace(initJSONData, "<orderID>", string(orderData.ID), -1)

		if container, err := linebot.UnmarshalFlexMessageJSON([]byte(initMenuJSON)); err != nil {
			log.Println("err:", err)
			return err
		} else if _, errorMsg := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("alt message", container),
		).Do(); errorMsg != nil {
			log.Println("reply token:", replyToken)
			log.Println("reply text err:", errorMsg)
			return errorMsg
		}
		// return err
	}

	// initMenuJSON := strings.Replace(initJSONData, "<groupID>", groupID, -1)
	// if container, err := linebot.UnmarshalFlexMessageJSON([]byte(initMenuJSON)); err != nil {
	// 	log.Println("err:", err)
	// 	return err
	// } else if _, errorMsg := app.bot.ReplyMessage(
	// 	replyToken,
	// 	linebot.NewFlexMessage("alt message", container),
	// ).Do(); errorMsg != nil {
	// 	log.Println("reply token:", replyToken)
	// 	log.Println("reply text err:", errorMsg)
	// 	return errorMsg
	// }
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
