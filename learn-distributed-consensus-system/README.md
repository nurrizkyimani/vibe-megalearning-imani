# M3/W1/D7 - Sat, 07 Mar 2026 (WIB)

## Distributed Consensus System — Full Curriculum

A comprehensive learning path covering both **distributed systems fundamentals** and  
**consensus mechanisms** for senior SWE roles at companies like Anthropic and OpenAI.  
Each topic is a runnable Go simulation with rich inline comments. The curriculum spans  
24 topics covering time & ordering, failure detection, quorums, consensus algorithms  
(Raft/Paxos), distributed transactions, CRDTs, and production operations. Topics build  
on each other — read in order.

---

## Curriculum Table

| # | Topic | Core Concept | Real World | Demo |
| --- | --- | --- | --- | --- |
| 01 | Consistent Hashing | Distribute data across nodes without reshuffling | DynamoDB, Redis Cluster, Cassandra, CDN | Num01 ✅ |
| 02 | Leader Election (Raft) | Cluster agrees on one coordinator, survives crashes | etcd, Kubernetes, CockroachDB, Consul | Num02 ✅ |
| 03 | Distributed Queues | Async messaging, partitions, consumer groups | Kafka, Pub/Sub, Kinesis, NATS | Num03 ✅ |
| 04 | Eventual vs Strong Consistency | CAP theorem tradeoff, convergence | DynamoDB, Spanner, Cassandra | Num04 ✅ |
| 05 | Replication Strategies | Primary-replica, multi-leader, leaderless, quorum | MySQL, CouchDB, Cassandra | Num05 ✅ |
| 06 | Distributed Transactions | 2PC, Saga pattern, atomic cross-service writes | Payment systems, order fulfillment | Num06 ✅ |
| 07 | Rate Limiting | Token bucket, leaky bucket, sliding window | Every public API, DDoS protection | Num07 ✅ |
| 08 | Circuit Breaker | Fault tolerance, cascading failure prevention | Netflix Hystrix, microservice mesh | Num08 ✅ |
| 09 | Gossip Protocol | Decentralized state propagation, O(log N) convergence | Cassandra, Consul, Bitcoin | Num09 ✅ |
| 10 | Distributed Locking | Redlock, fencing tokens, coordination | Payments dedup, job coordination | Num10 🔲 |
| 11 | Observability | Tracing consensus RPCs, split-brain detection | Prometheus, Jaeger, OpenTelemetry | Num11 🔲 |
| 12 | Sharding Strategies | Per-shard Raft groups, range partitioning | TiKV, CockroachDB, Instagram, YouTube | Num12 🔲 |
| 13 | Event Sourcing & CQRS | Distributed log as consensus output | Event-driven microservices, Kafka Streams | Num13 🔲 |
| 14 | Consensus Algorithms | Raft vs Paxos vs ZAB comparison | etcd, Consul, ZooKeeper | Num14 🔲 |
| 15 | CAP Theorem | Partition tolerance, CP vs AP trade-offs, PACELC | System design interviews, architecture decisions | Num15 🔲 |
| 16 | Logical Time & Clocks | Lamport timestamps, vector clocks, causality | Debugging race conditions, DynamoDB versioning | Num16 🔲 |
| 17 | Failure Detection | Phi Accrual detector, SWIM indirect probing | Cassandra, Consul Memberlist | Num17 🔲 |
| 18 | Quorums & Consistency | N/R/W models, sloppy quorums, hinted handoff | Tuning Riak/Dynamo latency vs consistency | Num18 🔲 |
| 19 | Atomic Broadcast | Total order broadcast, FIFO channels, ZAB | Theoretical basis of ZooKeeper/etcd | Num19 🔲 |
| 20 | Raft: Edge Cases | Log compaction, snapshots, membership changes | Productionizing Raft (preventing disk fill-up) | Num20 🔲 |
| 21 | Paxos (Synod) | Propose/Accept/Learn, single-value problem | Google Chubby, Spanner (Paxos Made Live) | Num21 🔲 |
| 22 | Multi-Paxos | Leader stickiness, optimizing round trips | High-throughput WAN consensus replication | Num22 🔲 |
| 23 | CRDTs (Conflict-Free) | G-Counter, PN-Counter, LWW-Register | Collaborative editing (Figma), Chat (Discord) | Num23 🔲 |
| 24 | Linearizability Testing | Jepsen testing basics, history verification | Proving your distributed system actually works | Num24 🔲 |

**Legend:** ✅ = Fully implemented | 🔲 = Skeleton stub (to be implemented)

**Cross-references for new curriculum topics:**

*   **Logical Time & Clocks** → See Num16
*   **Failure Detection** → See Num17 (also: heartbeats in Num09 Gossip)
*   **Quorums & Consistency** → See Num18 (also: quorum basics in Num04)
*   **CAP Theorem** → See Num15 (also: discussed in Num04)
*   **Atomic Broadcast** → See Num19
*   **Leader Election (Basic)** → See Num02 (Raft-based election covers this)
*   **Raft Basics** → See Num02 (leader election + log replication)
*   **Raft Edge Cases** → See Num20
*   **Paxos Synod** → See Num21
*   **Multi-Paxos** → See Num22
*   **Two-Phase Commit** → See Num06 (2PC fully covered)
*   **Sagas** → See Num06 (Saga pattern fully covered)
*   **Gossip** → See Num09
*   **CRDTs** → See Num23
*   **Linearizability Testing** → See Num24

---

## How to run

```
# run all topics (only uncommented ones in main.go will execute)
go run ./learn-distributed-consensus-system/
```

To run a specific topic, comment/uncomment the relevant line in `main.go`.

---

## Function signatures

```
func Num01ConsistentHashingDemo()          // ✅ fully implemented
func Num02LeaderElectionDemo()             // ✅ fully implemented
func Num03DistributedQueuesDemo()          // ✅ fully implemented
func Num04ConsistencyModelsDemo()          // ✅ fully implemented
func Num05ReplicationStrategiesDemo()      // ✅ fully implemented
func Num06DistributedTransactionsDemo()    // ✅ fully implemented
func Num07RateLimitingDemo()               // ✅ fully implemented
func Num08CircuitBreakerDemo()             // ✅ fully implemented
func Num09GossipProtocolDemo()             // ✅ fully implemented
func Num10DistributedLockingDemo()         // 🔲 stub — to be implemented
func Num11ObservabilityDemo()              // 🔲 stub — to be implemented
func Num12ShardingStrategiesDemo()         // 🔲 stub — to be implemented
func Num13EventSourcingDemo()              // 🔲 stub — to be implemented
func Num14ConsensusAlgorithmsDemo()        // 🔲 stub — to be implemented
func Num15CAPTheoremDemo()                 // 🔲 stub — to be implemented
func Num16LogicalClocksDemo()              // 🔲 stub — to be implemented
func Num17FailureDetectionDemo()           // 🔲 stub — to be implemented
func Num18QuorumsConsistencyDemo()         // 🔲 stub — to be implemented
func Num19AtomicBroadcastDemo()            // 🔲 stub — to be implemented
func Num20RaftEdgeCasesDemo()              // 🔲 stub — to be implemented
func Num21PaxosSynodDemo()                 // 🔲 stub — to be implemented
func Num22MultiPaxosDemo()                 // 🔲 stub — to be implemented
func Num23CRDTsDemo()                      // 🔲 stub — to be implemented
func Num24LinearizabilityTestingDemo()     // 🔲 stub — to be implemented
```

---

---

## Num01 — Consistent Hashing with Virtual Nodes

### The Problem

You have N servers and need to decide which server stores a given key.

**Naive approach:** `server = hash(key) % N`

The problem: when N changes (add or remove a server), almost **every** key remaps to a different server. This invalidates your entire cache and causes massive data rebalancing.

### The Concept

Imagine a circle (ring) with positions `0..2^32-1`.

*   Each server gets one or more positions on the ring by hashing its name
*   Each key gets a position by hashing its value
*   A key is owned by the **first server clockwise** from the key's position

```
         0
         |
   NodeC#1 (416M)
         |
   NodeC#0 (1086M)
   NodeC#2 (1088M)
   NodeA#0 (1159M)
         |
   NodeA#2 (2522M)
   NodeB#0 (2629M)
   NodeA#1 (2803M)
   NodeB#1 (2982M)
         |
   NodeB#2 (4018M)
         |
      4294967295 (wraps back to 0)
```

When you **add** a server: only the keys between the new server and its predecessor move.  
When you **remove** a server: only that server's keys move to its successor.  
All other keys are completely unaffected.

### Virtual Nodes

With only 1 position per server, load can be uneven — one server might own a large arc  
of the ring and handle 60% of all keys. Virtual nodes solve this:

*   Each physical server gets `replicas` positions on the ring (e.g. `NodeA#0`, `NodeA#1`, `NodeA#2`)
*   More positions = more even distribution of keys
*   In production (Cassandra, DynamoDB): **150–200 virtual nodes per server**

### What the demo shows

```
Step 1: Build ring with 3 nodes (NodeA, NodeB, NodeC) — 9 virtual positions total
Step 2: Route 6 sample keys — each finds its closest clockwise node
Step 3: Add NodeD — insert 3 more positions
Step 4: Route same 6 keys — only keys whose arc now has NodeD between them and
        their previous owner move. Others are completely unchanged.
Step 5: Remove NodeB — its 3 positions removed from ring
Step 6: Route same 6 keys — only NodeB's former keys redistribute. Others unchanged.
```

### Key Insight

```
Naive hashing (key % N):  add 1 server to a 3-server cluster → ~75% of keys move
Consistent hashing:       add 1 server to a 3-server cluster → ~25% of keys move (1/N)
```

### Real World Usage

| System | Uses consistent hashing for |
| --- | --- |
| DynamoDB | Deciding which partition node stores a key |
| Cassandra | Distributing rows across cluster nodes |
| Redis Cluster | Routing keys to the right shard |
| CDN (Cloudflare, Akamai) | Routing requests to nearest edge cache |
| Memcached | Client-side sharding across cache nodes |

