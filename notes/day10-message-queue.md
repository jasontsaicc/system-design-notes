# Day 10 — Message Queue & Async Processing

> Session 11 (Part 1): Chunk 1-4 教學 + Chunk 5 教學
> Session 12 (Part 2): Chunk 5-6 Gate ✅ + Design Exercise ✅ + Simon Drill ✅
> 下次繼續：Step F (Interview Drill) → Notes 補完

---

## One-Liner

> Message Queue 是 service 之間的非同步通訊機制，producer 把 message 丟進 queue，consumer 自己的節奏處理，達成 decoupling、buffering、resilience 三大好處。

---

## 核心概念

### 1. Why Async？Sync vs Async

**Synchronous 的問題：**
- 所有步驟串聯等待 → response time = 所有步驟加總
- 任一 downstream service 掛了 → 整條 chain 失敗
- 流量暴增 → downstream 來不及處理 → timeout → 用戶 retry → 重複處理

**Async（用 Queue）的三大好處：**

| 好處 | 說明 |
|------|------|
| **Decoupling** | Producer 不需要知道有哪些 consumer，加新 service 不需改 producer |
| **Buffering** | Queue 吸收流量高峰，consumer 按自己速度處理 |
| **Resilience** | Consumer 掛了，message 在 queue 裡等，不會丟失 |

### 2. Core Components

```
Producer → Broker (manages queues) → Consumer
```

- **Producer**：送出 message 的 service（例：Order Service）
- **Consumer**：接收並處理 message 的 service（例：Payment Service）
- **Broker**：管理 queue、路由 message（例：SQS, Kafka, RabbitMQ）
- **Topic / Queue**：用名稱分類的 message channel

**兩種 pattern：**

| Pattern | 行為 | 用途 |
|---------|------|------|
| **Point-to-Point** | 一個 message → 一個 consumer 取走 | Job processing |
| **Pub/Sub** | 一個 message → 所有 subscriber 都收到一份 | Event broadcasting（多個 service 都需要同一事件）|

**重點**：Pub/Sub 才能維持 decoupling。用 Point-to-Point 的話，producer 要分別送到每個 queue → 又 coupling 回去了。

### 3. SQS vs Kafka vs RabbitMQ

| | SQS | Kafka | RabbitMQ |
|---|---|---|---|
| 定位 | Managed queue (AWS) | Distributed log | Traditional broker |
| 最適合 | 簡單 job queue | 高吞吐 event streaming | 複雜 routing / priority |
| 消費後 message | **刪除** | **保留**（可 replay） | **刪除** |
| Ordering | FIFO mode only | Per-partition | Per-queue |
| Throughput | 中 | 非常高（millions/sec） | 中 |
| 維運成本 | 低（fully managed） | 高（Zookeeper/KRaft） | 中 |

**面試決策框架：**
- 簡單 async job → **SQS**
- 需要 event replay / analytics → **Kafka**
- 複雜 routing / priority → **RabbitMQ**
- 小團隊 + AWS → **SQS 先，之後有需要再升 Kafka**

**Kafka 最大差異**：message 消費後不會刪除，它是 **log** 不是 queue。可以 replay 任意時間點的 message。

### 4. Delivery Semantics（最重要！）

| Semantic | 意思 | 風險 | 適用場景 |
|----------|------|------|----------|
| **At-most-once** | 送一次不 retry | 可能 **丟失** message | Logging, metrics |
| **At-least-once** | 持續 retry 直到 ACK | 可能 **重複** 處理 | 大多數 production 系統 |
| **Exactly-once** | 每個 message 恰好處理一次 | ⚠️ 分散式系統中幾乎不可能 | 理想情況 |

**為什麼會重複（at-least-once）：**
```
Consumer 處理完 message → 準備 ACK → 網路斷了 → ACK 失敗
→ Broker 以為沒處理 → 重新 deliver → Consumer 又處理一次 😱
```

**為什麼 exactly-once 很難：**
- Broker 無法區分「consumer 處理完但 ACK 丟了」vs「consumer 處理到一半 crash 了」
- 這是 distributed systems 的根本問題（Two Generals' Problem）
- Kafka 的 "exactly-once" 其實是 idempotent producer + transactional consumer = **effectively-once**

