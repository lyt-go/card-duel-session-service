# 卡牌

纯 Go 标准库（`net/http`）实现的后端服务，零第三方依赖，标准分层（cmd/internal/pkg），开箱即跑。

## 运行

```bash
cd origin
go run ./cmd/server
# 默认监听 :8080，可用 PORT/ADDR/MAX_PAGE_SIZE 环境变量覆盖
```

## 业务实体

Card（卡牌）、Deck（牌组）、Game（对局）、PlayRecord（出牌记录）、Player（玩家）、Collection（卡册/收藏）

## API 一览

统一响应结构：`{"code":0,"message":"ok","data":...}`。

| 模块 | 接口 | 说明 |
|------|------|------|
| cards | POST /api/cards | 创建Card |
| cards | GET /api/cards | 列表查询（分页+筛选） |
| cards | GET /api/cards/{id} | 详情 |
| cards | PUT /api/cards/{id} | 更新 |
| cards | DELETE /api/cards/{id} | 删除 |
| decks | POST /api/decks | 创建Deck |
| decks | GET /api/decks | 列表查询（分页+筛选） |
| decks | GET /api/decks/{id} | 详情 |
| decks | PUT /api/decks/{id} | 更新 |
| decks | DELETE /api/decks/{id} | 删除 |
| decks | PATCH /api/decks/{id}/status | 状态流转 |
| games | POST /api/games | 创建Game |
| games | GET /api/games | 列表查询（分页+筛选） |
| games | GET /api/games/{id} | 详情 |
| games | PUT /api/games/{id} | 更新 |
| games | DELETE /api/games/{id} | 删除 |
| games | PATCH /api/games/{id}/status | 状态流转 |
| play-records | POST /api/play-records | 创建PlayRecord |
| play-records | GET /api/play-records | 列表查询（分页+筛选） |
| play-records | GET /api/play-records/{id} | 详情 |
| play-records | DELETE /api/play-records/{id} | 删除 |
| players | POST /api/players | 创建Player |
| players | GET /api/players | 列表查询（分页+筛选） |
| players | GET /api/players/{id} | 详情 |
| players | PUT /api/players/{id} | 更新 |
| players | DELETE /api/players/{id} | 删除 |
| players | PATCH /api/players/{id}/status | 状态流转 |
| collections | POST /api/collections | 创建Collection |
| collections | GET /api/collections | 列表查询（分页+筛选） |
| collections | GET /api/collections/{id} | 详情 |
| collections | DELETE /api/collections/{id} | 删除 |
| stats | GET /api/stats/overview | 全局统计总览 |

## 说明

- 数据存储为内存实现，重启即清空。
- cost/attack/health/turn/level/score/count 等均为整数；played_at 为 Unix 毫秒时间戳。
