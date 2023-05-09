package fall

import "context"

type Option func(f *Fall)

// Context 设置上下文。
func Context(ctx context.Context) Option {
    return func(f *Fall) { f.ctx = ctx }
}

// Seed 设置随机种子。
func Seed(seed int64) Option {
    return func(f *Fall) { f.seed = seed }
}

// Mode 设置掉落模式。
func Mode(mode DropMode) Option {
    return func(f *Fall) { f.mode = mode }
}
