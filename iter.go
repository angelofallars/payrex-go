package payrex

// Seq2 is the same as the iter.Seq2 type in Go 1.23+.
type Seq2[K, V any] func(yield func(K, V) bool)
