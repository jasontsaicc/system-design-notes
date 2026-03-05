# System Design Interview Preparation Curriculum

> Target: AWS, Google, and other big tech SD interviews
> Pace: ~1 hour/day, deep learning over memorization
> Approach: Discussion → PoC → Notes → Reflect

---

## Phase -1: Go Refresher (Day -5 to Day -1)

> **Goal**: Targeted Go review for building SD proof-of-concept projects
> NOT a generic Go course — focused on HTTP servers, concurrency, data structures, JSON, Docker
> **Time**: 1 week before Day 1, one topic per day

### Day -5: Go Fundamentals — Types, Structs, Error Handling
**Goal**: Model real-world entities in Go
- [x] Variables, types, zero values — why Go is strict about types
- [x] Structs — modeling a `URL` or `User` entity
- [x] Pointers — when and why (receiver methods)
- [x] Error handling — Go's explicit `if err != nil` philosophy
- [x] **Exercise**: In-memory phonebook (struct + map + CRUD functions)
- [x] **Output**: `projects/go-refresher/day01-fundamentals/main.go`
- [x] **Notes**: `notes/go-01-fundamentals.md`

### Day -4: Slices, Maps, Interfaces — Core Data Structures
**Goal**: Comfortable with Go's most-used data structures + polymorphism
- [x] Slices vs arrays — capacity, append, gotchas
- [x] Maps — Go's hash tables (directly relevant to SD!)
- [x] Interfaces — implicit implementation, why it matters
- [x] **Exercise**: Key-value store with `Get/Set/Delete` satisfying a `Store` interface
- [x] **Output**: `projects/go-refresher/day02-data-structures/`
- [x] **Notes**: `notes/go-02-data-structures.md`

### Day -3: Goroutines & Channels — Concurrency Foundations
**Goal**: Understand Go's concurrency model (critical for SD topics)
- [x] Goroutines — lightweight threads, `go` keyword
- [x] Channels — communication between goroutines
- [x] `sync.WaitGroup`, `sync.Mutex` — coordination primitives
- [x] **SD Connection**: Every SD system has concurrent components (multiple users, async processing)
- [x] **Exercise**: Producer-Consumer pattern (maps to Message Queue concepts in Phase 1)
- [x] **Output**: `projects/go-refresher/day03-concurrency/`
- [x] **Notes**: `notes/go-03-concurrency.md`

### Day -2: HTTP Server & JSON — Building APIs in Go
**Goal**: Build a simple REST API (backbone of every SD PoC)
- [x] `net/http` — handlers, routes, middleware concept
- [x] `context.Context` — request-scoped values, cancellation, timeouts (used in every HTTP handler)
- [x] JSON marshal/unmarshal with struct tags
- [x] Request parsing (query params, path params, body)
- [x] **Exercise**: REST API for the KV store (GET/PUT/DELETE `/keys/:key`)
- [x] **Output**: `projects/go-refresher/day04-http-server/`
- [x] **Notes**: `notes/go-04-http-server.md`

### Day -1: Testing & Docker — Making PoCs Real
**Goal**: Test code and containerize it
- [x] `go test` — table-driven tests (Go convention)
- [x] Benchmarking with `go test -bench`
- [x] Dockerfile for Go apps (multi-stage build)
- [x] Docker Compose basics
- [x] **Exercise**: Add tests to KV store API + Dockerize it
- [x] **Output**: `projects/go-refresher/day05-testing-docker/`
- [x] **Notes**: `notes/go-05-testing-docker.md`

---

## Phase 0: Thinking Framework (Day 1-3)

### Day 1: What SD Interviews Actually Test
- [x] Understand the 4 scoring dimensions (Problem Navigation, Design, Technical Depth, Trade-offs)
- [x] Analyze good vs bad answers for the same question
- [x] **Discussion**: Deconstruct a real SD interview rubric
- [x] **Notes**: `notes/day01-interview-rubric.md`

### Day 2: Back-of-Envelope Estimation
- [x] Powers of 2, latency numbers every engineer should know
- [x] Practice: Estimate YouTube daily storage, Twitter QPS
- [x] **PoC**: Build an estimation cheat sheet you actually understand
- [x] **Notes**: `notes/day02-estimation.md`

