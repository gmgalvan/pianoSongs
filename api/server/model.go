package server

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PianoSong struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Autor       string `json:"autor"`
	Description string `json:"description"`
}

// Store --
type Store struct {
	client *sql.DB
}

const (
	allSongs = `
		SELECT * FROM songs
	`
	createSong = `
		INSERT INTO songs (id, name, autor, description)
		VALUES ($1, $2, $3, $4)
	`
	getSong = `
		SELECT * from songs WHERE id=$1 
	`
	getAllSongs = `
		SELECT * from songs
	`
	updateSong = `
		UPDATE songs SET id=$1, name=$2, autor=$3, description=$4 WHERE id=$1
	`
	deleteSong = `
		DELETE FROM songs WHERE id=$1
	`
)

func (s *Store) create(pnS *PianoSong) error {
	id := uuid.New()
	_, err := s.client.Exec(createSong, id, pnS.Name, pnS.Autor, pnS.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) getOne(id string) (*PianoSong, error) {
	pS := &PianoSong{}
	err := s.client.QueryRow(getSong, id).Scan(&pS.Id, &pS.Autor, &pS.Name, &pS.Description)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("No result found with id %v", id)
			return nil, errors.New(errMsg)
		}
		return nil, err
	}
	return pS, nil
}

func (s *Store) getAll() ([]*PianoSong, error) {
	songs := make([]*PianoSong, 0)
	rows, err := s.client.Query(getAllSongs)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pS := PianoSong{}
		err := rows.Scan(&pS.Id, &pS.Autor, &pS.Name, &pS.Description)
		if err != nil {
			return nil, err
		}
		songs = append(songs, &pS)
	}
	return songs, nil
}

func (s *Store) update(id string, song *PianoSong) error {
	_, err := s.client.Exec(updateSong, id, song.Name, song.Autor, song.Description)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) delete(id string) error {
	_, err := s.client.Exec(deleteSong, id)
	if err != nil {
		return err
	}
	return nil
}
