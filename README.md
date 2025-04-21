# tdd_go_web_server
I've been requested to create a web sedrver where users can track how many games players have won.

Having the next API calls:
* `GET /players/{name}` - Indicating the total number of wins
* `POST /players/{name}` - Records or increment a win to that name on every subsequent post call
