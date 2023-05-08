package fall

import (
    "context"
    "time"

    "github.com/camry/g/glog"
)

// Fall 掉落对象。
type Fall struct {
    ctx  context.Context // 上下文。
    seed int64           // 随机种子
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
    glog.Debug("Fall New")
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
