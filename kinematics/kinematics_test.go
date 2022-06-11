package kinematics

import (
	"fmt"
	"math/big"
	"testing"
)

func TestComputesDisplacementGivenTimeEqualsThree(t *testing.T) {
	fn := GenDisplaceFn(10, 2, 1)
	got := fn(3)
	wanted := 52.0

	if big.NewFloat(got).Cmp(big.NewFloat(wanted)) != 0 {
		t.Errorf("got: %f; wanted: %f\n", got, wanted)
	} else {
		fmt.Printf("got: %f; wanted: %f\n", got, wanted)
	}

}

func TestComputesDisplacementGivenTimeEqualsFive(t *testing.T) {
	fn := GenDisplaceFn(10, 2, 1)
	got := fn(5)
	wanted := 136.0

	if big.NewFloat(got).Cmp(big.NewFloat(wanted)) != 0 {
		t.Errorf("got: %f; wanted: %f\n", got, wanted)
	} else {
		fmt.Printf("got: %f; wanted: %f\n", got, wanted)
	}

}
