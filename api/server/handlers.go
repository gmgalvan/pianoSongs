package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Home --
func (s *Svr) Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome"))
}

// AddNewSong --
func (s *Svr) AddNewSong(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "could not read body", http.StatusInternalServerError)
		}
		if req.Body != nil {
			defer req.Body.Close()
		}
		var song PianoSong
		err = json.Unmarshal(body, &song)
		if err != nil {
			http.Error(w, "could not unmarshall", http.StatusBadRequest)
		}
		err = s.Store.create(&song)
		if err != nil {
			http.Error(w, "could not store song", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Success in create song"))
	} else {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}

func validMethodToProcess(method string) bool {
	if method == http.MethodGet {
		return true
	}
	if method == http.MethodPut {
		return true
	}
	if method == http.MethodDelete {
		return true
	}
	return false
}

// ProcessSong --
func (s *Svr) ProcessSong(w http.ResponseWriter, req *http.Request) {
	id := strings.Replace(req.URL.Path, "/song/", "", 1)
	if !validMethodToProcess(req.Method) {
		http.Error(w, "method not Allowed", http.StatusMethodNotAllowed)
	}
	if req.Method == http.MethodGet {
		pS, err := s.getSong(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		resp, err := json.Marshal(pS)
		if err != nil {
			http.Error(w, "could not marshall song", http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		w.Header().Set("Content/Type", "application/json")
	}
	if req.Method == http.MethodPut {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "could not read body", http.StatusInternalServerError)
		}
		if req.Body != nil {
			defer req.Body.Close()
		}
		var song PianoSong
		err = json.Unmarshal(body, &song)
		if err != nil {
			http.Error(w, "could not unmarshall", http.StatusBadRequest)
		}
		err = s.updateSong(id, &song)
		if err != nil {
			http.Error(w, "could not update song", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Success in update song"))
	}
	if req.Method == http.MethodDelete {
		err := s.deleteSong(id)
		if err != nil {
			http.Error(w, "could not delete song", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Success in delete song"))
	}
}

// GetSong --
func (s *Svr) getSong(id string) (*PianoSong, error) {
	pS, err := s.getOne(id)
	if err != nil {
		return nil, err
	}
	return pS, nil
}

// GetAllSongs --
func (s *Svr) GetAllSongs(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		songs, err := s.getAll()
		if err != nil {
			http.Error(w, "could not get songs", http.StatusInternalServerError)
		}
		resp, err := json.Marshal(songs)
		if err != nil {
			http.Error(w, "could not marshall songs", http.StatusInternalServerError)
		}
		w.Write([]byte(resp))
		w.Header().Set("Content/Type", "application/json")
	} else {
		http.Error(w, "get only", http.StatusMethodNotAllowed)
	}
}

// UpdateSong --
func (s *Svr) updateSong(id string, song *PianoSong) error {
	err := s.update(id, song)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSong --
func (s *Svr) deleteSong(id string) error {
	err := s.delete(id)
	if err != nil {
		return err
	}
	return nil
}
