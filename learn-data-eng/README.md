# M3/W3/D4 - Thu, 19 Mar 2026 (WIB)

## Data Engineering — Full Curriculum

A structured learning path for Staff/Principal Data Engineers at top-tier tech companies like **TikTok (ByteDance), Grab, GoTo/Gojek, Meta, Stripe, Airbnb, Uber, and Shopee**. Each topic is a runnable Go simulation focusing on data pipelines, query engines, streaming systems, and warehouse internals.

**Target audience:** Engineers who can write SQL and build basic ETL pipelines, but need to learn how data systems work **at scale** — understanding query optimization, distributed execution, and platform architecture.

This curriculum is designed to make you the most sought-after data engineer in Southeast Asia and globally. After mastering all 22 topics, you will command Staff/Principal-level compensation ($140K–$280K USD equivalent) and be able to design, build, and operate data platforms that serve billions of queries per day.

---

## Curriculum Table

| # | Topic | Core Concept | Real World at TikTok/Grab | Demo |
| --- | --- | --- | --- | --- |
| **Domain 1: Ingestion & Transport** |   |   |   |   |
| 01 | SQL Query Execution | Parse → Plan → Optimize → Execute, predicate pushdown | Why `SELECT *` on 1 TB kills Athena; the plan BigQuery uses | Num01 ✅ |
| 02 | Batch Ingestion Pipeline | ETL vs ELT, idempotency, checkpointing, pagination | Daily S3 → Hive ingestion of 100M GoTo ride events | Num02 🔲 |
| 03 | Change Data Capture (CDC) | WAL tailing, MySQL binlog, Debezium, at-least-once | Real-time order syncing: MySQL → Kafka → data lake | Num03 🔲 |
| **Domain 2: Transport & Messaging** |   |   |   |   |
| 04 | Kafka Internals | Producer acks, ISR, partition assignment, compaction | Grab: 50M ride events/day, zero message loss | Num04 🔲 |
| 05 | Stream Processing & Windows | Event time vs processing time, watermarks, tumbling | TikTok: real-time feed ranking, live view counts | Num05 🔲 |
| 06 | Stream-Table Duality | KTable, changelog streams, materialized views, joins | Join click stream with user profile table in real-time | Num06 🔲 |
| **Domain 3: Storage & Formats** |   |   |   |   |
| 07 | Columnar Storage (Parquet) | Row vs columnar, RLE encoding, dictionary encoding | 10x query speedup by switching MySQL → Parquet on S3 | Num07 🔲 |
| 08 | Data Lake Table Formats | Iceberg/Delta Lake ACID, time travel, schema evolution | Grab: atomic 5 TB partition swap without downtime | Num08 🔲 |
| 09 | Data Replication Patterns | Full refresh vs incremental, SCD Type 1/2/3, upsert | GoTo: slowly changing dimension for driver history | Num09 🔲 |
| **Domain 4: Compute Engines** |   |   |   |   |
| 10 | Batch Processing (MapReduce) | Map/Shuffle/Reduce, combiner, data skew, speculation | The model running inside Spark/Hive/Flink | Num10 🔲 |
| 11 | Distributed Shuffle & Skew | External sort, broadcast join, salting for skew | Spark job spilling to disk and crashing at 3am | Num11 🔲 |
| 12 | Data Warehouse Internals | MPP architecture, distribution keys, zone maps | Redshift/BigQuery: why some queries are 100x faster | Num12 🔲 |
| 13 | Real-Time OLAP (Pinot/Druid) | Pre-aggregation, segment generation, star-tree index | TikTok: "How many views in last 60s?" across 1B videos | Num13 🔲 |
| **Domain 5: Orchestration & Quality** |   |   |   |   |
| 14 | Pipeline Orchestration (DAGs) | Dependency resolution, backfill, idempotency, SLA | Airflow managing 10,000+ daily DAG runs at TikTok | Num14 🔲 |
| 15 | Data Quality & Contracts | Schema validation, anomaly detection, dbt tests | "Why did the revenue dashboard show $0 at 9am?" | Num15 🔲 |
| 16 | Data Lineage & Catalog | Column-level lineage, impact analysis, OpenLineage | Tracing bad revenue metric back to broken upstream | Num16 🔲 |
| **Domain 6: Architecture & Governance** |   |   |   |   |
| 17 | Cost-Based Query Optimizer | Statistics, cardinality estimation, join reordering | Spark picked wrong join — ran 3h instead of 10min | Num17 🔲 |
| 18 | Data Mesh & Governance | Domain ownership, data products, schema registry | Grab: 50 domains, each team owns data contracts | Num18 🔲 |
| 19 | Multi-Tenant Data Platform | Resource isolation, quota management, cost attribution | "Team A's Spark job killed Team B's pipeline" | Num19 🔲 |
| 20 | Lakehouse Architecture | Lambda vs Kappa vs Lakehouse, unified batch+stream | Meta/Airbnb: why they killed Lambda for Lakehouse | Num20 🔲 |
| 21 | Data Security & Privacy | Column encryption, row-level security, PII masking | GoTo: handling PII under Indonesian PDPA law | Num21 🔲 |
| 22 | End-to-End Performance Tuning | Query profiling, partition pruning, Z-ordering, cache | Going from 4-hour Spark job to 15-minute | Num22 🔲 |

