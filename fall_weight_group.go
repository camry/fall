package fall

import (
    "github.com/camry/g/gerrors/gerror"

    "github.com/camry/fall/pb"
)

// runWeightGroup 执行权重掉落组式掉落。
func (f *Fall) runWeightGroup() ([]*pb.Item, error) {
    if len(f.TableWeightGroupMasters()) <= 0 {
        return nil, gerror.Newf(`runWeightGroup fall.TableWeightGroupMasters Is Empty`)
    }
    if len(f.TableWeightGroupSubsets()) <= 0 {
        return nil, gerror.Newf(`runWeightGroup fall.TableWeightGroupSubsets Is Empty`)
    }
    return f.dropAdvanceWeightGroup(false), nil
}

// dropAdvanceWeightGroup 权重掉落组式掉落（进阶掉落）。
// `enableAdvance` 进阶掉落是否可用？
func (f *Fall) dropAdvanceWeightGroup(enableAdvance bool) []*pb.Item {
    items := make([]*pb.Item, 0, 10)
    length := len(f.TableWeightGroupSubsets())
    tableWeightGroupMasters := make(map[int32]*pb.TableWeightGroupMaster, len(f.TableWeightGroupMasters()))
    tableWeightGroupSubsets := make(map[int32][]*pb.TableWeightGroupSubset, length)
    subsetTotalWeights := make(map[int32]int32, length)
    // 根据子集编号索引母集列表
    for _, master := range f.tableWeightGroupMasters {
        tableWeightGroupMasters[master.GetSubsetId()] = master
    }
    for _, subset := range f.TableWeightGroupSubsets() {
        // 汇总计算各个子集的总权重
        subsetTotalWeights[subset.GetSubsetId()] += subset.GetSubsetWeight()
        // 根据子集编号索引子集列表
        tableWeightGroupSubsets[subset.GetSubsetId()] = append(tableWeightGroupSubsets[subset.GetSubsetId()], subset)
    }
    // 根据掉落母集找到子集
    for _, master := range f.TableWeightGroupMasters() {
        // 排除进阶子集
        if master.GetIsNextSubset() {
            continue
        }
        // 判断母集是否产生进阶掉落
        subsetId := master.GetSubsetId()
        if enableAdvance {
            subsetId = f.nextSubset(master, tableWeightGroupMasters).GetSubsetId()
        }
        // 生成基于总权重的随机数
        subsetTotalWeight := subsetTotalWeights[subsetId]
        randWeight := int32(f.Rand().RangeInt(1, int(subsetTotalWeight)))
        // 判断掉落哪些物品
        var rangeWeightLast, rangeWeightCurrent int32
        for _, subset := range tableWeightGroupSubsets[subsetId] {
            isDrop := false
            rangeWeightCurrent += subset.GetSubsetWeight()
            if randWeight > rangeWeightLast && randWeight <= rangeWeightCurrent {
                isDrop = true
            }
            rangeWeightLast += subset.GetSubsetWeight()
            if isDrop {
                items = append(items, &pb.Item{
                    Type: subset.GetFallType(),
                    Id:   subset.GetFallTypeId(),
                    Num:  int32(f.Rand().RangeInt(int(subset.GetSubsetNumMin()), int(subset.GetSubsetNumMax()))),
                })
                break
            }
        }
    }
    return items
}

// nextSubset 进阶子集。
func (f *Fall) nextSubset(master *pb.TableWeightGroupMaster, tableWeightGroupMasters map[int32]*pb.TableWeightGroupMaster) *pb.TableWeightGroupMaster {
    master.AdvanceNum++
    y := int32(f.Rand().RangeInt(int(master.GetNextSubsetMin()), int(master.GetNextSubsetMax())))
    if master.GetAdvanceNum() >= y {
        if nextMaster, ok := tableWeightGroupMasters[master.GetNextSubsetId()]; ok {
            master.AdvanceNum = 0
            return f.nextSubset(nextMaster, tableWeightGroupMasters)
        }
    }
    return master
}
