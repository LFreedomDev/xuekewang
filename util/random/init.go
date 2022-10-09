package random

import (
	"math/rand"
)

func Init(seed int64) {
	rand.Seed(seed)
}