**Legend:** ✅ = Fully implemented | 🔲 = Skeleton stub (to be implemented)

**Cross-references for connections to other curricula:**

*   **Kafka Internals (Num04)** → See `learn-distributed-consensus-system` Num03 (Distributed Queues) — same Kafka, different depth
*   **MapReduce Model (Num10)** → See `learn-distributed-consensus-system` Num01 (Consistent Hashing) — partitioning strategy
*   **Data Warehouse Internals (Num12)** → See `learn-infra-eng` Num12 (CSI Storage) — data lives on storage layer
*   **Pipeline Orchestration (Num14)** → See `learn-infra-eng` Num04 (Operator Pattern) — Airflow on K8s uses operators
*   **Multi-Tenant Platform (Num19)** → See `learn-infra-eng` Num10/19 — resource quotas and cost attribution

---

## How to run

```
# run all topics (only uncommented ones in main.go will execute)
go run ./learn-data-eng/
```

To run a specific topic, comment/uncomment the relevant line in `main.go`.

---

## Function signatures

```
func Num01SQLQueryExecution()          // ✅ fully implemented
func Num02BatchIngestionPipeline()     // 🔲 stub — to be implemented
func Num03CDCAndStreaming()            // 🔲 stub — to be implemented
func Num04KafkaInternals()             // 🔲 stub — to be implemented
func Num05StreamProcessingWindows()    // 🔲 stub — to be implemented
func Num06StreamTableDuality()         // 🔲 stub — to be implemented
func Num07ColumnarStorageParquet()     // 🔲 stub — to be implemented
func Num08DataLakeTableFormats()       // 🔲 stub — to be implemented
func Num09DataReplicationPatterns()    // 🔲 stub — to be implemented
func Num10MapReduceModel()             // 🔲 stub — to be implemented
func Num11DistributedShuffleSkew()     // 🔲 stub — to be implemented
func Num12DataWarehouseInternals()     // 🔲 stub — to be implemented
func Num13RealtimeOLAP()               // 🔲 stub — to be implemented
func Num14PipelineOrchestration()      // 🔲 stub — to be implemented
func Num15DataQualityContracts()       // 🔲 stub — to be implemented
func Num16DataLineageCatalog()         // 🔲 stub — to be implemented
func Num17CostBasedQueryOptimizer()    // 🔲 stub — to be implemented
func Num18DataMeshGovernance()         // 🔲 stub — to be implemented
func Num19MultiTenantDataPlatform()    // 🔲 stub — to be implemented
func Num20LakehouseArchitecture()      // 🔲 stub — to be implemented
func Num21DataSecurityPrivacy()        // 🔲 stub — to be implemented
func Num22PerformanceTuning()          // 🔲 stub — to be implemented
```

