package etcd

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/coreos/etcd/client"

	"github.com/kryptn/confg/containers"
)

type EtcdSource struct {
	kapi client.KeysAPI
}

func (es *EtcdSource) Lookup(lookup string) (interface{}, bool) {
	resp, err := es.kapi.Get(context.Background(), lookup, nil)
	if err != nil {
		log.Printf("Error on etcd get %v", err)
		return nil, false
	}
	log.Printf("etcd %+v", resp.Node)
	return resp.Node.Value, true
}

func (es *EtcdSource) Gather(keys []*containers.Key) {
	for _, key := range keys {
		v, ok := es.Lookup(key.Lookup)
		key.Inject(v, ok)
	}

}

func Get(backend *containers.Backend) (*EtcdSource, error) {
	if backend.Source != "etcd" {
		return nil, errors.New("source.etcd invalid backend")
	}
	es := EtcdSource{}

	cfg := client.Config{
		Endpoints:               backend.Hosts,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	es.kapi = client.NewKeysAPI(etcdClient)

	return &es, nil
}
