# Day 10 — Message Queue & Async Processing（Part 2）

> Session 12: Chunk 5-6 Feynman Gate ✅ + Design Exercise ✅ + Simon Drill ✅
> 下次繼續：Step F (Interview Drill) → Notes 補完（Trade-off, Scale trigger, DevOps angle）

---

## Chunk 5 Feynman Gate: Idempotency ✅

**小杰的錯誤做法 vs 正確做法：**

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

---

## Chunk 6: Observability Mini ✅

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

**核心 pattern：fast sync ack + async heavy processing** — 適用於 order processing, email, video transcoding, notifications。

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

## 🔴 My Mistakes & Misconceptions

| What I Thought | Reality | Why I Was Wrong |
|---|---|---|
| Simon Drill: Why Async 只記得 "fast response" | 三個好處：decoupling, buffering, resilience | 從 user 角度想得到 fast，但 system architecture 角度的 decoupling 和 buffering 沒內化 |
| Simon Drill: delivery semantics 說成 "most, least, excely" | 完整名稱：at-most-once, at-least-once, exactly-once | 術語要講完整，面試印象差很多 |
| 設計練習不知道怎麼開始 | 拆成小問題一步步推：先決定 sync vs async → 再決定 idempotency 放哪 | 面試也是一樣，不需要一次畫完，一步一步 reason through |

---

## 🗣️ English Practice

| My Answer | English Polish |
|---|---|
| idempotency is use queue maybe send more time message but consumer will just prosses one time | Idempotency means that even if the queue delivers the same message multiple times, the consumer will only process it once — preventing duplicate charges. |
| worker will get dulicate message | The Worker is the one receiving duplicate messages from the queue, so it should be responsible for the idempotency check. |
| payment processing is not good idea, because at-most-once only send massage once maybe will losing massage | At-most-once is terrible for payment processing because if a message is lost, the user paid but their order never gets processed. Duplicates are fixable with idempotency, but lost payments are unrecoverable. |
| use SQS or Kafka can thinks about trade-off like the cost, infra | The choice between SQS and Kafka depends on whether you need message replay. SQS for simple job queues, Kafka when you need event replay or streaming analytics. |

---

## 待完成（下次 session 繼續）

- [ ] Step F: Interview Drill（flash sale order system — 4-step framework）
- [ ] Step G: Notes 補完（Trade-off, Scale trigger, DevOps angle）
- [ ] Step H: Progress Update
