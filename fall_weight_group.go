package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runWeightGroup 执行权重掉落组式掉落。
func (f *Fall) runWeightGroup() ([]*pb.Item, error) {
    return nil, gerror.Newf(`WeightGroupMode Is Not Implement`)
}