### Day 3: Your SD Answer Framework (4-Step Method)
- [x] Step 1: Clarify requirements & scope (functional + non-functional)
- [x] Step 2: High-level design (API + data model + architecture)
- [x] Step 3: Deep dive into core components
- [x] Step 4: Scale, trade-offs, monitoring
- [x] **Time budget for 45-min SD interview**:
  - [0-5 min] Clarify requirements & scope (DO NOT SKIP)
  - [5-10 min] Back-of-envelope estimation (if needed)
  - [10-20 min] High-level design (API + data model + architecture diagram)
  - [20-35 min] Deep dive into 1-2 core components (interviewer may guide)
  - [35-45 min] Scale, trade-offs, monitoring, failure handling
- [x] **Handling unknown problems**: Decompose into known building blocks
  - Every system = Storage + API + unique domain logic
  - Practice: "Design a parking lot" → Database + Rate Limiting + API Design
- [x] **Scope negotiation**: "I'll focus on X. Should I also cover Y, or dive deeper into X?"
- [x] **Standard Whiteboard Diagram Template** — use this 8-block skeleton for EVERY design:
  ```
  Client → DNS → LB → API Server → Cache → Database
                                  ↘ Queue → Worker
  ```
  - Every diagram starts with these 8 blocks. Add more for the problem, never remove any.
  - Consistent layout = readable by interviewer in 5 seconds
  - Annotate each arrow with protocol/data: `HTTP`, `gRPC`, `async`, `pub/sub`
  - Highlight the "interesting" blocks for this specific problem (bold/color/star)
- [x] Common diagram components: boxes, arrows, databases, queues, caches
- [x] Practice: Draw "Design URL Shortener" architecture in 5 minutes using the standard template
- [x] **Practice**: Answer "Design URL Shortener" using the framework (dry run)
- [x] **Notes**: `notes/day03-framework.md`

---

## Phase 1: Core Building Blocks (Day 4-16)

> **📡 Observability Mini** — Apply to EVERY topic in Phase 1. This is your DevOps competitive advantage.
>
> | Element | What to define |
> |---------|---------------|
> | **SLIs** | Availability, latency (P50/P99), error rate for this component |
> | **SLO target** | e.g., 99.9% availability, P99 < 200ms |
> | **Alerts** | Burn-rate on error budget; saturation (CPU/mem/connections) |
> | **Dashboards** | 3 graphs: throughput (QPS), latency distribution, error rate |

### Day 4-5: Load Balancer & Reverse Proxy
- [x] DNS fundamentals — resolution flow, TTL, A/AAAA/CNAME records
- [x] DNS-based load balancing: Route 53 weighted, latency-based, failover routing
- [x] L4 vs L7 load balancing — when to use which
- [x] Algorithms: Round Robin, Least Connections, IP Hash, Weighted
- [x] Health checks, sticky sessions, connection draining
- [x] 🔧 **DevOps Angle**: ALB vs NLB real-world differences, target group health checks
- [x] **📡 Observability Mini**: SLI = backend healthy %; SLO = 99.95% avail; Alert = burn-rate on 5xx; Dashboard = active connections, request rate, response time
- [x] **PoC**: Nginx L4/L7 LB with Docker Compose, observe behavior difference
- [x] **Notes**: `notes/day04-05-load-balancer.md`

### Day 6-7: Caching & CDN Strategies
- [ ] Cache levels: Browser → CDN → App → DB
- [ ] Patterns: Cache-Aside, Write-Through, Write-Behind, Read-Through
- [ ] Eviction: LRU, LFU, TTL
- [ ] Cache invalidation — the hard problem
- [ ] CDN deep dive: Edge caching vs Origin pull
- [ ] Pull CDN vs Push CDN — trade-off
- [ ] CDN cache invalidation strategies (TTL, purge, versioned URLs)
- [ ] 🔧 **DevOps Angle**: ElastiCache cluster mode vs non-cluster, operational gotchas
- [ ] 🔧 **DevOps Angle**: CloudFront behaviors, origin shield, cache hit ratio monitoring
- [ ] **📡 Observability Mini**: SLI = cache hit ratio, origin latency; SLO = hit ratio > 90%, P99 < 50ms; Alert = hit ratio drop > 10% in 5min; Dashboard = hit/miss ratio, eviction rate, memory usage
- [ ] **PoC**: Redis cache layer, measure latency with/without cache
- [ ] **Notes**: `notes/day06-07-caching.md`

