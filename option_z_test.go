package fall_test

import (
    "context"
    "testing"

    "github.com/camry/fall"
    "github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
    ctx := context.Background()
    f := fall.New(fall.Context(ctx))
    assert.Equal(t, f.Ctx(), ctx)
}

func TestSeed(t *testing.T) {
    f := fall.New(fall.Seed(10000))
    assert.Equal(t, int64(10000), f.Seed())
}
