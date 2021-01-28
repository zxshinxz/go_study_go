package post_manager

import (
	_ "github.com/lib/pq"
	. "go.etcd.io/etcd/client"
	"golang.org/x/net/context"
	"log"
	"time"
)

type EtcdPostStore struct {
	kapi KeysAPI
}

func NewEtcdPostStore(host string, port int, username string, password string) (store *EtcdPostStore, err error) {
	cfg := Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := NewKeysAPI(c)
	store = &EtcdPostStore{kapi}
	return
}

func (s *EtcdPostStore) Post(title string, content string) (err error) {
	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := s.kapi.Set(context.Background(), "posts", title+":"+content, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}

	return
}

func (s *EtcdPostStore) GetPosts() (posts string, err error) {
	posts = ""
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp, err := s.kapi.Get(context.Background(), "posts", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
		posts = resp.Node.Value
	}

	return
}

