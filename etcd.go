package etcd

import (
	"github.com/mchirico/go.etcd/pkg/etcdutils"
	"log"

	"fmt"
	"time"
)

func D(subkey string) {

	e, cancel := etcdutils.NewETC("server")
	defer cancel()

	now := time.Now()
	msg := fmt.Sprintf("{email: %q, timestamp: %q}", subkey, now)
	key := fmt.Sprintf("aibot-gmail/%s", subkey)
	key2 := fmt.Sprintf("aibot-gmail/last/%s", subkey)

	e.PutWithLease(key, msg, 300*600)
	e.PutWithLease(key2, msg, 300*600)

}

func Query() {

	e, cancel := etcdutils.NewETC("server")
	defer cancel()


	result, _ := e.GetWithPrefix("aibot-gmail")

	for i, v := range result.Kvs {
		log.Printf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}

}

func Delete() {

	e, cancel := etcdutils.NewETC("server")
	defer cancel()

	e.DeleteWithPrefix("aibot-gmail")

	result, _ := e.GetWithPrefix("aibot-gmail")

	for i, v := range result.Kvs {
		log.Printf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}

}
