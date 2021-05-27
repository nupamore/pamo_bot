# Pamobot (Beta) guide

### [Discord bot invitation link](https://discordapp.com/oauth2/authorize?client_id=502450494380179461&permissions=522304&scope=bot)

## command
- ***$help***  
Show commands list

- ***$t [target] [text]***  
Translation function. Write down the language code[target] and sentence[text] to be translated.  
Used the AWS Translate API.  
https://docs.aws.amazon.com/translate/latest/dg/what-is.html#what-is-languages

- ***$image [username]***  
Randomly shows one of the images stored in the archive. (See ***Crawl*** below)

- ***$dice [max]***  
Gets a random number with [max] as the maximum.

- ***$crawl [on/off]***  
Save new images uploaded to the channel(not server) to [Archive](bot.nupa.moe).  
This is a command that monitors only one channel in a server and allows only administrator privileges.  
***$crawl past*** stores past images of the channel in the archive.

## [Archive](http://bot.nupa.moe/)  
Log in with your Discord account, and a list of servers containing the bot is displayed.  
You can erase images you post or images on servers you are an administrator.  
- Random image API  
    https://v.nupa.moe/image/[discord_server_id]
