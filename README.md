# alexa
Types, interfaces and utilities for interacting with the Amazon Alexa Service

## Verification of Alexa Requests
Alexa places a number of demands on a hosted skill which [largely revolve around providing
encryption and verification of the request](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/developing-an-alexa-skill-as-a-web-service).
This requirement is the primary motivation for this package.


### Development
While developing locally in a development environment this is less than ideal. An acceptable method
of by-stepping this requirement is to make use of [ngrok](https://ngrok.com/).

* Start your Go program, taking note of the port that it is listening on
* Start ngrok with the command `ngrok http <Go Port Number>

Take note of the HTTP address and update the Alexa developer console for the skill under test in
the _Configuration_ section, setting the certificate in the _SSL Certificate_ section to "_My
development endpoint is a sub-domain of a domain that has a wildcard certificate from a
certificate authority_"

```
ngrok by @inconshreveable                                                                     (Ctrl+C to quit)

Session Status                online
Version                       2.2.4
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://e22449b2.ngrok.io -> localhost:9000
Forwarding                    https://e22449b2.ngrok.io -> localhost:9000 <<<<< This Entry

Connections                   ttl     opn     rt1     rt5     p50     p90
```
