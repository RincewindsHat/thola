package database

import (
	"context"
	"encoding/json"
	"github.com/dgraph-io/badger/v2"
	"github.com/RincewindsHat/thola/internal/device"
	"github.com/RincewindsHat/thola/internal/network"
	"github.com/RincewindsHat/thola/internal/parser"
	"github.com/RincewindsHat/thola/internal/tholaerr"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type badgerDatabase struct {
	db *badger.DB
}

func (d *badgerDatabase) SetDeviceProperties(_ context.Context, ip string, data device.Device) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	JSONData, err := parser.ToJSON(data)
	if err != nil {
		return errors.Wrap(err, "failed to marshall response")
	}
	entry := badger.Entry{
		Key:       []byte("DeviceInfo-" + ip),
		Value:     JSONData,
		ExpiresAt: uint64(time.Now().Add(cacheExpiration).Unix()),
	}

	err = txn.SetEntry(&entry)
	if err != nil {
		return errors.Wrap(err, "failed to store identify data")
	}

	err = txn.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to store identify data")
	}
	return nil
}

func (d *badgerDatabase) GetDeviceProperties(_ context.Context, ip string) (device.Device, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte("DeviceInfo-" + ip))
	if err != nil {
		return device.Device{}, tholaerr.NewNotFoundError("cannot find cache entry")
	}

	value, err := item.ValueCopy(nil)
	if err != nil {
		return device.Device{}, errors.Wrap(err, "failed to get value from db item")
	}

	data := device.Device{}
	err = json.Unmarshal(value, &data)
	if err != nil {
		return device.Device{}, errors.Wrap(err, "failed to unmarshall device properties")
	}
	return data, nil
}

func (d *badgerDatabase) SetConnectionData(_ context.Context, ip string, data network.ConnectionData) error {
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	JSONData, err := parser.ToJSON(data)
	if err != nil {
		return errors.Wrap(err, "failed to marshall connection data")
	}
	entry := badger.Entry{
		Key:       []byte("ConnectionData-" + ip),
		Value:     JSONData,
		ExpiresAt: uint64(time.Now().Add(cacheExpiration).Unix()),
	}

	err = txn.SetEntry(&entry)
	if err != nil {
		return errors.Wrap(err, "failed to store connection data")
	}

	err = txn.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to store connection data")
	}
	return nil
}

func (d *badgerDatabase) GetConnectionData(_ context.Context, ip string) (network.ConnectionData, error) {
	txn := d.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte("ConnectionData-" + ip))
	if err != nil {
		return network.ConnectionData{}, tholaerr.NewNotFoundError("cannot find cache entry")
	}

	value, err := item.ValueCopy(nil)
	if err != nil {
		return network.ConnectionData{}, errors.Wrap(err, "failed to get value from db item")
	}

	data := network.ConnectionData{}
	err = json.Unmarshal(value, &data)
	if err != nil {
		return network.ConnectionData{}, errors.Wrap(err, "failed to unmarshall connectionData")
	}
	return data, nil
}

func (d *badgerDatabase) CheckConnection(_ context.Context) error {
	if d.db.IsClosed() {
		return errors.New("badger db is closed")
	}
	return nil
}

func (d *badgerDatabase) CloseConnection(ctx context.Context) error {
	log.Ctx(ctx).Debug().Msg("closing connection to built-in database")
	return d.db.Close()
}
