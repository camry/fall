package fall_test

import (
    "context"
    "testing"

    "github.com/camry/g/gutil/grand/v2"
    "github.com/stretchr/testify/assert"

    "github.com/camry/fall"
)

func TestContext(t *testing.T) {
    ctx := context.Background()
    f := fall.New(fall.Context(ctx))
    assert.Equal(t, f.Ctx(), ctx)
}

func TestSeed(t *testing.T) {
    rand := grand.NewRand(10000)
    f := fall.New(fall.Rand(rand))
    assert.Equal(t, rand, f.Rand())
}

func TestMode(t *testing.T) {
    f := fall.New(fall.Mode(fall.AdvanceMode))
    assert.Equal(t, fall.AdvanceMode, f.Mode())
}
