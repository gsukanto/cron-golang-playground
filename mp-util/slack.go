package mp_util

import (
	"fmt"
	"os"

	"github.com/bluele/slack"
)

const (
	webhookUrl = "https://hooks.slack.com/services/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"
	mokaLogo   = ":moka_logo:"
)

func getEnvironment() string {
	var environmet string
	switch os.Getenv("ENV") {
	case "development":
		environmet = "Local"
	case "develop-stage":
		environmet = "Development"
	case "staging":
		environmet = "Staging"
	case "production":
		environmet = "Production"
	default:
		environmet = "Unknown"
	}
	return environmet
}

func SendChat(from string, title string, color string, channel string, text string) int {
	hook := slack.NewWebHook(webhookUrl)
	err := hook.PostMessage(&slack.WebHookPostPayload{
		Username:    from,
		IconEmoji:   mokaLogo,
		Attachments: []*slack.Attachment{{Color: color, Title: title, Text: text, MarkdownIn: []string{"text"}}},
		Channel:     channel})

	if err != nil {
		panic(err)
		return -1
	}

	return 0
}

func SlackLowInventorySummary(businessCount int64, emailSent int64, emailFailed int64) int {
	from := "Moka Report"
	channel := "#site-reliability"
	title := fmt.Sprintf("Low Inventory Email Report - %s - %s", getEnvironment(), GetJakartaTimeNow())

	summary := fmt.Sprintf("```Business count : %v\nEmail Sent     : %v\nEmail Failed   : %v```",
		businessCount, emailSent, emailFailed)

	var attachmentColor string
	if emailFailed > 0 {
		attachmentColor = "danger"
	} else {
		attachmentColor = "good"
	}

	statusCode := SendChat(from, title, attachmentColor, channel, summary)
	return statusCode
}

func SlackDailySalesSummary(businessCount int64, emailSent int64, emailFailed int64) int {
	from := "Moka Report"
	channel := "#site-reliability"
	title := fmt.Sprintf("Daily Sales Email Report - %s - %s", getEnvironment(), GetJakartaTimeNow())

	summary := fmt.Sprintf("```Business count : %v\nEmail Sent     : %v\nEmail Failed   : %v```",
		businessCount, emailSent, emailFailed)

	var attachmentColor string
	if emailFailed > 0 {
		attachmentColor = "danger"
	} else {
		attachmentColor = "good"
	}

	statusCode := SendChat(from, title, attachmentColor, channel, summary)
	return statusCode
}