---

---

## Num02 — Leader Election (Raft basics)

### The Problem

Distributed systems need **exactly one node** to coordinate writes — the leader.  
This ensures writes are serialized and consistent. But what happens when that node dies?  
You need the cluster to automatically elect a new leader without human intervention  
and without two nodes both thinking they are the leader (**split-brain**).

### The Concept

Raft uses two key ideas to solve this:

**1\. Terms — the logical clock**

A term is an integer that increments every time a new election starts. It acts as  
a generation counter. Any message from a node with a lower term is ignored — this  
is how a stale leader that recovers after a crash is automatically rejected by the cluster.

```
Term 1: Node1 is leader
Node1 crashes
Term 2: Node3 wins election, becomes leader
Node1 recovers — it's still at term 1
Node1 sends heartbeat — all followers see term=1 < their term=2 → reject
Node1 steps down automatically
```

**2\. Majority Quorum**

To win an election, a candidate needs votes from `N/2 + 1` nodes (majority).

```
5-node cluster: majority = 3  → tolerates 2 crashes
3-node cluster: majority = 2  → tolerates 1 crash
```

This guarantees only ONE node can win per term — two candidates cannot both  
get 3 votes out of 5.

### Node States

```
FOLLOWER  — default state, listens for heartbeats from the leader
              if no heartbeat received before timeout → becomes CANDIDATE

CANDIDATE — increments term, votes for itself, requests votes from all nodes
              if majority votes received → becomes LEADER
              if another leader announces itself → reverts to FOLLOWER
              if split vote (timeout) → increment term, retry

LEADER    — sends heartbeats every ~100ms to prevent followers from timing out
              handles all write requests
              if it sees a message with higher term → steps down to FOLLOWER
```

### Vote Granting Rules

A follower grants its vote to a candidate if:

1.  The candidate's term >= the follower's current term
2.  The follower has not already voted for someone else in this term

### What the demo shows

```
Step 1: 5 nodes start up — all FOLLOWER at term=0
Step 2: Node1 election timeout fires → becomes CANDIDATE, increments to term=1
Step 3: Node1 requests votes — all 4 followers grant votes (term 0 < 1, no prior vote)
Step 4: Node1 gets 5/5 votes (needed 3) → becomes LEADER for term=1
Step 5: Node1 crashes — followers stop receiving heartbeats
Step 6: Node3 times out → becomes CANDIDATE for term=2
Step 7: Node3 gets 4/5 votes (Node1 crashed) → becomes new LEADER for term=2
```

### Split Vote — What happens when two candidates tie?

In real Raft, multiple followers can time out near-simultaneously and both become  
candidates. If the votes split evenly (e.g. 2-2 in a 4-node cluster), nobody wins.  
Each candidate waits a **random** timeout before retrying with a higher term.  
The randomness ensures one fires before the other and wins cleanly.

This is why Raft uses random election timeouts (150–300ms) — it dramatically reduces  
the probability of repeated split votes.

### Real World Usage — Infrastructure

| System | Uses leader election for |
| --- | --- |
| etcd | Kubernetes stores all cluster state here — must have one authoritative leader |
| CockroachDB | Per-range Raft groups — automatic primary failover per shard |
| TiKV | Distributed storage layer of TiDB |
| Consul | Service discovery coordination |

### Real World Usage — Non-Infrastructure (Product Layer)

| Product scenario | Why leader election |
| --- | --- |
| **Google Docs / Notion** | One server is authority for a document — serializes concurrent edits. If it dies, new leader picks up from last saved state |
| **Multiplayer game servers** | One authoritative server per match for positions, scores, physics. Crashes trigger election, new server resumes from checkpoint |
| **Slack / WhatsApp message ordering** | One node per channel assigns sequence numbers to messages. "Connecting..." = waiting for new leader election |
| **Flash sale / inventory reservation** | Only the leader commits inventory writes — prevents 10 app servers all selling the same last item |
| **Push notification deduplication** | Only the leader dispatches notifications — others are warm standbys, preventing duplicate sends |
| **Financial reconciliation / billing** | Exactly-once execution — new leader reads the log to see what was already processed before resuming |

### The Pattern Across All Examples

Every example has the same three requirements:

| Requirement | What breaks without leader election |
| --- | --- |
| **Exactly-once execution** | Billing runs 3x, notifications sent 3x, item sold 3x |
| **Ordered writes** | Document corruption, chat messages out of order, game state diverges |
| **Automatic recovery** | Human has to promote a new leader manually at 3am |

### Why This Makes You a Better Engineer

Most engineers know leader election exists but not _why_ terms matter or what split-brain  
means. In a system design interview, saying:

> "I'd run 3 coordinator nodes with Raft-based leader election — the majority quorum  
> guarantees only one leader per term, and a stale leader that recovers will step down  
> automatically when it sees a higher term"

signals you understand failure modes, not just the happy path. That is the difference  
between a mid-level and a senior answer.

The concept also maps directly to application bugs — the TOCTOU race where a late  
cancellation signal overwrites a completed order status is the same class of problem  
as a stale leader trying to commit writes after a new leader was elected. Both are  
about who has authority and how you detect that authority has changed.

---

_Next: Num03 — Distributed Queues (Kafka concepts)_

---

---

## Num03 — Distributed Queues (Kafka concepts)

### The Problem

You have multiple services that need to communicate asynchronously at scale.  
A direct HTTP call is synchronous — the caller blocks until the receiver responds,  
and if the receiver is down, the message is lost. You need:

*   **Decoupling** — producer doesn't care if consumer is slow or temporarily down
*   **Durability** — messages survive crashes, can be replayed
*   **Parallelism** — multiple consumers process messages simultaneously
*   **Ordering** — related events (same order, same customer) must stay in sequence

### Core Concepts

**TOPIC**  
A named stream of messages. Like a named Go channel, but persistent on disk  
and replayable. Example: `"orders"`, `"payments"`, `"user-events"`.

**PARTITION**  
A topic is split into N ordered partitions. Each partition is an append-only log.  
Messages within a partition are strictly ordered. Across partitions, ordering  
is NOT guaranteed. Partitions are the unit of parallelism.

```
Topic "orders"
├── P0: [dave:placed, dave:paid]
├── P1: [alice:placed, alice:paid, alice:shipped]
└── P2: [bob:placed, carol:placed, bob:paid, carol:paid]
```

**OFFSET**  
Every message in a partition gets a sequential integer (0, 1, 2, ...).  
Consumers track their own offset — they control their own position in the log.  
This enables replay, resume after crash, and backfill.

**PARTITION KEY**  
Producers choose which partition a message goes to by hashing a key  
(e.g. customer ID). Same key always goes to same partition — guaranteeing  
ordering per key. Different keys may share a partition but are interleaved.

```
hash("customer:alice") % 3 = P1   ← all alice events always go to P1
hash("customer:bob")   % 3 = P2   ← all bob events always go to P2
```

**CONSUMER GROUP**  
A set of consumers that share the work of reading a topic. Each partition is  
assigned to exactly ONE consumer in the group at a time. This prevents  
duplicate processing while enabling parallelism.

```
Consumer group "billing-service":
  Consumer1 → P0 (reads dave's events)
  Consumer2 → P1 (reads alice's events)
  Consumer3 → P2 (reads bob's and carol's events)
```

### What the demo shows

```
Step 1: Create topic "orders" with 3 partitions
Step 2: ProducerA and ProducerB write 9 messages using customer ID as partition key
        — same customer always lands in same partition regardless of which producer
Step 3: Show partition state — append-only logs per partition
Step 4: Consumer group "billing-service" reads — 3 consumers, 1 partition each,
        all reading in parallel, each advancing its own offset
Step 5: Consumer2 crashes — P1 rebalances to Consumer1
        — Consumer1 inherits Consumer2's offset, resumes with no data loss
Step 6: Consumer3 resets offset to 0 on P2 — replays all messages from beginning
```

### Consumer crash and rebalance

```
Before crash:
  Consumer1 → P0 (offset=2, exhausted)
  Consumer2 → P1 (offset=3, exhausted)   ← crashes here
  Consumer3 → P2 (offset=3, exhausted)

After rebalance:
  Consumer1 → P0 + P1 (inherits offset=3 for P1, finds no new messages)
  Consumer3 → P2
```

No messages lost, no messages reread. The offset acts as a bookmark — wherever  
Consumer2 was when it crashed, Consumer1 picks up from exactly that point.

In real Kafka, offsets are committed to a special internal topic called  
`__consumer_offsets`. Consumers periodically commit their current offset.  
If a consumer crashes before committing, it may reprocess the last few messages  
(at-least-once delivery). To get exactly-once, you need idempotent consumers.

### Replay

Unlike traditional queues (RabbitMQ, SQS) where messages are deleted after  
consumption, Kafka retains messages for a configurable period (hours, days, forever).  
Any consumer can reset its offset to replay:

| Use case | How |
| --- | --- |
| Bug fix | Deploy fixed consumer, reset offset to 0, reprocess all events |
| New service | New analytics service joins, reads full history from beginning |
| Audit | Compliance reads all payment events from the last 7 years |
| A/B testing | Run two versions of consumer on same partition, compare results |

### Parallelism scaling

```
1 partition  = max 1 consumer active at a time (no parallelism)
3 partitions = max 3 consumers active (our demo)
10 partitions = max 10 consumers active
100 partitions = max 100 consumers active
```

Adding more consumers than partitions is wasteful — extra consumers sit idle  
since each partition can only be assigned to one consumer per group.

### Real World Usage

| System | Uses Kafka/queues for |
| --- | --- |
| **Uber** | Trip events, driver location updates, surge pricing triggers |
| **LinkedIn** | Activity feed, notifications, analytics pipeline |
| **Netflix** | User activity, recommendation pipeline, error tracking |
| **Anthropic/OpenAI** | Request logging, usage metering, async eval pipelines |
| **Banks** | Transaction events, fraud detection, audit log |

