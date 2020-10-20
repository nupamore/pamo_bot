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
Save new images uploaded to the channel(not server) to [Archive](vrc.nupa.moe).  
This is a command that monitors only one channel in a server and allows only administrator privileges.  
***$crawl past*** stores past images of the channel in the archive.

## [Archive](http://vrc.nupa.moe/)  
Log in with your Discord account, and a list of servers containing the bot is displayed.  
When you hover over the image, you will see a delete button.  
Only images uploaded by yourself are displayed, and server administrators can do all of them.  
- Random image API  
    https://vrc.nupa.moe/randomImage/[discord_server_id]
