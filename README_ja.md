# パモボット(ベータ)日本語説明書

### [ディスコードボット初代リンク](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

## コマンド
- ***$help***  
コマンドリストを見る。

- ***$kj(jk、ke、ek、kc、ck、je、ej、fe、ef、fk、kf、sk、ks、ke2、ek2、kj2、jk2)***  
翻訳機能。翻訳する言語の最初の文字と目的言語の最初の文字を記します。 (jk=Japanese→Korean)  
KakaoとPapago APIを使用します。  
(Korean、 Japanese、 English、 Chinese、 Spanish、 French)

- ***$image***  
アーカイブに保存されたイメージのいずれかをランダムで表示します。 (下記 ***Crawl***参考)

- ***$dice 6***  
ランダム数字を取得します。 (***$dice 100*** = 最大100まで)

- ***$crawl on***  
当該チャンネル(サーバーではない)に新しくアップロードされるイメージを[アーカイブ](vrc.nupa.moe)に保存します。  
1つのサーバーで1つのチャンネルのみを監視し、管理者権限のみ可能なコマンドです。  
***$crawloff***したらリアルタイム監視をやめます。

## [アーカイブ](http://vrc.nupa.moe/)  
ディスコードアカウントでログインし、ボットが入っているサーバーのリストが表示されます。  
イメージにマウスを上げると削除ボタンが表示されます。  
自分がアップロードした画像のみ表示され、サーバー管理者は全て可能です。
- ランダムイメージ API  
    https://vrc.nupa.moe/randomImage/{ serverId }