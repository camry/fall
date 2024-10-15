package fall

import (
    "context"

    "github.com/camry/g/gutil/grand/v2"

    "github.com/camry/fall/pb"
)

type Option func(f *Fall)

// Context 设置上下文。
func Context(ctx context.Context) Option {
    return func(f *Fall) { f.ctx = ctx }
}

// Rand 设置种子随机数 GRand。
func Rand(rand *grand.GRand) Option {
    return func(f *Fall) { f.rand = rand }
}

// Mode 设置掉落模式。
func Mode(mode DropMode) Option {
    return func(f *Fall) { f.mode = mode }
}

// TablePercents 设置配置表（逐个百分比掉落列表）。
func TablePercents(tablePercents []*pb.TablePercent) Option {
    return func(f *Fall) { f.tablePercents = tablePercents }
}

// TableWeightGroupMasters 设置配置表（权重掉落组式掉落母集进阶掉落列表）。
// 注：只支持一个母集对应的子集列表。
func TableWeightGroupMasters(tableWeightGroupMasters []*pb.TableWeightGroupMaster) Option {
    return func(f *Fall) { f.tableWeightGroupMasters = tableWeightGroupMasters }
}

// TableWeightGroupSubsets 设置配置表（权重掉落组式掉落子集列表）。
// 注：只支持一个母集对应的子集列表。
func TableWeightGroupSubsets(tableWeightGroupSubsets []*pb.TableWeightGroupSubset) Option {
    return func(f *Fall) { f.tableWeightGroupSubsets = tableWeightGroupSubsets }
}

// TableVats 设置配置表（木桶原理掉落列表）。
func TableVats(tableVats []*pb.TableVat) Option {
    return func(f *Fall) { f.tableVats = tableVats }
}
