package alexa_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/alexa"
	"time"
)

func TestRequestUnmarshal(t *testing.T) {
	Convey(`Given I have a JSON request object`, t, func() {
		jsonRequest := []byte(`{
					"version": "1.0",
  					"session": {
    					"sessionId": "9077cdd6-49c7-4114-9176-6b99ac320962"
  					},
  					"context": {
    					"System": {
      						"application": {
        						"applicationId": "2f28fc87-e422-41d4-a003-b3f990fc5f8e"
      						}
      					}
  					},
  					"request": {}
				}`)

		Convey(`When I unmarshal the request into a Request struct`, func() {
			req := &alexa.Request{}
			if err := json.Unmarshal(jsonRequest, req); err != nil {
				panic(err)
			}

			Convey(`Then the Version field will be set correctly`, func() {
				So(req.Version, ShouldEqual, "1.0")
			})

			Convey(`Then the Session ID field will be set correctly`, func() {
				So(req.Session.ID, ShouldEqual, "9077cdd6-49c7-4114-9176-6b99ac320962")
			})

			Convey(`Then the Context System Application ID will be set correctly`, func() {
				So(req.Context.System.Application.ID, ShouldEqual, "2f28fc87-e422-41d4-a003-b3f990fc5f8e")
			})

		})
	})
}

func TestSessionUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON session object`, t, func() {
		jsonSession := []byte(`
			{
    			"new": true,
    			"sessionId": "9077cdd6-49c7-4114-9176-6b99ac320962",
    			"application": {
      				"applicationId": "8f2d5b26-9abe-45f2-a2d4-d80f70133461"
    			},
    			"attributes": {
      				"name": "Test Name"
    			},
    			"user": {
      				"userId": "string",
        			"permissions": {
          				"consentToken": "string"
      					},
      				"accessToken": "string"
    			}
  			}
  			`)

		Convey(`When I unmarshal the session into a Session struct`, func() {
			session := &alexa.Session{}
			if err := json.Unmarshal(jsonSession, session); err != nil {
				panic(err)
			}

			Convey(`Then the New field will be set correctly`, func() {
				So(session.New, ShouldBeTrue)
			})

			Convey(`Then the ID field will be set correctly`, func() {
				So(session.ID, ShouldEqual, "9077cdd6-49c7-4114-9176-6b99ac320962")
			})

			Convey(`Then the "name" Attribute field will be set correctly`, func() {
				So(session.Attributes[`name`].(string), ShouldEqual, "Test Name")
			})

			Convey(`Then the Application ID field will be set correctly`, func() {
				So(session.Application.ID, ShouldEqual, "8f2d5b26-9abe-45f2-a2d4-d80f70133461")
			})
		})
	})
}

func TestApplicationUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON application object`, t, func() {
		jsonApp := []byte(`
			{
				"applicationID": "8f2d5b26-9abe-45f2-a2d4-d80f70133461"
			}`)

		Convey(`When I unmarshal the application into an Application struct`, func() {
			app := &alexa.Application{}
			if err := json.Unmarshal(jsonApp, app); err != nil {
				panic(err)
			}

			Convey(`Then the ID will be set correctly`, func() {
				So(app.ID, ShouldEqual, "8f2d5b26-9abe-45f2-a2d4-d80f70133461")
			})
		})
	})
}

func TestUserUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON user object`, t, func() {
		jsonUser := []byte(`
			{
      			"userId": "25fff62e-4455-4570-9b90-fc46c6fbeebc",
        		"permissions": {
          			"consentToken": "e3b28067-44e6-44fa-9e2d-5f9272bf5a84"
      				},
      			"accessToken": "d8c7cee5-f6b0-4d9a-9b7b-64cfecb1852d"
    		}
		`)

		Convey(`When I unmarshal the user into a User struct`, func() {
			user := &alexa.User{}
			if err := json.Unmarshal(jsonUser, user); err != nil {
				panic(err)
			}

			Convey(`Then the ID field will be set correctly`, func() {
				So(user.ID, ShouldEqual, "25fff62e-4455-4570-9b90-fc46c6fbeebc")
			})

			Convey(`Then the AccessToken field will be set correctly`, func() {
				So(user.AccessToken, ShouldEqual, "d8c7cee5-f6b0-4d9a-9b7b-64cfecb1852d")
			})

			Convey(`Then the Permissions ConsentToken field will be set correctly`, func() {
				So(user.Permissions.ConsentToken, ShouldEqual, "e3b28067-44e6-44fa-9e2d-5f9272bf5a84")
			})
		})
	})
}

func TestPermissionUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON permissions object`, t, func() {
		jsonPermission := []byte(`
			{
				"consentToken": "e3b28067-44e6-44fa-9e2d-5f9272bf5a84"
			}`)

		Convey(`When I unmarshal the permission into an Permission struct`, func() {
			permission := &alexa.Permission{}
			if err := json.Unmarshal(jsonPermission, permission); err != nil {
				panic(err)
			}

			Convey(`Then the ConsentToken will be set correctly`, func() {
				So(permission.ConsentToken, ShouldEqual, "e3b28067-44e6-44fa-9e2d-5f9272bf5a84")
			})
		})
	})
}

func TestSystemUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON System object`, t, func() {
		jsonSystem := []byte(`
			{
				"application": {
					"applicationId": "8ca0fb84-7c63-4133-9216-f25345303c34"
				},
				"user": {
					"userId": "54033e0f-4a8c-4d49-b6d9-4326def9032c"
				},
				"device": {
					"deviceId": "0b71fe12-f499-45eb-a4e8-76d330dd7098"
				},
				"apiEndpoint": "552d8213-d22c-475e-a7eb-072dd32729fc"
			}
		`)

		Convey(`When I unmarshal the system into a System struct`, func() {
			system := &alexa.System{}
			if err := json.Unmarshal(jsonSystem, system); err != nil {
				panic(err)
			}

			Convey(`Then the Application ID will be set correctly`, func() {
				So(system.Application.ID, ShouldEqual, "8ca0fb84-7c63-4133-9216-f25345303c34")
			})

			Convey(`Then the User ID will be set correctly`, func() {
				So(system.User.ID, ShouldEqual, "54033e0f-4a8c-4d49-b6d9-4326def9032c")
			})

			Convey(`Then the Device ID will be set correctly`, func() {
				So(system.Device.ID, ShouldEqual, "0b71fe12-f499-45eb-a4e8-76d330dd7098")
			})

			Convey(`Then the APIEndpoint field will be set correctly`, func() {
				So(system.APIEndpoint, ShouldEqual, "552d8213-d22c-475e-a7eb-072dd32729fc")
			})
		})
	})
}

func TestDeviceUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON permissions object`, t, func() {
		jsonDevice := []byte(`
			{
				"deviceId": "0b71fe12-f499-45eb-a4e8-76d330dd7098",
				"supportedInterfaces": {
						"AudioPlayer": {}
				}
			}`)

		Convey(`When I unmarshal the permission into an Device struct`, func() {
			device := &alexa.Device{}
			if err := json.Unmarshal(jsonDevice, device); err != nil {
				panic(err)
			}

			Convey(`Then the ID field will be set correctly`, func() {
				So(device.ID, ShouldEqual, "0b71fe12-f499-45eb-a4e8-76d330dd7098")
			})

			Convey(`Then the HasAudioPlayerSupport method should return true`, func() {
				So(device.HasAudioPlayerSupport(), ShouldBeTrue)
			})
		})
	})
}

func TestAudioPlayerUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON AudioPlayer object`, t, func() {
		jsonAudioPlayer := []byte(`
			 {
      			"token": "fdfa1d60-a826-4eda-bea1-6da272acfa63",
      			"offsetInMilliseconds": 1234567890,
      			"playerActivity": "PAUSED"
    		 }
		`)

		Convey(`When I unmarshal the audio player into an AudioPlayer struct`, func() {
			audioPlayer := &alexa.AudioPlayer{}
			if err := json.Unmarshal(jsonAudioPlayer, audioPlayer); err != nil {
				panic(nil)
			}

			Convey(`Then the Token field will be set correctly`, func() {
				So(audioPlayer.Token, ShouldEqual, "fdfa1d60-a826-4eda-bea1-6da272acfa63")
			})

			Convey(`Then the OffsetInMilliseconds field will be set correctly`, func() {
				So(audioPlayer.OffsetInMilliseconds, ShouldEqual, 1234567890)
			})

			Convey(`Then the PlayerActivity fields will be set correctly`, func() {
				So(audioPlayer.PlayerActivity, ShouldEqual, alexa.AudioPlayerPausedState)
			})
		})
	})
}

func TestLaunchRequestUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON launch request`, t, func() {
		jsonRequest := []byte(`
			{
			  "type": "LaunchRequest",
			  "requestId": "a4101467-5c17-438b-a3ec-a35d901214ff",
			  "timestamp": "2017-08-01T15:03:44Z",
			  "locale": "en-GB"
			}
		`)

		Convey(`When I unmarshal the request into a LaunchRequest struct`, func() {
			req := &alexa.LaunchRequest{}
			if err := json.Unmarshal(jsonRequest, req); err != nil {
				panic(err)
			}

			Convey(`Then the ID field will be set correctly`, func() {
				So(req.ID, ShouldEqual, "a4101467-5c17-438b-a3ec-a35d901214ff")
			})

			Convey(`Then the timestamp field will be set correctly`, func() {
				So(req.Timestamp, ShouldResemble, time.Date(2017, 8, 01, 15, 3, 44, 0, time.UTC))
			})

			Convey(`Then the Locale field will be set correctly`, func() {
				So(req.Locale, ShouldEqual, "en-GB")
			})
		})
	})
}

func TestIntentRequestUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON intent request`, t, func() {
		jsonRequest := []byte(`
			{
				"type": "IntentRequest",
				"requestId": "c4763de6-1e7d-4842-99d5-29fe1836b577",
				"timestamp": "2017-08-01T15:03:44Z",
				"dialogState": "STARTED",
				"locale": "en-GB",
				"intent": {
					"name": "IntentName"
			  	}
			}
		`)

		Convey(`When I unmarshal the request into a IntentRequest struct`, func() {
			req := &alexa.IntentRequest{}
			if err := json.Unmarshal(jsonRequest, req); err != nil {
				panic(err)
			}

			Convey(`Then the ID field will be set correctly`, func() {
				So(req.ID, ShouldEqual, "c4763de6-1e7d-4842-99d5-29fe1836b577")
			})

			Convey(`Then the DialogState field will be set correctly`, func() {
				So(req.DialogState, ShouldEqual, alexa.DialogStateStarted)
			})

			Convey(`Then the Intent Name field will be set correctly`, func() {
				So(req.Intent.Name, ShouldEqual, "IntentName")
			})
		})
	})
}

func TestIntentUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON intent`, t, func() {
		jsonIntent := []byte(`
			{
				"name": "IntentName",
				"confirmationStatus": "CONFIRMED",
				"slots": {
					"SlotName": {
						"name": "ThisIsTheSlotName"
					}
				}
			}
		`)

		Convey(`When I unmarshal the request into an Intent struct`, func() {
			intent := &alexa.Intent{}
			if err := json.Unmarshal(jsonIntent, intent); err != nil {
				panic(err)
			}

			Convey(`Then the Name field will be correctly set`, func() {
				So(intent.Name, ShouldEqual, "IntentName")
			})

			Convey(`Then the ConfirmationStatus field will be correctly set`, func() {
				So(intent.ConfirmationStatus, ShouldEqual, alexa.ConfirmationStatusStateConfirmed)
			})

			Convey(`Then name of the slot in SlotName will be correctly set`, func() {
				So(intent.Slots["SlotName"].Name, ShouldEqual, "ThisIsTheSlotName")
			})
		})
	})
}

func TestSlotUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON slot`, t, func() {
		jsonSlot := []byte(`
			{
				"name": "SlotName",
				"value": "String value",
				"confirmationStatus": "CONFIRMED",
				"resolutions": {
					"resolutionsPerAuthority": [
						{
							"authority": "string",
							"status": {
								"code": "string"
							},
							"values": [
								{
									"value": {
										"name": "string",
										"id": "string"
									}
								}
							]
						}
					]
				}
			}
		`)

		Convey(`When I unmarshal the request into an Intent struct`, func() {
			slot := &alexa.Slot{}
			if err := json.Unmarshal(jsonSlot, slot); err != nil {
				panic(err)
			}

			Convey(`Then the Name field will be correctly set`, func() {
				So(slot.Name, ShouldEqual, "SlotName")
			})

			Convey(`Then the Value field will be correctly set`, func() {
				So(slot.Value, ShouldEqual, "String value")
			})

			Convey(`Then the ConfirmationStatus field will be correctly set`, func() {
				So(slot.ConfirmationStatus, ShouldEqual, alexa.ConfirmationStatusStateConfirmed)
			})

		})
	})
}

func TestResolutionUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON resolution`, t, func() {
		jsonResolution := []byte(`
			{
				"resolutionsPerAuthority": [
					{
						"authority": "cf7beac0-4b44-4d4d-9031-228b47556173"
					}
				]
			}
		`)

		Convey(`When I unmarshal the request into an Resolution struct`, func() {
			resolution := &alexa.Resolution{}
			if err := json.Unmarshal(jsonResolution, resolution); err != nil {
				panic(err)
			}

			Convey(`Then the Name field of the first Authority will be set correctly`, func() {
				So(resolution.Authorities[0].Name, ShouldEqual, "cf7beac0-4b44-4d4d-9031-228b47556173")
			})
		})
	})
}

func TestResolutionAuthorityUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON resolution authority`, t, func() {
		jsonResolution := []byte(`
			{
				"authority": "dd1ec03b-70cb-4628-9489-0c4c2efc1361",
				"status": {
					"code": "ER_SUCCESS_MATCH"
				},
				"values": [
					{
						"value": {
							"name": "cd21ea5f-2b81-41a8-8263-e5e7e51ad6b8"
						}
					}
				]
			}
		`)

		Convey(`When I unmarshal the request into an ResolutionAuthority struct`, func() {
			resolution := &alexa.ResolutionAuthority{}
			if err := json.Unmarshal(jsonResolution, resolution); err != nil {
				panic(err)
			}

			Convey(`Then the Name field will be correctly set`, func() {
				So(resolution.Name, ShouldEqual, "dd1ec03b-70cb-4628-9489-0c4c2efc1361")
			})

			Convey(`Then the Status Code field will be correctly set`, func() {
				So(resolution.Status.Code, ShouldEqual, alexa.ResolutionStatusCodeMatch)
			})

			Convey(`Then the first Values Name will be correctly set`, func() {
				So(resolution.Values[0].Value.Name, ShouldEqual, "cd21ea5f-2b81-41a8-8263-e5e7e51ad6b8")
			})
		})
	})
}

func TestResolutionStatusUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON resolution status`, t, func() {
		jsonStatusCode := []byte(`
			{
				"code": "ER_SUCCESS_MATCH"
			}
		`)

		Convey(`When I unmarshal the status code into an ResolutionStatusCode struct`, func() {
			status := &alexa.ResolutionStatus{}
			if err := json.Unmarshal(jsonStatusCode, status); err != nil {
				panic(err)
			}

			Convey(`Then the Code field will be correctly set`, func() {
				So(status.Code, ShouldEqual, alexa.ResolutionStatusCodeMatch)
			})

		})
	})
}

func TestResolutionValueUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON resolution value`, t, func() {
		jsonValue := []byte(`
			{
				"value": {
					"name": "c72a04e0-3fbb-4234-a418-dcb9650bd0aa"
				}
			}
		`)

		Convey(`When I unmarshal the request into an ResolutionValue struct`, func() {
			value := &alexa.ResolutionValue{}
			if err := json.Unmarshal(jsonValue, value); err != nil {
				panic(err)
			}

			Convey(`Then the Value Name field will be correctly set`, func() {
				So(value.Value.Name, ShouldEqual, "c72a04e0-3fbb-4234-a418-dcb9650bd0aa")
			})

		})
	})
}

func TestResolutionValueDetailUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON resolution value detail`, t, func() {
		jsonValue := []byte(`
			{
				"name": "c72a04e0-3fbb-4234-a418-dcb9650bd0aa",
				"id": "9510a7ce-63c5-4d90-9998-c302fdfe5839"
			}
		`)

		Convey(`When I unmarshal the request into an ResolutionValueDetail struct`, func() {
			value := &alexa.ResolutionValueDetail{}
			if err := json.Unmarshal(jsonValue, value); err != nil {
				panic(err)
			}

			Convey(`Then the Name field will be correctly set`, func() {
				So(value.Name, ShouldEqual, "c72a04e0-3fbb-4234-a418-dcb9650bd0aa")
			})

			Convey(`The the ID field will be correctly set`, func() {
				So(value.ID, ShouldEqual, "9510a7ce-63c5-4d90-9998-c302fdfe5839")
			})
		})
	})
}

func TestSessionEndedRequestUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON session ended request`, t, func() {
		jsonRequest := []byte(`
			{
				"type": "SessionEndedRequest",
				"requestId": "e53a801d-033b-4d91-936e-4baff168b45b",
				"timestamp": "2017-08-01T15:03:44Z",
				"reason": "EXCEEDED_MAX_REPROMPTS",
				"error": {
					"type": "string",
					"message": "string"
			  	}
			}

		`)

		Convey(`When I unmarshal the request into a SessionEndedRequest struct`, func() {
			req := &alexa.SessionEndedRequest{}
			if err := json.Unmarshal(jsonRequest, req); err != nil {
				panic(err)
			}

			Convey(`Then the ID field will be set correctly`, func() {
				So(req.ID, ShouldEqual, "e53a801d-033b-4d91-936e-4baff168b45b")
			})

			Convey(`Then the Reason field will be set correctly`, func() {
				So(req.Reason, ShouldEqual, alexa.SessionEndedReasonUserExceededMaxReprompts)
			})
		})
	})
}

func TestSessionErrorUnmarshalling(t *testing.T) {
	Convey(`Given I have a JSON session error`, t, func() {
		jsonError := []byte(`
			{
				"type": "INVALID_RESPONSE",
				"message": "835fa683-b03f-4e1d-a27f-b4c5a8125986"
			}
		`)

		Convey(`When I unmarshal the request into an ResolutionValueDetail struct`, func() {
			error := &alexa.SessionError{}
			if err := json.Unmarshal(jsonError, error); err != nil {
				panic(err)
			}

			Convey(`Then the Type field will be correctly set`, func() {
				So(error.Type, ShouldEqual, alexa.SessionErrorTypeInvalidResponse)
			})

			Convey(`The the Message field will be correctly set`, func() {
				So(error.Message, ShouldEqual, "835fa683-b03f-4e1d-a27f-b4c5a8125986")
			})
		})
	})
}
