# Piano Song simple server

Http api for piano songs, simple example using golang and postgres

## setting project locally

- Run postgres database locally in your console
```console
    docker run -p 5432:5432 -e POSTGRES_USER=test_user -e POSTGRES_PASSWORD=yourPassword -e POSTGRES_DB=PianoSongs postgres
```

- Run the app using go
```console
    export APP_DB_USERNAME=test_user
    export APP_DB_PASSWORD=yourPassword
    export APP_DB_NAME=PianoSongs

    cd api

    go mod tidy && go mod vendor

    go test ./...

    go run main.go
```

## Interacting
- create one song:

```console
curl -H "Content-Type: application/json" -X POST -d '{"name":"Habanera Carmen", "autor":"Bizet", "description": ""}' http://localhost:8081/song
```

- get one song by id
```console
curl -X GET http://localhost:8081/song/1
```

- get all songs
```console
curl -X GET http://localhost:8081/songs
```

- get all songs on csv
```console
curl -H "Accept: text/csv" -X GET http://localhost:8081/songs
```

- update one song:
```console
curl -H "Accept: application/json" -X PUT -d '{"name":"Habanerax Carmen", "autor":"Bizet", "description": ""}' http://localhost:8081/song/1
```

- delete one song by id
```console
curl -X DELETE http://localhost:8081/song/1
```