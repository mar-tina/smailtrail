# smailtrail

To run the app. - Clone the repositiory - Go to [Google auth](https://developers.google.com/gmail/api/quickstart/go) and click on Enable Gmail API. - Mv The provided credentials.json file to the cmd directory in the root of the repo where the main.go is

There is a backend service and a frontend service. To run the backend

    `cd cmd`
    `go run main.go`

To run the frontend

    - cd web
    - Install the dependencies...
    - yarn | yarn install
    - yarn start

The main functionality of the app is to fetch Gmail Messages. Parse the body and extract the unsubscribe link from the body text.

Caveats

    - Some pages are not text/html formatted. Makes pinpointing the <a> tag with
    the unsubscribe difficult
    - Some services do not provide an explicit "unsubscribe" label for their links.

Things Not Done Yet:

    - Pagination.
    - As at now Fetching data from the API requires you to refresh page to see new content.
    This makes it a game of guesswork to see if you have new content.

How To navigate the app

The initial load of the page takes you to the homepage .
![Homepage Screenshot](/screenshots/home.png)

To be able to load subscriptions navigate to Manage Your Account. If you had not followed the above
process of configuring the credentials.json . The page will load forever. Make sure the credentials.json
file is in the cmd directory next to main.go
![Authpage Screenshot](/screenshots/auth.png)

Navigate back to the subscriptions page. Click the fetch from API button to load messages from Gmail.
The messages are loaded per page. Current limit is set to 50. Then refresh the page after load is complete
to see your links
![Authpage Screenshot](/screenshots/subs.png)

Here is the overall flow
![Overall Flow](/screenshots/smailtrail.gif)

Feel Free to raise issues and bug reports
