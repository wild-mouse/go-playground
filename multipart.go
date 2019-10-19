package filesapi

import (
	"mime/multipart"

	files "github.com/wild-mouse/go-playground/gen/files"
)

// FilesAddDecoderFunc implements the multipart decoder for service "files"
// endpoint "add". The decoder must populate the argument p after encoding.
func FilesAddDecoderFunc(mr *multipart.Reader, p **files.AddPayload) error {
	// Add multipart request decoder logic here
	return nil
}

// FilesAddEncoderFunc implements the multipart encoder for service "files"
// endpoint "add".
func FilesAddEncoderFunc(mw *multipart.Writer, p *files.AddPayload) error {
	// Add multipart request encoder logic here
	return nil
}
