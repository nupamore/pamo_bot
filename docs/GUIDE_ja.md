# パモボット(ベータ)日本語説明書

### [ディスコードボット初代リンク](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

## コマンド
- ***$help***  
コマンドリストを見る。

- ***$t [target] [text]***  
翻訳機能。翻訳する言語コード[target]と文字[text]を記します。  
AWS Translate APIを使用します。  
https://docs.aws.amazon.com/translate/latest/dg/what-is.html#what-is-languages

- ***$image [username]***  
アーカイブに保存されたイメージのいずれかをランダムで表示します。 (下記 ***Crawl***参考)

- ***$dice 6***  
最大[max]のランダム数字を取得します。

- ***$crawl [on/off]***  
当該チャンネル(サーバーではない)に新しくアップロードされるイメージを[アーカイブ](bot.nupa.moe)に保存します。  
1つのサーバーで1つのチャンネルのみを監視し、管理者権限のみ可能なコマンドです。    
***$crawl past***は当該チャンネルの過去のイメージをアーカイブに保存します。

## [アーカイブ](http://bot.nupa.moe/)  
ディスコードアカウントでログインして、ボットが入っているサーバーのリストが表示されます。  
自分が上げたイメージや、自分が管理者でいるサーバーのイメージを消すことができます。  
- ランダムイメージ API  
    https://v.nupa.moe/image/[discord_server_id]
