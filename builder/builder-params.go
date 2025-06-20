package main

import (
	"fmt"
	"strings"
)

type email struct {
	from    string
	to      string
	subject string
	body    string
}

type EmailBuilder struct {
	email email
	err   error
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if b.err != nil {
		return b
	}
	if !strings.Contains(from, "@") {
		b.err = fmt.Errorf("From should contain an @")
		return b
	}
	b.email.from = from
	return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
	if b.err != nil {
		return b
	}
	if !strings.Contains(to, "@") {
		b.err = fmt.Errorf("From should contain an @")
		return b
	}
	b.email.to = to
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	if b.err != nil {
		return b
	}
	b.email.subject = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	if b.err != nil {
		return b
	}
	b.email.body = body
	return b
}

func sendMailEmailImpl(email *email) {
	fmt.Printf("Email sent to `%s` by `%s` with subject `%s` and body as `%s`\n", email.to, email.from, email.subject, email.body)
}

type build func(*EmailBuilder)

func SendMail(action build) error {
	emailBuilder := EmailBuilder{}
	action(&emailBuilder)
	if emailBuilder.err != nil {
		return emailBuilder.err
	}
	sendMailEmailImpl(&emailBuilder.email)
	return nil
}

func TestBuilderParams() {
	err := SendMail(func(eb *EmailBuilder) {
		eb.
			From("abc@gmail.com").
			To("def@yahoo.com").
			Subject("First email").
			Body("Hello, how are you?")
	})
	if err != nil {
		fmt.Printf("Error sending email: %v", err.Error())
	}
	err = SendMail(func(eb *EmailBuilder) {
		eb.
			From("abcgmail.com").
			To("def@yahoo.com").
			Subject("First email").
			Body("Hello, how are you?")
	})
	if err != nil {
		fmt.Printf("Error sending email: %v", err.Error())
	}
}
