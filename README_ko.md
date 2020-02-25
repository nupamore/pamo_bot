# 파모봇(베타) 한국어 설명서

### [디스코드 봇 초대 링크](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

## 명령어
- ***$help***  
    명령어 목록 보기

- ***$kj (jk, ke, ek, kc, ck, je, ej, fe, ef, fk, kf, sk, ks, ke2, ek2, kj2, jk2)***  
    번역기능. 번역할 언어의 첫 글자와 목적 언어 첫 글자를 적으면 됩니다. (jk = Japanese -> Korean)    
    카카오와 파파고 API를 사용합니다.  
    (Korean, Japanese, English, Chinese, Spanish, French)

- ***$image***  
    아카이브에 저장된 이미지 중 하나를 랜덤으로 보여줍니다. (아래 ***Crawl***참고)

- ***$dice 6***  
    랜덤 숫자를 가져옵니다. (***$dice 100*** = 최대 100까지)

- ***$crawl on***  
    해당 채널(서버 아님)에 새로 올라오는 이미지를 [아카이브](vrc.nupa.moe)에 저장합니다.  
    한 서버에서 하나의 채널만 감시하고, 관리자 권한만 가능한 명령어입니다.  
    ***$crawl off*** 하면 실시간 감시를 그만둡니다.  
    ***$crawl past*** 해당 채널의 과거 이미지들을 아카이브에 저장합니다.  

## [아카이브](http://vrc.nupa.moe/)
디스코드 계정으로 로그인하며, 봇이 들어가 있는 서버의 리스트가 표시됩니다.  
이미지에 마우스를 올리면 삭제 버튼이 표시됩니다.  
자신이 올린 이미지만 표시되며, 서버 관리자는 전부 가능합니다.
- 랜덤이미지 API  
    https://vrc.nupa.moe/randomImage/{ serverId }