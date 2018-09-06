package etcd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/client"

	"github.com/kryptn/confg/containers"
)

type EtcdSource struct {
	kapi client.KeysAPI
}

func (es *EtcdSource) Lookup(lookup string) (interface{}, error) {
	resp, err := es.kapi.Get(context.Background(), lookup, nil)
	if err != nil {
		errorMsg := fmt.Sprintf("Error on etcd get %v", err)
		return nil, errors.New(errorMsg)
	}
	log.Printf("etcd %+v", resp.Node)
	return resp.Node.Value, nil
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
