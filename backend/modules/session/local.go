package session

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)


// ローカルにセッションファイル作成
func (s *Session)CreateSession(project string) error {
	if project != "" {
		err := s.DSCreateSession(project); if err != nil {
			return err
		}
		return nil
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

// ローカルのセッションファイル削除
func DeleteSession(ID string, project string) error {
	if project != "" {
		err := DSDeleteSession(ID, project);if err != nil {
			return err
		}
	}
	filepath := fmt.Sprintf("./session/%v.txt", ID)
	os.Remove(filepath)
	return nil
}

// ローカルのセッションファイル読み込み
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

// セッションファイルの有無確認
func CheckSession(ID string, project string) (Session, error) {
	var s Session
	if ID == "" {
		err := errors.New("no cookie")
		return s, err
	}
	if project != "" {
		s, err := DSGetSession(ID, project); if err != nil {
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