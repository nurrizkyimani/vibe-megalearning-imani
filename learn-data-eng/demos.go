package main

import (
	"fmt"
	"strings"
	"time"
)

// =============================================================================
// Num01SQLQueryExecution
//
// TOPIC: SQL Query Execution — Parse → Plan → Optimize → Execute
//   - Parsing: tokenize SQL string → Abstract Syntax Tree (AST)
//   - Analysis: resolve table/column names, type checking
//   - Logical Plan: relational algebra tree (Filter → Join → Project → Scan)
//   - Physical Plan: choose join algorithms (hash join vs nested loop)
//   - Optimization: predicate pushdown, join reordering, constant folding
//   - Execution: row-by-row vs vectorized (columnar batch) processing
//
// Real world: Why `SELECT *` on a 1 TB table kills Athena performance
// TikTok/Grab: Understanding the query plan that BigQuery/Presto uses
// =============================================================================
func Num01SQLQueryExecution() {
	fmt.Println("============================================================")
	fmt.Println("  Num01 -- SQL Query Execution")
	fmt.Println("============================================================")
	fmt.Println()

	// ── Example Query ────────────────────────────────────────────────────────
	query := `
SELECT u.name, COUNT(*) as order_count, SUM(o.amount) as total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.status = 'completed' AND o.amount > 100
GROUP BY u.name
HAVING COUNT(*) > 5
ORDER BY total_spent DESC
LIMIT 10
`
	fmt.Println("[SQL Query]")
	fmt.Println(query)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 1: Parsing (Tokenization + AST) ────────────────────────────────
	// [SQL Parser] Convert raw string → Abstract Syntax Tree
	// Tokens: SELECT, u.name, COUNT, FROM, users, ...
	fmt.Println("── Step 1: Parsing (SQL → Abstract Syntax Tree) ────────────")
	fmt.Println()
	fmt.Println("[Tokenization]")
	tokens := []string{"SELECT", "u.name", "COUNT(*)", "FROM", "users", "JOIN", "orders", "WHERE", "o.status='completed'"}
	fmt.Printf("  Tokens: %s\n", strings.Join(tokens, ", "))
	fmt.Println()
	fmt.Println("[Abstract Syntax Tree (AST)]")
	fmt.Println("  Query")
	fmt.Println("  ├── SelectList")
	fmt.Println("  │   ├── Column(u.name)")
	fmt.Println("  │   ├── AggregateFunc(COUNT(*))")
	fmt.Println("  │   └── AggregateFunc(SUM(o.amount))")
	fmt.Println("  ├── FromClause")
	fmt.Println("  │   ├── Table(users) AS u")
	fmt.Println("  │   └── Join(INNER)")
	fmt.Println("  │       └── Table(orders) AS o")
	fmt.Println("  ├── WhereClause")
	fmt.Println("  │   ├── AND")
	fmt.Println("  │   │   ├── Equals(o.status, 'completed')")
	fmt.Println("  │   │   └── GreaterThan(o.amount, 100)")
	fmt.Println("  ├── GroupByClause")
	fmt.Println("  │   └── Column(u.name)")
	fmt.Println("  ├── HavingClause")
	fmt.Println("  │   └── GreaterThan(COUNT(*), 5)")
	fmt.Println("  └── OrderByClause")
	fmt.Println("      └── Column(total_spent) DESC")
	fmt.Println()
	fmt.Println("Key insight: AST is a tree representation of the query")
	fmt.Println("             Each node is a SQL operation (Filter, Join, Aggregate)")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 2: Analysis (Semantic Validation) ──────────────────────────────
	// [Semantic Analyzer] Resolve table/column names, check types
	fmt.Println("── Step 2: Analysis (Semantic Validation) ──────────────────")
	fmt.Println()
	fmt.Println("[Catalog Lookup]")
	fmt.Println("  users table:")
	fmt.Println("    - id        INT      PRIMARY KEY")
	fmt.Println("    - name      VARCHAR  NOT NULL")
	fmt.Println("    - email     VARCHAR")
	fmt.Println("    - created_at TIMESTAMP")
	fmt.Println()
	fmt.Println("  orders table:")
	fmt.Println("    - id        INT      PRIMARY KEY")
	fmt.Println("    - user_id   INT      FOREIGN KEY → users.id")
	fmt.Println("    - amount    DECIMAL")
	fmt.Println("    - status    VARCHAR")
	fmt.Println("    - order_date TIMESTAMP")
	fmt.Println()
	fmt.Println("[Type Checking]")
	fmt.Println("  ✓ u.id matches o.user_id       (both INT)")
	fmt.Println("  ✓ o.status is VARCHAR          (can compare with 'completed')")
	fmt.Println("  ✓ o.amount is DECIMAL          (can compare with 100)")
	fmt.Println("  ✓ COUNT(*) returns INT         (can compare with 5)")
	fmt.Println()
	fmt.Println("Key insight: Analysis phase catches errors like:")
	fmt.Println("             'column does not exist', 'type mismatch'")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 3: Logical Plan (Relational Algebra) ───────────────────────────
	// [Query Planner] Convert AST → Logical Plan (relational operators)
	fmt.Println("── Step 3: Logical Plan (Relational Algebra) ───────────────")
	fmt.Println()
	fmt.Println("[Logical Plan Tree]")
	fmt.Println("  Limit(10)")
	fmt.Println("  └── Sort(total_spent DESC)")
	fmt.Println("      └── Filter(COUNT(*) > 5)              [HAVING]")
	fmt.Println("          └── Aggregate(u.name, COUNT(*), SUM(o.amount))")
	fmt.Println("              └── Project(u.name, o.amount)")
	fmt.Println("                  └── Join(u.id = o.user_id)")
	fmt.Println("                      ├── Filter(o.status = 'completed' AND o.amount > 100)")
	fmt.Println("                      │   └── Scan(orders)")
	fmt.Println("                      └── Scan(users)")
	fmt.Println()
	fmt.Println("Key insight: Logical plan is database-agnostic")
	fmt.Println("             It says 'what to do', not 'how to do it'")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 4: Optimization (Predicate Pushdown) ───────────────────────────
	// [Optimizer] Apply rules to make query faster
	fmt.Println("── Step 4: Optimization (Predicate Pushdown) ───────────────")
	fmt.Println()
	fmt.Println("[BEFORE Optimization]")
	fmt.Println("  Join(u.id = o.user_id)")
	fmt.Println("  ├── Scan(orders)          [10,000,000 rows]")
	fmt.Println("  └── Scan(users)           [1,000,000 rows]")
	fmt.Println("  ↓")
	fmt.Println("  Filter(o.status = 'completed' AND o.amount > 100)")
	fmt.Println()
	fmt.Println("Problem: Join processes 10M × 1M = 10 trillion row comparisons!")
	fmt.Println()
	fmt.Println("[AFTER Optimization — Predicate Pushdown]")
	fmt.Println("  Join(u.id = o.user_id)")
	fmt.Println("  ├── Filter(o.status = 'completed' AND o.amount > 100)")
	fmt.Println("  │   └── Scan(orders)      [10,000,000 rows → 50,000 after filter]")
	fmt.Println("  └── Scan(users)           [1,000,000 rows]")
	fmt.Println()
	fmt.Println("Now: Join processes 50K × 1M = 50 billion comparisons")
	fmt.Println("     Speedup: 200x faster!")
	fmt.Println()
	fmt.Println("Key insight: Predicate pushdown = apply filters BEFORE joins")
	fmt.Println("             This is the most common optimization in SQL engines")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 5: Physical Plan (Choose Algorithms) ───────────────────────────
	// [Physical Planner] Choose concrete algorithms (hash join vs nested loop)
	fmt.Println("── Step 5: Physical Plan (Choose Algorithms) ───────────────")
	fmt.Println()
	fmt.Println("[Physical Plan]")
	fmt.Println("  TopN(10, total_spent DESC)           [use heap, not full sort]")
	fmt.Println("  └── Filter(COUNT(*) > 5)")
	fmt.Println("      └── HashAggregate(u.name)")
	fmt.Println("          └── Project(u.name, o.amount)")
	fmt.Println("              └── HashJoin(u.id = o.user_id)")
	fmt.Println("                  ├── Build: Scan(users)")
	fmt.Println("                  │          [build hash table: user_id → name]")
	fmt.Println("                  └── Probe: Filter(o.status='completed', o.amount>100)")
	fmt.Println("                             └── Scan(orders)")
	fmt.Println()
	fmt.Println("[Join Algorithm Choice]")
	fmt.Println("  Option 1: Nested Loop Join")
	fmt.Println("    - For each order, scan all users")
	fmt.Println("    - Complexity: O(orders × users) = 10M × 1M = 10 trillion ops")
	fmt.Println("    - Use when: one table is tiny (< 100 rows)")
	fmt.Println()
	fmt.Println("  Option 2: Hash Join (CHOSEN)")
	fmt.Println("    - Build hash table on smaller table (users)")
	fmt.Println("    - Probe with larger table (orders)")
	fmt.Println("    - Complexity: O(orders + users) = 10M + 1M = 11M ops")
	fmt.Println("    - Use when: tables fit in memory")
	fmt.Println()
	fmt.Println("Key insight: Hash join is 1000x faster than nested loop")
	fmt.Println("             for large tables that fit in memory")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 6: Execution (Row-by-Row vs Vectorized) ────────────────────────
	// [Execution Engine] Process the query
	fmt.Println("── Step 6: Execution (Row-by-Row vs Vectorized) ────────────")
	fmt.Println()
	fmt.Println("[Row-by-Row Execution (Traditional)]")
	fmt.Println("  for each row in orders:")
	fmt.Println("    if row.status == 'completed' and row.amount > 100:")
	fmt.Println("      hash_table_probe(row.user_id)")
	fmt.Println("      emit(row)")
	fmt.Println()
	fmt.Println("  Problem: 1 row at a time = poor CPU cache utilization")
	fmt.Println("           50,000 function calls, 50,000 if-statements")
	fmt.Println()
	fmt.Println("[Vectorized Execution (Modern — Presto, DuckDB, ClickHouse)]")
	fmt.Println("  for each batch of 1024 rows:")
	fmt.Println("    filter_batch(rows, status='completed', amount>100)")
	fmt.Println("    hash_probe_batch(rows, hash_table)")
	fmt.Println("    emit(batch)")
	fmt.Println()
	fmt.Println("  Benefit: Process 1024 rows in one pass")
	fmt.Println("           CPU cache stays warm, SIMD instructions")
	fmt.Println("           10x faster than row-by-row")
	fmt.Println()
	fmt.Println("Key insight: Vectorized execution = columnar batching")
	fmt.Println("             This is why Parquet + DuckDB is 10x faster than CSV + MySQL")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Step 7: Simulation Results ──────────────────────────────────────────
	fmt.Println("── Step 7: Query Results ────────────────────────────────────")
	fmt.Println()
	fmt.Println("[Execution Stats]")
	fmt.Println("  Rows scanned (orders):        10,000,000")
	fmt.Println("  Rows after filter:               50,000")
	fmt.Println("  Rows scanned (users):         1,000,000")
	fmt.Println("  Rows after join:                 50,000")
	fmt.Println("  Rows after aggregation:           1,200")
	fmt.Println("  Rows after HAVING:                  300")
	fmt.Println("  Rows returned (LIMIT 10):            10")
	fmt.Println()
	fmt.Println("[Sample Results]")
	fmt.Println("  ┌──────────────────┬─────────────┬─────────────┐")
	fmt.Println("  │ name             │ order_count │ total_spent │")
	fmt.Println("  ├──────────────────┼─────────────┼─────────────┤")
	fmt.Println("  │ Alice Johnson    │          47 │   $234,890  │")
	fmt.Println("  │ Bob Williams     │          39 │   $198,450  │")
	fmt.Println("  │ Carol Martinez   │          35 │   $187,320  │")
	fmt.Println("  │ David Brown      │          28 │   $156,780  │")
	fmt.Println("  │ Emma Davis       │          26 │   $145,230  │")
	fmt.Println("  │ Frank Garcia     │          23 │   $134,560  │")
	fmt.Println("  │ Grace Lee        │          21 │   $128,940  │")
	fmt.Println("  │ Henry Wilson     │          19 │   $119,870  │")
	fmt.Println("  │ Iris Anderson    │          17 │   $112,340  │")
	fmt.Println("  │ Jack Thomas      │          15 │   $108,920  │")
	fmt.Println("  └──────────────────┴─────────────┴─────────────┘")
	fmt.Println()
	fmt.Println("[Query Execution Time]")
	fmt.Println("  Without optimization (filter after join):   180 seconds")
	fmt.Println("  With optimization (predicate pushdown):      0.9 seconds")
	fmt.Println("  Speedup: 200x faster")
	fmt.Println()
	time.Sleep(500 * time.Millisecond)

	// ── Summary ──────────────────────────────────────────────────────────────
	fmt.Println("============================================================")
	fmt.Println("  Key Insights:")
	fmt.Println("  1. SQL execution has 6 phases:")
	fmt.Println("     Parse → Analyze → Logical Plan → Optimize → Physical Plan → Execute")
	fmt.Println("  2. Predicate pushdown = filter BEFORE join (200x speedup)")
	fmt.Println("  3. Hash join is 1000x faster than nested loop for large tables")
	fmt.Println("  4. Vectorized execution = process 1024 rows at once (10x speedup)")
	fmt.Println("  5. Understanding EXPLAIN PLAN is the #1 SQL debugging skill")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num02BatchIngestionPipeline
//
// TOPIC: Batch Ingestion — ETL vs ELT, Idempotency, Checkpointing
//   - stub — implement when teaching
//
// Real world: Daily S3 → Hive ingestion of 100M GoTo ride events
// =============================================================================
func Num02BatchIngestionPipeline() {
	// stub — implement when teaching
}

// =============================================================================
// Num03CDCAndStreaming
//
// TOPIC: Change Data Capture (CDC) — WAL tailing, Binlog, Debezium
//   - stub — implement when teaching
//
// Real world: Real-time order syncing: MySQL → Kafka → data lake
// =============================================================================
func Num03CDCAndStreaming() {
	// stub — implement when teaching
}

// =============================================================================
// Num04KafkaInternals
//
// TOPIC: Kafka Internals — Producer acks, ISR, partition assignment, compaction
//   - stub — implement when teaching
//
// Real world: Grab: 50M ride events/day, zero message loss
// =============================================================================
func Num04KafkaInternals() {
	// stub — implement when teaching
}

// =============================================================================
// Num05StreamProcessingWindows
//
// TOPIC: Stream Processing & Windows — Event time, watermarks, tumbling/sliding
//   - stub — implement when teaching
//
// Real world: TikTok: real-time feed ranking, live view counts
// =============================================================================
func Num05StreamProcessingWindows() {
	// stub — implement when teaching
}

// =============================================================================
// Num06StreamTableDuality
//
// TOPIC: Stream-Table Duality — KTable, changelog streams, stateful joins
//   - stub — implement when teaching
//
// Real world: Join click stream with user profile table in real-time
// =============================================================================
func Num06StreamTableDuality() {
	// stub — implement when teaching
}

// =============================================================================
// Num07ColumnarStorageParquet
//
// TOPIC: Columnar Storage (Parquet) — Row vs columnar, RLE, dictionary encoding
//   - stub — implement when teaching
//
// Real world: 10x query speedup by switching MySQL → Parquet on S3
// =============================================================================
func Num07ColumnarStorageParquet() {
	// stub — implement when teaching
}

// =============================================================================
// Num08DataLakeTableFormats
//
// TOPIC: Data Lake Table Formats — Iceberg/Delta Lake ACID, time travel
//   - stub — implement when teaching
//
// Real world: Grab: atomic 5 TB partition swap without pipeline downtime
// =============================================================================
func Num08DataLakeTableFormats() {
	// stub — implement when teaching
}

// =============================================================================
// Num09DataReplicationPatterns
//
// TOPIC: Data Replication Patterns — Full refresh, incremental, SCD Type 1/2/3
//   - stub — implement when teaching
//
// Real world: GoTo: slowly changing dimension for driver onboarding history
// =============================================================================
func Num09DataReplicationPatterns() {
	// stub — implement when teaching
}

// =============================================================================
// Num10MapReduceModel
//
// TOPIC: Batch Processing (MapReduce) — Map/Shuffle/Reduce, combiner, data skew
//   - stub — implement when teaching
//
// Real world: The model running inside Spark/Hive
// =============================================================================
func Num10MapReduceModel() {
	// stub — implement when teaching
}

// =============================================================================
// Num11DistributedShuffleSkew
//
// TOPIC: Distributed Shuffle & Skew — External sort, broadcast join, salting
//   - stub — implement when teaching
//
// Real world: Spark job spilling to disk and crashing at 3am
// =============================================================================
func Num11DistributedShuffleSkew() {
	// stub — implement when teaching
}

// =============================================================================
// Num12DataWarehouseInternals
//
// TOPIC: Data Warehouse Internals — MPP, distribution keys, sort keys, zone maps
//   - stub — implement when teaching
//
// Real world: Redshift/BigQuery: why some queries are 100x faster
// =============================================================================
func Num12DataWarehouseInternals() {
	// stub — implement when teaching
}

// =============================================================================
// Num13RealtimeOLAP
//
// TOPIC: Real-Time OLAP (Pinot/Druid) — Pre-aggregation, star-tree index
//   - stub — implement when teaching
//
// Real world: TikTok: "How many views in last 60 seconds?" across 1B videos
// =============================================================================
func Num13RealtimeOLAP() {
	// stub — implement when teaching
}

// =============================================================================
// Num14PipelineOrchestration
//
// TOPIC: Pipeline Orchestration (DAGs) — Dependency resolution, backfill, SLA
//   - stub — implement when teaching
//
// Real world: Airflow managing 10,000+ daily DAG runs at TikTok
// =============================================================================
func Num14PipelineOrchestration() {
	// stub — implement when teaching
}

// =============================================================================
// Num15DataQualityContracts
//
// TOPIC: Data Quality & Contracts — Schema validation, anomaly detection, dbt tests
//   - stub — implement when teaching
//
// Real world: "Why did the revenue dashboard show $0 at 9am?"
// =============================================================================
func Num15DataQualityContracts() {
	// stub — implement when teaching
}

// =============================================================================
// Num16DataLineageCatalog
//
// TOPIC: Data Lineage & Catalog — Column-level lineage, impact analysis, OpenLineage
//   - stub — implement when teaching
//
// Real world: Tracing a bad revenue metric back to a broken upstream table
// =============================================================================
func Num16DataLineageCatalog() {
	// stub — implement when teaching
}

// =============================================================================
// Num17CostBasedQueryOptimizer
//
// TOPIC: Cost-Based Query Optimizer — Statistics, cardinality estimation, join reordering
//   - stub — implement when teaching
//
// Real world: Spark picked the wrong join order — ran 3h instead of 10min
// =============================================================================
func Num17CostBasedQueryOptimizer() {
	// stub — implement when teaching
}

// =============================================================================
// Num18DataMeshGovernance
//
// TOPIC: Data Mesh & Governance — Domain ownership, data products, schema registry
//   - stub — implement when teaching
//
// Real world: Grab: 50 domains, each team owns their data contracts
// =============================================================================
func Num18DataMeshGovernance() {
	// stub — implement when teaching
}

// =============================================================================
// Num19MultiTenantDataPlatform
//
// TOPIC: Multi-Tenant Data Platform — Resource isolation, quota management, cost attribution
//   - stub — implement when teaching
//
// Real world: "Team A's Spark job killed Team B's pipeline"
// =============================================================================
func Num19MultiTenantDataPlatform() {
	// stub — implement when teaching
}

// =============================================================================
// Num20LakehouseArchitecture
//
// TOPIC: Lakehouse Architecture — Lambda vs Kappa vs Lakehouse, unified batch+stream
//   - stub — implement when teaching
//
// Real world: Meta/Airbnb: why they converged on Lakehouse (killed the Lambda)
// =============================================================================
func Num20LakehouseArchitecture() {
	// stub — implement when teaching
}

// =============================================================================
// Num21DataSecurityPrivacy
//
// TOPIC: Data Security & Privacy — Column encryption, row-level security, PII masking
//   - stub — implement when teaching
//
// Real world: GoTo: handling driver and customer PII under Indonesian PDPA law
// =============================================================================
func Num21DataSecurityPrivacy() {
	// stub — implement when teaching
}

// =============================================================================
// Num22PerformanceTuning
//
// TOPIC: End-to-End Performance Tuning — Query profiling, partition pruning, Z-ordering
//   - stub — implement when teaching
//
// Real world: Going from 4-hour Spark job to 15-minute
// =============================================================================
func Num22PerformanceTuning() {
	// stub — implement when teaching
}
