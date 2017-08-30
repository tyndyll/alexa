package alexa

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRequestUnmarshal(t *testing.T) {
	Convey(`When I unmarshal the LaunchRequest JSON into a Request struct`, t, func() {
		request := &Request{}
		if err := json.Unmarshal(launchRequestJSON, request); err != nil {
			panic(err)
		}

		Convey(`Then the Version field will be set to 1.0`, func() {
			So(request.Version, ShouldEqual, `1.0`)
		})

		Convey(`Then the Session New field will be set correctly`, func() {
			So(request.Session.New, ShouldBeTrue)
		})

		Convey(`Then the Session ID field will be set correctly`, func() {
			So(request.Session.ID, ShouldEqual, `7547f227-9a55-4974-94cf-ed6fdd1cc631`)
		})

		Convey(`Then the Session Attributes map will be set correctly`, func() {
			So(request.Session.Attributes[`4c96a4a6-0082-463c-90dd-3108c8ae787f`], ShouldEqual, `d69e7caa-c029-4187-83e8-bdb2987baad7`)
		})

		Convey(`Then the Session ApplicationID will be set correctly`, func() {
			So(request.Session.ApplicationID, ShouldEqual, `a3a5bb41-e0c9-438b-a8f9-7f2e4de84641`)
		})

		Convey(`Then the Session User ID will be set correctly`, func() {
			So(request.Session.User.ID, ShouldEqual, `32549228-5f2d-4749-98c2-fe3a394366ad`)
		})

		Convey(`Then the Session User AccessToken will be set correctly`, func() {
			So(request.Session.User.AccessToken, ShouldEqual, `67b433cd-1a8a-4a53-ae66-402ca5a03fe1`)
		})

		Convey(`Then the Session User Permissions ConsentToken will be set correctly`, func() {
			So(request.Session.User.Permissions.ConsentToken, ShouldEqual, `c6cc6b93-3f7b-4b88-ad47-97186886cd72`)
		})

		Convey(`Then the Context System ApplicationID will be set correctly`, func() {
			So(request.Context.System.ApplicationID, ShouldEqual, `a3a5bb41-e0c9-438b-a8f9-7f2e4de84641`)
		})

		Convey(`Then the Context System User ID will be set correctly`, func() {
			So(request.Context.System.User.ID, ShouldEqual, `32549228-5f2d-4749-98c2-fe3a394366ad`)
		})

		Convey(`Then the Context System User AccessToken will be set correctly`, func() {
			So(request.Context.System.User.AccessToken, ShouldEqual, `67b433cd-1a8a-4a53-ae66-402ca5a03fe1`)
		})

		Convey(`Then the Context System User Permissions ConsentToken will be set correctly`, func() {
			So(request.Context.System.User.Permissions.ConsentToken, ShouldEqual, `c6cc6b93-3f7b-4b88-ad47-97186886cd72`)
		})

		Convey(`Then the Context System Device ID will be set correctly`, func() {
			So(request.Context.System.Device.ID, ShouldEqual, `05cdef34-7fbf-4be5-9051-09125301f935`)
		})

		Convey(`Then the Context System Device Supported Interfaces should contain AudioPlayerSupported`, func() {
			_, found := request.Context.System.Device.SupportedInterfaces[AudioPlayerSupported]
			So(found, ShouldBeTrue)
		})

		Convey(`Then the Context System APIEndpoint field will be set correctly`, func() {
			So(request.Context.System.APIEndpoint, ShouldEqual, `26909ef5-2c76-4929-9309-90352a425ef4`)
		})

		Convey(`Then the Detail type will be a LaunchRequest struct`, func() {
			So(request.Detail, ShouldHaveSameTypeAs, &LaunchRequest{})
		})

		Convey(`Then the Detail GetID will be set correctly`, func() {
			So(request.Detail.GetID(), ShouldEqual, "58eeabff-a19a-4fd1-87f1-7c45789769ad")
		})
	})

	Convey(`When I unmarshal the IntentRequest JSON into a Request struct`, t, func() {
		request := &Request{}
		if err := json.Unmarshal(intentRequestJSON, request); err != nil {
			panic(err)
		}

		Convey(`Then the Detail type will be a LaunchRequest struct`, func() {
			So(request.Detail, ShouldHaveSameTypeAs, &IntentRequest{})
		})

		Convey(`Then the Detail ID will be set correctly`, func() {
			So(request.Detail.GetID(), ShouldEqual, "c4763de6-1e7d-4842-99d5-29fe1836b577")
		})
	})

	Convey(`When I unmarshal the SessionEndedRequest JSON into a Request struct`, t, func() {
		request := &Request{}
		if err := json.Unmarshal(sessionEndedRequestJSON, request); err != nil {
			panic(err)
		}

		Convey(`Then the Detail type will be a SessionEndedRequest struct`, func() {
			So(request.Detail, ShouldHaveSameTypeAs, &SessionEndedRequest{})
		})

		Convey(`Then the Detail ID will be set correctly`, func() {
			So(request.Detail.GetID(), ShouldEqual, "amzn1.echo-api.request.65d4c1e0-1013-40fd-9312-9b7fa462e0a9")
		})
	})
}

