package alexa

import (
	"encoding/json"
	"time"
)

// APIEndpointAddress is a known of base URI's for calls for device address data to the Alexa service
type APIEndpointAddress string

// AudioPlayerState indicates the last known state of audio playback
type AudioPlayerState string

// ConfirmationStatusState
type ConfirmationStatusState string

// DialogState indicates the status of a multi-turn dialog.
type DialogState string

// RequestTypeName indicates the types of request in the Alexa JSON request
type RequestTypeName string

// ResolutionStatusCode indicates the results of attempting to resolve the user utterance against the defined slot types.
type ResolutionStatusCode string

// SessionEndedReason describes why the session ended.
type SessionEndedReason string

// SessionErrorType describes the type of error that occurred
type SessionErrorType string

// SupportedInterfaces is a supported interface as used as a key in the Device.SupportedInterfaces field
type SupportedInterfaces string

const (
	// AudioPlayerSupported indicates the device supports streaming audio
	AudioPlayerSupported SupportedInterfaces = `AudioPlayer`
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

	ConfirmationStatusStateNone      ConfirmationStatusState = "NONE"
	ConfirmationStatusStateConfirmed ConfirmationStatusState = "CONFIRMED"
	ConfirmationStatusStateDenied    ConfirmationStatusState = "DENIED"

	// DialogStateStarted indicates that dialog has just started
	DialogStateStarted DialogState = "STARTED"
	// DialogStateInProgress indicates that dialog is in progress
	DialogStateInProgress DialogState = "IN_PROGRESS"
	// DialogStateCompleted indicates that dialog has has completed
	DialogStateCompleted DialogState = "COMPLETED"

	ResolutionStatusCodeMatch     ResolutionStatusCode = "ER_SUCCESS_MATCH"
	ResolutionStatusCodeNoMatch   ResolutionStatusCode = "ER_SUCCESS_NO_MATCH"
	ResolutionStatusCodeTimeout   ResolutionStatusCode = "ER_ERROR_TIMEOUT"
	ResolutionStatusCodeException ResolutionStatusCode = "ER_ERROR_EXCEPTION"

	// SessionErrorTypeInvalidResponse indicates that the response was invalid
	SessionErrorTypeInvalidResponse SessionErrorType = "INVALID_RESPONSE"
	// SessionErrorTypeDeviceCommunicationError indicates that there were problems communicating with the device
	SessionErrorTypeDeviceCommunicationError SessionErrorType = "DEVICE_COMMUNICATION_ERROR"
	// SessionErrorTypeInternalError indicates that there was an error with Alexa
	SessionErrorTypeInternalError SessionErrorType = "INTERNAL_ERROR"

	// USAPIEndpointAddress is the base URI for US calls for device address data
	USAPIEndpointAddress APIEndpointAddress = `https://api.amazonalexa.com/`
	// UKDEEndpointAddress is the base URI for UK or DE calls for device address data
	UKDEAPIEndpointAddress APIEndpointAddress = `https://api.eu.amazonalexa.com`

	// LaunchRequestType indicates a request opening a skill, or which has no intents
	LaunchRequestType RequestTypeName = `LaunchRequest`
	// IntentRequestType indicates a request for a skill with an intent
	IntentRequestType RequestTypeName = `IntentRequest`
	// SessionEndedRequestType indicates a currently open session is closed
	SessionEndedRequestType RequestTypeName = `SessionEndedRequest`
	// HelpRequestType is a built in request type indicating that the user has requested help
	HelpRequestType RequestTypeName = `AMAZON.HelpIntent`
	// StopRequestType is a build in request type indicating that the interaction is stopped
	StopRequestType RequestTypeName = `AMAZON.StopIntent`

	AudioPlayerRequestType = `AudioPlayerReq`
)

