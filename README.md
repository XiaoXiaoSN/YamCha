# LINE BOT - yamcha

[![Build Status](https://travis-ci.org/XiaoXiaoSN/YamCha.svg?branch=master)](https://travis-ci.org/XiaoXiaoSN/YamCha)
https://github.com/<OWNER>/<REPOSITORY>/workflows/<WORKFLOW_NAME>/badge.svg

## Getting Start
```
make init
make run
```

## TODO
- implement filebase storage (or not)
- 錯誤處理，回傳自定義錯誤訊息

```
docker build . yamcha
docker run --rm \
    -p 18180:18180 \
    -e LINECORP_PLATFORM_CHANNEL_SERVERURL=https://28966197.ngrok.io \
    -e LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=a43bc61e3433ec2023d96141ad42bf6b \
    -e LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=9RCh7QaXx25rPTgyeDuqF1jjSzexpcSQubLkbxjfjYRM7gTeQkekJsY2QEexXqNX1Lf3aQK3W3W0tvV6G1ESA19dPJfGM/rmiXlahQpKEZgeRg6FXOZ0obiPFUdWSIrAuskBZAIuyT1O9Zu9IRBaAQdB04t89/1O/w1cDnyilFU= \
    yamcha
```