---

## Num01 — SQL Query Execution

### The Problem

You write a SQL query. You press Enter. The database returns results.

**But what happened in between?**

Most data engineers can write SQL. Almost none understand what the database does with that SQL string. This is why:

*   You write `SELECT * FROM orders` on a 1 TB table and Athena times out
*   A join that should take 10 seconds takes 10 minutes
*   Two queries that look identical have 100x different runtimes

Understanding query execution is the **#1 differentiator** between a mid-level data engineer and a Staff/Principal engineer.

### The Concept

SQL execution has **6 phases**:

```
1. Parsing       — SQL string → Abstract Syntax Tree (AST)
2. Analysis      — Resolve table/column names, type checking
3. Logical Plan  — AST → Relational algebra tree (Filter, Join, Project, Scan)
4. Optimization  — Apply rules (predicate pushdown, join reordering)
5. Physical Plan — Choose algorithms (hash join vs nested loop)
6. Execution     — Process rows (row-by-row vs vectorized batching)
```

#### Phase 1: Parsing

The SQL string is tokenized and converted into an Abstract Syntax Tree (AST):

```
SELECT u.name, COUNT(*) 
FROM users u 
JOIN orders o ON u.id = o.user_id 
WHERE o.status = 'completed'
```

Becomes:

```
Query
├── SelectList
│   ├── Column(u.name)
│   └── AggregateFunc(COUNT(*))
├── FromClause
│   ├── Table(users) AS u
│   └── Join(INNER)
│       └── Table(orders) AS o
└── WhereClause
    └── Equals(o.status, 'completed')
```

#### Phase 2: Analysis

The analyzer looks up table schemas in the catalog:

```
users:
  - id (INT, PRIMARY KEY)
  - name (VARCHAR)
  - email (VARCHAR)

orders:
  - id (INT, PRIMARY KEY)
  - user_id (INT, FOREIGN KEY → users.id)
  - status (VARCHAR)
  - amount (DECIMAL)
```

Type checking:

*   ✓ `u.id` matches `o.user_id` (both INT)
*   ✓ `o.status` is VARCHAR (can compare with 'completed')

#### Phase 3: Logical Plan

The AST is converted into a logical plan (relational algebra):

```
Aggregate(u.name, COUNT(*))
└── Project(u.name)
    └── Join(u.id = o.user_id)
        ├── Filter(o.status = 'completed')
        │   └── Scan(orders)
        └── Scan(users)
```

This is **database-agnostic** — it says **what** to do, not **how** to do it.

#### Phase 4: Optimization — Predicate Pushdown

**BEFORE optimization:**

```
Join(u.id = o.user_id)
├── Scan(orders)           [10,000,000 rows]
└── Scan(users)            [1,000,000 rows]
↓
Filter(o.status = 'completed')
```

Problem: Join processes **10M × 1M = 10 trillion row comparisons**, then filters.

**AFTER optimization (predicate pushdown):**

```
Join(u.id = o.user_id)
├── Filter(o.status = 'completed')
│   └── Scan(orders)       [10M rows → 50K after filter]
└── Scan(users)            [1M rows]
```

Now: Join processes **50K × 1M = 50 billion comparisons**.

**Speedup: 200x faster!**

This is the **most common optimization** in SQL engines.

#### Phase 5: Physical Plan

The optimizer chooses concrete algorithms:

```
HashJoin(u.id = o.user_id)
├── Build: Scan(users)      [build hash table: user_id → row]
└── Probe: Filter(o.status = 'completed')
           └── Scan(orders)
```

**Join algorithm choices:**

| Algorithm | Complexity | Use when |
| --- | --- | --- |
| Nested Loop Join | O(orders × users) = 10M × 1M = 10 trillion | One table is tiny (\< 100 rows) |
| Hash Join | O(orders + users) = 11M | Tables fit in memory |
| Sort-Merge Join | O(n log n) | Data already sorted on join key |

**Hash join is 1000x faster** than nested loop for large tables.

#### Phase 6: Execution — Vectorized vs Row-by-Row

