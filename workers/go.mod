module scrapper

go 1.15

replace (
	covid-19-alert-to-slack/common/slackUtil => ../common/slackUtil
)

require (
	github.com/aws/aws-lambda-go v1.6.0
	github.com/cespare/reflex v0.3.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	golang.org/x/net v0.0.0-20201207224615-747e23833adb
	covid-19-alert-to-slack/common/slackUtil v0.0.1
)
