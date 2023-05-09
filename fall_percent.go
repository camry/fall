package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runPercent 执行逐个百分比掉落。
func (f *Fall) runPercent() ([]*pb.Item, error) {
    if len(f.tablePercents) <= 0 {
        return nil, gerror.Newf(`fall.TablePercents Is Empty`)
    }
    var items []*pb.Item
    for _, v := range f.TablePercents() {
        if f.Rand().HitProb(v.GetProb()) {
            items = append(items, &pb.Item{Id: v.GetId(), Num: uint32(f.Rand().Intn(int(v.GetMax()-v.GetMin()))) + v.GetMin()})
        }
    }
    return items, nil
}