### Day 8-9: Database Selection
- [ ] SQL (PostgreSQL) vs NoSQL (DynamoDB/MongoDB) vs NewSQL (CockroachDB)
- [ ] When to pick what — decision framework
- [ ] Indexing deep dive: B-tree (read-optimized, PostgreSQL) vs LSM-tree (write-optimized, Cassandra/RocksDB)
- [ ] Denormalization — when and why to sacrifice normalization for performance
- [ ] Sharding intro — hash-based vs range-based (Day 15-16 solves resharding with consistent hashing)
- [ ] Connection pooling — why and how (pgbouncer, HikariCP concepts)
- [ ] Write-Ahead Log (WAL) — how databases ensure durability
- [ ] Read replicas — scaling reads vs consistency trade-off
- [ ] **📐 Data Model Design Template** (apply to every Phase 3 problem):
  - **Entities**: Core domain objects
  - **Access patterns**: Top 3-5 read/write queries
  - **Partition key**: What distributes data evenly?
  - **Secondary index**: Which queries need GSI?
  - **Hot partition risk**: Celebrity/viral/time-bucket skew — mitigation?
  - **Backfill & migration**: Schema evolution without downtime?
- [ ] **📡 Observability Mini**: SLI = query latency, replication lag; SLO = P99 read < 10ms, lag < 1s; Alert = lag > 5s, connection pool exhaustion; Dashboard = QPS by query type, slow query count, pool utilization
- [ ] **PoC**: Same problem (e.g., user feed) with SQL vs NoSQL, compare trade-offs
- [ ] **Notes**: `notes/day08-09-database.md`

### Day 10-11: Message Queue & Async Processing
- [ ] Why async? Decoupling, buffering, peak handling
- [ ] SQS vs Kafka vs RabbitMQ — positioning
- [ ] At-least-once, at-most-once, exactly-once semantics
- [ ] Dead letter queue, retry strategies
- [ ] 🔧 **DevOps Angle**: SQS FIFO vs Standard, DLQ monitoring in CloudWatch
- [ ] **📡 Observability Mini**: SLI = processing latency, DLQ depth; SLO = 99.9% processed < 30s; Alert = DLQ depth > 0, consumer lag > threshold; Dashboard = in-flight msgs, processing time, DLQ count
- [ ] **PoC**: Producer-consumer with Redis/SQS, simulate failures
- [ ] **Notes**: `notes/day10-11-message-queue.md`

### Day 12-13: API Design
- [ ] REST vs gRPC vs GraphQL — trade-off matrix
- [ ] Pagination: Offset vs Cursor
- [ ] Versioning strategies
- [ ] Idempotency in API design
- [ ] **📡 Observability Mini**: SLI = API success rate, P99 per endpoint; SLO = 99.9% success, P99 < 300ms; Alert = error spike > 1%, P99 > 500ms; Dashboard = QPS by endpoint, latency heatmap, error by status
- [ ] **PoC**: Design & implement a small API, discuss design choices
- [ ] **Notes**: `notes/day12-13-api-design.md`

### Day 14: Security & Authentication Patterns
- [ ] Authentication: JWT vs Session-based — stateless vs stateful trade-off
- [ ] OAuth 2.0 flow basics (Authorization Code Grant)
- [ ] API Key management — when to use vs token auth
- [ ] HTTPS, encryption at rest / in transit
- [ ] SD Interview context: "How do you authenticate users in this system?"
- [ ] 🔧 **DevOps Angle**: AWS Cognito, ALB OIDC integration, secrets rotation
- [ ] **📡 Observability Mini**: SLI = auth success rate, token latency; SLO = 99.99% auth avail; Alert = auth failure spike, signing latency > 100ms; Dashboard = login success/failure, OAuth completion, token refresh rate
- [ ] **Notes**: `notes/day14-security-auth.md`

### Day 15-16: Consistent Hashing & Data Partitioning
- [ ] Why simple hash mod N fails
- [ ] Consistent hashing with virtual nodes
- [ ] Range-based vs Hash-based partitioning
- [ ] **PoC**: Implement consistent hashing algorithm from scratch
- [ ] **📡 Observability Mini**: SLI = key distribution variance, rebalance duration; SLO = max skew < 15%; Alert = node join/leave, skew > 20%; Dashboard = keys per node, virtual node distribution, rebalance progress
- [ ] **Notes**: `notes/day15-16-consistent-hashing.md`

### 🎯 Checkpoint: Mini-Mock (Day 16 afternoon)
> **15-minute drill**: Pick any building block from Phase 1. Explain it using the 4-step framework as if in an interview.
> **Self-check**: Can you explain WHY you'd pick this component, not just WHAT it does?

---

## Phase 2: Distributed Systems Core (Day 17-26)

