package session

import (
	"context"
	"cloud.google.com/go/datastore"
)


// データストアにセッション作成
func (s *Session)DSCreateSession(project string) error {
	s.GenerateSessionID()
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return err
	}
	defer dsClient.Close()

	k := datastore.NameKey("Session", s.ID, nil)

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
	key := datastore.NameKey("Session", ID, nil)
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
	key := datastore.NameKey("Session", ID, nil)
	err = dsClient.Delete(ctx, key);if err != nil {
		return err
	}
	return nil
}