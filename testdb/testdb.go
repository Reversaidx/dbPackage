package testdb

import (
	"errors"
	"fmt"
)

type BackendDatabase interface {
	New() error
	Put([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Delete([]byte) error
	Close()
	Stats() string
	Flush() error
}

func (d *Database) New() error {
	d.items = make(map[string][]byte)
	return nil

}

var ErrKeyNotFound = errors.New("The key is not found")

func (d *Database) Put(bytes []byte, bytes2 []byte) error {
	_, ok := d.items[string(bytes)]
	if !ok {
		d.items[string(bytes)] = bytes2
		d.Stat.inserts++
	} else {
		d.items[string(bytes)] = bytes2
		d.Stat.updates++
	}
	d.items[string(bytes)] = bytes2
	return nil

}

func (d *Database) Get(bytes []byte) ([]byte, error) {
	value, ok := d.items[string(bytes)]
	if !ok {
		d.Stat.miss++
		return nil, ErrKeyNotFound
	} else {
		d.Stat.hit++
		return value, nil
	}
}

func (d *Database) Delete(bytes []byte) error {
	_, ok := d.items[string(bytes)]
	d.Stat.deletes++
	if !ok {
		return ErrKeyNotFound
	} else {
		delete(d.items, string(bytes))

		return nil
	}
}

func (d *Database) Close() {
	d.items = nil
}

func (d *Database) Stats() string {
	return fmt.Sprintf("inserts:%v\nupdates:%v\ndeletes:%v\nmiss:%v\nhit:%v\n", d.Stat.inserts, d.Stat.updates, d.Stat.deletes, d.Stat.miss, d.Stat.hit)
}

func (d *Database) Flush() error {
	d.items = make(map[string][]byte)
	return nil
}

type Database struct {
	items map[string][]byte
	Stat
}

type Stat struct {
	inserts int // no key
	updates int // when key
	deletes int
	miss    int // get to key
	hit     int // get failed
}