**Row-by-Row (Traditional):**

```python
for row in orders:
    if row.status == 'completed':
        probe_hash_table(row.user_id)
        emit(row)
```

Problem: 1 row at a time = poor CPU cache utilization.

**Vectorized (Modern — Presto, DuckDB, ClickHouse):**

```python
for batch of 1024 rows:
    filter_batch(batch, status='completed')
    hash_probe_batch(batch, hash_table)
    emit(batch)
```

Benefit: Process 1024 rows in one pass. CPU cache stays warm. **10x faster.**

### What the demo shows

```
Step 1: Parsing — SQL string → AST
Step 2: Analysis — resolve table/column names, type checking
Step 3: Logical Plan — relational algebra tree
Step 4: Optimization — predicate pushdown (200x speedup)
Step 5: Physical Plan — choose hash join over nested loop
Step 6: Execution — vectorized (1024 rows/batch) vs row-by-row
Step 7: Results — 10 rows returned after filtering 10M rows
```

### Key Insight

```
Predicate pushdown = filter BEFORE join (not after)
Hash join is 1000x faster than nested loop for large tables
Vectorized execution = process 1024 rows at once (10x speedup)
Understanding EXPLAIN PLAN is the #1 SQL debugging skill
```

### Real World Usage

| System | Uses query execution for |
| --- | --- |
| **Presto/Trino** | Vectorized execution, predicate pushdown to Parquet readers |
| **Spark SQL** | Catalyst optimizer (cost-based join reordering) |
| **BigQuery** | Dremel execution engine (columnar scan + shuffle) |
| **DuckDB** | Vectorized query engine (10x faster than SQLite) |
| **ClickHouse** | Vectorized execution on compressed columnar blocks |

### Interview Tips

**"Why is my query slow?"**

*   Start with: "Let me look at the EXPLAIN PLAN"
*   Check for: full table scans, nested loop joins on large tables, no predicate pushdown
*   Example answer: "The join is using nested loop instead of hash join because statistics are stale. Run ANALYZE TABLE."

**"What is predicate pushdown?"**

*   "Applying filters as early as possible in the query plan — before joins or aggregations"
*   "Without pushdown, you join 10M rows then filter. With pushdown, you filter to 50K rows then join. 200x faster."

**"Explain vectorized execution."**

*   "Traditional engines process 1 row at a time. Vectorized engines process batches of 1024 rows at once."
*   "Benefits: better CPU cache utilization, SIMD instructions, lower function call overhead. 10x speedup."

**"How would you optimize this query?" (given a slow query)**

*   Step 1: Check EXPLAIN PLAN for full table scans
*   Step 2: Add indexes on join keys and WHERE columns
*   Step 3: Check if statistics are up-to-date (ANALYZE TABLE)
*   Step 4: Rewrite to push predicates before joins
*   Step 5: If data skew, consider broadcast join or bucketing

---

## Num02 — Batch Ingestion Pipeline

### The Problem

**To be implemented** — this section will cover ETL vs ELT trade-offs, idempotency patterns, checkpointing for fault tolerance, and handling schema evolution.

---

## Num03 — Change Data Capture (CDC)

### The Problem

**To be implemented** — this section will cover WAL tailing, MySQL binlog parsing, Debezium connectors, at-least-once vs exactly-once delivery, and handling schema changes.

---

## Num04 — Kafka Internals

### The Problem

**To be implemented** — this section will cover producer acknowledgment modes (acks=0/1/all), in-sync replica (ISR) logic, consumer group rebalancing, partition assignment strategies, and log compaction.

---

## Num05 — Stream Processing & Windows

### The Problem

**To be implemented** — this section will cover event time vs processing time, watermarks for handling late data, tumbling/sliding/session windows, and exactly-once semantics in Flink/Kafka Streams.

---

## Num06 — Stream-Table Duality

### The Problem

**To be implemented** — this section will cover KTable vs KStream, changelog streams, materialized views, stateful joins between streams and tables, and Kafka Streams topologies.

---

