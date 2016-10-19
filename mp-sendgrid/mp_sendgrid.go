package mp_sendgrid

import (
	"fmt"
	"os"

	mpModel "../mp-model"

	"github.com/sendgrid/sendgrid-go"
)

var (
	apiKey = os.Getenv("SENDGRID_API_KEY")
	host   = "https://api.sendgrid.com"
)

func makeSubstitutionKey(key string) string {
	return fmt.Sprintf("||%s||", key)
}

func makeFromEmail(from string) string {
	return fmt.Sprintf("\"from\": { \"email\": \"%s\" }", from)
}

func makeTemplateId(id string) string {
	return fmt.Sprintf("\"template_id\": \"%s\"", id)
}

func makeSubject(subject string) string {
	return fmt.Sprintf("\"subject\": \"%s\"", subject)
}

func makeToEmails(listRecipientEmail []string) string {
	toEmail := "\"to\": [ "
	for _, v := range listRecipientEmail {
		toEmail = toEmail + fmt.Sprintf("{ \"email\": \"%s\" }, ", v)
	}

	return toEmail + "]"
}

func makeSubstitutions(listSubstitutions []mpModel.Substitude) string {
	substitutions := "\"substitutions\": { "
	for _, v := range listSubstitutions {
		substitutions = substitutions + fmt.Sprintf("\"%s\": \"%s\",",
			makeSubstitutionKey(v.Key),
			v.Value)
	}

	return substitutions + "}"
}

func makePersonalizations(data mpModel.EmailHeader) string {
	return fmt.Sprintf("\"personalizations\": [ { %s, %s, %s } ]",
		makeSubject(data.Subject),
		makeToEmails(data.ListRecipientEmail),
		makeSubstitutions(data.ListSubstitutions))
}

func makeBody(fromEmail string, templateId string, personalData mpModel.EmailHeader) string {
	return fmt.Sprintf("{ %s, %s, %s }",
		makeFromEmail(fromEmail),
		makeTemplateId(templateId),
		makePersonalizations(personalData))
}

func mailSend(body string) int {
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", host)
	request.Method = "POST"
	request.Body = []byte(body)

	response, err := sendgrid.API(request)

	if err != nil {
		panic(err)
	}

	return response.StatusCode
}

func SendLowInventory(substitude []mpModel.Substitude, emails []string) int {
	subject := "Low Inventory - Daily"
	fromEmail := "no-reply@mokapos.com"
	templateId := os.Getenv("TEMPLATE_ID_LOW_INVENTORY")

	EmailHeader := mpModel.EmailHeader{subject, emails, substitude}
	body := makeBody(fromEmail, templateId, EmailHeader)

	return mailSend(body)
}

func SendDailySales(substitude []mpModel.Substitude, emails []string) int {
	subject := "Sales Summary - Daily"
	fromEmail := "no-reply@mokapos.com"
	templateId := os.Getenv("TEMPLATE_ID_DAILY_SALES")

	EmailHeader := mpModel.EmailHeader{subject, emails, substitude}
	body := makeBody(fromEmail, templateId, EmailHeader)

	return mailSend(body)
}
