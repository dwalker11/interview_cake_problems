package dict

import (
	lists "dwalker/lib/list"
)

type UALdict[K comparable, V any] struct {
	list lists.AList[*KVPair[K, V]]
}

func (d UALdict[K, V]) Find(key K) (V, bool) {
	for d.list.MoveToStart(); d.list.CurrPos() < d.list.Length(); d.list.Next() {
		temp := d.list.GetValue()
		if key == temp.key {
			return temp.val, true
		}
	}

	var zero V
	return zero, false
}

func (d UALdict[K, V]) Insert(key K, value V) {
	temp := &KVPair[K, V]{key: key, val: value}
	d.list.Append(temp)
}

func (d UALdict[K, V]) Remove(key K) V {
	temp, ok := d.Find(key)

	if ok {
		d.list.Remove()
	}

	return temp
}

func (d UALdict[K, V]) RemoveAny() V {
	d.list.MoveToEnd()
	d.list.Prev()
	e := d.list.Remove()

	return e.val
}

func (d UALdict[K, V]) Clear() {
	d.list.Clear()
}

func (d UALdict[K, V]) Size() int {
	return d.list.Length()
}
