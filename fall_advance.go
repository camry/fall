package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runAdvance 执行进阶掉落。
func (f *Fall) runAdvance() ([]*pb.Item, error) {
    return nil, gerror.Newf(`AdvanceMode Is Not Implement`)
}
