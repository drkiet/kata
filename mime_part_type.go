package kata

type MimePartType struct {
	Body string `json:"body,omitempty"`
	BodyRawRef string `json:"body_raw_ref,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
}
