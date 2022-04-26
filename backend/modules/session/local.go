package session

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
)



func (s *Session)CreateSession(project string) error {
	if project != "" {
		err := s.DSCreateSession(project)
		return err
	}
	s.GenerateSessionID()
	// session用ファイル作成
	// sessionフォルダの有無判定
	_, err := os.Stat("session")
	if err != nil {
		os.Mkdir("session", 0777)
	}
	filepath := fmt.Sprintf("./session/%v.txt", s.ID)
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)

	if err := enc.Encode(s); err != nil {
		return err
	}
	return nil
}

func GetSession(ID string, project string) (Session, error) {
	if project != "" {
		s, err := DSGetSession(ID, project);if err != nil {
			return s, err
		}
	}
	var s Session
	filepath := fmt.Sprintf("./session/%v.txt", ID)
	err := s.ReadSession(filepath);if err != nil {
		return s, err
	}
	return s, nil
}

func DeleteSession(ID string, project string) (Session, error) {
	if project != "" {
		s, err := DSDeleteSession(ID, project);if err != nil {
			return s, err
		}
	}
	var s Session
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return s, err
	}
	defer dsClient.Close()
	key := datastore.NameKey("Session", s.ID, nil)
	err = dsClient.Delete(ctx, key);if err != nil {
		return s, err
	}
	return s,nil
}

func (s *Session) ReadSession(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(f)
	defer f.Close()
	if err := dec.Decode(&s); err != nil {
		return err
	}
	return nil
}

func CheckSession(ID string, project string) (Session, error) {
	var s Session
	if ID == "" {
		err := errors.New("no cookie")
		return s, err
	}
	if project != "" {
		s, err := DSGetSession(project, ID); if err != nil {
			return s, err
		}
		return s, nil
	}
	filepath := fmt.Sprintf("./session/%v.txt", ID)
	err := s.ReadSession(filepath);if err != nil {
		return s, err
	}
	return s, nil
}