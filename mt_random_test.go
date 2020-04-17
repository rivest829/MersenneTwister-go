package mt_random_test

import (
	"fmt"
	mt_random "gitlab.dianchu.cc/ants/blueprint/infrastructure/tools/mt-random"
	"testing"
)

func TestNewMersenneTwister(t *testing.T) {
	mt := mt_random.NewMersenneTwister(0)
	for i := 0; i < 10; i++ {
		fmt.Println(mt.Rand())
	}
}
