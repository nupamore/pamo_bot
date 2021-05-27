# 파모봇(베타) 한국어 설명서

### [디스코드 봇 초대 링크](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

## 명령어
- ***$help***  
명령어 목록 보기

- ***$t [target] [text]***  
번역기능. 번역할 언어코드[target]와 문장[text]을 적으면 됩니다.  
AWS Translate API를 사용합니다.  
https://docs.aws.amazon.com/translate/latest/dg/what-is.html#what-is-languages

- ***$image [username]***  
아카이브에 저장된 이미지 중 하나를 랜덤으로 보여줍니다. (아래 ***Crawl***참고)

- ***$dice [max]***  
[max]를 최대치로 하는 랜덤 숫자를 가져옵니다.

- ***$crawl [on/off]***  
해당 채널(서버 아님)에 새로 올라오는 이미지를 [아카이브](bot.nupa.moe)에 저장합니다.
한 서버에서 하나의 채널만 감시하고, 관리자 권한만 가능한 명령어입니다.  
***$crawl past***는 해당 채널의 과거 이미지들을 아카이브에 저장합니다.

## [아카이브](http://bot.nupa.moe/)
디스코드 계정으로 로그인하며, 봇이 들어가 있는 서버의 리스트가 표시됩니다.  
자신이 올린 이미지나 자신이 관리자로 있는 서버의 이미지를 지울 수 있습니다.  
- 랜덤이미지 API  
    https://v.nupa.moe/image/[discord_server_id]
