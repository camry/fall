package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runAdvance 执行进阶掉落。
func (f *Fall) runAdvance() ([]*pb.Item, error) {
    if len(f.TableWeightGroupMasters()) <= 0 {
        return nil, gerror.Newf(`runAdvance fall.TableWeightGroupMasters Is Empty`)
    }
    if len(f.TableWeightGroupSubsets()) <= 0 {
        return nil, gerror.Newf(`runAdvance fall.TableWeightGroupSubsets Is Empty`)
    }
    return f.dropAdvanceWeightGroup(true), nil
}