### Day 17-18: CAP Theorem in Practice
- [ ] CAP is about network partitions — not a daily choice
- [ ] Real-world: DynamoDB (AP), Zookeeper (CP), Spanner (technically CP but...)
- [ ] PACELC model as a more practical framework
- [ ] **Discussion**: Analyze real AWS services through CAP lens
- [ ] **Notes**: `notes/day17-18-cap-theorem.md`

### Day 19-20: Consistency Models
- [ ] Strong, Eventual, Causal, Read-your-writes
- [ ] Quorum: W + R > N
- [ ] Vector clocks, conflict resolution
- [ ] **PoC**: Simulate eventual consistency behavior
- [ ] **Notes**: `notes/day19-20-consistency.md`

### Day 21-22: Replication & Leader Election
- [ ] Single-leader, multi-leader, leaderless replication
- [ ] Raft consensus algorithm (simplified)
- [ ] Redis Sentinel vs Cluster — real trade-offs
- [ ] **Discussion**: When do you need consensus? When is it overkill?
- [ ] Service Discovery basics: Consul, Kubernetes DNS, AWS Cloud Map — how services find each other
- [ ] **Notes**: `notes/day21-22-replication.md`

### Day 23-24: Rate Limiting Algorithms (Local, Single-Node) & Circuit Breaker
- [ ] Token Bucket, Sliding Window, Fixed Window — algorithm-level understanding
- [ ] Distributed rate limiting **challenges** (to be solved as a system in Phase 3)
- [ ] Circuit Breaker pattern (Closed → Open → Half-Open)
- [ ] Bulkhead pattern
- [ ] **PoC**: Implement Token Bucket + Circuit Breaker in Go
- [ ] **Notes**: `notes/day23-24-rate-limiting.md`

### 📋 Reference: Distributed Systems Kill Pack (No Extra Day — Just a Cheat Sheet)

> These 3 concepts appear in almost every Phase 3 problem. Study them here, apply them everywhere.
> Save as `notes/distributed-kill-pack.md`

- [ ] **Idempotency Key**:
  - Scope: per-user? per-request? per-API-call?
  - TTL: How long to remember? (24h for payments, 5min for retries)
  - Storage: Redis with TTL? DB unique constraint?
  - Behavior: Same key → return cached response, never re-execute side effects
- [ ] **Deduplication**:
  - Message ID: Who generates it? Client vs server vs infra (SQS MessageDeduplicationId)
  - Dedup window: 5 min? Exactly-once is an illusion — at-least-once + idempotent handler is reality
  - Where to dedup: Consumer side (check before process) vs infra side (SQS FIFO, Kafka compaction)
- [ ] **Distributed Lock**:
  - When needed: Leader election, exclusive resource access, preventing double-spend
  - Lease-based: Lock with TTL, auto-release on crash (Redlock, DynamoDB conditional write)
  - Fencing token: Monotonic counter to detect stale lock holders
  - Failure modes: Split-brain, clock skew, GC pause — what happens when lock holder dies?
  - When NOT to lock: Prefer optimistic concurrency (CAS/versioning) for high-throughput paths

### Day 25: Observability Consolidation & SD Interview Strategy (Practiced Since Day 4)
- [ ] Review all Phase 1 Observability Mini entries — identify cross-component patterns
- [ ] Metrics, Logs, Traces — three pillars (consolidation, not first exposure)
- [ ] Advanced: Distributed tracing across microservices (Jaeger, X-Ray)
- [ ] SLI/SLO/SLA — formalize definitions, error budget calculation
- [ ] Structured logging: correlation IDs, log levels, searchability
- [ ] How to discuss monitoring in SD interview: lead with SLIs → dashboards → alerting philosophy
- [ ] 🔧 **DevOps Angle**: Prometheus/Grafana/ELK, SLO-based alerting, on-call runbook design — your real experience
- [ ] **Exercise**: Pick one Phase 1 topic, build complete observability story (SLI → SLO → alert → dashboard → runbook)
- [ ] **Notes**: `notes/day25-observability.md`

### Day 26: Bloom Filter, Gossip Protocol & Advanced Distributed Concepts
- [ ] **Bloom Filter**: Space-efficient probabilistic membership testing
  - Use cases: Web crawler URL dedup, cache key checking, spell checker
  - False positive rate vs space trade-off
- [ ] **Gossip Protocol**: How nodes discover and share state
  - Use cases: Cassandra membership, DynamoDB failure detection
  - Epidemic vs anti-entropy propagation
- [ ] **Distributed Transactions (overview)**: 2PC/3PC basics (deep dive in Payment System)
- [ ] **Notes**: `notes/day26-advanced-distributed.md`

