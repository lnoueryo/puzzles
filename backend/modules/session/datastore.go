package session

import (
	"context"
	"os"

	"cloud.google.com/go/datastore"
)

var sessionKey string

func init() {
	sessionKey = os.Getenv("SESSION_KEY"); if sessionKey == "" {
		sessionKey = "Session"
	}
}

// データストアにセッション作成
func (s *Session)DSCreateSession(project string) error {
	s.GenerateSessionID()
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return err
	}
	defer dsClient.Close()

	k := datastore.NameKey(sessionKey, s.ID, nil)

	if _, err := dsClient.Put(ctx, k, s); err != nil {
		return err
	}
	return nil
}

// データストアのセッション取得
func DSGetSession(ID string, project string) (Session, error) {
	var s Session
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return s, err
	}
	defer dsClient.Close()
	key := datastore.NameKey(sessionKey, ID, nil)
	if err := dsClient.Get(ctx, key, &s); err != nil {
		return s, err
	}

	return s, nil
}

// データストアのセッション削除
func DSDeleteSession(ID string, project string) error {
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return err
	}
	defer dsClient.Close()
	key := datastore.NameKey(sessionKey, ID, nil)
	err = dsClient.Delete(ctx, key);if err != nil {
		return err
	}
	return nil
}