// AudioPlayer provides the current state for the AudioPlayer interface.
//
// NOTE: AudioPlayer is included on all customer-initiated requests (such as requests made by voice or using a
// remote control), but includes the details about the playback (token and offsetInMilliseconds) only when sent to
// a skill that was most recently playing audio.
type AudioPlayer struct {
	// Token represents the audio stream described by this AudioPlayer object.
	//
	// You provide this token when sending the Play directive. This is only included in the AudioPlayer object when your
	// skill was the skill most recently playing audio on the device.
	Token string `json:"token"`

	// OffsetInMilliseconds identifies a track’s offset in milliseconds at the time the request was sent. This is 0 if
	// the track is at the beginning. This is only included in the AudioPlayer object when your skill was the skill
	// most recently playing audio on the device.
	OffsetInMilliseconds int64 `json:"offsetInMilliseconds"`

	// ActivityState indicates the last known state of audio playback
	ActivityState AudioPlayerState `json:"playerActivity"`
}

// BaseRequestType contains fields that are common across all RequestTypes
type BaseRequestType struct {
	// Timestamp provides the date and time when Alexa sent the request as an ISO 8601 formatted string.
	//
	// This value is used to verify the request when hosting your skill as a web service.
	Timestamp time.Time `json:"timestamp"`

	// ID represents a unique identifier for the specific request.
	ID string `json:"requestId"`

	// Locale is a string indicating the user’s locale. For example: en-US.
	Locale string `json:"locale"`
}

// GetID returns the unique ID for the request
func (request *BaseRequestType) GetID() string {
	return request.ID
}

// GetTimestamp returns the supplied time for the request
func (request *BaseRequestType) GetTimestamp() time.Time {
	return request.Timestamp
}

// GetLocale returns the supplied locale for the request
func (request *BaseRequestType) GetLocale() string {
	return request.Locale
}

// Context provides your skill with information about the current state of the Alexa service and device at the time
// the request is sent to your service.
//
// This is included on all requests. For requests sent in the context of a session (LaunchRequest and IntentRequest),
// the context object duplicates the user and application information that is also available in the Request.Session
// field.
type Context struct {
	// System provides information about the current state of the Alexa service and the device interacting with your
	// skill.
	System *System `json:"system"`

	// AudioPlayer provides the current state for the AudioPlayer interface.
	//
	// NOTE: AudioPlayer is included on all customer-initiated requests (such as requests made by voice or using a
	// remote control), but includes the details about the playback (token and offsetInMilliseconds) only when sent to
	// a skill that was most recently playing audio.
	//
	// Requests that are not customer-initiated, such as the AudioPlayer requests do not include the AudioPlayer object in
	// the context. For these requests, the request type indicates the current state (for example, the request
	// AudioPlayer.PlaybackStarted indicates that the playback has started) and details about the state are part of the
	// request object.
	AudioPlayer *AudioPlayer `json:"AudioPlayer"`
}

// Device provides information about the device used to send a request.
type Device struct {
	// ID to uniquely identify a device
	ID string `json:"deviceId"`

	// SupportedInterfaces lists each interface that the device supports. For example, if SupportedInterfaces includes
	// the key AudioPlayerSupported, then you know that the device supports streaming audio using the Alexa AudioPlayer
	// interface.
	SupportedInterfaces map[SupportedInterfaces]interface{}
}

// Request is an implementation of a JSON Alexa request. The definition of the request is made by Amazon and defined in
// the Alexa Documentation. This request will be available in the body of POST request from the Alexa service
//
// All requests include the context, and request objects at the top level. The session object is included for
// all standard requests, but it is not included for AudioPlayer, VideoApp, or PlaybackController requests.
type Request struct {
	// Version is the specifier for the request with the value. It is currently always defined as: “1.0”
	Version string `json:"version"`

	// Session provides additional context associated with the request.
	//
	// NOTE: The session is included for all standard requests, but it is not included for AudioPlayer, VideoApp, or
	// PlaybackController requests.
	Session *RequestSession `json:"session"`

	// Context provides your skill with information about the current state of the Alexa service and device at the time
	// the request is sent to your service.
	//
	// This is included on all requests. For requests sent in the context of a session (LaunchRequest and
	// IntentRequest), the context object duplicates the user and application information that is also available in the
	// session.
	Context *Context `json:"context"`

	// Request provides the details of the user’s request. There are several different request types available
	// * LaunchRequest
	// * IntentRequest
	// * SessionEndedRequest
	// * AudioPlayer Requests
	// * Display.RenderTemplate Requests
	// * Display.ElementSelected Requests
	// * VideoApp Requests
	// * PlaybackController Requests
	Request RequestType
}

