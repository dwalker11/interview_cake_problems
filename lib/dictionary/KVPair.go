package dict

type KVPair[K comparable, V any] struct {
	key K
	val V
}
