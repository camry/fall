package fall

import (
    "context"
    "github.com/camry/g/gutil/grand"
    "time"

    "github.com/camry/g/glog"
)

type DropMode int8

const (
    PercentMode     DropMode = iota + 1 // 逐个百分比掉落
    WeightGroupMode                     // 权重掉落组式掉落
    AdvancedMode                        // 进阶掉落
    VatMode                             // 木桶原理掉落
)

// Fall 掉落对象。
type Fall struct {
    ctx  context.Context // 上下文。
    seed int64           // 随机种子。
    rand *grand.GRand    // 种子随机数对象。
    mode DropMode        // 掉落模式。
}

// New 新建掉落对象。
func New(opts ...Option) *Fall {
    f := &Fall{
        ctx:  context.Background(),
        seed: time.Now().UnixNano(),
    }
    for _, opt := range opts {
        opt(f)
    }
    f.rand = grand.NewRand(f.seed)
    glog.Debug("Fall New Finished")
    return f
}

// Ctx 上下文。
func (f *Fall) Ctx() context.Context {
    return f.ctx
}

// Seed 随机种子。
func (f *Fall) Seed() int64 {
    return f.seed
}

// Rand 种子随机数对象。
func (f *Fall) Rand() *grand.GRand {
    return f.rand
}

// Mode 掉落模式。
func (f *Fall) Mode() DropMode {
    return f.mode
}
