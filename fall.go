package fall

import (
    "context"

    "github.com/camry/g/gerrors/gerror"
    "github.com/camry/g/gutil/grand/v2"

    "github.com/camry/fall/pb"
)

type DropMode int8

const (
    PercentMode     DropMode = iota + 1 // 逐个百分比掉落
    WeightGroupMode                     // 权重掉落组式掉落
    AdvanceMode                         // 进阶掉落
    VatMode                             // 木桶原理掉落
)

// Fall 掉落对象。
type Fall struct {
    ctx                     context.Context              // 上下文。
    seed                    int64                        // 随机种子。
    rand                    *grand.GRand                 // 种子随机数对象。
    mode                    DropMode                     // 掉落模式。
    tablePercents           []*pb.TablePercent           // 配置表（逐个百分比掉落列表）。
    tableWeightGroupMasters []*pb.TableWeightGroupMaster // 配置表（权重掉落组式掉落母集进阶掉落列表）。
    tableWeightGroupSubsets []*pb.TableWeightGroupSubset // 配置表（权重掉落组式掉落子集列表）。
    tableVats               []*pb.TableVat               // 配置表（木桶原理掉落列表）。
}

// New 新建掉落对象。
func New(opts ...Option) *Fall {
    f := &Fall{
        ctx: context.Background(),
    }
    for _, opt := range opts {
        opt(f)
    }
    return f
}

// Ctx 上下文。
func (f *Fall) Ctx() context.Context {
    return f.ctx
}

// Rand 种子随机数对象。
func (f *Fall) Rand() *grand.GRand {
    return f.rand
}

// Mode 掉落模式。
func (f *Fall) Mode() DropMode {
    return f.mode
}

// TablePercents 配置表（逐个百分比掉落列表）。
func (f *Fall) TablePercents() []*pb.TablePercent {
    return f.tablePercents
}

// TableWeightGroupMasters 配置表（权重掉落组式掉落母集进阶掉落列表）。
// `tableWeightGroupMasters` 格式 [MasterId_SubsetId]*pb.TableWeightGroupMaster
func (f *Fall) TableWeightGroupMasters() []*pb.TableWeightGroupMaster {
    return f.tableWeightGroupMasters
}

// TableWeightGroupSubsets 配置表（权重掉落组式掉落子集列表）。
func (f *Fall) TableWeightGroupSubsets() []*pb.TableWeightGroupSubset {
    return f.tableWeightGroupSubsets
}

// TableVats 配置表（木桶原理掉落列表）。
func (f *Fall) TableVats() []*pb.TableVat {
    return f.tableVats
}

// Run 开始执行。
func (f *Fall) Run() ([]*pb.Item, error) {
    switch f.Mode() {
    case PercentMode:
        return f.runPercent()
    case WeightGroupMode:
        return f.runWeightGroup()
    case AdvanceMode:
        return f.runAdvance()
    case VatMode:
        return f.runVat()
    default:
        return nil, gerror.Newf(`Unknown DropMode %d`, f.Mode())
    }
}
