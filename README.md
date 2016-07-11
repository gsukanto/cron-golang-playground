# moka-cron

**This project is a microservice for sending cron for daily summary and low inventory email.**

# Installation

## Setup Golang Environment
please refer to [golang getting-started](https://golang.org/doc/install)

## Setup Environment Variables
update your environment with your [SENDGRID_API_KEY](https://app.sendgrid.com/settings/api_keys).
get you api-key from wilson

```bash
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore
source ./sendgrid.env
```

## Install Package

`go get github.com/sendgrid/sendgrid-go`

```go
import "github.com/sendgrid/sendgrid-go"
```



## Dependencies
- The SendGrid [SDK](https://github.com/sendgrid/sendgrid-go)
