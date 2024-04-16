package fall

import (
    "github.com/camry/g/gerrors/gerror"

    "github.com/camry/fall/pb"
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
                AutoId: v.GetAutoId(),
                Type:   v.GetFallType(),
                Id:     v.GetFallTypeId(),
                Num:    int32(f.Rand().RangeInt(int(v.GetNumMin()), int(v.GetNumMax()))),
            })
        }
    }
    return items, nil
}
