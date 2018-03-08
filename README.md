# People -> Pizzas

![Travis CI status](https://travis-ci.org/snyderks/people-to-pizzas.svg?branch=master)

![People to Pizzas Icon](icon.svg)

Take people and get pizzas.

## Setup
1. You need [Go](https://golang.org) to compile and run the app—this was written in `1.9.2`.
2. `go get github.com/snyderks/people-to-pizzas`
3. Install dependencies using [godep](https://github.com/tools/godep) (deprecated, but required for Heroku).
4. Create a file called `config.json` in the directory. Use `testConfig.json` for the format.
    * `slack-token` is a valid Slack API token.
    * `sheets-key` is a valid Google Sheets API token.
    * `sheet-id` is the spreadsheet ID of a **public** Google Sheet.
5. Run tests with `go test ./...`
6. `go build`
7. Run the resulting executable.
8. Set up the Slack slash command to take one parameter: a number.