## Num07 — Columnar Storage (Parquet)

### The Problem

**To be implemented** — this section will cover row-oriented vs column-oriented storage, run-length encoding (RLE), dictionary encoding, predicate pushdown to storage, and why Parquet is 10x faster than CSV.

---

## Num08 — Data Lake Table Formats

### The Problem

**To be implemented** — this section will cover Iceberg/Delta Lake ACID transactions on object storage, time travel, schema evolution, partition evolution, and compaction strategies.

---

## Num09 — Data Replication Patterns

### The Problem

**To be implemented** — this section will cover full refresh vs incremental replication, slowly changing dimensions (SCD Type 1/2/3), upsert/merge patterns, and handling deletes.

---

## Num10 — Batch Processing (MapReduce)

### The Problem

**To be implemented** — this section will cover the Map/Shuffle/Reduce programming model, combiner pattern for pre-aggregation, handling data skew, speculative execution, and how Spark implements this model.

---

## Num11 — Distributed Shuffle & Skew

### The Problem

**To be implemented** — this section will cover external sort for large datasets, shuffle partition tuning, broadcast join optimization, salting for data skew, and debugging Spark shuffle spill.

---

## Num12 — Data Warehouse Internals

### The Problem

**To be implemented** — this section will cover Massively Parallel Processing (MPP) architecture, distribution keys (hash vs round-robin), sort keys, zone maps, query compilation vs interpretation, and why Redshift/BigQuery are fast.

---

## Num13 — Real-Time OLAP (Pinot/Druid)

### The Problem

**To be implemented** — this section will cover pre-aggregation strategies, segment generation from streaming data, star-tree index for sub-second queries, and handling 1B records with \< 100ms latency.

---

## Num14 — Pipeline Orchestration (DAGs)

### The Problem

**To be implemented** — this section will cover Directed Acyclic Graphs (DAGs), dependency resolution, backfill strategies, idempotent task design, SLA alerting, and how Airflow schedules 10,000+ daily tasks.

---

## Num15 — Data Quality & Contracts

### The Problem

**To be implemented** — this section will cover schema validation (Great Expectations), anomaly detection (row count spikes), dbt tests, data contracts (SLAs between teams), and debugging "why did the revenue dashboard show $0?"

---

## Num16 — Data Lineage & Catalog

### The Problem

**To be implemented** — this section will cover column-level lineage, impact analysis ("which dashboards break if I change this table?"), OpenLineage standard, and DataHub/Amundsen catalog metadata.

---

## Num17 — Cost-Based Query Optimizer

### The Problem

**To be implemented** — this section will cover statistics collection (histograms, NDV), cardinality estimation, join reordering (left-deep vs bushy trees), cost model tuning, and why Spark picked the wrong join order.

---

## Num18 — Data Mesh & Governance

### The Problem

**To be implemented** — this section will cover domain ownership model, data products as APIs, schema registry (Avro/Protobuf), federated governance, and how Grab manages 50 domains with data contracts.

---

## Num19 — Multi-Tenant Data Platform

### The Problem

**To be implemented** — this section will cover resource isolation (YARN queues, Spark dynamic allocation), quota management, cost attribution/chargeback, workload management, and preventing "Team A's Spark job killed Team B's pipeline."

---

## Num20 — Lakehouse Architecture

### The Problem

**To be implemented** — this section will cover Lambda vs Kappa vs Lakehouse architectures, unified batch+stream processing, open table formats (Iceberg/Delta), and why Meta/Airbnb killed the Lambda architecture.

---

## Num21 — Data Security & Privacy

### The Problem

**To be implemented** — this section will cover column-level encryption, row-level security (RLS), PII masking/tokenization, GDPR/PDPA compliance, and handling driver/customer PII under Indonesian data protection law.

---

## Num22 — End-to-End Performance Tuning

### The Problem

**To be implemented** — this section will cover query profiling (Spark UI, EXPLAIN ANALYZE), partition pruning, Z-ordering, caching layers (Delta Cache, Alluxio), index strategies, and going from a 4-hour Spark job to 15 minutes.

