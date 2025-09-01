package store

import "github.com/hashicorp/raft"

var prefixRedisAddr = []byte("___redisAddr")

func GetRedisAddrByNodeID(store raft.StableStore, lid raft.ServerID) (string, error) {
	v, err := store.Get(append(prefixRedisAddr, []byte(lid)...))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func SetRedisAddrByNodeID(store raft.StableStore, lib raft.ServerID, addr string) error {
	return store.Set(append(prefixRedisAddr, []byte(lib)...), []byte(addr))
}
