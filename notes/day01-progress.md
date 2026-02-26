# Day 01 - What SD Interviews Actually Test (進度)

## Status: 🔄 In Progress

## 已完成
- A. 複習 Day -1 ✅
- B. 引入（廚藝比賽比喻）✅
- C. 核心教學（4 Dimensions）✅
  - Dimension 1: Problem Navigation — 問對問題釐清需求
  - Dimension 2: System Design — 畫合理架構圖
  - Dimension 3: Technical Depth — 深入解釋 why + how + 限制
  - Dimension 4: Trade-off Analysis — 分析優缺點與替代方案
  - 口訣：問 → 畫 → 深 → 比

## 未完成（下次繼續）
- D. 討論練習 — URL Shortener 題目已出，需用 4 Dimensions 評分回答 A vs 回答 B
- E. Voice Drill
- F. 筆記 (notes/day01-interview-rubric.md)
- G. 更新進度 (CURRICULUM.md + curriculum-roadmap.md)

## User 卡住 / 要記錄到筆記的點
- 擔心問問題後被問到不會的（如問了群聊但不熟群聊）→ 問問題是展示思考能力，不是承諾都會
- 擔心沒實際工作經驗的技術講不出深度（如 WebSocket）→ 理解原理就夠，知道 what problem / why this / downsides
- 被問 4 Dimensions 時一開始說不出來，需要提示才回想起來

## D. 討論練習題目（下次要做）

### 題目：Design a URL Shortener

**回答 A：**
用一個 server 接收長網址，產生短網址存到資料庫。用戶訪問短網址時，server 從資料庫查出長網址，redirect 過去。資料庫用 MySQL。

**回答 B：**
先確認一下：預期的 QPS 大概多少？短網址需要過期嗎？需要自訂短碼嗎？

假設 100M URLs、讀寫比 100:1。我會用一個 Application Service 處理建立和查詢，DB 用 NoSQL 因為 schema 簡單而且讀取量大。短碼用 Base62 encode，長度 7 位可以產生 62^7 ≈ 3.5 trillion 組合，足夠用。

讀取量大所以加一層 Redis cache，cache hit 就不用查 DB，降低 latency。

Trade-off：Base62 是用 counter 遞增的，好處是不會碰撞，但缺點是短碼可預測。如果安全性重要，可以改用 random hash + collision check，但會增加 DB write。

### 任務：用 4 Dimensions 分別評分這兩個回答
