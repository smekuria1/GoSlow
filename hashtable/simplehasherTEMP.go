package hashtable

import (
	"fmt"
	"hash/fnv"
)

// Helper hash function
func Hash(key interface{}) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return int(h.Sum32())
}
