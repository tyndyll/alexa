package alexa

// LaunchRequest is an object that represents a user request to an Alexa skill, where the user invokes the skill with
// the invocation name, but does not provide any command mapping to an intent e.g. "Alexa, talk to MyGoSkill"
//
// For skills that just do one thing e.g. tells trivia, the service can take action without requesting more information
// from the user. Services that need more information from the user may need to respond with a prompt.
//
// A LaunchRequest always starts a new session.
//
// A skill can respond to LaunchRequest with any combination of:
//   * Standard response properties (OutputSpeech, Card, and Reprompt).
//   * Any AudioPlayer directives.
type LaunchRequest struct {
	*BaseRequestType
}

func (request *LaunchRequest) GetType() RequestTypeName {
	return LaunchRequestType
}