**面試金句：**
> "True exactly-once delivery is impossible in distributed systems. What we actually do is at-least-once delivery combined with idempotent processing to achieve effectively-once semantics."

**Production 選擇：at-least-once** — 因為 duplicate 可以用 idempotency 處理，但 lost message 是不可逆的。

### 5. Idempotency ✅

**Idempotency** = 同一操作執行多次，結果跟執行一次一樣。

**解法：Idempotency Key**
```
Producer 產生 unique key: "order-123-payment"
Consumer 收到 message:
  → 查 DB/Redis：這個 key 處理過了嗎？
    → YES → skip，回傳之前的結果
    → NO → 處理 → 存 key → 回傳結果
```

| 設計要素 | 做法 |
|----------|------|
| Key 格式 | `{entity}-{id}-{action}` |
| 儲存 | Redis（快速查詢 + TTL）或 DB table |
| TTL | 跟 message retention 對齊（例如 14 天） |

**⚠️ 小杰的錯誤做法 vs 正確做法：**

| 小杰：separate flag | 正確：atomic transaction |
|---|---|
| 先處理，再設 `processed = true` | idempotency key 和 business logic 在**同一個 transaction** |
| Crash between process & flag → 重複處理 | Atomic：兩者一起成功或一起 rollback |
| Race condition with multiple consumers | Unique constraint on key → 第二次 write 自動失敗 |

```sql
-- atomic idempotency（正確做法）
BEGIN TRANSACTION;
  INSERT INTO processed_messages (idempotency_key) VALUES ('order-123-abc');
  -- duplicate key → transaction fails → skip
  UPDATE account SET balance = balance - 100 WHERE user_id = 42;
COMMIT;
```

**面試金句：** "The dedup check and the business operation must be atomic — one transaction, not two separate steps."

### 6. Observability Mini ✅

| Element | Message Queue |
|---------|--------------|
| **SLIs** | Queue depth, consumer lag, processing latency (P50/P99), error rate (DLQ count) |
| **SLO target** | Consumer lag < 30s at P99, error rate < 0.1%, queue depth 不能無限增長 |
| **Alerts** | 🔴 Consumer lag > 60s, 🔴 DLQ count spike, 🟡 Queue depth 持續增長 5+ min |
| **Dashboards** | 1. Throughput (messages in vs out/sec) 2. Latency distribution 3. DLQ trend + queue depth |

**Consumer lag** = queue 最重要的 metric。代表 consumer 跟不上 producer 的程度。
- Kafka: latest produced offset - latest consumed offset (per partition)
- SQS: `ApproximateNumberOfMessagesVisible`

**3 AM Queue Depth Alert — Debug Decision Tree：**
```
Queue depth growing 📈
  ├─ Producer throughput spiked? → traffic surge → scale consumers
  ├─ Consumer throughput dropped?
  │   ├─ Consumer errors up? → check DLQ → fix bug / rollback
  │   ├─ Consumer latency up? → downstream slow (DB? API?)
  │   └─ Consumers crashed? → check pod count → restart
  └─ Both normal? → partition imbalance (Kafka) / visibility timeout (SQS)
```

---

## Design Exercise: Order Processing with MQ ✅

**架構圖：**
```
Client ─→ DNS ─→ LB ─→ API Server ─→ Cache (order status)
                            │                    ↑
                      (async)│              (update status)
                            ▼                    │
                         Queue ──→ Worker ──→ Database
                            │         │
                            │    [idempotency check]
                            ▼
                           DLQ (failed messages)
```

**設計決策：**
1. **API Server** sync 回覆 "order received"（快），heavy work 丟 queue（async）
2. **Idempotency check 在 Worker**（因為 Worker 才是收到重複 message 的人）
3. **DLQ** 接 failed messages → alert → 人工調查

---

## 🔴 My Mistakes & Misconceptions

