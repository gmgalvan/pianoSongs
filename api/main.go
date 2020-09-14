package main

import (
	"lab/pianoSong/api/server"
	"log"
	"net/http"
)

func main() {
	// Init database
	db, err := server.StartDB()
	if err != nil {
		log.Fatalf("Error in initializing Postgresql: [%v]", err)
	}

	// Start server config
	mux := http.NewServeMux()
	cfgServer := server.ConfigServer{
		Addr: ":8080",
		Mux:  mux,
	}

	// create server and initilize schema or migrations
	svr := server.NewServer(db, cfgServer)
	err = svr.InitDBSchema()
	if err != nil {
		log.Fatalf("Error in initializing schema in Postgresql: [%v]", err)
	}

	// handling routes
	mux.HandleFunc("/", svr.Home)
	mux.HandleFunc("/song", svr.AddNewSong)
	mux.HandleFunc("/song/", svr.ProcessSong)
	mux.HandleFunc("/songs", svr.GetAllSongs)

	log.Fatal(svr.HTTPSrv.ListenAndServe())
}