### 🎯 Checkpoint: Mid-Mock (Day 26 afternoon)
> **30-minute mock**: Design a complete system using Phase 1 building blocks + Phase 2 distributed concepts.
> **Suggested topic**: "Design a distributed key-value store" (ties together caching, hashing, replication, consistency)

---

## Notes Template（所有 Phase 適用）

> 每份筆記 (`notes/dayXX-*.md`) 使用以下 template。元素依 Phase 逐步增加，不需要一開始就填完。

### 基礎元素（所有 Phase）

| Element | Format | Example |
|---------|--------|---------|
| **One-liner** | "X is a system that..." | "A URL shortener maps long URLs to short codes for sharing" |
| **Trade-off** | "We chose X over Y because..." | "We chose base62 over MD5 because we need shorter URLs and can handle collisions" |
| **Scale trigger** | "At N scale, we need..." | "At 10K writes/sec, we need database sharding to avoid single-node bottleneck" |
| **DevOps angle** | "In production, I'd monitor..." | "I'd set up CloudWatch alarms on 4xx/5xx rates and P99 latency" |

### 進階元素（Phase 1+ PoC topics）

- **Capacity & cost estimation** _(practice the numbers story)_:
  - **Traffic assumptions**: DAU, peak QPS, payload size
  - **Storage growth**: Daily new data (GB), yearly total (TB)
  - **Cost top 2**: Compute vs storage vs data transfer — who's biggest, why, how to cut 30%
  - Example: "10M DAU, 1K QPS peak, 500 GB/day images. Biggest cost = S3 storage; lifecycle policy saves 40%"

### 完整元素（Phase 3+）

- **Failure modes**: Address at least 3:
  - [ ] Dependency down (upstream/downstream fails)
  - [ ] Latency spike (P99 blows up)
  - [ ] Partial outage (one AZ or shard down)
  - [ ] Message duplication (at-least-once side effect)
  - [ ] Data corruption (write conflict, stale cache as truth)
  - [ ] Thundering herd (cache stampede, retry storm)
- **Abuse & security**: List at least 3 attack vectors + countermeasures:
  - [ ] Credential stuffing → device fingerprint / CAPTCHA / adaptive rate limit
  - [ ] Replay attack → nonce / idempotency key / timestamp window
  - [ ] Spam / bot abuse → reputation score / greylisting / content moderation pipeline
  - [ ] Data exfiltration → field-level encryption / audit log / access control
  - [ ] DDoS / resource exhaustion → WAF / auto-scaling / circuit breaker
  - Example: "Replay attacks on payment API; we use idempotency keys with a 24h TTL window"

### 每次必填

- **🔴 My Mistakes & Misconceptions** — Record errors and wrong intuitions from the session:

  | What I Thought | Reality | Why I Was Wrong |
  |---|---|---|

---

## Phase 3: Classic SD Problems (Day 27-53)

> Format: Day 1 = Discussion & Design | Day 2 = PoC + Whiteboard Diagram + Notes
> Complex problems get 3 days: Day 1 = Discussion | Day 2 = Design + Diagram | Day 3 = PoC
>
> **📐 Data Model Reminder**: For every problem, fill out the template from Day 8-9:
> Entities → Access patterns → Partition key → Secondary index → Hot partition risk → Backfill strategy
>
> **🎨 Diagram Reminder**: Start every whiteboard with the 8-block skeleton from Day 3:
> Client → DNS → LB → API → Cache → DB → Queue → Worker. Add more, never remove.
>
> **🔩 PoC Production Hooks** — Every PoC must include these 3 non-negotiables:
> 1. **Metrics endpoint**: `/metrics` or log latency distribution (P50/P99/P999) per request
> 2. **Failure injection**: A flag/env var to simulate timeout, retry storm, or duplicate messages
> 3. **Load test script**: A one-liner with `vegeta` or `hey` (e.g., `echo "GET http://localhost:8080/api" | vegeta attack -rate=1000 -duration=10s | vegeta report`)
>    - Run baseline → enable failure injection → compare P99 and error rate
>    - Save output as `load-test-results.md` in each project dir
> Goal: Be able to say "I hit 1K RPS, P99 was 12ms, then I turned on failure injection and saw retry storm push P99 to 800ms." Not toy code.

---

### ⭐ Tier 1: Must Do (8 problems, 19 days, Day 27-45)

> These 8 problems cover 90% of core SD concepts. Master these first.

