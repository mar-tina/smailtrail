# smailtrail

To run the app.
    - Clone the repositiory
    - Go to [Google auth](https://developers.google.com/gmail/api/quickstart/go) and click on Enable Gmail API.
    - Mv The provided credentials.json file to the cmd directory in the root of the repo where the main.go is

There is a backend service and a frontend service. To run the backend

    `cd cmd`
    `go run main.go`

To run the frontend

    `cd web`
    `yarn start`

The main functionality of the app is to fetch Gmail Messages. Parse the body and extract the unsubscribe link from the body text.

Caveats
    - Some links dont direct you immediately to the unsubscribe page
    