---

---

## What to Expect in a Staff/Principal Data Engineer Interview

After mastering all 22 topics, here's what your interviews will look like at companies like TikTok, Grab, GoTo, Meta, Stripe, Airbnb, and Uber.

### System Design Round (60–90 min)

**Typical prompt:**

> _"Design a data platform for a ride-hailing company with 100M daily active users. Handle real-time analytics, batch reporting, and ML feature pipelines."_

**What they're testing:**

Can you decompose this into layers and justify architectural choices?

*   **Ingestion layer** (Num02, Num03) — CDC from MySQL → Kafka, batch S3 ingestion
*   **Transport layer** (Num04) — Kafka for event streaming (50M ride events/day)
*   **Stream processing** (Num05, Num06) — Real-time aggregations (Flink), KTables for user state
*   **Storage layer** (Num07, Num08) — Data lake on S3 with Iceberg/Delta Lake
*   **Compute layer** (Num10, Num12) — Spark for batch, Presto for ad-hoc queries
*   **Real-time OLAP** (Num13) — Pinot for dashboard queries (\< 100ms latency)
*   **Orchestration** (Num14) — Airflow DAGs, SLA monitoring
*   **Quality & Lineage** (Num15, Num16) — dbt tests, OpenLineage for impact analysis
*   **Governance** (Num18, Num21) — Data mesh (each domain owns contracts), PII masking
*   **Architecture** (Num20) — Lakehouse (unified batch+stream, no Lambda)

**The one answer that separates Staff from Senior:**

> "I'd use a Lakehouse architecture (Num20) instead of Lambda. Why? Lambda requires maintaining two separate pipelines (batch + stream) with different codebases. With Lakehouse (Iceberg on S3 + Spark Structured Streaming), I write the pipeline once and it handles both batch backfill and real-time updates. Trade-off: Iceberg transactions add 100ms latency vs raw Kafka, but that's acceptable for analytics (not user-facing APIs). This reduces maintenance burden by 50% and eliminates batch/stream inconsistency bugs."

This answer shows you've **operated** systems at scale, not just read blog posts.

### SQL & Query Optimization Round (45–60 min)

**Typical prompts:**

_"This query takes 3 hours. Optimize it."_

**Expected approach:**

*   Run `EXPLAIN` to see the query plan
*   Check for full table scans → add indexes or partition pruning
*   Check for nested loop joins on large tables → rewrite to hash join
*   Check for Cartesian products (missing join condition)
*   Check statistics freshness → run `ANALYZE TABLE`
*   Apply predicate pushdown manually if optimizer missed it

_"Write a SQL query to find the top 10 users by total order amount, excluding canceled orders."_

**Expected answer:**

_"Explain the difference between WHERE and HAVING."_

**Expected answer:**

*   `WHERE` filters rows **before** aggregation (applies to raw rows)
*   `HAVING` filters groups **after** aggregation (applies to aggregate results like `COUNT(*)`, `SUM()`)
*   Example: `WHERE amount > 100` filters orders, `HAVING COUNT(*) > 5` filters groups

### Data Modeling Round (45 min)

**Typical prompts:**

_"Design a star schema for an e-commerce analytics warehouse."_

**Expected answer:**

*   **Fact table:** `orders` (order\_id, user\_id, product\_id, amount, order\_date)
*   **Dimension tables:**
    *   `users` (user\_id, name, email, signup\_date)
    *   `products` (product\_id, name, category, price)
    *   `dates` (date\_id, date, year, month, day, quarter)
*   Foreign keys from fact table to dimension tables
*   Why star schema: simple joins, fast aggregations, easy for BI tools

_"How do you handle slowly changing dimensions (SCD)?"_

**Expected answer (Num09):**

*   **Type 1:** Overwrite old value (no history) — `UPDATE users SET city = 'New York'`
*   **Type 2:** Create new row with version (full history) — add `valid_from`, `valid_to`, `is_current` columns
*   **Type 3:** Add column for old value (limited history) — `previous_city` column
*   Most common: Type 2 for auditing, Type 1 for non-critical attributes

