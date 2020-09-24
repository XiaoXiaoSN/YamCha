# LINE BOT - yamcha

[![Build Status](https://travis-ci.org/XiaoXiaoSN/YamCha.svg?branch=master)](https://travis-ci.org/XiaoXiaoSN/YamCha)
![deploy heroku](https://github.com/XiaoXiaoSN/yamcha/workflows/deploy%20heroku/badge.svg)

## Getting Started
啟動需要四個環境變數，需要到 [heroku 設定](https://dashboard.heroku.com/apps/yamcha/settings) Settings > Config Vars
- LINECORP_PLATFORM_CHANNEL_SERVERURL
- LINECORP_PLATFORM_CHANNEL_CHANNELSECRET
- LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN
- MYSQL_DSN

### Docker directly
```
docker build . -t yamcha
docker run --rm \
    -p 18180:18180 \
    -e LINECORP_PLATFORM_CHANNEL_SERVERURL=https://28966197.ngrok.io \
    -e LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=a43bc61e3433ec2023d96141ad42bf6b \
    -e LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=9RCh7QaXx25rPTgyeDuqF1jjSzexpcSQubLkbxjfjYRM7gTeQkekJsY2QEexXqNX1Lf3aQK3W3W0tvV6G1ESA19dPJfGM/rmiXlahQpKEZgeRg6FXOZ0obiPFUdWSIrAuskBZAIuyT1O9Zu9IRBaAQdB04t89/1O/w1cDnyilFU= \
    - e MYSQL_DSN=root:root@tcp(mysql.example.com)/db_name?charset=utf8&parseTime=true \
    yamcha
```

### VSCode Remote Container
這是個令人興奮的功能，用 docker 包裝開發環境免去了繁複的環境設定
只需要在 VSCode 安裝 `Remote - Containers` 插件

雖然這不是一個複雜的專案，但是個好的開始，對吧 ^__^

## Deploy New Version
push master 後會自動部署到 heroku

## Future
- migrate
- gorm v2
- dev container 自帶 mysql
- dev container mysql 自帶測試資料
- test
- viper
- pkg/error

## Learn More
This is a golang [Line bot](https://github.com/line/line-bot-sdk-go) project
