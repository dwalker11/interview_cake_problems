package dict

import (
	lists "dwalker/lib/list"
)

type SALdict[K comparable, T any] struct {
	list lists.SAList[*KVPair[K, T]]
}

func (d SALdict[K, T]) Find(key K) (T, bool) {
	l := -1
	r := d.list.Length()

	for l+1 != r {
		i := (l + r) / 2
		d.list.MoveToPos(i)
		temp := d.list.GetValue()

		if key < temp.key {
			r = i
		}

		if key == temp.key {
			return temp.val, true
		}

		if key > temp.key {
			l = i
		}
	}

	var zero T
	return zero, false
}

func (d SALdict[K, T]) Insert(key K, val T) {
	temp := &KVPair[K, T]{key: key, val: val}
	d.list.Insert(temp)
}

func (d SALdict[K, T]) Remove(key K) T {
	temp, ok := d.Find(key)

	if ok {
		d.list.Remove()
	}

	return temp
}

func (d SALdict[K, T]) RemoveAny() T {
	d.list.MoveToEnd()
	d.list.Prev()
	e := d.list.Remove()

	return e.val
}

func (d SALdict[K, T]) Clear() {
	d.list.Clear()
}

func (d SALdict[K, T]) Size() int {
	return d.list.Length()
}
