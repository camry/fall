package fall_test

import (
    "testing"

    "github.com/camry/g/gutil/grand/v2"
    "github.com/stretchr/testify/assert"

    "github.com/camry/fall"
    "github.com/camry/fall/pb"
)

func TestNew(t *testing.T) {
    seed := int64(10000)
    f := fall.New(fall.Rand(grand.NewRand(uint64(seed))))
    assert.Equal(t, 1, f.Rand().RangeInt(1, 5))
}

func TestRand(t *testing.T) {
    seed := int64(10000)
    f := fall.New(fall.Rand(grand.NewRand(uint64(seed))))
    assert.True(t, f.Rand().HitProb(0.80))
}

func TestRunPercent(t *testing.T) {
    f := fall.New(
        fall.Rand(grand.NewRand(1683699333882771600)),
        fall.Mode(fall.PercentMode),
        fall.TablePercents([]*pb.TablePercent{
            {AutoId: 1, FallType: 1, FallTypeId: 1, Prob: 0.15, NumMin: 1, NumMax: 2},
            {AutoId: 2, FallType: 1, FallTypeId: 2, Prob: 0.20, NumMin: 1, NumMax: 3},
            {AutoId: 3, FallType: 1, FallTypeId: 3, Prob: 0.25, NumMin: 1, NumMax: 4},
            {AutoId: 4, FallType: 1, FallTypeId: 4, Prob: 0.30, NumMin: 1, NumMax: 5},
        }),
    )
    items, err := f.Run()
    assert.Nil(t, err)
    v := []*pb.Item{
        {AutoId: 2, Type: 1, Id: 2, Num: 2},
        {AutoId: 3, Type: 1, Id: 3, Num: 3},
        {AutoId: 4, Type: 1, Id: 4, Num: 2},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestRunWeightGroup(t *testing.T) {
    f := fall.New(
        fall.Rand(grand.NewRand(1683699333882771600)),
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
        {Type: 1, Id: 0, Num: 1267},
        {Type: 2, Id: 2, Num: 3},
        {Type: 3, Id: 3, Num: 1},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestRunAdvance(t *testing.T) {
    f := fall.New(
        fall.Rand(grand.NewRand(1683699333882771600)),
        fall.Mode(fall.AdvanceMode),
        fall.TableWeightGroupMasters([]*pb.TableWeightGroupMaster{
            {MasterId: 101, SubsetId: 1001, NextSubsetId: 1002, NextSubsetMin: 2, NextSubsetMax: 3, AdvanceNum: 2},
            {MasterId: 101, SubsetId: 1002, IsNextSubset: true, NextSubsetId: 1003, NextSubsetMin: 3, NextSubsetMax: 5, AdvanceNum: 1},
            {MasterId: 101, SubsetId: 1003, IsNextSubset: true, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0, AdvanceNum: 0},
            {MasterId: 101, SubsetId: 1004, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0, AdvanceNum: 0},
            {MasterId: 101, SubsetId: 1005, NextSubsetId: 0, NextSubsetMin: 0, NextSubsetMax: 0, AdvanceNum: 0},
        }),
        fall.TableWeightGroupSubsets([]*pb.TableWeightGroupSubset{
            {AutoId: 1, SubsetId: 1001, FallType: 1, FallTypeId: 0, SubsetNumMin: 1, SubsetNumMax: 2000, SubsetWeight: 2000},
            {AutoId: 2, SubsetId: 1002, FallType: 1, FallTypeId: 0, SubsetNumMin: 5000, SubsetNumMax: 10000, SubsetWeight: 2000},
            {AutoId: 3, SubsetId: 1003, FallType: 2, FallTypeId: 1, SubsetNumMin: 20, SubsetNumMax: 20, SubsetWeight: 1},
            {AutoId: 4, SubsetId: 1004, FallType: 4, FallTypeId: 1, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {AutoId: 5, SubsetId: 1004, FallType: 4, FallTypeId: 2, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {AutoId: 6, SubsetId: 1004, FallType: 4, FallTypeId: 3, SubsetNumMin: 1, SubsetNumMax: 3, SubsetWeight: 1000},
            {AutoId: 7, SubsetId: 1005, FallType: 5, FallTypeId: 1, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 2000},
            {AutoId: 8, SubsetId: 1005, FallType: 5, FallTypeId: 2, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 2000},
            {AutoId: 9, SubsetId: 1005, FallType: 5, FallTypeId: 3, SubsetNumMin: 1, SubsetNumMax: 1, SubsetWeight: 4000},
        }),
    )
    items, err := f.Run()
    assert.Nil(t, err)
    t.Log(items)
    v := []*pb.Item{
        {AutoId: 2, Type: 1, Id: 0, Num: 9287},
        {AutoId: 4, Type: 4, Id: 1, Num: 2},
        {AutoId: 9, Type: 5, Id: 3, Num: 1},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetAutoId(), item.GetAutoId())
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}

func TestVat(t *testing.T) {
    f := fall.New(
        fall.Rand(grand.NewRand(1683699333882771600)),
        fall.Mode(fall.VatMode),
        fall.TableVats([]*pb.TableVat{
            {VatId: 10001, FallType: 101, FallTypeId: 1001, VatNumMin: 1, VatNumMax: 1, ExpectNum: 1000, ExistingNum: 0},
            {VatId: 10001, FallType: 102, FallTypeId: 2001, VatNumMin: 1, VatNumMax: 1, ExpectNum: 1000, ExistingNum: 500},
            {VatId: 10001, FallType: 103, FallTypeId: 3001, VatNumMin: 1, VatNumMax: 1, ExpectNum: 1000, ExistingNum: 600},
            {VatId: 10001, FallType: 104, FallTypeId: 4001, VatNumMin: 1, VatNumMax: 1, ExpectNum: 1000, ExistingNum: 111},
        }),
    )
    items, err := f.Run()
    assert.Nil(t, err)
    v := []*pb.Item{
        {Type: 101, Id: 1001, Num: 1},
        {Type: 104, Id: 4001, Num: 1},
    }
    for i, item := range items {
        assert.Equal(t, v[i].GetType(), item.GetType())
        assert.Equal(t, v[i].GetId(), item.GetId())
        assert.Equal(t, v[i].GetNum(), item.GetNum())
    }
}
