package alexa

// IntentRequest is an object that represents a request a user makes to a skill that maps to an intent. The request
// object sent to your service includes the specific intent and any defined slot values.
//
// An intent represents a high-level action that fulfills a user’s spoken request. Intents can optionally have
// arguments called slots that collect additional information needed to fulfill the user’s request.
//
// Note that an IntentRequest can either start a new session or continue an existing session, depending on how the user
// begins interacting with the skill:
//   e.g. The user asks Alexa a question or states a command, all in a single phrase. This sends a new IntentRequest
//         and starts a new session:
//         User: Alexa, Ask GoDoc for ResponseWriter.
//
// In this case, no LaunchRequest is ever sent to your service. The IntentRequest starts session instead.
//
// Once a session is already open, the user states a command that maps to one of the intents defined in your voice
// interface. This sends the IntentRequest within the existing session:
//   e.g. User: Alexa, talk to GoDoc # (This command sends the service a LaunchRequest and opens a new session)
//        Alexa: You can ask for a function, a struct or an interface. Which will it be? (response to the LaunchRequest)
//        User: Give me the doc for ReadAll (This command sends the service an IntentRequest with the existing,
//                open session)
//
// Note that Alexa provides a collection of built-in intents for very common actions, such as stopping, canceling,
// and providing help. If you choose to use these intents, your code for handling IntentRequest requests needs to handle
// them as well.
//
// If you are using the skill builder (beta) and you have created a dialog model, the IntentRequest includes a
// DialogState property. You can use this to determine the current status of the conversation with the user and return
// the Dialog.Delegate directive if the conversation is not yet complete.
type IntentRequest struct {
	*BaseRequestType

	// DialogState indicates the status of a multi-turn dialog. This property is included if the skill meets the
	// requirements to use the Dialog directives
	DialogState DialogState `json:"dialogState"`

	// Intent represents what the action the user wants to perform
	Intent *Intent `json:"intent"`
}

type Intent struct {
	// Name represents the name of the intent. It is set in the Alexa Developer console
	Name string `json:"name"`

	// ConfirmationStatus indicates whether the user has explicitly confirmed or denied the entire intent.
	ConfirmationStatus ConfirmationStatusState `json:"confirmationStatus"`

	// Slots is a map of key-value pairs that further describes what the user meant based on a predefined intent schema.
	// The map can be empty.
	Slots map[string]*Slot `json:"slots"`
}

func (request *IntentRequest) GetType() RequestTypeName {
	return IntentRequestType
}
