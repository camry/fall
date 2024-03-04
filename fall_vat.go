package fall

import (
	"github.com/camry/fall/pb"
	"github.com/camry/g/gerrors/gerror"
)

// runVat 执行木桶原理掉落。
func (f *Fall) runVat() ([]*pb.Item, error) {
	if len(f.TableVats()) <= 0 {
		return nil, gerror.Newf(`runVat fall.TableVats Is Empty`)
	}
	items := make([]*pb.Item, 0, 10)
	// 计算掉落权重
	var totalWeight int32
	for _, vat := range f.TableVats() {
		// 计算空缺量 = 预期量 - 现有量
		vat.VacancyNum = vat.GetExpectNum() - vat.GetExistingNum()
		if vat.GetVacancyNum() < 10 {
			vat.VacancyNum = 10
		}
		// 计算空缺率
		vat.VacancyProb = float32(vat.GetVacancyNum() / vat.GetExpectNum())
		// 算出额外系数
		switch {
		case vat.GetVacancyProb() >= 0.90:
			vat.Coefficient = 4
		case vat.GetVacancyProb() >= 0.70 && vat.GetVacancyProb() < 0.90:
			vat.Coefficient = 2
		default:
			vat.Coefficient = 1
		}
		// 得出掉落权重
		vat.VatWeight = vat.GetVacancyNum() * vat.GetCoefficient()
		totalWeight += vat.GetVatWeight()
	}
	if totalWeight > 0 {
		for _, vat := range f.TableVats() {
			if f.Rand().Hit(int(vat.GetVatWeight()), int(totalWeight)) {
				items = append(items, &pb.Item{
					Type: vat.GetFallType(),
					Id:   vat.GetFallTypeId(),
					Num:  int32(f.Rand().RangeInt(int(vat.GetVatNumMin()), int(vat.GetVatNumMax()))),
				})
			}
		}
	}
	return items, nil
}
