module scrapper

go 1.15

require (
	github.com/aws/aws-lambda-go v1.6.0
	github.com/cespare/reflex v0.3.0 // indirect
	github.com/joho/godotenv v1.3.0
	golang.org/x/net v0.0.0-20201207224615-747e23833adb
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.8
)

//require (
//	covid-19-alert-to-slack/slack v0.0.1
//	covid-19-alert-to-slack/scrapper v0.0.1
//)
//replace covid-19-alert-to-slack/slack => ./../pkg/slack
//replace covid-19-alert-to-slack/scrapper => ./../pkg/scrappe
