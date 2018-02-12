package alexa

import "encoding/json"

const (
	// AudioPlayerPlaybackStartedType is sent when Alexa begins playing the audio stream previously sent in a Play
	// directive. This lets your skill verify that playback began successfully.
	AudioPlayerPlaybackStartedType = `AudioPlayer.PlaybackStarted`
	// AudioPlayerPlaybackFinishedType is sent when the stream Alexa is playing comes to an end on its own.
	AudioPlayerPlaybackFinishedType = `AudioPlayer.PlaybackFinished`
	// AudioPlayerPlaybackStoppedType is sent when Alexa stops playing an audio stream in response to a voice request or
	// an AudioPlayer directive.
	AudioPlayerPlaybackStoppedType = `AudioPlayer.PlaybackStopped`
	// AudioPlayerPlaybackNearlyFinishedType is sent when the currently playing stream is nearly complete and the
	// device is ready to receive a new stream.
	AudioPlayerPlaybackNearlyFinishedType = `AudioPlayer.PlaybackNearlyFinished`
	// AudioPlayerPlaybackFailedType is sent when Alexa encounters an error when attempting to play a stream.
	AudioPlayerPlaybackFailedType = `AudioPlayer.PlaybackFailed `

	PauseIntentType    = `AMAZON.PauseIntent`
	ResumeIntentType   = `AMAZON.ResumeIntent`
	CancelIntentType   = `AMAZON.CancelIntent`
	LoopOffIntentType  = `AMAZON.LoopOffIntent`
	LoopOnIntentType   = `AMAZON.LoopOnIntent`
	NextIntentType     = `AMAZON.NextIntent`
	PreviousIntentType = `AMAZON.PreviousIntent`
	RepeatIntentType   = `AMAZON.RepeatIntent`
	ShuffleOffIntent   = `AMAZON.ShuffleOffIntent`
	ShuffleOnIntent    = `AMAZON.ShuffleOnIntent`
	StartOverIntent    = `AMAZON.StartOverIntent`
)

type AudioPlayerRequest struct {
	*BaseRequestType

	// Reason describes why the session ended.
	Reason SessionEndedReason `json:"reason"`
}

func (request *AudioPlayerRequest) GetType() RequestTypeName {
	return SessionEndedRequestType
}


type AudioDirective struct {
	Type string `json:"type"`
	PlayBehavior string `json:"playBehavior"`
	AudioItem *AudioStream `json:"audioItem"`
}

type AudioStream struct {
	URL string `json:"url"`
	Token string `json:"token"`
	ExpectedPreviousToken string `json:"expectedPreviousToken"`
	Offset int64 `json:"offsetInMilliseconds"`
}

func (stream *AudioStream) MarshalJSON() ([]byte, error) {
	type streamAlias AudioStream
	aux := struct {
		Stream *streamAlias `json:"stream"`
	} { (*streamAlias)(stream) }
	return json.Marshal(&aux)
}