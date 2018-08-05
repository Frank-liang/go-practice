package ns

import (
	"path"

	etcd "github.com/coreos/etcd/client"
)

func NewEtcdRegistry(kAPI etcd.KeysAPI, keyPrefix string) *EtcdRegistry {
	return &EtcdRegistry{
		kAPI:      kAPI,
		keyPrefix: keyPrefix,
	}
}

// EtcdRegistry fulfils the Registry interface and uses etcd as a backend
type EtcdRegistry struct {
	kAPI      etcd.KeysAPI
	keyPrefix string
}

func (r *EtcdRegistry) prefixed(p ...string) string {
	return path.Join(r.keyPrefix, path.Join(p...))
}

func (r *EtcdRegistry) suffixed(p ...string) string {
	return path.Join(p...)
}

func isEtcdError(err error, code int) bool {
	eerr, ok := err.(etcd.Error)
	return ok && eerr.Code == code
}