### Debugging Round (30–45 min)

**Typical prompt:**

> _"The daily revenue report shows $0 at 9am. How do you debug this?"_

**Expected debugging flow:**

**Check data lineage (Num16):** Trace the revenue metric back through transformations

*   Dashboard reads from `revenue_daily` table
*   `revenue_daily` is populated by `calculate_revenue` DAG
*   That DAG reads from `orders` table

**Check pipeline status (Num14):** Is the DAG running?

*   Check Airflow: `calculate_revenue` DAG failed at 8:30am
*   Error: "Table `orders` is empty"

**Check upstream data (Num03):** Why is `orders` empty?

*   Check CDC pipeline: Debezium connector is down
*   MySQL binlog hasn't been read since yesterday

**Check data quality (Num15):** Add checks to prevent this

*   Add dbt test: `assert row_count(orders) > 1000`
*   Add Airflow sensor: wait for upstream `orders` table before running

**The answer that impresses:**

> "I'd trace the lineage from revenue dashboard → revenue\_daily table → calculate\_revenue DAG → orders table using OpenLineage (Num16). Check if the DAG ran (Airflow UI), then check if upstream data arrived (row count on orders). If orders is empty, check CDC pipeline health (Num03). For prevention, add data quality checks (Num15): assert row\_count > threshold, assert max(order\_date) is within last 2 hours, and alert on SLA breach. This 3-layer defense (lineage + quality checks + SLA monitoring) catches 95% of data bugs before they reach dashboards."

### Coding Round (60–90 min)

**Typical prompt:**

> _"Write a Python/Go script to read a 10 GB CSV file, deduplicate rows by user\_id, and write to Parquet."_

**Expected solution (Python):**

```python
import pandas as pd

# Read CSV in chunks (avoid OOM)
chunks = []
for chunk in pd.read_csv('input.csv', chunksize=100000):
    # Deduplicate within chunk
    chunk = chunk.drop_duplicates(subset=['user_id'], keep='last')
    chunks.append(chunk)

# Concatenate and deduplicate across chunks
df = pd.concat(chunks).drop_duplicates(subset=['user_id'], keep='last')

# Write to Parquet
df.to_parquet('output.parquet', compression='snappy', index=False)
```

**Follow-up:** "What if the file is 1 TB and doesn't fit in memory?"

**Expected answer:**

*   Use Spark: `spark.read.csv('input.csv').dropDuplicates(['user_id']).write.parquet('output.parquet')`
*   Spark partitions the data, deduplicates in parallel, and spills to disk if needed

### Compensation Signal

At firms like GoTo, Grab, TikTok, Meta, Stripe:

| Level | Knows | Compensation (USD equivalent) |
| --- | --- | --- |
| **Mid-Level** | Writes SQL, builds basic ETL | $80K–$120K |
| **Senior** | Optimizes queries, designs pipelines | $120K–$180K |
| **Staff** | Designs platforms, understands internals | $160K–$240K |
| **Principal** | Architectures data strategy, org-wide impact | $200K–$300K |

**The differentiators that command Staff/Principal offers:**

**Real-time OLAP (Num13)** — Pinot/Druid expertise is rare in SEA. Every tech company wants sub-second dashboards but \< 5% of engineers can architect it. Worth $20K–$40K in negotiation leverage.

**Data Mesh (Num18)** — Understanding domain ownership and federated governance separates platform architects from pipeline builders. This is what gets you Principal title.

**Cost-Based Optimization (Num17)** — Knowing _why_ Spark chose a bad plan and _how_ to fix it (statistics, join hints) is the #1 interview differentiator at Meta/Stripe.

### The One Answer That Proves You're Staff/Principal-Level

**Question:** _"How would you reduce the cost of our data platform by 50%?"_

**Senior answer:** "Use Spot instances for Spark."

**Staff/Principal answer:**

