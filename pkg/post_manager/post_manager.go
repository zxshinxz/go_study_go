package post_manager

import (
	"errors"
	om "github.com/the-gigi/delinkcious/pkg/object_model"
)

type PostManager struct {
	store om.PostManager
}

func NewPostManager(store om.PostManager) (om.PostManager, error) {
	if store == nil {
		return nil, errors.New("store can't be nil")
	}
	return &PostManager{store: store}, nil
}

func (m *PostManager) Post(title string, content string) (err error) {
	if title == "" || content == "" {
		err = errors.New("followed and follower can't be empty")
		return
	}

	return m.store.Post(title, content)
}

func (m *PostManager) GetPosts() (string, error) {
	return m.store.GetPosts()
}
