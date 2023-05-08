package fall_test

import (
    "github.com/camry/fall"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNew(t *testing.T) {
    seed := int64(10000)
    f := fall.New(fall.Seed(seed))
    assert.Equal(t, seed, f.Seed())
}

func TestRand(t *testing.T) {
    seed := int64(10000)
    f := fall.New(fall.Seed(seed))
    assert.True(t, f.Rand().HitProb(0.80))
}