#### Day 27-28: URL Shortener ★★☆
- Key concepts: Hashing, base62, collision handling, read-heavy system
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/url-shortener/`
- [ ] **Notes**: `notes/day27-28-url-shortener.md`

#### Day 29-30: Unique ID Generator ★★☆ ⭐NEW
- Key concepts: Snowflake algorithm, coordination-free generation, clock synchronization
- Why it matters: Sub-problem in almost EVERY other SD design
- Clock-specific challenges (interviewers WILL drill these):
  - [ ] **Clock drift**: NTP sync failure, monotonic vs wall clock — what if clock goes backward?
  - [ ] **Multi-region ordering**: IDs from US-East vs EU-West — global ordering possible? At what cost?
  - [ ] **Sequence overflow**: Sequence bits wrap within same millisecond — what happens?
  - [ ] **Worker ID allocation**: Zookeeper? Config? What if two workers get same ID?
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/id-generator/`
- [ ] **Notes**: `notes/day29-30-id-generator.md`

#### Day 31-32: Distributed Rate Limiter (Multi-Node, Redis-Backed) ★★★
- Key concepts: Distributed sliding window, Redis Lua scripts, race conditions
- Note: Builds on Day 23-24 algorithm knowledge — now design the **distributed service**
- Multi-tenant & quota design:
  - [ ] **Hierarchical quotas**: per-user / per-API-key / per-IP / per-org — how do limits compose?
  - [ ] **Burst vs sustained**: Token bucket burst vs sliding window smooth — which fits which?
  - [ ] **Soft limit vs hard limit**: 429 reject vs queue-and-delay (throttle) — UX vs protection trade-off
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/rate-limiter/`
- [ ] **Notes**: `notes/day31-32-rate-limiter.md`

#### Day 33-34: Notification System ★★★
- Key concepts: Push vs Pull, priority queue, retry, multi-channel (SMS/Email/Push)
- 🔧 **DevOps Angle**: SQS/SNS integration, retry with exponential backoff, DLQ monitoring
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/notification-system/`
- [ ] **Notes**: `notes/day33-34-notification.md`

#### Day 35-37: Chat System (WhatsApp-like) ★★★★ (3 days)
- Key concepts: WebSocket, message storage, presence, read receipts, group chat fan-out
- [ ] Real-time transport comparison: WebSocket vs SSE vs Long Polling — decision matrix
- [ ] Day 35: **Discussion** — requirements, API, data model
- [ ] Day 36: **Design** + **Whiteboard Diagram** — connection management, message flow
- [ ] Day 37: **PoC**: `projects/chat-system/`
- [ ] **PoC Scope**: WebSocket echo server + broadcast to all clients (~60 lines Go)
- [ ] **Notes**: `notes/day35-37-chat.md`

#### Day 38-39: Distributed Cache ★★★
- Key concepts: Consistent hashing (revisit!), cache invalidation at scale, thundering herd
- 🔧 **DevOps Angle**: ElastiCache cluster operations, failover behavior, eviction monitoring
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/distributed-cache/`
- [ ] **Notes**: `notes/day38-39-distributed-cache.md`

#### Day 40-42: News Feed (Twitter/Facebook) ★★★★ (3 days)
- Key concepts: Fan-out on write vs read, ranking algorithm, timeline generation, cache warming
- [ ] Day 40: **Discussion** — fan-out strategies, trade-offs at different scales
- [ ] Day 41: **Design** + **Whiteboard Diagram** — celebrity problem, hybrid approach
- [ ] Day 42: **PoC**: `projects/news-feed/`
- [ ] **PoC Scope**: In-memory fan-out-on-write simulation (~50 lines Go)
- [ ] **Notes**: `notes/day40-42-news-feed.md`

#### Day 43-45: Payment System ★★★★ (3 days)
- Key concepts: Idempotency, SAGA pattern, exactly-once processing, reconciliation
- Uses: Distributed Transactions (2PC from Day 26), eventual consistency (Day 19-20)
- [ ] Day 43: **Discussion** — payment flow, failure modes, idempotency keys
- [ ] Day 44: **Design** + **Whiteboard Diagram** — SAGA orchestration vs choreography
- [ ] Day 45: **PoC**: `projects/payment-system/`
- [ ] **PoC Scope**: Idempotency key handler — same response for duplicate requests (~40 lines Go)
- [ ] **Notes**: `notes/day43-45-payment.md`

---

### ⭐ Tier 2: Should Do (4 problems, 8 days, Day 46-53)

> Important topics with slightly lower interview frequency. Your DevOps background makes Metrics & Logging a **strong scoring opportunity**.

#### Day 46-47: Metrics & Logging System 🔧 YOUR DEVOPS ADVANTAGE
- Key concepts: Time-series DB, aggregation, sampling, data pipeline
- 🔧 **DevOps Angle**: This is YOUR territory — Prometheus, Grafana, ELK, CloudWatch
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/metrics-system/`
- [ ] **Notes**: `notes/day46-47-metrics.md`

