package alexa

import "time"

// AudioPlayerState is the state an AudioPlayer request may have
type AudioPlayerState string

type DialogState string

type ConfirmationStatusState string

type ResolutionStatusCode string

type SessionEndedReason string

type SessionErrorType string

const (
	audioPlayerSupported = `AudioPlayer`

	// AudioPlayerIdleState indicates nothing was playing, no enqueued items.
	AudioPlayerIdleState AudioPlayerState = "IDLE"
	// AudioPlayerPausedState indicates the stream was paused.
	AudioPlayerPausedState AudioPlayerState = "PAUSED"
	// AudioPlayerPlayingState indicates the stream was playing.
	AudioPlayerPlayingState AudioPlayerState = "PLAYING"
	// AudioPlayerBufferUnderrunState indicates that there was a buffer underrun
	AudioPlayerBufferUnderrunState AudioPlayerState = "BUFFER_UNDERRUN"
	// AudioPlayerFinishedState indicates the stream was finished playing.
	AudioPlayerFinishedState AudioPlayerState = "FINISHED"
	// AudioPlayerStoppedState indicates the stream was interrupted
	AudioPlayerStoppedState AudioPlayerState = "STOPPED"

	// DialogStateStarted indicates that dialog has just started
	DialogStateStarted DialogState = "STARTED"
	// DialogStateInProgress indicates that dialog is in progress
	DialogStateInProgress DialogState = "IN_PROGRESS"
	// DialogStateCompleted indicates that dialog has has completed
	DialogStateCompleted DialogState = "COMPLETED"

	ConfirmationStatusStateNone      ConfirmationStatusState = "NONE"
	ConfirmationStatusStateConfirmed ConfirmationStatusState = "CONFIRMED"
	ConfirmationStatusStateDenied    ConfirmationStatusState = "DENIED"

	ResolutionStatusCodeMatch     ResolutionStatusCode = "ER_SUCCESS_MATCH"
	ResolutionStatusCodeNoMatch   ResolutionStatusCode = "ER_SUCCESS_NO_MATCH"
	ResolutionStatusCodeTimeout   ResolutionStatusCode = "ER_ERROR_TIMEOUT"
	ResolutionStatusCodeException ResolutionStatusCode = "ER_ERROR_EXCEPTION"

	SessionEndedReasonUserInitiated            SessionEndedReason = "USER_INITIATED"
	SessionEndedReasonError                    SessionEndedReason = "ERROR"
	SessionEndedReasonUserExceededMaxReprompts SessionEndedReason = "EXCEEDED_MAX_REPROMPTS"

	SessionErrorTypeInvalidResponse          SessionErrorType = "INVALID_RESPONSE"
	SessionErrorTypeDeviceCommunicationError SessionErrorType = "DEVICE_COMMUNICATION_ERROR"
	SessionErrorTypeInternalError            SessionErrorType = "INTERNAL_ERROR"
)

// Request is an implementation of a JSON Alexa request
type Request struct {
	// Version specifier for the request with the value defined as: “1.0”
	Version string `json:"version"`
	// Session provides additional context associated with the request.
	Session *Session `json:"session"`

	Context *Context `json:"context"`
}

// Session provides additional context associated with a request. Standard request types (LaunchRequest, IntentRequest,
// and SessionEndedRequest) include the session object.
//
// Requests from interfaces such as AudioPlayer and PlaybackController are not sent in the context of a session, so
// they do not include the session object.
type Session struct {
	// Indicates whether this is a new session. Returns true for a new session or false for an existing session.
	New bool `json:"new"`
	// ID represents a unique identifier per a user’s active session.
	// Note: An ID is consistent for multiple subsequent requests for a user and session. If the session ends for
	// a user, then a new unique ID value is provided for subsequent requests for the same user.
	ID string `json:"sessionId"`
	// Attributes is a map of key-value pairs.
	//
	// The key is a string that represents the name of the attribute. The value is an type that represents the value of
	// the attribute. It should be cast to the appropriate type
	//
	// When returning your Response, you can include data you need to persist during the session in the
	// SessionAttributes property. The attributes you provide are then passed back to your skill on the next request.
	Attributes map[string]interface{} `json:"attributes"`
	// Application contains application specific data
	Application *Application `json:"application"`
}