// UnmarshalJSON implements the json.Unmarshaler interface for the Request type. It extracts the Request data and
// populates the Data field with an appropriate RequestData type
func (request *Request) UnmarshalJSON(b []byte) error {
	type requestAlias Request
	aux := &struct {
		RequestDetail json.RawMessage `json:"request"`
		*requestAlias
	}{
		requestAlias: (*requestAlias)(request),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return request.determineRequestType(aux.RequestDetail)
}

func (request *Request) determineRequestType(b []byte) error {
	identifier := &struct {
		Type RequestTypeName `json:"type"`
	}{}

	if err := json.Unmarshal(b, identifier); err != nil {
		return err
	}

	switch identifier.Type {
	case LaunchRequestType:
		request.Request = &LaunchRequest{}
	case IntentRequestType:
		request.Request = &IntentRequest{}
	case SessionEndedRequestType:
		request.Request = &SessionEndedRequest{}
	}

	return json.Unmarshal(b, request.Request)
}

// RequestSession provides additional context associated with a request. Standard request types (LaunchRequest,
// IntentRequest, and SessionEndedRequest) include the session object.
//
// Requests from interfaces such as AudioPlayer and PlaybackController are not sent in the context of a session, so
// they do not include the session object.
type RequestSession struct {
	// IsNew indicates whether this is a new session. True indicates a new session, False indicates an existing session.
	IsNew bool `json:"new"`

	// ID represents a unique identifier per a user’s active session.
	//
	// NOTE: An ID is consistent for multiple subsequent requests for a user and session. If the session ends for
	// a user, then a new unique ID value is provided for subsequent requests for the same user.
	ID string `json:"sessionId"`

	// Attributes is a map of key-value pairs. The attributes map is empty for requests where a new session has started
	// with the property new set to true.
	//
	// The key is a string that represents the name of the attribute. The value is an type that represents the value of
	// the attribute. It should be cast to the appropriate type
	//
	// When returning your Response, you can include data you need to persist during the session in the
	// SessionAttributes property. The attributes you provide are then passed back to your skill on the next request.
	Attributes map[string]interface{} `json:"attributes"`

	// ApplicationID represents the Application ID associated in Alexa. The skill’s application ID is displayed on the
	// Skill Information page in the developer portal. This information is also available in the System.Application
	// field
	ApplicationID string

	// User describes the user making the request. This is a user from the perspective of the Alexa system
	User *User `json:"user"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for the RequestSession type. It extracts the application
// object and pulls out the ID without having to build a public type
func (session *RequestSession) UnmarshalJSON(b []byte) error {
	type sessionAlias RequestSession
	aux := &struct {
		Application struct {
			ID string `json:"applicationID"`
		} `json:"application"`
		*sessionAlias
	}{
		sessionAlias: (*sessionAlias)(session),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	session.ApplicationID = aux.Application.ID
	return nil
}

// RequestType is the interface used in the Request struct to hold an instance of a specific instance of the requested
// request
//
// Known implementations are:
//   * LaunchIntent
//   * IntentRequest
//   * SessionEndedRequest
type RequestType interface {
	GetType() RequestTypeName
	GetID() string
	GetTimestamp() time.Time
	GetLocale() string
}

// A Resolutions object representing the results of resolving the words captured from the user’s utterance.
//
// This is included for slots that use a custom slot type or a built-in slot type that you have extended with your own
// values. Note that resolutions is not included for built-in slot types that you have not extended
type Resolution struct {
	Authorities []*ResolutionAuthority `json:"resolutionsPerAuthority"`
}

type ResolutionAuthority struct {
	Name   string `json:"authority"`
	Status *struct {
		Code ResolutionStatusCode `json:"code"`
	} `json:"status"`
	Values []*ResolutionValue `json:"values"`
}

type ResolutionValue struct {
	Value *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"value"`
}

// SessionError is object providing more information about the error that occurred.
type SessionError struct {
	// Type indicates the type of error that occurred
	Type    SessionErrorType `json:"type"`
	Message string           `json:"message"`
}

type Slot struct {
	// Name represents the name of the slot.
	Name string `json:"name"`

	// Value represents the value the user spoke for the slot. This is the actual value the user spoke, not necessarily
	// the canonical value or one of the synonyms defined for the entity.
	//
	// NOTE: AMAZON.LITERAL slot values sent to your service are always in all lower case.
	Value string `json:"value"`

	// ConfirmationStatus indicates whether the user has explicitly confirmed or denied the value of this slot.
	ConfirmationStatus ConfirmationStatusState `json:"confirmationStatus"`

	// Resolutions represents the results of resolving the words captured from the user’s utterance.
	//
	// This is included for slots that use a custom slot type or a built-in slot type that you have extended with your
	// own values. Note that resolutions is not included for built-in slot types that you have not extended.
	Resolutions *Resolution `json:"resolutions"`
}

// System provides information about the current state of the Alexa service and the device interacting with your skill.
type System struct {
	// ApplicationID represents the Application ID associated in Alexa. The skill’s application ID is displayed on the
	// Skill Information page in the developer portal.
	ApplicationID string

	// User describes the user making the request. This is a user from the perspective of the Alexa system
	User *User `json:"user"`

	// Device provides information about the device used to send the request.
	Device *Device `json:"device"`

	// APIEndpoint references the correct base URI to refer to by region. The base URI for US calls for device address
	// data is: https://api.amazonalexa.com/. The base URI for UK and DE calls for device address data is:
	// https://api.eu.amazonalexa.com.
	APIEndpoint APIEndpointAddress `json:"apiEndpoint"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for the RequestSession type. It extracts the application
// object and pulls out the ID without having to build a public type
func (system *System) UnmarshalJSON(b []byte) error {
	type systemAlias System
	aux := &struct {
		Application struct {
			ID string `json:"applicationID"`
		} `json:"application"`
		*systemAlias
	}{
		systemAlias: (*systemAlias)(system),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	system.ApplicationID = aux.Application.ID
	return nil
}

// User describes the user making a request from the perspective of the Alexa system.
type User struct {
	// ID is a unique identifier for the user who made the request. The length of this identifier can vary, but is
	// never more than 255 characters. The ID is automatically generated when a user enables the skill in the Alexa app.
	//
	// NOTE: A user who disables and re-enables a skill will have a new identifier generated
	ID string `json:"userId"`

	// AccessToken is a token identifying the user in another system. This is only provided if the user has successfully
	// linked their account.
	AccessToken string `json:"accessToken"`

	// PermissionConsentToken is a token allowing the skill access to information that the customer has consented to
	// provide, such as address information.
	Permissions *UserPermission `json:"permissions"`
}

// UserPermission contain details allowing the skill access to information that the customer has consented to provide,
// such as address information.
type UserPermission struct {
	// ConsentToken is a provided token for accessing customer information
	ConsentToken string `json:"consentToken"`
}

/*

func (d *Device) HasAudioPlayerSupport() bool {
	_, hasSupport := d.SupportedInterfaces[audioPlayerSupported]
	return hasSupport
}

*/
