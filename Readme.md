# Piano Song simple server

Http api for piano songs, simple example using golang and postgres

## Setting project locally

- Dependencies:
    1. Docker
    2. Docker Compose (optional)
    3. Golang
    4. Postgres (optional)
    5. kubernetes (WIP)

- Tested on Fedora 32

### Option 1
- Run postgres database locally with docker or [install](https://www.postgresql.org/docs/9.3/tutorial-install.html) postgres locally
```console
    docker run -d -p 5432:5432 -e POSTGRES_USER=test_user -e POSTGRES_PASSWORD=yourPassword -e POSTGRES_DB=PianoSongs postgres
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

### Option 2
- Run the app with docker
```console
    cd api

    docker build -t gmgalvan/api-piano-songs . --build-arg app_db_username=test_user --build-arg app_db_password=yourPassword --build-arg app_db_name=PianoSongs
    
    docker run -d --net=host gmgalvan/api-piano-songs
```

- Run postgres database locally with docker
```console
    docker run -d -p 5432:5432 -e POSTGRES_USER=test_user -e POSTGRES_PASSWORD=yourPassword -e POSTGRES_DB=PianoSongs postgres
```

### Option 3
- Run with docker compose

Create a .env file for app credentials
```console
echo -e "APP_DB_USERNAME=test_user\nAPP_DB_PASSWORD=yourPassword\nAPP_DB_NAME=PianoSongs" > .env
```

Create env_db for postgres credentials
```console
echo -e "POSTGRES_USER=test_user\nPOSTGRES_PASSWORD=yourPassword\nPOSTGRES_DB=PianoSongs" > .env_db
```

Run with docker compose cli
```console
    docker-compose up
```

## Interacting
- create one song:

```console
curl -H "Content-Type: application/json" -X POST -d '{"name":"Habanera Carmen", "autor":"Bizet", "description": ""}' http://localhost:8080/song
```

- get one song by id
```console
curl -X GET http://localhost:8080/song/1
```

- get all songs
```console
curl -X GET http://localhost:8080/songs
```

- get all songs on csv
```console
curl -H "Accept: text/csv" -X GET http://localhost:8080/songs
```

- update one song:
```console
curl -H "Accept: application/json" -X PUT -d '{"name":"Habanerax Carmen", "autor":"Bizet", "description": ""}' http://localhost:8080/song/1
```

- delete one song by id
```console
curl -X DELETE http://localhost:8080/song/1
```

## Deploy with K8s (WIP)
- Build docker image

```console
    cd api
    docker build -t gmgalvan/api-piano-songs-v0.0.1 .
    docker run -d -p 8080:8080 gmgalvan/api-piano-songs
```