| What I Thought | Reality | Why I Was Wrong |
|---|---|---|
| Async 的好處只想到速度（buffering） | 三個好處：decoupling, buffering, resilience | 只從 user 角度想，沒從 system architecture 角度想 |
| 知道答案是 Pub/Sub 但說不出 why | Point-to-Point 需要 producer 知道所有 consumer → coupling | 沒想到 Point-to-Point 的 coupling 問題 |
| 以為 SQS 也可以 replay message | SQS 消費後 message 就刪了，只有 Kafka 保留 | 混淆了 retention（未消費）和 replay（已消費可重讀） |
| 以為 exactly-once 很難是因為「慢」 | 根本原因是 distributed systems 無法區分 ACK 丟失 vs consumer crash | 把 performance 問題跟 fundamental impossibility 搞混 |
| Simon Drill: Why Async 只記得 "fast response" | 三個好處：decoupling, buffering, resilience | 從 user 角度想得到 fast，但 system architecture 角度的 decoupling 和 buffering 沒內化 |
| Simon Drill: delivery semantics 說成 "most, least, excely" | 完整名稱：at-most-once, at-least-once, exactly-once | 術語要講完整，面試印象差很多 |
| 設計練習不知道怎麼開始 | 拆成小問題一步步推：先決定 sync vs async → 再決定 idempotency 放哪 | 面試也是一樣，不需要一次畫完，一步一步 reason through |

---

## 🎤 How to Say It in Interview

**Opening (30 sec):**
> "A message queue enables asynchronous communication between services. The producer sends messages to a broker, and consumers process them independently. The key benefits are decoupling, buffering during traffic spikes, and resilience when downstream services fail."

**When asked "Why not synchronous?":**
> "Synchronous calls create tight coupling and cascading failures. If any downstream service is slow or down, the entire chain fails. With a queue, the producer returns immediately, and messages are safely buffered until consumers can process them."

**When asked "How do you handle duplicate messages?":**
> "In production, we use at-least-once delivery because losing messages is unrecoverable. To prevent duplicate processing, we implement idempotency — each message carries a unique key, and the consumer checks if it's already been processed before acting on it."

**When asked "SQS or Kafka?":**
> "It depends on the use case. SQS for simple job queues with minimal ops overhead. Kafka when you need event replay, high throughput, or streaming analytics. Start simple with SQS and evolve to Kafka when the need arises."

---

## 🗣️ English Practice

| My Answer | English Polish |
|---|---|
| synchronous will wait all step finish, but if use queue will like todo list save to queue and get number | Synchronous processing waits for all steps to finish before responding. But with a queue, it's like a to-do list — the request is saved to the queue and the user gets a confirmation number immediately. |
| is pub/sub | It's Pub/Sub, because all three services — Payment, Inventory, and Notification — need to receive the same event. With Point-to-Point, the producer would need to send separately to each, which re-introduces coupling. |
| if aws and team only 5 persion and don't want spend to many time i would pick SQS because SQS is managed AWS | If we're on AWS with a small team of 5 and don't want to spend too much time on operations, I'd pick SQS because it's fully managed — we don't need to care about infrastructure, and it scales easily within the AWS ecosystem. |
| At-least-once delivery because in production systems can dulicate but not loses | At-least-once delivery, because in production systems you can handle duplicates with idempotency, but you can't afford to lose messages. |
| payment processing is not good idea, because at-most-once only send massage once maybe will losing massage | At-most-once is terrible for payment processing because if a message is lost, the user paid but their order never gets processed. Duplicates are fixable with idempotency, but lost payments are unrecoverable. |
| use SQS or Kafka can thinks about trade-off like the cost, infra | The choice between SQS and Kafka depends on whether you need message replay. SQS for simple job queues, Kafka when you need event replay or streaming analytics. |
| worker will get duplicate message | The Worker is the one receiving duplicate messages from the queue, so it's responsible for the idempotency check. |

---

## Simon Drill Recall Gaps

| Chunk | Recalled | Gap |
|---|---|---|
| 1. Why Async | fast response | 漏了 **decoupling** 和 **buffering** |
| 2. Core Components | ✅ | — |
| 3. SQS/Kafka/RabbitMQ | 方向對但模糊 | 要用 **event streaming + replay** 定位 Kafka，**complex routing** 定位 RabbitMQ |
| 4. Delivery Semantics | "most, least, excely" | 講完整名稱：at-most-once, at-least-once, exactly-once |
| 5. Idempotency | ✅ concept | 補充 **how**：atomic transaction with idempotency key |
| 6. Observability | ✅ triage 對 | 記住關鍵字 **consumer lag** |

---

## 待完成（下次 session 繼續）

- [ ] Step F: Interview Drill（flash sale order system）
- [ ] Step G: Notes 補完（Trade-off, Scale trigger, DevOps angle）
- [ ] Step H: Progress Update
