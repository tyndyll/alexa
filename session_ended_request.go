package alexa

const (
	// SessionEndedReasonUserInitiated indicates the user explicitly ended the session.
	SessionEndedReasonUserInitiated SessionEndedReason = "USER_INITIATED"
	// SessionEndedReasonError indicates an  error occurred that caused the session to end.
	SessionEndedReasonError SessionEndedReason = "ERROR"
	// SessionEndedReasonUserExceededMaxReprompts indicates the user either did not respond or responded with an
	// utterance that did not match any of the intents defined in your voice interface
	SessionEndedReasonUserExceededMaxReprompts SessionEndedReason = "EXCEEDED_MAX_REPROMPTS"
)

// SessionEndedRequest is an object that represents a request made to an Alexa skill to notify that a session was ended.
// Your service receives a SessionEndedRequest when a currently open session is closed for one of the following reasons:
//  * The user says “exit”.
//  * The user does not respond or says something that does not match an intent defined in your voice interface while
// 		the device is listening for the user’s response.
//  * An error occurs.
//
// NOTE: Setting the shouldEndSession flag to true in your response also ends the session. In this case, your service
// does not receive a SessionEndedRequest.
//
// NOTE: Your skill cannot return a response to SessionEndedRequest.
type SessionEndedRequest struct {
	*BaseRequestType

	// Reason describes why the session ended.
	Reason SessionEndedReason `json:"reason"`
}

func (request *SessionEndedRequest) GetType() RequestTypeName {
	return SessionEndedRequestType
}