### Connection to what you already know

Your `orderCh` in `learn-go-intermediate` is essentially a single-partition Kafka topic:

| Your `orderCh` | Kafka |
| --- | --- |
| In-memory Go channel | Persistent on disk |
| No offset — message gone after read | Offset tracked, replay possible |
| Single channel, no partitions | N partitions for parallelism |
| Workers compete for same channel (fan-out) | Each consumer owns a dedicated partition |
| No consumer groups | Consumer groups coordinate shared reading |

The fan-out pattern (5 workers reading one channel) maps to having 5 consumers  
in a group with 5 partitions — each worker gets exclusive ownership of its partition.

---

_Next: Num04 — Eventual vs Strong Consistency_

---

---

## Num04 — Eventual vs Strong Consistency

### The Problem

You have 3 nodes storing the same data. A client writes `user:alice -> "premium"`.  
Which nodes have to confirm the write before we tell the client "done"?  
And if we tell the client "done" before all nodes confirm — what happens when  
the next read hits a node that hasn't received the write yet?

### PART A: Strong Consistency

The coordinator sends the write to all 3 nodes and waits for all 3 to confirm  
before ACKing the client.

```
[Primary   ]  WRITE  "user:alice" -> "premium"  ver=1
[Secondary1]  WRITE  "user:alice" -> "premium"  ver=1
[Secondary2]  WRITE  "user:alice" -> "premium"  ver=1

ALL confirmed. ACK sent to client.

Read from any node:
[Primary   ]  READ -> "premium"   ← correct
[Secondary1]  READ -> "premium"   ← correct
[Secondary2]  READ -> "premium"   ← correct
```

All 3 WRITE lines appear before the ACK. Only then does the client get a response.  
Any read — no matter which node the load balancer routes it to — returns `"premium"`.

**Cost:** write latency = latency of the slowest node. If Secondary2 is slow or  
unreachable, the write blocks until it responds.

### PART B: Eventual Consistency

The coordinator writes to Primary only and ACKs the client immediately.  
Secondaries receive the write asynchronously in the background.

```
[Primary]  WRITE  "user:alice" -> "premium"  ver=1
ACK sent to client.

Read immediately after:
[Primary   ]  READ -> "premium"   ← correct
[Secondary1]  READ -> "free"      ← STALE
[Secondary2]  READ -> "free"      ← STALE
```

Secondary1 and Secondary2 still hold the old value `"free"` because the  
background replication hasn't fired yet. This is called a **stale read**.

Then background replication fires:

```
[Secondary1]  WRITE  "user:alice" -> "premium"  ver=1
[Secondary2]  WRITE  "user:alice" -> "premium"  ver=1

Read again:
[Primary   ]  READ -> "premium"   ← converged
[Secondary1]  READ -> "premium"   ← converged
[Secondary2]  READ -> "premium"   ← converged
```

All nodes eventually agree. The time between the write and convergence is  
called **replication lag** — usually milliseconds, but can be seconds under load.

**Benefit:** write returned immediately without waiting for secondaries.  
System stays available even if a secondary is slow or temporarily down.

### PART C: Tunable Quorum (W + R > N)

Cassandra and DynamoDB let you choose W and R per request:

```
W = how many nodes must confirm the write before ACK
R = how many nodes must respond to a read
N = total number of nodes (3 in our case)

Rule: W + R > N  →  write set and read set overlap  →  consistent
      W + R ≤ N  →  no guaranteed overlap            →  eventual
```

```
Config                   W   R   W+R   Result
W=1 R=1 (fastest)        1   1    2    EVENTUAL   — stale reads possible
W=2 R=2 (balanced)       2   2    4    CONSISTENT — guaranteed overlap
W=3 R=1 (strong write)   3   1    4    CONSISTENT — guaranteed overlap
W=1 R=3 (strong read)    1   3    4    CONSISTENT — guaranteed overlap
```

Why does W+R > N guarantee consistency? If you write to 2 nodes and read from  
2 nodes out of 3 total, at least 1 node must be in both sets. That overlapping  
node has the latest write — so the read always sees it.

### CAP Theorem

```
C — Consistency:  every read sees the latest write
A — Availability: every request gets a non-error response
P — Partition tolerance: system works despite network splits
```

Network partitions always happen. So the real choice is:

*   **CP** — during a split, refuse requests to stay correct (Strong)
*   **AP** — during a split, serve possibly stale data to stay available (Eventual)

| System | Model | Notes |
| --- | --- | --- |
| DynamoDB | AP / Eventual | `ConsistentRead=true` per-request for strong |
| Spanner | CP / Strong | TrueTime API for global clock |
| Cassandra | Tunable | ONE / QUORUM / ALL per query |
| ZooKeeper | CP / Strong | Coordination primitive, not general storage |
| Redis Cluster | AP / Eventual | Async replication by default |

### The Real-World Bug This Explains

```
1. User upgrades to premium — write succeeds on Primary
2. Page redirects to dashboard
3. Dashboard reads from Secondary1
4. Secondary1 hasn't received the write yet — shows "free" tier
5. User is confused — they just paid
```

Not a code bug. A **consistency model mismatch**. Fix options:

*   Route reads-after-writes to Primary for that session
*   Use strong consistency (`ConsistentRead=true`) for the premium feature gate
*   Design the UI to tolerate eventual consistency (loading state, retry)

---

_Next: Num05 — Replication Strategies_

---

---

## Num05 — Replication Strategies

### The Problem

You decided to copy your data across 3 nodes (Num04 explained why). Now the question is: **who is allowed to accept a write, and how does that write reach the other nodes?**

Three different answers to that question give you three different strategies.

### PART A: Single Leader (Primary-Replica)

One node is the leader. All writes go to it. It then copies the write to the other nodes (replicas). Replicas serve reads only.

```
Client writes "plan=premium"
  → Primary writes it            ✓
  → Primary replicates to Replica1 (sync — confirmed before ACK)  ✓
  → Primary replicates to Replica2 (async — happens in background)
  → ACK to client

Reads immediately after:
  Primary   → "premium"   ✓
  Replica1  → "premium"   ✓  (sync, already has it)
  Replica2  → (empty)     ✗  (async, hasn't arrived yet)

Later, Replica2 catches up → "premium" ✓
```

Then Primary crashes:

```
Primary CRASHED — writes are now BLOCKED
Replica1 must be promoted to new Primary (failover)
Until promotion completes: no writes accepted
```

**When to use:** most apps. MySQL, Postgres, MongoDB all default to this. Simple to reason about. The async replica lag is the same stale read problem from Num04.

---

### PART B: Multi Leader

Two nodes both accept writes. Useful when you have two data centers and want each to accept local writes without waiting for a cross-DC round trip.

```
DC1 user writes "plan=premium"    → Leader-DC1 writes it  ver=1
DC2 user writes "plan=enterprise" → Leader-DC2 writes it  ver=1

Leaders sync with each other — CONFLICT:
  Leader-DC1 has "premium"    ver=1
  Leader-DC2 has "enterprise" ver=1
  Same key, same version, different values
```

Conflict resolution — Last Write Wins (LWW): whichever write had the later timestamp survives. DC2 wins:

```
All nodes converge to "enterprise" ver=2
```

**The risk:** if both users wrote at the exact same millisecond, LWW silently discards one of them. The DC1 user's upgrade to "premium" is lost with no error. This is why multi-leader is hard — you must design your app to handle conflicts explicitly.

**When to use:** cross-DC active-active setups where write latency across DCs is unacceptable. Google Docs uses a more sophisticated version of this (operational transforms / CRDTs instead of LWW).

---

### PART C: Leaderless (Dynamo-style)

No leader at all. The client writes to all nodes simultaneously. It only needs W nodes to confirm. Reads also go to all nodes — the client picks the response with the highest version.

```
Client writes "plan=premium" to all nodes (W=2 needed):
  NodeA → "premium" ver=1  ✓
  NodeB → "premium" ver=1  ✓
  NodeC → (slow, pending)
W=2 confirmed → ACK to client

Client reads immediately (R=2 needed):
  NodeA → "premium" ver=1   ✓
  NodeB → "premium" ver=1   ✓
  NodeC → "free"    ver=0   ← stale, ignored

Client picks highest version: "premium" ver=1
```

Then read repair — the client pushes the latest value back to NodeC:

```
NodeC → "premium" ver=1  (healed)
```

W+R=5 > N=3 → guaranteed that at least one node in the read set also confirmed the write → always correct.

**When to use:** systems that need no single point of failure and can tolerate more complex read logic. DynamoDB and Cassandra both support this model.

---

### Strategy Comparison

| Strategy | Who accepts writes | Conflict possible? | Single point of failure? | Real world |
| --- | --- | --- | --- | --- |
| Single Leader | Primary only | No | Yes (Primary) | MySQL, Postgres, MongoDB |
| Multi Leader | Multiple nodes | Yes | No | CouchDB, Google Docs (CRDT) |
| Leaderless | Any node | No (quorum) | No | DynamoDB, Cassandra, Riak |

### Connection to Num04

*   Single Leader async replication = the stale read problem from Num04 Part B
*   Leaderless W+R>N = exactly the quorum rule from Num04 Part C
*   Multi Leader conflict = what happens when you have eventual consistency without a single source of truth

---

_Next: Num06 — Distributed Transactions (2PC)_

---

## Topic 06 — Distributed Transactions

### The Problem

You're building an e-commerce system split into microservices:

*   **OrderService** creates orders
*   **PaymentService** charges cards
*   **InventoryService** reserves stock

When a customer clicks "Buy Now", you need to:

1.  Create the order
2.  Charge the payment
3.  Reserve inventory

**What if step 2 succeeds but step 3 fails?** You charged the customer but have no inventory to ship!

You need **atomicity across services** — all three succeed or all three fail together.

### Two-Phase Commit (2PC)