#### Day 48-49: Search Autocomplete ★★★
- Key concepts: Trie, ranking, Elasticsearch, prefix matching at scale
- Extended pipeline (interviewers probe beyond Trie):
  - [ ] **Hot term updates**: Real-time trending injection vs batch rebuild — freshness vs compute
  - [ ] **Typo tolerance**: Edit distance (Levenshtein), phonetic matching — where in pipeline?
  - [ ] **Tokenization & language**: CJK segmentation, stemming, stop words — trie structure impact
  - [ ] **Offline build + online serving**: Periodic trie rebuild from analytics → deploy; online handles trending overlay
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/search-autocomplete/`
- [ ] **Notes**: `notes/day48-49-autocomplete.md`

#### Day 50-51: Web Crawler ★★★ ⭐NEW
- Key concepts: BFS/DFS crawl strategy, URL frontier, politeness, dedup (Bloom filter!)
- Connects to: Bloom Filter (Day 26), Distributed Task scheduling concepts
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/web-crawler/`
- [ ] **Notes**: `notes/day50-51-web-crawler.md`

#### Day 52-53: Proximity Service (Nearby) ★★★ ⭐NEW
- Key concepts: Geohash, QuadTree, spatial indexing, location-based search
- Common in: Uber, Yelp, Google Maps style interviews
- [ ] **Design** + **Whiteboard Diagram**
- [ ] **PoC**: `projects/proximity-service/`
- [ ] **Notes**: `notes/day52-53-proximity.md`

---

### 📦 Tier 3: Nice to Have (Skip unless extra time)

> These topics are either too domain-specific or their core concepts are already covered by Tier 1/2 problems.

| Topic | Why Deprioritized | Covered By |
|-------|-------------------|------------|
| Video Streaming (YouTube/Netflix) | Domain-specific (transcoding, adaptive bitrate) | CDN in caching, async in notification |
| Cloud Storage (Google Drive) | Domain-specific (file sync, chunking) | Consistency in distributed cache, dedup in crawler |
| Distributed Task Scheduler | Concepts overlap | Notification (priority queue) + Payment (distributed locking) |
| Ticket/Booking System | Concurrency focus | Payment (idempotency) + Rate Limiter (distributed counting) |

---

## Phase 4: Advanced & Mock Interviews (Day 54-61)

### Day 54-55: Trade-off Analysis Deep Dive
- [ ] Practice specific trade-off scenarios (5 min each side, then state your decision):
  - SQL vs NoSQL: "You chose DynamoDB, but now need cross-entity JOINs?"
  - Push vs Pull: "Fan-out-on-write, but a user has 10M followers?"
  - Consistency vs Availability: "Eventual consistency in chat — user sends message but doesn't see it?"
  - Sync vs Async: "Payment confirmation in a queue — user stares at loading screen. SLA?"
  - Cache vs Fresh Data: "Cache TTL 5 min, but stale prices cost money?"
  - Cost vs Performance: "50 Redis nodes, finance says cut 40% cost. What do you sacrifice?"
- [ ] **🪤 Trap & Pivot Drills** — Practice graceful pivots when your initial design hits a wall:
  - "Design a chat system" → You pick polling → Interviewer: "Now support 1M concurrent users" → Pivot to WebSocket
  - "Design a URL shortener" → You pick auto-increment ID → Interviewer: "Now make it multi-region" → Pivot to distributed ID (Snowflake)
  - "Design a payment system" → You pick synchronous flow → Interviewer: "Timeout at 30s, but payment takes 2 min" → Pivot to async + webhook callback
  - **How to pivot gracefully**: "Good point — at this scale, [original] breaks because [specific reason]. Let me switch to [alternative] which handles [new constraint] because [technical why]."
  - **Anti-pattern**: Don't say "oh you're right, let me start over." Instead, treat it as **evolution**, not replacement
- [ ] "What if requirements change?" exercises — revisit 2 designs with new constraints
- [ ] Cost estimation for your designs (ties in with your AWS/DevOps background)
- [ ] 🔧 **DevOps Angle**: Multi-AZ cost, data transfer pricing, reserved vs on-demand

