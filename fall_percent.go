package fall

import (
    "github.com/camry/fall/pb"
    "github.com/camry/g/gerrors/gerror"
)

// runPercent 执行逐个百分比掉落。
func (f *Fall) runPercent() ([]*pb.Item, error) {
    if len(f.TablePercents()) <= 0 {
        return nil, gerror.Newf(`runPercent fall.TablePercents Is Empty`)
    }
    items := make([]*pb.Item, 0, 10)
    for _, v := range f.TablePercents() {
        if f.Rand().HitProb(v.GetProb()) {
            items = append(items, &pb.Item{
                Type: v.GetType(),
                Id:   v.GetId(),
                Num:  uint32(f.Rand().RangeInt(int(v.GetMin()), int(v.GetMax()))),
            })
        }
    }
    return items, nil
}
