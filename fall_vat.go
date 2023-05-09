package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runVat 执行木桶原理掉落。
func (f *Fall) runVat() ([]*pb.Item, error) {
    return nil, gerror.Newf(`VatMode Is Not Implement`)
}