### Day 56-57: Mock Interview Round 1 (Recorded)
- [ ] 45-minute **strictly timed** mock interview (I play interviewer)
- [ ] **Record the entire session** (screen + voice) — non-negotiable
- [ ] Detailed feedback and scoring against the 4 dimensions from Day 1
- [ ] **Self-review**: Watch at 1.5x — note every pause >10s, every vague statement, every missed trade-off
- [ ] Targeted improvement plan based on recording review
- [ ] Practice "thinking aloud" — narrate thought process, never go silent >10 sec
- [ ] **Trap mode**: Interviewer will deliberately agree with one suboptimal choice, then later add a constraint that breaks it — practice the pivot
- [ ] Recognize interviewer hints — if they redirect, follow the signal
- [ ] Recovering from mistakes: "Actually, I realize this has a problem with X. Let me reconsider..."

### Day 58-59: Weak Spot Reinforcement
- [ ] Review all notes, identify patterns in mistakes
- [ ] Re-do 2-3 designs where you struggled
- [ ] Practice "Interview Language" — articulate trade-offs clearly in 2-3 sentences

### Day 60-61: Final Mock Interview (Brutal Mode)
- [ ] Full simulation: 2 back-to-back SD interviews, **45 min each, strictly timed**
- [ ] **Record both sessions** — no exceptions
- [ ] Interviewer will interrupt, change requirements mid-design, challenge trade-offs aggressively
- [ ] **Double trap**: Interviewer sets two sequential traps — first design choice leads to a second trap. Practice chaining pivots without losing composure
- [ ] **Self-review**: Watch both, compare with Day 56-57 recording — measure improvement
- [ ] Score yourself on 4 dimensions before reading feedback
- [ ] Final review and interview-day tips
- [ ] 🔧 **DevOps Angle**: Practice weaving in operational experience naturally

---

## Language & Tools

- **PoC Language**: **Go** — chosen for concurrency model, performance, and interview relevance
- **Diagrams**: Mermaid + hand-drawn whiteboard practice
- **Infrastructure**: Docker Compose for local PoC environments
- **Notes**: Markdown in `notes/` directory

---

## Daily Routine

```
A. [5 min]  複習 — Review yesterday's notes + check mistakes
B. [3 min]  引入 — Daily topic intro (analogy / scenario)
C. [12 min] 核心教學 — Topic discussion, Feynman method (理解確認穿插其中)
D. [20 min] 動手做 — Hands-on PoC or design exercise
E. [5 min]  Voice Drill — Record yourself: Recall 元素 (play back, refine)
F. [5 min]  整理筆記 — Write notes (using Notes Template)
G. [5 min]  更新進度 — Update progress + preview tomorrow
     Total ≈ 55 min（留 5 min buffer）
```

---

## Weekly Review（每週六，30 min）

> 防止遺忘的間隔複習機制。課程越往後，早期 building blocks 越容易忘。
> **核心原則：先回憶，再驗證。**不翻 notes 就能講出來的才是真正學會的。

### Recall 元素（依 Phase 調整）

| Phase | 必答元素 | 滿分 |
|-------|---------|------|
| Phase 0 | One-liner, Trade-off | 2/2 |
| Phase 1 | One-liner, Trade-off, Scale trigger, DevOps angle | 4/4 |
| Phase 2 | 同 Phase 1 + Capacity | 5/5 |
| Phase 3+ | 全部 6 元素（+ Failure modes, Security） | 6/6 |

- 目標：滿分 = 已掌握，≤ 滿分的 2/3 = 下週重點複習

### Step 1：Blind Recall（15 min）
- 隨機挑 3 個已學 topics（本週 1 個 + 過去 2 個）
- **不看 notes**，嘗試口述每個 topic 的 Recall 元素（依上方 Phase 表格）
- 計分：每個元素能講出來 = ✅，卡住 = ❌

### Step 2：Gap Check（10 min）
- 打開 notes 對答案，找出遺漏和錯誤
- 特別注意：你以為對但其實錯的（最危險的盲點）
- 標記需要加強的概念

### Step 3：Quick Drill（5 min）
- 挑一個本週最弱的概念，重新口述一次直到流暢
- Phase 3 開始後：加入 1 題 Phase 3 問題的 whiteboard 30-second sketch（不看 notes，純記憶畫出架構）

---

## Resources (Reference, not primary learning)

- "System Design Interview" by Alex Xu (Vol 1 & 2)
- "Designing Data-Intensive Applications" by Martin Kleppmann (DDIA)
- ByteByteGo YouTube channel
- Real-world architecture blogs: AWS Architecture Blog, Google Cloud Blog