// Context provides your skill with information about the current state of the Alexa service and device at the time
// the request is sent to your service. This is included on all requests. For requests sent in the context of a session
// (LaunchRequest and IntentRequest), the context object duplicates the user and application information that is also
// available in the Request.Session field.
type Context struct {
	System      *System      `json:"system"`
	AudioPlayer *AudioPlayer `json:"AudioPlayer"`
}

// Application contains Alexa application information. It can be used to verify that a Request was intended for your
// service
type Application struct {
	// ID represents the application ID for your skill. A skill’s application ID is displayed on the Skill
	// Information page in the developer portal.
	ID string `json:"applicationID"`
}

// User is an object that describes the user making a request
type User struct {
	// ID is a unique identifier for the user who made the request. The length of this identifier can vary, but is
	// never more than 255 characters. The ID is automatically generated when a user enables the skill in the Alexa app.
	ID string `json:"userId"`
	// AccessToken is a token identifying the user in another system. This is only provided if the user has successfully
	// linked their account.
	AccessToken string `json:"accessToken"`
	// Permissions contains a consentToken allowing the skill access to information that the customer has consented to
	// provide, such as address information.
	Permissions *Permission `json:"permissions"`
}

// Permission contain deteails allowing the skill access to information that the customer has consented to provide,
// such as address information.
type Permission struct {
	// ConsentToken is a provided token for accessing customer information
	ConsentToken string `json:"consentToken"`
}

// System provides information about the current state of the Alexa service and the device interacting with your skill.
type System struct {
	// Application containing an application ID. This is used to verify that the request was intended for your service
	Application *Application `json:"application"`
	// User describes the user making the request
	User *User `json:"user"`
	// Device provides information about the device used to send the request.
	Device *Device `json:"device"`
	// APIEndpoint references the correct base URI to refer to by region. The base URI for US calls for device address
	// data is: https://api.amazonalexa.com/. The base URI for UK and DE calls for device address data is:
	// https://api.eu.amazonalexa.com.
	APIEndpoint string `json:"apiEndpoint"`
}

// Device provides information about the device used to send a request.
type Device struct {
	// ID to uniquely identify a device
	ID string `json:"deviceId"`

	SupportedInterfaces map[string]interface{}
}

func (d *Device) HasAudioPlayerSupport() bool {
	_, hasSupport := d.SupportedInterfaces[audioPlayerSupported]
	return hasSupport
}

type AudioPlayer struct {
	Token                string           `json:"token"`
	OffsetInMilliseconds int64            `json:"offsetInMilliseconds"`
	PlayerActivity       AudioPlayerState `json:"playerActivity"`
}

// LaunchRequest is an object that represents that a user made a request to an Alexa skill, but did not provide a
// specific intent.
type LaunchRequest struct {
	// Represents a unique identifier for the specific request.
	ID string `json:"requestId"`
	// Provides the date and time when Alexa sent the request as an ISO 8601 formatted string
	Timestamp time.Time `json:"timestamp"`
	// A string indicating the user’s locale. For example: en-US.
	Locale string `json:"locale"`
}

type IntentRequest struct {
	*LaunchRequest
	DialogState DialogState `json:"dialogState"`
	Intent      *Intent     `json:"intent"`
}

type Intent struct {
	Name               string                  `json:"name"`
	ConfirmationStatus ConfirmationStatusState `json:"confirmationStatus"`
	Slots              map[string]*Slot        `json:"slots"`
}

type Slot struct {
	Name               string                  `json:"name"`
	Value              string                  `json:"value"`
	ConfirmationStatus ConfirmationStatusState `json:"confirmationStatus"`
}

// A Resolutions object representing the results of resolving the words captured from the user’s utterance.
//
// This is included for slots that use a custom slot type or a built-in slot type that you have extended with your own
// values. Note that resolutions is not included for built-in slot types that you have not extended
type Resolution struct {
	Authorities []*ResolutionAuthority `json:"resolutionsPerAuthority"`
}

type ResolutionAuthority struct {
	Name   string             `json:"authority"`
	Status *ResolutionStatus  `json:"status"`
	Values []*ResolutionValue `json:"values"`
}

type ResolutionStatus struct {
	Code ResolutionStatusCode `json:"code"`
}

type ResolutionValue struct {
	Value *ResolutionValueDetail `json:"value"`
}

type ResolutionValueDetail struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SessionEndedRequest struct {
	*LaunchRequest
	Reason SessionEndedReason `json:"reason"`
}

type SessionError struct {
	Type    SessionErrorType `json:"type"`
	Message string           `json:"message"`
}
