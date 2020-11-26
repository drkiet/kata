package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type EmailMessage struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	IsMultipart bool `json:"is_multipart" binding:"required"`
	Date time.Time `json:"date,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	FromRef string `json:"from_ref,omitempty"`
	SenderRef string `json:"sender_ref,omitempty"`
	ToRefs []string `json:"to_refs,omitempty"`
	CcRefs []string `json:"cc_refs,omitempty"`
	BccRefs []string `json:"bcc_refs,omitempty"`
	MessageId string `json:"message_id,omitempty"`
	Subject string `json:"subject,omitempty"`
	ReceivedLines []string `json:"received_lines,omitempty"`
	AdditionalHeaderFields map[string]interface{} `json:"additional_header_fields,omitempty"`
	Body string `json:"body,omitempty"`
	BodyMultipart []MimePartType `json:"body_multipart,omitempty"`
	RawEmailRef string `json:"raw_email_ref,omitempty"`
}

func printEmailMessage(emailMessage EmailMessage) {
	fmt.Println("Email Address:\n", marshalEmailMessage(emailMessage))
}

func marshalEmailMessage(emailMessage EmailMessage) (jsonData string){
	data, e := json.MarshalIndent(emailMessage, "", "  ")
	check(e)
	return string(data)
}

func unmarshalEmailMessage(obj json.RawMessage) (emailMessage EmailMessage) {
	json.Unmarshal(obj, &emailMessage)
	return emailMessage
}