2PC is a distributed transaction protocol that guarantees atomicity:

```
                    ┌──────────────┐
                    │ Coordinator  │
                    └──────┬───────┘
                           │
         ┌─────────────────┼─────────────────┐
         │                 │                 │
         ▼                 ▼                 ▼
  ┌──────────┐      ┌──────────┐     ┌──────────┐
  │  Order   │      │ Payment  │     │Inventory │
  │ Service  │      │ Service  │     │ Service  │
  └──────────┘      └──────────┘     └──────────┘
```

#### Phase 1: PREPARE (Voting)

Coordinator asks each participant: **"Can you commit?"**

```
Coordinator → OrderService:     PREPARE
OrderService:                   checks resources, acquires local lock
OrderService → Coordinator:     VOTE_COMMIT ✓

Coordinator → PaymentService:   PREPARE
PaymentService:                 validates card, reserves amount
PaymentService → Coordinator:   VOTE_COMMIT ✓

Coordinator → InventoryService: PREPARE
InventoryService:               checks stock available
InventoryService → Coordinator: VOTE_COMMIT ✓
```

**Key point:** Each participant acquires locks and prepares to commit, but **does not commit yet**.

#### Phase 2: COMMIT or ABORT (Decision)

**If ALL voted YES:**

```
Coordinator → all participants: COMMIT
Each participant:               commits local transaction, releases lock
```

**If ANY voted NO:**

```
Coordinator → all participants: ABORT
Each participant:               rolls back prepared state, releases lock
```

### The 2PC Blocking Problem

**Critical flaw:** What if the coordinator crashes between Phase 1 and Phase 2?

```
Phase 1 complete:
  OrderService:    "I voted YES, holding lock on order #7890..."
  PaymentService:  "I voted YES, holding lock on $1200..."
  InventoryService:"I voted YES, holding lock on 1 laptop..."

💥 Coordinator crashes

Participants:
  ❌ Cannot COMMIT (no commit message received)
  ❌ Cannot ABORT (coordinator might have told others to commit)
  ⏳ BLOCKED FOREVER (holding locks, can't proceed)
```

**Impact:**

*   Resources locked indefinitely
*   Other transactions blocked waiting for these locks
*   System grinds to a halt

**Why this happens:**

*   Participants who voted YES are **committed to the coordinator's decision**
*   They cannot decide independently (would break atomicity)
*   The coordinator is a **single point of failure**

### Solutions to the Blocking Problem

1.  **Three-Phase Commit (3PC):** Adds a "pre-commit" phase so participants can timeout and abort safely
2.  **Paxos/Raft:** Replicate the coordinator itself so failure doesn't block
3.  **Saga Pattern:** Avoid distributed locks entirely (see below)

---

### Saga Pattern: The Alternative Approach

**Core idea:** Don't use distributed locks. Instead:

1.  Break the transaction into a **sequence of local transactions**
2.  Each service commits **immediately** (no prepare phase)
3.  If any step fails, run **compensating transactions** to undo previous steps

#### Saga: Happy Path

```
Step 1 — OrderService:
  ✓ Create order #7893 (status=PENDING)
  ✓ Commit locally
  → Emit event: OrderCreated

Step 2 — PaymentService:
  ← Receives: OrderCreated
  ✓ Charge $1200
  ✓ Commit locally
  → Emit event: PaymentCharged

Step 3 — InventoryService:
  ← Receives: PaymentCharged
  ✓ Reserve 1 laptop (stock 5→4)
  ✓ Commit locally
  → Emit event: ItemReserved

Step 4 — OrderService:
  ← Receives: ItemReserved
  ✓ Update order #7893 → CONFIRMED
```

**Key differences from 2PC:**

*   No coordinator, no locks across services
*   Each service commits its work **immediately**
*   Uses events (async) instead of RPC (sync)

#### Saga: Failure Path with Compensation

```
Step 1 — OrderService:
  ✓ Create order #7894 (status=PENDING)
  → Emit: OrderCreated

Step 2 — PaymentService:
  ✓ Charge $1200
  → Emit: PaymentCharged

Step 3 — InventoryService:
  ✗ Reserve failed (out of stock)
  → Emit: ItemReservationFailed

── Compensating Transactions (rollback) ──

Compensate Step 2 — PaymentService:
  ← Receives: ItemReservationFailed
  ✓ REFUND $1200 to customer
  → Emit: PaymentRefunded

Compensate Step 1 — OrderService:
  ← Receives: PaymentRefunded
  ✓ Update order #7894 → CANCELLED
```

**Trade-offs:**