> "Multi-layered cost optimization (Num19, Num22):
> 
> **1\. Storage tier (Num07, Num08):** Switch from raw Parquet to Iceberg with Z-ordering (Num22). This enables aggressive partition pruning — queries scan 10% of data instead of 100%. Estimated savings: 30% on storage I/O costs.
> 
> **2\. Compute tier (Num10, Num11):** Tune Spark shuffle partitions to reduce shuffle spill (Num11). Most jobs at 10TB scale need 2000–4000 partitions, not the default 200. This reduces disk I/O and cluster size. Use broadcast joins for tables \< 1 GB (Num17). Estimated savings: 40% on Spark compute.
> 
> **3\. Query layer (Num01, Num12):** Add materialized views for common aggregations (Num06) — daily revenue, user cohorts. These pre-compute expensive joins. Cache hot data in Alluxio (Num22). Estimated savings: 60% on Presto query costs.
> 
> **4\. Orchestration tier (Num14):** Deduplicate redundant DAGs — found 3 teams computing the same 'active users' metric independently. Consolidate into one shared dataset. Estimated savings: 20% on pipeline duplication.
> 
> **5\. Multi-tenancy (Num19):** Implement chargeback. Once teams see their costs, they self-optimize. Set resource quotas per team to prevent runaway jobs.
> 
> Combined savings: 50–70%. This is how Meta reduced Presto costs by $100M/year."

This answer shows you've **optimized production systems at scale** — that's the Principal level.

---

---

## Why This Curriculum Makes You Hireable

### The Market Reality (2026)

Most "data engineers" know:

*   Write SQL queries
*   Build ETL pipelines with Airflow
*   Maybe some Spark

**They cannot:**

*   Explain why their query is slow (query plan, predicate pushdown)
*   Design a real-time analytics pipeline (Kafka + Flink + Pinot)
*   Debug "revenue dashboard shows $0" (lineage, quality, CDC lag)
*   Reduce data platform costs by 50% (partitioning, caching, optimizer tuning)

These 22 topics cover **exactly the gaps** between "I write SQL" and "I architect data platforms."

### The Interview Delta

After this curriculum, when an interviewer asks:

> _"Design a data platform for 100M daily active users"_

**Before:** "Uh... Airflow... and Spark... and maybe a data lake?"

**After:** "Lakehouse architecture (Num20) with CDC (Num03) → Kafka (Num04) → Flink (Num05) for real-time, Spark (Num10) for batch, Iceberg (Num08) on S3 for unified storage, Pinot (Num13) for sub-second dashboards, Airflow (Num14) for orchestration, dbt tests (Num15) for quality, OpenLineage (Num16) for impact analysis, and data mesh (Num18) for domain ownership. Let me draw the architecture and explain the failure modes of each layer."

That's the difference between a $120K offer and a $220K offer.

### The Unique Value of Topics 13, 18, 20, 21, 22

These 5 topics are **missing from 90% of data engineering learning resources**. They're also the topics that:

*   **Real-Time OLAP (Num13):** Every company wants Pinot/Druid but \< 5% of DEs in SEA can architect it. Early-mover advantage.
*   **Data Mesh (Num18):** The shift from centralized to federated data platforms. This is what gets you Principal title — understanding organizational architecture.
*   **Lakehouse (Num20):** The convergence of data lake + warehouse. Meta, Airbnb, Stripe all moved here. Knowing _why_ separates you.
*   **Security/Privacy (Num21):** GDPR/PDPA compliance is now table stakes at top firms. Indonesian companies (GoTo, Grab, TikTok) explicitly require this.
*   **Performance Tuning (Num22):** The synthesizing topic — going from 4-hour job to 15-minute is a real interview case study at Grab/Stripe.

These are **not toy topics**. They're the difference between Senior and Principal.

---

## License

This curriculum is open for educational use. Go build platforms. Go get Staff/Principal offers at TikTok, Grab, GoTo, Meta, Stripe.

```
SELECT user_id, SUM(amount) as total_spent
FROM orders
WHERE status != 'canceled'
GROUP BY user_id
ORDER BY total_spent DESC
LIMIT 10
```