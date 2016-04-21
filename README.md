# slack-translator-bot
##### overview
demo Golang backend code for a Slack /slash command to translate text from one language to another.

i wrote this because i wanted 2x new /slash commands:

`/e2j <english text>` to translate English to Japanese

`/j2e <japanese text>` to translate Japanese to English

##### example run
assuming slash command `/e2j` has been setup to point to your appengine instance, if a user entered `/e2j the rain in spain falls mainly on the plane` the response would be something like:

![ScreenShot](http://i.imgur.com/zu8pKFc.png)
