# fall

游戏内的掉落包

## 游戏内的掉落方式

- [x] 逐个百分比掉落
- [x] 权重掉落组式掉落
- [ ] 进阶掉落
- [ ] 木桶原理掉落

## protoc 指令

```bash
protoc --go_opt=paths=source_relative --go_out=. pb/*.proto
```