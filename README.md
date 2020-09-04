# go-api-mongo

Created with GO, MUX and MongoDB

AddPerson() => Add new person route `POST /person`

GetPeople() => Get people route `GET /people`

GetPerson() => Get person by ID `GET /person/{id}`

DeletePerson() => Delete person by ID `DELETE /person/{id}`

UpdatePerson() => Update person by ID `PUT /person/{id}`

## Deploy app to Heroku

Create a Procfile, add `web: <app-name>`

go mod init <github-repo>

## Run locally

`go build` to build project, compile packages and dependencies

`go run .` to run the project