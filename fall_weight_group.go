package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runWeightGroup 执行权重掉落组式掉落。
func (f *Fall) runWeightGroup() ([]*pb.Item, error) {
    if len(f.tableWeightGroupMasters) <= 0 {
        return nil, gerror.Newf(`fall.TableWeightGroupMasters Is Empty`)
    }
    if len(f.tableWeightGroupSubsets) <= 0 {
        return nil, gerror.Newf(`fall.TableWeightGroupSubsets Is Empty`)
    }
    items := make([]*pb.Item, 0, 10)
    length := len(f.TableWeightGroupSubsets())
    tableWeightGroupSubsets := make(map[uint32][]*pb.TableWeightGroupSubset, length)
    subsetTotalWeights := make(map[uint32]uint32, length)
    for _, v := range f.TableWeightGroupSubsets() {
        subsetTotalWeights[v.GetSubsetId()] += v.GetSubsetWeight()
        tableWeightGroupSubsets[v.GetSubsetId()] = append(tableWeightGroupSubsets[v.GetSubsetId()], v)
    }
    for _, master := range f.TableWeightGroupMasters() {
        subsetTotalWeight := subsetTotalWeights[master.GetSubsetId()]
        randWeight := uint32(f.Rand().RangeInt(1, int(subsetTotalWeight)))
        var rangeWeightLast, rangeWeightCurrent uint32
        for i, subset := range tableWeightGroupSubsets[master.GetSubsetId()] {
            isAdd := false
            rangeWeightCurrent += subset.GetSubsetWeight()
            if i > 0 {
                if randWeight > rangeWeightLast && randWeight < rangeWeightCurrent {
                    isAdd = true
                }
            } else {
                if randWeight <= subset.GetSubsetWeight() {
                    isAdd = true
                }
            }
            rangeWeightLast += subset.GetSubsetWeight()
            if isAdd {
                items = append(items, &pb.Item{
                    Type: subset.GetFallType(),
                    Id:   subset.GetFallTypeId(),
                    Num:  uint32(f.Rand().RangeInt(int(subset.GetSubsetNumMin()), int(subset.GetSubsetNumMax()))),
                })
            }
        }
    }
    return items, nil
}
