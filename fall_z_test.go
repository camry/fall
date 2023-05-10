package fall_test

import (
    "testing"

    "github.com/camry/fall"
    "github.com/camry/fall/pb"
    "github.com/stretchr/testify/assert"
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
        fall.Seed(1683699333882771600),
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
        {Type: 1, Id: 2, Num: 2},
        {Type: 1, Id: 3, Num: 3},
        {Type: 1, Id: 4, Num: 2},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestRunWeightGroup(t *testing.T) {
    f := fall.New(
        fall.Seed(1683699333882771600),
        fall.Mode(fall.WeightGroupMode),
        fall.TableWeightGroupMasters([]*pb.TableWeightGroupMaster{
            {MasterId: 101, SubsetId: 1001, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0},
            {MasterId: 101, SubsetId: 1002, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0},
            {MasterId: 101, SubsetId: 1003, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0},
        }),
        fall.TableWeightGroupSubsets([]*pb.TableWeightGroupSubset{
            {SubsetId: 1001, FallType: 1, FallTypeId: 0, SubsetNumMin: 1, SubsetNumMax: 2000, SubsetWeight: 2000},
            {SubsetId: 1002, FallType: 2, FallTypeId: 1, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {SubsetId: 1002, FallType: 2, FallTypeId: 2, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {SubsetId: 1002, FallType: 2, FallTypeId: 3, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {SubsetId: 1003, FallType: 3, FallTypeId: 1, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 2000},
            {SubsetId: 1003, FallType: 3, FallTypeId: 2, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 2000},
            {SubsetId: 1003, FallType: 3, FallTypeId: 3, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 4000},
        }),
    )
    items, err := f.Run()
    assert.Nil(t, err)
    v := []*pb.Item{
        {Type: 1, Id: 0, Num: 1811},
        {Type: 2, Id: 2, Num: 2},
        {Type: 3, Id: 1, Num: 1},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestRunAdvance(t *testing.T) {
    // TODO TestRunAdvance
}

func TestVat(t *testing.T) {
    // TODO TestVat
}
