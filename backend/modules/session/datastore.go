package session

import (
	"context"
	"fmt"
	"cloud.google.com/go/datastore"
)



func (s *Session)DSCreateSession(project string) error {
	s.GenerateSessionID()
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, project);if err != nil {
		return err
	}
	defer dsClient.Close()

	k := datastore.NameKey("Session", s.ID, nil)

	if _, err := dsClient.Put(ctx, k, s); err != nil {
		fmt.Print(err,1)
		return err
	}
	return nil
}

func DSGetSession(project string, ID string) (Session, error) {
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

func DSDeleteSession(project string, ID string) error {
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