*   ✓ No blocking — each service works independently
*   ✓ No single point of failure (no coordinator)
*   ✗ **Temporary inconsistency** — other users could briefly see order #7894 before it's cancelled
*   ✗ Compensating transactions must be **idempotent** (can be retried safely)
*   ✗ Some operations cannot be compensated (e.g., sending an email — you can send an apology, but can't unsend)

---

### 2PC vs Saga: When to Use What

| Property | Two-Phase Commit (2PC) | Saga Pattern |
| --- | --- | --- |
| **Atomicity** | Strong (ACID) — all or nothing, no intermediate states visible | Eventual (BASE) — temporary partial states visible |
| **Isolation** | Locks held during prepare → no dirty reads | No locks → other transactions see partial updates |
| **Blocking** | YES — participants block if coordinator crashes | NO — each step is independent |
| **Coordinator failure** | Catastrophic — blocks entire transaction | No impact — services continue independently |
| **Complexity** | Protocol is complex (2 phases, voting) | Business logic is complex (compensating txns) |
| **Latency** | Higher (synchronous, waits for all votes) | Lower (asynchronous, event-driven) |
| **Use when** | Strong consistency required, bounded # of participants | High availability needed, many services involved |

### Real-World Usage

| System | Pattern | Why |
| --- | --- | --- |
| **Traditional databases** | 2PC (XA transactions) | Single DB engine coordinates, strong ACID guarantees |
| **Banking core systems** | 2PC | Regulatory requirements demand strong consistency |
| **Microservices (Uber, Netflix)** | Saga | Availability > consistency, services evolve independently |
| **E-commerce order flows** | Saga | User experience tolerates eventual consistency |
| **Payment gateways (Stripe)** | Hybrid | 2PC for critical ledger updates, Saga for notifications |

### Connection to Previous Topics

*   **Num04 (Consistency Models):** 2PC = strong consistency, Saga = eventual consistency
*   **Num05 (Replication):** Multi-leader conflicts are similar to Saga's compensating transaction problem — both deal with concurrent updates
*   **Num02 (Leader Election):** 2PC's coordinator is the "leader" — if it fails, you need Raft/Paxos to elect a new one

---

### Interview Tips

**"How do you handle transactions across microservices?"**

*   Start with Saga pattern (more common in modern systems)
*   Mention compensating transactions explicitly
*   Acknowledge trade-offs: eventual consistency, complexity in rollback logic

**"What's the problem with Two-Phase Commit?"**

*   Focus on the **blocking problem** and coordinator as single point of failure
*   Mention 3PC or Paxos as solutions

**"Design a payment system"**

*   Use Saga for the order flow (order → payment → fulfillment)
*   Use 2PC or strong consistency for the actual ledger writes within the payment service
*   This hybrid approach is what Stripe, Square, and PayPal do

---

_Next: Num07 — Rate Limiting_

---

## Topic 07 — Rate Limiting

### The Problem

You've built a public API. Users start sending requests:

*   A buggy client sends **1000 requests/second** in an infinite loop
*   A malicious attacker floods your API with **100,000 requests/second** (DDoS)
*   One whale customer hogs all server resources, starving other users

Without rate limiting, your servers crash, your database melts, and your site goes down.

**Solution:** Limit each user to N requests per time window (e.g., 100 requests per minute).

---

### Algorithm 1: Token Bucket

**Mental model:** Imagine a physical bucket that holds tokens. Tokens drip in at a fixed rate. Each request eats one token. When the bucket is empty, reject requests.

```
Bucket (capacity=5, refill=2 tokens/sec)

t=0  [●●●●●]  ← full, 5 tokens
     Request → [●●●●_]  allowed ✓
     Request → [●●●__]  allowed ✓
     Request → [●●___]  allowed ✓
     Request → [●____]  allowed ✓
     Request → [_____]  allowed ✓
     Request → [_____]  REJECTED ✗ (empty!)

t=1  [●●___]  ← refilled +2 tokens (1 sec * 2 tokens/sec)
     Request → [●____]  allowed ✓
```

**Properties:**

*   **Allows bursts:** If bucket is full, you can fire 5 requests instantly
*   **Smooth long-term rate:** Over 10 seconds, you can only send 20 requests (2/sec \* 10)
*   **O(1) memory:** Just track `(tokens, last_refill_time)`

**Real-world usage:**

*   **GitHub API:** 5000 requests/hour (allows bursts, smooth average)
*   **Stripe API:** 100 requests/sec
*   **AWS API Gateway:** Token bucket with burstable quotas

---

### Algorithm 2: Fixed Window Counter

**Mental model:** Reset a counter every minute. Count requests. If count > limit, reject.

```
Limit = 100 requests/min

Minute 0:  count = 0 → 40 requests → count = 40 ✓
Minute 1:  count = 0 (reset!) → 50 requests → count = 50 ✓
Minute 2:  count = 0 (reset!) → 120 requests → count = 120 ✗ (rejected)
```

**The Boundary Burst Exploit:**

```
Timeline (limit = 100/min):

Minute 0, second 59:  90 requests → count = 90 ✓
                       ↓
Minute 1, second 00:  90 requests → count = 90 ✓ (counter reset!)
                       ↑
                   boundary

Result: 180 requests in 2 seconds — both windows allowed them!
```

This is a **real vulnerability** exploited in production systems.

**Properties:**

*   ✓ **Dead simple:** Just 1 counter + 1 timestamp
*   ✓ **O(1) memory**
*   ✗ **Boundary burst problem:** Attacker can send 2x limit by hitting window boundaries

**Real-world usage:**

*   Internal microservices with coarse limits (e.g., "don't DOS the payment service")
*   Systems where boundary bursts are acceptable

---

### Algorithm 3: Sliding Window Log

**Mental model:** Keep a log of all request timestamps. For each new request, count how many timestamps fall within the last N seconds. If count >= limit, reject.

```
Limit = 5 requests per 10 seconds

t=1   log=[1]          count=1  ✓
t=2   log=[1,2]        count=2  ✓
t=3   log=[1,2,3]      count=3  ✓
t=4   log=[1,2,3,4]    count=4  ✓
t=5   log=[1,2,3,4,5]  count=5  ✓
t=6   log=[1,2,3,4,5]  count=5  ✗ REJECTED (window full)

t=12  Evict t=1,2 (>10 sec old)
      log=[3,4,5]      count=3  ✓
      log=[3,4,5,12]   count=4  ✓
```

**Properties:**

*   ✓ **No boundary burst:** Rolling window is always accurate
*   ✓ **Precise:** Guarantees exactly N requests per window
*   ✗ **O(N) memory:** Must store all timestamps in the window (can be expensive at scale)

**Optimization — Sliding Window Counter (hybrid):**  
Instead of storing every timestamp, keep only 2 counters:

*   `countCurrentWindow` (e.g., current minute)
*   `countPreviousWindow` (e.g., previous minute)

Estimate the rolling count using weighted average based on how far into the current window you are:

```
estimatedCount = countPreviousWindow * (1 - percentIntoCurrentWindow) 
                 + countCurrentWindow
```

This reduces memory from O(N) → O(1) with slight inaccuracy.

**Real-world usage:**

*   **Cloudflare Workers:** Sliding window log (high scale)
*   **Redis ZSET:** Store timestamps as sorted set, use `ZREMRANGEBYSCORE` to evict old entries

---

### Algorithm 4: Distributed Rate Limiting

**The problem:**

You have 3 API servers behind a load balancer. Each server has its own in-memory rate limiter (limit=10).

```
User sends 30 requests:
  → Load balancer round-robins
    → Server1 sees 10 requests → allows all 10 ✓
    → Server2 sees 10 requests → allows all 10 ✓
    → Server3 sees 10 requests → allows all 10 ✓

Total: 30 requests allowed (limit was 10!)
```

User bypassed the rate limit by being routed to multiple servers!

**The solution: Shared counter in Redis**

```
All servers share a single Redis counter:

Server1: INCR ratelimit:user123 → 1  ✓
Server2: INCR ratelimit:user123 → 2  ✓
Server3: INCR ratelimit:user123 → 3  ✓
...
Server1: INCR ratelimit:user123 → 10 ✓
Server2: GET  ratelimit:user123 → 10 (>= limit) ✗ REJECTED
Server3: GET  ratelimit:user123 → 10 (>= limit) ✗ REJECTED
```

All servers see the **same counter** → globally consistent limit.

**Why Redis?**

*   ✓ **Atomic operations:** `INCR` is atomic, no race conditions
*   ✓ **Built-in TTL:** `EXPIRE` key after 60 seconds → auto-reset for time windows
*   ✓ **Fast:** In-memory, microsecond latency
*   ✓ **Lua scripts:** Atomic multi-step operations (refill + decrement for token bucket)

**Real-world usage:**

*   **Kong API Gateway:** Redis-backed rate limiting
*   **Cloudflare:** Distributed Workers with shared Durable Objects
*   **Nginx:** `limit_req_zone` (shared memory zone across workers)

---

### Comparison Table

| Algorithm | Memory | Allows Bursts? | Boundary Safe? | Distributed? | Real-World |
| --- | --- | --- | --- | --- | --- |
| **Token Bucket** | O(1) | ✓ Yes | ✓ Yes | Needs Redis | GitHub, Stripe, AWS |
| **Fixed Window** | O(1) | ✗ No | ✗ No (exploit!) | Needs Redis | Internal APIs |
| **Sliding Window** | O(N) | ✗ No | ✓ Yes | Needs Redis | Cloudflare, Redis |

---

### When to Use What

| Scenario | Algorithm | Why |
| --- | --- | --- |
| **Public API** (user-facing) | Token Bucket + Redis | Burst-friendly, smooth average rate |
| **DDoS protection** | Fixed Window + Redis | Dead simple, "good enough" for blocking attacks |
| **Strict compliance** (SLA) | Sliding Window + Redis | Accurate, no boundary exploits |
| **Internal microservices** | Fixed Window (local) | Coarse protection, no need for precision |

---

### Implementation: Token Bucket in Redis (Production Pattern)

Most production systems use **Token Bucket with Redis Lua script** for atomicity:

```
-- Lua script executed atomically in Redis
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local bucket = redis.call('HMGET', key, 'tokens', 'last_refill')
local tokens = tonumber(bucket[1]) or capacity
local last_refill = tonumber(bucket[2]) or now

-- Refill tokens based on time passed
local elapsed = now - last_refill
local tokens_to_add = elapsed * refill_rate
tokens = math.min(capacity, tokens + tokens_to_add)

-- Try to consume 1 token
if tokens >= 1 then
  tokens = tokens - 1
  redis.call('HMSET', key, 'tokens', tokens, 'last_refill', now)
  redis.call('EXPIRE', key, 3600) -- TTL 1 hour
  return 1 -- allowed
else
  return 0 -- rejected
end
```

This ensures **atomicity** — refill + decrement happen in one operation, no race conditions.

---

### Connection to Previous Topics

*   **Num04 (Consistency):** Rate limiting is eventually consistent — a user might briefly exceed limit due to Redis replication lag
*   **Num06 (Distributed Transactions):** Redis Lua scripts = local ACID transaction for rate limit check
*   **Num03 (Queues):** Rate limiting can use Kafka for async throttling (queue requests, process at fixed rate)

---

### Interview Tips

**"How do you rate-limit a public API?"**

*   Start with Token Bucket (industry standard)
*   Mention Redis for distributed systems
*   Highlight the burst-friendly property

**"What's the problem with Fixed Window?"**

*   Boundary burst exploit (2x limit at window edges)
*   Draw the timeline diagram

**"How does Redis rate limiting scale?"**

*   Redis Cluster shards by key (user ID) → horizontal scale
*   Lua scripts ensure atomicity
*   `EXPIRE` for auto-cleanup

**"Design rate limiting for 1M requests/sec"**

*   Token Bucket algorithm
*   Redis Cluster (100 shards)
*   Lua script for atomic refill + decrement
*   Fallback: local in-memory limiter if Redis unavailable (graceful degradation)

---

_Next: Num08 — Circuit Breaker Pattern_

---

## Topic 08 — Circuit Breaker Pattern

### The Problem

Imagine this cascade of failures:

```
1. Payment Service database crashes
2. Payment Service becomes slow (every request times out after 30 seconds)
3. API Gateway keeps calling Payment Service → threads blocked waiting
4. API Gateway's 100-thread pool exhausted → now IT can't serve ANY requests
5. Load balancer marks API Gateway as down, routes traffic to other instances
6. Those instances also exhaust their threads → entire system collapses
```

This is a **cascading failure** — one service's problem spreads like wildfire through the entire system.

**Circuit Breaker prevents this:** After N failures, STOP calling the downstream service. Fail fast. Save your threads. Survive.

---

### The Mental Model: Electrical Circuit Breaker

Your house has a circuit breaker box. When too much current flows (e.g., you plug in 10 space heaters):

1.  **Breaker trips** → power cuts off immediately
2.  Prevents fire → saves the house
3.  You fix the problem (unplug space heaters)
4.  **Reset the breaker** → power flows again

Same concept in software:

1.  **Circuit trips** → stop calling downstream service
2.  Prevents thread exhaustion → saves your service
3.  Downstream service recovers (database back online)
4.  **Circuit closes** → resume normal operation

---

### The Three States

```
           ┌─────────────────────┐
           │      CLOSED         │  Normal operation
           │  (requests flow)    │  Track failure count
           └──────────┬──────────┘
                      │
                      │ N failures
                      ▼
           ┌─────────────────────┐
           │       OPEN          │  Fail fast
           │  (reject requests)  │  Don't call downstream
           └──────────┬──────────┘
                      │
                      │ timeout expires
                      ▼
           ┌─────────────────────┐
           │     HALF_OPEN       │  Testing recovery
           │  (try 1 request)    │  
           └──────────┬──────────┘
                      │
            ┌─────────┴─────────┐
            │                   │
         success             failure
            │                   │
            ▼                   ▼
         CLOSED               OPEN
```

**CLOSED:** Normal state. Requests go through. Count failures.

*   After 3 failures → transition to OPEN

**OPEN:** Circuit is tripped. Reject all requests immediately (fail fast). Don't call downstream service.

*   After 5 seconds → transition to HALF\_OPEN

**HALF\_OPEN:** Testing if downstream recovered. Allow 1 request through.

*   If succeeds → transition to CLOSED (recovered!)
*   If fails → transition back to OPEN (still broken)

---

### Step-by-Step Example

**Initial state: CLOSED**

```
Request 1 → Payment Service: ✓ SUCCESS
  Circuit: CLOSED, failure_count=0

Request 2 → Payment Service: ✓ SUCCESS
  Circuit: CLOSED, failure_count=0

Request 3 → Payment Service: ✗ TIMEOUT
  Circuit: CLOSED, failure_count=1

Request 4 → Payment Service: ✗ TIMEOUT
  Circuit: CLOSED, failure_count=2

Request 5 → Payment Service: ✗ TIMEOUT
  Circuit: CLOSED, failure_count=3 → THRESHOLD BREACHED
  Circuit: CLOSED → OPEN
```

---

**Circuit is OPEN**

```
Request 6 → Circuit Breaker: ✗ REJECTED (fail fast, no call made)
  Response time: 1ms (instead of 30-second timeout!)

Request 7 → Circuit Breaker: ✗ REJECTED
  Response time: 1ms

... (5 seconds pass) ...
```

---

**Timeout expired → HALF\_OPEN**

```
Request 8 → Circuit: OPEN → HALF_OPEN (testing recovery)
           → Payment Service: ✗ TIMEOUT (still broken)
  Circuit: HALF_OPEN → OPEN (test failed, reopen circuit)

... (5 more seconds pass) ...

Request 9 → Circuit: OPEN → HALF_OPEN
           → Payment Service: ✓ SUCCESS (database recovered!)
  Circuit: HALF_OPEN → CLOSED (recovery confirmed!)

Request 10 → Payment Service: ✓ SUCCESS
  Circuit: CLOSED, failure_count=0
  Normal operation resumed
```

---

### Why This Saves Your System

**Without Circuit Breaker:**

```
Payment Service down (30-second timeout per request)
API Gateway has 100 threads

Request 1: thread 1 blocked for 30 seconds
Request 2: thread 2 blocked for 30 seconds
...
Request 100: thread 100 blocked

Request 101: NO THREADS AVAILABLE → API Gateway can't accept request
User gets 503 error even for endpoints that DON'T need Payment Service!
```

**With Circuit Breaker:**

```
Requests 1-3: fail, each takes 30 seconds (threads recover after timeout)
Circuit opens after 3rd failure

Request 4: Circuit Breaker rejects in 1ms → thread immediately available
Request 5: Circuit Breaker rejects in 1ms → thread immediately available
...

API Gateway survives!
- Other endpoints (User Service, Search, etc.) still work
- Only Payment-dependent endpoints fail fast
- User gets instant error instead of 30-second timeout
```

---

### Real-World: Combining with Other Patterns

Circuit Breaker is usually used with:

#### 1\. **Timeout**

```
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()
response, err := paymentService.Call(ctx, request)
```

Kill slow requests → don't let them block threads forever.

#### 2\. **Retry (with exponential backoff)**

```
Attempt 1: fail (network glitch)
Wait 100ms
Attempt 2: fail
Wait 200ms
Attempt 3: fail → Circuit Breaker counts this as 1 failure
```

Transient failures get retried. Permanent failures trigger circuit breaker.

#### 3\. **Bulkhead (thread pool isolation)**

```
Payment Service: 10 threads
User Service: 20 threads
Search Service: 30 threads
```

If Payment's 10 threads are exhausted, User and Search still have their pools → partial degradation instead of total collapse.

---

### Netflix Hystrix Flow

```
Request arrives
  ↓
Is Circuit OPEN?
  Yes → Return fallback (cached response / error) [1ms]
  No  → Continue
  ↓
Check bulkhead (thread pool)
  Pool full? → Return fallback
  Pool available? → Get thread
  ↓
Call downstream service with timeout (2 seconds)
  ↓
Success? → Reset failure counter
Timeout? → Retry 1x with backoff
Failure? → Increment counter
  ↓
Failure count >= threshold?
  Yes → OPEN circuit
  No  → Keep CLOSED
```

---

### Monitoring (Critical for Production!)

You MUST monitor these metrics:

```
# Gauge: 0=CLOSED, 1=OPEN, 2=HALF_OPEN
circuit_breaker_state{service="payment"} 1

# Counter: total failures
circuit_breaker_failures{service="payment"} 47

# Counter: rejected requests (circuit open)
circuit_breaker_rejected{service="payment"} 1203

# Histogram: downstream latency
payment_service_latency_seconds_p99 4.5
```

**Alert rules:**

```
ALERT CircuitOpen
  IF circuit_breaker_state == 1 FOR > 1 minute
  MESSAGE "Payment Service circuit has been OPEN for 1 minute"

ALERT HighRejectionRate
  IF rate(circuit_breaker_rejected[1m]) > 100
  MESSAGE "Circuit Breaker rejecting >100 req/s"
```

---

### Configuration Parameters

| Parameter | Typical Value | Meaning |
| --- | --- | --- |
| `failureThreshold` | 3-10 | Open circuit after N failures |
| `successThreshold` | 2-5 | Close circuit after N successes in HALF\_OPEN |
| `timeout` | 5-30 seconds | Wait this long before trying HALF\_OPEN |
| `requestTimeout` | 1-5 seconds | Kill slow downstream calls after this |
| `maxConcurrentRequests` | 10-100 | Bulkhead: thread pool size |

**Tuning tips:**

*   **Low** `**failureThreshold**` (e.g., 3) → sensitive, opens quickly → more false positives
*   **High** `**failureThreshold**` (e.g., 20) → tolerant, opens slowly → more failures slip through
*   **Short** `**timeout**` (e.g., 5s) → test recovery frequently → faster recovery
*   **Long** `**timeout**` (e.g., 60s) → test recovery rarely → less load on sick downstream

---

### Connection to Previous Topics

*   **Num07 (Rate Limiting):** Both protect against overload. Rate limiting = protect FROM users. Circuit Breaker = protect FROM downstream failures.
*   **Num06 (Distributed Transactions):** Saga pattern uses Circuit Breaker to avoid calling failed services during compensating transactions.
*   **Num04 (Consistency):** Circuit Breaker prioritizes **availability** over **consistency** (CAP theorem) — better to return stale cache than block.

---

### Interview Tips

**"How do you prevent cascading failures?"**

*   Start with Circuit Breaker (the gold standard)
*   Mention timeout + retry + bulkhead combo
*   Give the thread exhaustion example

**"What are the Circuit Breaker states?"**

*   CLOSED (normal), OPEN (tripped), HALF\_OPEN (testing)
*   Draw the state transition diagram
*   Explain the threshold logic

**"Design a resilient payment API"**

**"What metrics would you monitor?"**

*   Circuit state (CLOSED/OPEN/HALF\_OPEN)
*   Failure count
*   Rejection rate (requests blocked by circuit)
*   P99 latency of downstream calls

---

### Real-World Libraries

| Library | Language | Notes |
| --- | --- | --- |
| **Netflix Hystrix** | Java | Original implementation (now in maintenance mode) |
| **Resilience4j** | Java | Modern alternative, lighter weight |
| **Polly** | .NET | Fluent API, supports Circuit Breaker + Retry |
| **sony/gobreaker** | Go | Simple, production-ready |
| **Istio / Envoy** | Service Mesh | Circuit Breaker built into sidecar proxy |

**Production recommendation:** Use a service mesh (Istio/Linkerd) — circuit breaker logic moves to sidecar, no code changes needed.

---

## Topic 09: Gossip Protocol

### The Problem

**Scenario:** You have a 1,000-node distributed database (like Cassandra). How do all nodes learn when:

*   A new node joins the cluster?
*   A node crashes or becomes unreachable?
*   Configuration changes (e.g., "maintenance mode enabled")?

**Bad Approach #1:** Central coordinator broadcasts to all nodes

*   ❌ Single point of failure (coordinator crashes = no updates)
*   ❌ Bottleneck (coordinator must send 1,000 messages)

**Bad Approach #2:** Every node broadcasts to every other node

*   ❌ O(N²) network messages (1,000 nodes = 1 million messages!)
*   ❌ Network storm

**Gossip Protocol Solution:**

*   ✅ Each node picks K random peers and shares state
*   ✅ Information spreads exponentially: O(log N) rounds to full convergence
*   ✅ No coordinator — fully decentralized
*   ✅ Fault-tolerant — works even if 50% of nodes fail

---

### Real-World Analogy

**Gossip Protocol = How Rumors Spread in School**

Day 0: Alice learns a secret → 1 person knows

Day 1: Alice tells 2 friends (Bob, Carol) → 3 people know

Day 2: Each of those 3 tells 2 more friends → 9 people know

Day 3: Those 9 tell 2 each → 27 people know

**Result:** In log₂(N) days, the entire school (N students) knows the secret!

---

### How It Works (4 Core Mechanisms)

#### 1\. **Membership Propagation**

```
Node-5 joins cluster
  ↓
Node-1 learns about Node-5 (direct contact)
  ↓
Round 1: Node-1 gossips to 2 random peers → 3 nodes know
Round 2: Those 3 gossip to 2 each → ~7 nodes know
Round 3: Those 7 gossip to 2 each → ~15 nodes know
  ↓
After log₂(100) ≈ 7 rounds → All 100 nodes know Node-5 exists
```

**Key insight:** Exponential spread! Each round doubles the number of informed nodes.

#### 2\. **Failure Detection (Heartbeat Counters)**

```
Each node maintains:
  Node-1: heartbeat = 5
  Node-2: heartbeat = 5
  Node-3: heartbeat = 3  ← Stuck! Node-3 crashed
  Node-4: heartbeat = 5
```

Every round:

1.  Increment own heartbeat
2.  Gossip heartbeat counters with 2 random peers
3.  Merge counters (take max for each node)

**Detection:** If Node-X's heartbeat hasn't increased in 2+ rounds → Node-X suspected failed

#### 3\. **State Synchronization**

```
Node-1 stores: {user:1 → "Alice" (v2)}
Node-2 stores: {user:2 → "Bob" (v1)}

Round 1: Node-1 and Node-2 gossip
  → Node-1 receives user:2 → "Bob" (v1)
  → Node-2 receives user:1 → "Alice" (v2)

Result: Both nodes now have both keys (eventual consistency)
```

**Conflict resolution:** Use version numbers (or vector clocks) to keep the latest value.

#### 4\. **Anti-Entropy (Background Repair)**

```
Node-1 hash: 0xABC123 (hash of all keys)
Node-2 hash: 0xABC123 (same)
Node-3 hash: 0xDEF456 (different! missing some keys)

→ Node-1 and Node-3 gossip full state to sync
→ Node-3 hash becomes 0xABC123
```

Used by: DynamoDB (Merkle trees), Cassandra (repair service)

---

### Gossip Protocol Parameters

| Parameter | Typical Value | Trade-off |
| --- | --- | --- |
| **Fanout (K)** | 2-3 peers per round | Higher K = faster spread, more network load |
| **Gossip Interval** | 1 second | Shorter interval = faster convergence, more CPU/network |
| **Timeout Threshold** | 3 missed heartbeats | Lower threshold = faster failure detection, more false positives |
| **Max Payload Size** | 1-10 KB | Larger payload = fewer rounds, more bandwidth per message |

**Example:** Cassandra uses K=3, interval=1s → 10,000 node cluster converges in ~13 seconds.

---

### Visual: Gossip Spread (10 Nodes, Fanout=2)

```
Round 0: Node-1 learns about new node
  [●] [○] [○] [○] [○] [○] [○] [○] [○] [○]
   1 knows

Round 1: Node-1 tells 2 peers
  [●] [●] [●] [○] [○] [○] [○] [○] [○] [○]
   3 know

Round 2: Those 3 tell 2 each (6 new, but some overlap)
  [●] [●] [●] [●] [●] [●] [●] [○] [○] [○]
   7 know

Round 3: Those 7 tell 2 each
  [●] [●] [●] [●] [●] [●] [●] [●] [●] [●]
   All 10 know! (in just 3 rounds)
```

**Formula:** After R rounds, ~min(N, K^R) nodes know

*   K=2, R=3 → 2³ = 8 nodes
*   K=2, R=7 → 2⁷ = 128 nodes

---

### Gossip vs Other Approaches

```
┌───────────────┬────────────────┬─────────────────┬──────────────────┐
│ Approach      │ Latency        │ Fault Tolerance │ Network Load     │
├───────────────┼────────────────┼─────────────────┼──────────────────┤
│ Gossip        │ O(log N)       │ ⭐⭐⭐⭐⭐        │ O(N*K) per round │
│               │ ~7 rounds      │ No SPOF         │ Moderate         │
│               │ (N=100)        │                 │                  │
├───────────────┼────────────────┼─────────────────┼──────────────────┤
│ Broadcast     │ O(1) - instant │ ⭐               │ O(N²) messages   │
│ (all-to-all)  │ 1 round        │ Source fail =   │ Network storm    │
│               │                │ lost update     │                  │
├───────────────┼────────────────┼─────────────────┼──────────────────┤
│ Central Hub   │ O(1) - 2 hops  │ ⭐⭐             │ O(N) messages    │
│               │ Node→Hub→All   │ Hub = SPOF      │ Efficient        │
├───────────────┼────────────────┼─────────────────┼──────────────────┤
│ Consensus     │ O(1) + quorum  │ ⭐⭐⭐⭐          │ O(N) messages    │
│ (Raft/Paxos)  │ Requires acks  │ Survives N/2-1  │ Requires leader  │
│               │                │ failures        │                  │
└───────────────┴────────────────┴─────────────────┴──────────────────┘
```

---

### Real-World Example: Cassandra Ring Membership

**Scenario:** 100-node Cassandra cluster. Node-42 joins.

```
Step 1: Node-42 contacts seed node (Node-1)
  → Node-1 learns: "Node-42 owns tokens 500-599"

Step 2: Gossip propagation (every 1 second)
  Round 1: Node-1 gossips to 3 random nodes → 4 know
  Round 2: Those 4 gossip to 3 each → ~13 know
  Round 3: Those 13 gossip → ~40 know
  Round 4: Those 40 gossip → ~100 know ✓

Total time: 4 seconds

Step 3: All nodes update their routing tables
  Tokens 500-599: Node-42 (previously owned by Node-27)
  → Clients now send writes for those tokens to Node-42
```

**Key benefit:** No downtime! Cluster continues serving requests during gossip.

---

### Real-World Use Cases

#### 1\. **Cassandra (Membership + Failure Detection)**

```
Every node maintains:
  - Ring topology (which node owns which token range)
  - Node status (UP / DOWN / LEAVING)
  - Schema versions

Gossip interval: 1 second
Failure detection: 3 missed heartbeats → mark node as DOWN

Result: 1000-node cluster detects failures in ~3 seconds
```

#### 2\. **Consul (Service Discovery)**

```
Consul agent on each server:
  - Gossips service health checks
  - Uses SWIM protocol (reduces false positives)

Example:
  Web-Server-1 fails health check
  → Consul agent gossips: "Web-Server-1 DOWN"
  → All agents update service registry in ~1 second
  → Load balancers stop routing to Web-Server-1
```

#### 3\. **Redis Cluster (Node Discovery)**

```
Every Redis node:
  - Gossips cluster topology (which node owns which hash slots)
  - Gossips with 3 random nodes per second

16,384 hash slots across 6 nodes:
  Node-1: slots 0-2729
  Node-2: slots 2730-5460
  ...

New node joins → all nodes learn topology in ~3 gossip rounds
```

#### 4\. **AWS DynamoDB (Anti-Entropy)**

```
Background process:
  1. Build Merkle tree of all keys (hash tree)
  2. Compare tree with peer nodes
  3. If mismatch detected → sync only differing keys

Example:
  Node-A hash: 0xABC
  Node-B hash: 0xABC ✓
  Node-C hash: 0xDEF ✗
  → Node-A and Node-C gossip to repair inconsistency
```

---

### SWIM Protocol (Advanced Gossip)

**Problem:** Simple gossip has false positives (network delays look like node failures)

**SWIM Solution:** Indirect probing

```
Node-1 suspects Node-3 failed (missed heartbeat)
  ↓
Node-1 asks Node-2: "Can you ping Node-3?"
  ↓
Node-2 pings Node-3 → Success!
  ↓
Node-1 realizes: Network partition, not failure
  ↓
Node-1 does NOT mark Node-3 as DOWN
```

**Used by:** Consul, Memberlist (HashiCorp library)

---

### Implementation Details

#### Pseudocode: Basic Gossip

```
type Node struct {
    ID string
    KnownMembers map[string]bool
    Heartbeats map[string]int
}

func (n *Node) GossipRound() {
    // Increment own heartbeat
    n.Heartbeats[n.ID]++
    
    // Pick K random peers
    peers := n.pickRandomPeers(K)
    
    for _, peer := range peers {
        // Send current state to peer
        peer.receive(n.KnownMembers, n.Heartbeats)
        
        // Receive peer's state
        n.merge(peer.KnownMembers, peer.Heartbeats)
    }
    
    // Detect failures
    currentRound := n.Heartbeats[n.ID]
    for nodeID, lastHeartbeat := range n.Heartbeats {
        if currentRound - lastHeartbeat > 3 {
            n.markAsFailed(nodeID)
        }
    }
}

func (n *Node) merge(peerMembers map[string]bool, peerHeartbeats map[string]int) {
    // Add any new members
    for member := range peerMembers {
        n.KnownMembers[member] = true
    }
    
    // Take max heartbeat for each node
    for nodeID, hb := range peerHeartbeats {
        if hb > n.Heartbeats[nodeID] {
            n.Heartbeats[nodeID] = hb
        }
    }
}
```

---

### Trade-offs & Limitations

#### ✅ Pros

1.  **Scalability:** Works with 10,000+ nodes (O(log N) convergence)
2.  **Fault tolerance:** No single point of failure, works with 50%+ nodes down
3.  **Simplicity:** No complex coordination or leader election
4.  **Decentralized:** No coordinator needed

#### ❌ Cons

1.  **Eventual consistency:** Takes multiple rounds (seconds) to converge
2.  **Network overhead:** Continuous background chatter (K messages per node per second)
3.  **Redundant messages:** Same update may be sent multiple times
4.  **Not suitable for strong consistency:** Can't guarantee linearizability

#### When to Use Gossip

*   ✅ Membership tracking (who's in the cluster?)
*   ✅ Failure detection (who's alive?)
*   ✅ Configuration propagation (non-critical settings)
*   ✅ Metrics aggregation (approximate counts, averages)

#### When NOT to Use Gossip

*   ❌ Financial transactions (need strong consistency)
*   ❌ Leader election (use Raft/Paxos instead)
*   ❌ Distributed locks (use etcd/ZooKeeper)
*   ❌ Real-time updates (gossip has latency)

---

### Interview Tips

**"How does Cassandra detect node failures?"**

**"Why not use broadcast instead of gossip?"**

**"How do you handle network partitions?"**

**"Design a distributed monitoring system (1000 servers)"**

**"What's the difference between gossip and consensus?"**

---

### Monitoring & Observability

**Key Metrics to Track:**

```
Gossip Health:
- gossip_rounds_total (counter)
- gossip_messages_sent_total (counter)
- gossip_messages_received_total (counter)
- gossip_merge_conflicts_total (counter)

Convergence:
- cluster_membership_size (gauge) ← should be same across all nodes
- gossip_convergence_time_seconds (histogram) ← how long to full sync

Failure Detection:
- node_suspected_failed_total (counter)
- false_positive_rate (gauge) ← nodes marked DOWN but actually UP
```

**Alert Examples:**

```
1. Cluster split-brain:
   IF node_1.cluster_size = 50 AND node_51.cluster_size = 50
   → Network partition! Two separate clusters

2. Slow convergence:
   IF gossip_convergence_time > 10 seconds
   → Network issues or fanout too low

3. High false positive rate:
   IF false_positive_rate > 5%
   → Increase timeout threshold or use SWIM protocol
```

---

### Further Reading

*   **Cassandra gossip internals:** [DataStax Docs](https://docs.datastax.com/en/cassandra-oss/3.x/cassandra/architecture/archGossipAbout.html)
*   **SWIM protocol paper:** "SWIM: Scalable Weakly-consistent Infection-style Process Group Membership Protocol" (Cornell, 2002)
*   **HashiCorp Memberlist:** Production-ready gossip library (Go)
*   **Epidemic Algorithms:** "Epidemic Algorithms for Replicated Database Maintenance" (Xerox PARC, 1987)

---

## Num10 — Distributed Locking

**Topic:** Redlock algorithm, fencing tokens, coordination patterns

_To be implemented. See Num10DistributedLockingDemo() stub in demos.go_

---

## Num11 — Observability in Distributed Consensus Systems

**Topic:** Tracing Raft RPCs, detecting split-brain via metrics, visualizing consensus rounds

_To be implemented. See Num11ObservabilityDemo() stub in demos.go_

---

## Num12 — Sharding with Per-Shard Consensus Groups

**Topic:** Per-shard Raft groups (TiKV, CockroachDB), range partitioning, cross-shard 2PC

_To be implemented. See Num12ShardingStrategiesDemo() stub in demos.go_

---

## Num13 — Event Sourcing: The Distributed Log IS the Consensus Output

**Topic:** Raft log as event store, CQRS on top of consensus, replay and time travel

_To be implemented. See Num13EventSourcingDemo() stub in demos.go_

---

## Num14 — Consensus Algorithms Comparison

**Topic:** Raft vs Paxos vs ZAB — when to use which, trade-offs, production considerations

_To be implemented. See Num14ConsensusAlgorithmsDemo() stub in demos.go_

---

## Num15 — CAP Theorem & Trade-offs

**Topic:** Partition simulation, CP vs AP systems, PACELC extension, split-brain scenarios

_To be implemented. See Num15CAPTheoremDemo() stub in demos.go_

---

## Num16 — Logical Time & Clocks

**Topic:** Lamport timestamps, vector clocks, happened-before relation, causality tracking

_To be implemented. See Num16LogicalClocksDemo() stub in demos.go_

**Core Concepts:**

*   Physical clocks unreliable (clock skew, NTP drift)
*   Lamport timestamps: single counter guaranteeing happened-before ordering
*   Vector clocks: detect concurrent events, used in DynamoDB/Riak versioning
*   Applications: debugging race conditions, conflict detection in eventual consistency

**Real-world:** DynamoDB versioning, Riak dotted version vectors, distributed tracing

---

## Num17 — Failure Detection

**Topic:** Phi Accrual detector, SWIM indirect probing, reducing false positives

_To be implemented. See Num17FailureDetectionDemo() stub in demos.go_

**Core Concepts:**

*   Basic heartbeat problems: network delays cause false positives
*   Phi Accrual: statistical model of heartbeat arrivals, outputs suspicion level (0-100%)
*   SWIM protocol: indirect probing (ask peer to ping suspected node)
*   Trade-off: detection speed vs false positive rate

**Real-world:** Cassandra (Phi Accrual), Consul (SWIM via Serf), Akka cluster

---

## Num18 — Quorums & Consistency

**Topic:** N/R/W models, sloppy quorums, hinted handoff (Riak/Dynamo style)

_To be implemented. See Num18QuorumsConsistencyDemo() stub in demos.go_

**Core Concepts:**

*   Quorum: W + R > N guarantees read sees latest write (overlap property)
*   Tuning: W=1,R=N (fast writes) vs W=N,R=1 (fast reads)
*   Sloppy quorum: write to backup nodes if primary unavailable (AP in CAP)
*   Hinted handoff: backup node transfers data back to primary later

**Real-world:** Riak (tunable N/R/W), DynamoDB (configurable consistency), Cassandra (ONE/QUORUM/ALL)

---

## Num19 — Atomic Broadcast

**Topic:** Total-order broadcast, FIFO channels, ZAB (ZooKeeper Atomic Broadcast)

_To be implemented. See Num19AtomicBroadcastDemo() stub in demos.go_

**Core Concepts:**

*   Atomic broadcast: all nodes deliver messages in same total order
*   Properties: validity, uniform agreement, integrity, total order
*   ZAB: leader-based total order broadcast (basis of ZooKeeper)
*   State machine replication: same inputs in same order → same state

**Real-world:** ZooKeeper (ZAB), etcd (Raft log = total order), theoretical foundation of consensus

---

## Num20 — Raft: Edge Cases

**Topic:** Log compaction, snapshots, membership changes (joint consensus)

_To be implemented. See Num20RaftEdgeCasesDemo() stub in demos.go_

**Core Concepts:**

*   Log compaction: prevent disk exhaustion (Raft log grows forever)
*   Snapshots: compact log into state snapshot, discard old entries
*   InstallSnapshot RPC: chunk-based transfer for large snapshots
*   Membership changes: joint consensus (C\_old,new → C\_new) prevents split-brain
*   Other: leadership transfer, PreVote optimization

**Real-world:** etcd (log compaction), Consul (snapshot management), TiKV (region snapshots)

---

## Num21 — Paxos (Synod)

**Topic:** Propose/Promise/Accept/Learn, single-value consensus problem

_To be implemented. See Num21PaxosSynodDemo() stub in demos.go_

**Core Concepts:**

*   Original consensus algorithm (Lamport 1989), notoriously hard to understand
*   Three roles: Proposer, Acceptor, Learner
*   Phase 1 (Prepare/Promise): "can I propose with ballot N?"
*   Phase 2 (Accept): send value, acceptors accept if no higher promise
*   Safety: quorum intersection ensures only one value chosen
*   The "single value problem": Paxos only decides one value, not a log

**Real-world:** Google Chubby, Spanner, "Paxos Made Live" paper

---

## Num22 — Multi-Paxos

**Topic:** Leader stickiness, skipping Phase 1, high-throughput WAN consensus

_To be implemented. See Num22MultiPaxosDemo() stub in demos.go_

**Core Concepts:**

*   Basic Paxos runs 2 phases per value (4 RTTs) — very slow
*   Multi-Paxos: elect stable leader, skip Phase 1 for subsequent values
*   Leader runs Phase 1 once with "infinity ballot", then only Phase 2 (2 RTTs/value)
*   Log replication: use Multi-Paxos to agree on sequence of values
*   WAN optimization: batching multiple values in single Phase 2
*   Relation to Raft: Raft ≈ Multi-Paxos with stronger leader + simpler protocol

**Real-world:** Spanner (Multi-Paxos for WAN replication), Google F1, CockroachDB (originally)

---

## Num23 — CRDTs (Conflict-Free Replicated Data Types)

**Topic:** G-Counter, PN-Counter, LWW-Register, convergence without coordination

_To be implemented. See Num23CRDTsDemo() stub in demos.go_

**Core Concepts:**

*   Merge concurrent updates without coordination or conflicts
*   CRDT property: all replicas converge by applying all updates
*   G-Counter: grow-only counter, merge = sum (monotonic)
*   PN-Counter: positive-negative counter (P - N), supports inc/dec
*   LWW-Register: last-write-wins using timestamp
*   Types: state-based (merge state) vs op-based (commutative operations)
*   Trade-off: eventual consistency, cannot enforce strong invariants

**Real-world:** Figma (collaborative editing), Discord (presence), Riak, Automerge, Redis Enterprise

---

## Num24 — Linearizability Testing

**Topic:** Jepsen methodology, history verification, proving correctness

_To be implemented. See Num24LinearizabilityTestingDemo() stub in demos.go_

**Core Concepts:**

*   How do you PROVE your distributed system is correct?
*   Linearizability: strongest consistency — operations appear atomic and instantaneous
*   Jepsen: testing framework (run workload, inject faults, verify history)
*   Linearization check: build happened-before graph, check valid total order
*   Common bugs: split-brain, stale reads, lost writes, zombie processes

**Real-world:** Jepsen found bugs in MongoDB, Redis, Cassandra, etcd, Consul. Kyle Kingsbury's jepsen.io, chaos engineering (Netflix Chaos Monkey)

---

_Curriculum complete: 24 topics covering distributed systems fundamentals and consensus mechanisms._

```
Gossip:
- Eventual consistency (takes multiple rounds)
- No coordination (fully decentralized)
- High availability (works even if 50% nodes fail)
- Use case: Membership, metrics, caches

Consensus (Raft/Paxos):
- Strong consistency (quorum agreement)
- Requires leader + quorum (N/2+1 nodes)
- Moderate availability (requires majority)
- Use case: Leader election, distributed locks, logs

Trade-off: Gossip = AP (available + partition-tolerant)
           Consensus = CP (consistent + partition-tolerant)
```

```
Bad: All servers send metrics to central collector
→ Collector bottleneck: 1000 servers × 10 metrics/sec = 10,000 writes/sec

Good: Gossip-based aggregation
→ Each server gossips metrics to 3 random peers
→ Use CRDTs (commutative counters) for mergeable state
→ Converges in ~10 rounds (10 seconds)
→ No central bottleneck

Example metrics:
- Request count: CRDT counter (sum)
- CPU usage: Gossip + average
- Error rate: Gossip + percentage
```

```
- Gossip continues within each partition
- Nodes in partition A can't reach partition B
- When partition heals → gossip resumes across partitions
- Conflict resolution: use version numbers or vector clocks

Example:
Partition A writes: key=1 value="A" (v2)
Partition B writes: key=1 value="B" (v2) ← conflict!
→ Use vector clocks: [(A,2), (B,2)] = siblings
→ Application resolves conflict (e.g., last-write-wins or merge)
```

```
Broadcast = O(N²) messages (each node sends to all N nodes)
Gossip = O(N*K) per round, O(N*K*log N) total

Example (N=1000, K=3):
- Broadcast: 1,000,000 messages (network storm!)
- Gossip: 3,000 per round × 10 rounds = 30,000 messages

Gossip is 33× more efficient!
```

```
- Each node maintains heartbeat counters for all other nodes
- Nodes gossip every 1 second (share heartbeat counters)
- If Node-X's heartbeat hasn't increased in 3+ rounds → suspected failed
- Indirect probing (SWIM): ask peers to confirm before marking DOWN
- Phi Accrual Failure Detector: uses statistical model (not fixed threshold)
```

```
Client → API Gateway
         ├─ Circuit Breaker (fail after 3 timeouts)
         ├─ Timeout (2 seconds)
         ├─ Retry (1x with 100ms backoff)
         ├─ Bulkhead (10 threads for payment calls)
         └─ Fallback (return cached exchange rate or error)
```