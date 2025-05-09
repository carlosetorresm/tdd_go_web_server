# tdd_go_web_server
I've been requested to create a web server where users can track how many games players have won.
All the data will be saved as a json format file in the root of the project.

To execute the project we can run the next command:
  `go run ./cmd/webserver/main.go`

Having the next API calls:
* `GET /players/{name}` - Indicating the total number of wins
* `POST /players/{name}` - Records or increment a win to that name on every subsequent post call

The product owner requested a new end point where we return a json with every player with their wins
* `GET /league` - Returns a List of players with their wins

The product owner requested that the players record could also be saved making use of a CLI.

To execute the CLI project we can run the next command:
  `go run ./cmd/cli/main.go`

  The product owner requested that we implemented a versión of Texas Holdem, also to be implemented as a webview in the web server.
* `get /game` - Opens a page in the browser to play a version of Texas Holdem.