# slack-translator-bot
##### overview
demo Golang backend code for a Slack /slash command to translate text from one language to another.

i wrote this because i wanted 2x new /slash commands:

`/e2j <english text>` to translate English to Japanese

`/j2e <japanese text>` to translate Japanese to English

##### example run
assuming slash command `/e2j` has been setup in your slack as a new slash command to translate English to Japanese, you've pointed it to your appengine instance, if a user entered `/e2j the rain in spain falls mainly on the plane` the response would be something like:

![ScreenShot](http://i.imgur.com/zu8pKFc.png)

similarly, assuming `/j2e` has been setup as above (but for Japanese to English), if a user entered `/j2e スペインの雨は、平面上に、主に落ちます` the response would be something like: 

![ScreenShot](http://i.imgur.com/anBWGyE.png)

the demo code as you see it here is only setup for English <--> Japanese, but you can see how easy it would be to tailor to your specific needs/reqs.

##### how-to setup
instructions are for English <--> Japanese

1. the code as-is is for google app engine. create a new appengine app, do what you need to do
2. get your google translate api key. follow the steps here: https://cloud.google.com/translate/v2/getting_started
3. create your new /slash commands `/e2j` and `/j2e`. configure them to GET to the URIs `/translate/en_ja` and `/translate/ja_en` at the FQDN of your appengine app from step 1.
4. change the appropriate values in `app.yaml` (`GOOGLE_TRANSLATE_API_KEY` and `SLACK_VERIFY_TOKEN`). make sure your application name is correct (etc)
5. `$ go build && goapp deploy`
6. profit

