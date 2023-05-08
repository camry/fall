package fall_test

import (
    "testing"
    "time"

    "github.com/camry/fall"
    "github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
    seed := time.Now().UnixNano()
    f := fall.New(fall.Seed(seed))
    assert.Equal(t, seed, f.Seed())
}