var launchRequestJSON = []byte(`
{
	"version": "1.0",
	"session": {
		"new": true,
		"sessionId": "7547f227-9a55-4974-94cf-ed6fdd1cc631",
		"application": {
			"applicationId": "a3a5bb41-e0c9-438b-a8f9-7f2e4de84641"
		},
		"attributes": {
			"4c96a4a6-0082-463c-90dd-3108c8ae787f": "d69e7caa-c029-4187-83e8-bdb2987baad7"
		},
		"user": {
			"userId": "32549228-5f2d-4749-98c2-fe3a394366ad",
			"accessToken": "67b433cd-1a8a-4a53-ae66-402ca5a03fe1",
			"permissions": {
				"consentToken": "c6cc6b93-3f7b-4b88-ad47-97186886cd72"
			}
		}
	},
	"context": {
		"System": {
			"application": {
				"applicationId": "a3a5bb41-e0c9-438b-a8f9-7f2e4de84641"
			},
			"user": {
				"userId": "32549228-5f2d-4749-98c2-fe3a394366ad",
				"accessToken": "67b433cd-1a8a-4a53-ae66-402ca5a03fe1",
				"permissions": {
					"consentToken": "c6cc6b93-3f7b-4b88-ad47-97186886cd72"
				}
			},
			"device": {
				"deviceId": "05cdef34-7fbf-4be5-9051-09125301f935",
				"supportedInterfaces": {
					"AudioPlayer": {}
				}
			},
			"apiEndpoint": "26909ef5-2c76-4929-9309-90352a425ef4"
		},
		"AudioPlayer": {
			"token": "b505bf6a-1c80-4430-b9b9-581b0536ca41",
			"offsetInMilliseconds": 0,
			"playerActivity": "string"
		}
	},
	"request": {
		"type": "LaunchRequest",
		"requestId": "58eeabff-a19a-4fd1-87f1-7c45789769ad"
	}
}
`)

var intentRequestJSON = []byte(`
{
	"version": "1.0",
	"request": {
		"type": "IntentRequest",
		"requestId": "c4763de6-1e7d-4842-99d5-29fe1836b577",
		"timestamp": "2017-08-01T15:03:44Z",
		"dialogState": "STARTED",
		"locale": "en-GB",
		"intent": {
			"name": "IntentName"
		}
	}
}
`)

var sessionEndedRequestJSON = []byte(`
{
	"version": "1.0",
	"request": {
		"type": "SessionEndedRequest",
		"requestId": "amzn1.echo-api.request.65d4c1e0-1013-40fd-9312-9b7fa462e0a9",
		"timestamp": "2017-08-29T20:48:16Z",
		"locale": "en-GB",
		"reason": "EXCEEDED_MAX_REPROMPTS"
	}
}
`)
