package fall_test

import (
    "github.com/camry/fall"
    "github.com/camry/fall/pb"
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

func TestRunPercent(t *testing.T) {
    f := fall.New(
        fall.Seed(1683614541708899000),
        fall.Mode(fall.PercentMode),
        fall.TablePercents([]*pb.TablePercent{
            {Type: 1, Id: 1, Prob: 0.15, Min: 1, Max: 2},
            {Type: 1, Id: 2, Prob: 0.20, Min: 1, Max: 3},
            {Type: 1, Id: 3, Prob: 0.25, Min: 1, Max: 4},
            {Type: 1, Id: 4, Prob: 0.30, Min: 1, Max: 5},
        }),
    )
    items, err := f.Run()
    assert.Nil(t, err)
    v := []*pb.Item{
        {Type: 1, Id: 3, Num: 1},
        {Type: 1, Id: 4, Num: 2},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestRunWeightGroup(t *testing.T) {
    // TODO TestRunWeightGroup
}

func TestRunAdvance(t *testing.T) {
    // TODO TestRunAdvance
}

func TestVat(t *testing.T) {
    // TODO TestVat
}
