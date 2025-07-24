package errors

import "fmt"

type ErrorResponse struct {
	MessageID string
	Err       error
}

func (e *ErrorResponse) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("MessageID: %s, Error: %v", e.MessageID, e.Err)
	}
	return fmt.Sprintf("MessageID: %s, Error: unknown", e.MessageID)
}

func (e *ErrorResponse) ErrorString() string {
	return fmt.Sprintf("MessageId: %s, Error:%v", e.MessageID, e.Err)
}

func (e *ErrorResponse) NewErrorResponse(messageId string, err error) *ErrorResponse {
	return &ErrorResponse{
		MessageID: messageId,
		Err:       err,
	}
}
