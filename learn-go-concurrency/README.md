# M3/W5/D1 - Fri, 03 Apr 2026 (WIB)

## Go Concurrency - Full Curriculum

A structured learning path to master concurrency in Go using runnable demos.
Each topic is mapped to one `NumXX...` function in `main.go`, so you can learn
progressively from goroutine basics to production concurrency patterns.

Target audience: Go developers who can write sequential programs and want to
be confident with goroutines, synchronization, cancellation, and performance.

---

## Curriculum Table

| # | Topic | Core Concept | Real World Usage | Demo |
| --- | --- | --- | --- | --- |
| 01 | Goroutines | Lightweight concurrent execution | Background jobs, async API fan-out | Num01 ✅ |
| 02 | Channels (Basic) | Safe data handoff between goroutines | Pipeline stages and worker communication | Num02 ✅ |
| 03 | Channels (Buffered) | Queueing without immediate receiver | Burst smoothing, limited in-memory queues | Num03 ✅ |
| 04 | WaitGroup | Wait for many goroutines to finish | Batch processing, graceful shutdown | Num04 ✅ |
| 05 | Mutex | Protect shared mutable state | Counters, shared cache writes | Num05 ✅ |
| 06 | Select | Multiplex channel operations | Timeouts, cancellation, non-blocking flows | Num06 🔲 |
| 07 | Channel Patterns | Fan-in, fan-out, pipelines, tee | Stream processing and ETL workers | Num07 🔲 |
| 08 | Context | Cancellation, deadlines, request scope | HTTP handlers, RPC propagation | Num08 🔲 |
| 09 | RWMutex | Many readers + few writers lock strategy | Read-heavy in-memory stores | Num09 🔲 |
| 10 | sync.Once | Run initialization exactly once | Singleton setup, lazy initialization | Num10 🔲 |
| 11 | sync.Map | Concurrent map with optimized access paths | Shared caches under high contention | Num11 🔲 |
| 12 | Atomic | Lock-free primitives for simple state | Metrics counters, state flags | Num12 🔲 |
| 13 | Worker Pool | Bounded parallel job processing | Task queues, background workers | Num13 🔲 |
| 14 | errgroup | Error propagation across goroutines | Parallel API calls with fail-fast behavior | Num14 🔲 |
| 15 | Memory Model | Happens-before guarantees, visibility | Correct synchronization reasoning | Num15 🔲 |
| 16 | Race Detector | Detect data races with tooling | CI safety checks for concurrent code | Num16 🔲 |
| 17 | Semaphore | Limit concurrent work | DB connection caps, rate-limited APIs | Num17 🔲 |
| 18 | GMP Scheduler | Go runtime scheduling model | Tuning latency/throughput in prod services | Num18 🔲 |
| 19 | Benchmarking | Measure concurrency performance | Compare lock/channel/atomic strategies | Num19 🔲 |

Legend: ✅ = implemented | 🔲 = stub (to be implemented)

---

## How to run

```bash
# run all demos currently called in main.go
go run ./learn-go-concurrency/
```

To run a specific topic, comment/uncomment the relevant function call in `main.go`.

---

## Function signatures

```go
func Num01GoroutineDemo()       // ✅ implemented
func Num02ChannelsBasicDemo()   // ✅ implemented
func Num03ChannelsBufferedDemo() // ✅ implemented
func Num04WaitGroupDemo()       // ✅ implemented
func Num05MutexDemo()           // ✅ implemented
func Num06SelectDemo()          // 🔲 stub
func Num07ChannelPatternsDemo() // 🔲 stub
func Num08ContextDemo()         // 🔲 stub
func Num09RWMutexDemo()         // 🔲 stub
func Num10SyncOnceDemo()        // 🔲 stub
func Num11SyncMapDemo()         // 🔲 stub
func Num12AtomicDemo()          // 🔲 stub
func Num13WorkerPoolDemo()      // 🔲 stub
func Num14ErrGroupDemo()        // 🔲 stub
func Num15MemoryModelDemo()     // 🔲 stub
func Num16RaceDetectorDemo()    // 🔲 stub
func Num17SemaphoreDemo()       // 🔲 stub
func Num18GMPSchedulerDemo()    // 🔲 stub
func Num19BenchmarkingDemo()    // 🔲 stub
```

---

## Num01 - Goroutine Basics

### The Problem

If `main` exits too early, goroutines may never complete.

### The Concept

`go func() { ... }()` starts a new goroutine. The runtime schedules it
independently from `main`.

### What the demo shows

1. Starts one goroutine that prints `"hello world from goroutine"`.
2. Uses `time.Sleep(1 * time.Second)` so `main` does not exit immediately.
3. Prints `"main done"` at the end.

Expected output shape:

```text
hello world from goroutine
main done
```

---

## Num02 - Channels (Basic)

### The Problem

We need a safe way to pass data between goroutines without shared memory races.

### The Concept

An unbuffered channel synchronizes sender and receiver:
send blocks until a receiver is ready.

### What the demo shows

1. Creates `ch := make(chan string)`.
2. A goroutine sends `"data from goroutine pipeline"`.
3. Main receives from channel and prints the value.

Expected output shape:

```text
received: data from goroutine pipeline
```

---

## Num03 - Channels (Buffered)

### The Problem

Sometimes a producer should enqueue values before consumers read them.

### The Concept

A buffered channel allows up to `cap(ch)` sends without immediate receiver.

### What the demo shows

1. Creates `ch := make(chan int, 2)`.
2. Sends `10` and `20` without blocking.
3. Receives and prints both values in FIFO order.

Expected output:

```text
10
20
```

---

## Num04 - WaitGroup

### The Problem

Main goroutine needs to wait until all worker goroutines complete.

### The Concept

`sync.WaitGroup` coordinates a set of goroutines with:
`Add(n)`, `Done()`, and `Wait()`.

### What the demo shows

1. Spawns 3 workers.
2. Each worker prints `"worker <id> done"` and calls `Done()`.
3. `wg.Wait()` blocks until all workers finish.
4. Prints `"all worker one after wg wait"`.

Note: Worker print order is non-deterministic.

---

## Num05 - Mutex

### The Problem

Multiple goroutines incrementing shared state can cause race conditions.

### The Concept

`sync.Mutex` ensures only one goroutine updates shared data at a time.

### What the demo shows

1. Starts 1000 goroutines.
2. Each goroutine locks the mutex, increments `counter`, prints, then unlocks.
3. Waits for completion with `WaitGroup`.
4. Prints final counter value.

Expected invariant:

```text
final counter: 1000
```

---

## Next implementation order

Recommended order to continue filling stubs:

1. `Num06SelectDemo`
2. `Num08ContextDemo`
3. `Num13WorkerPoolDemo`
4. `Num14ErrGroupDemo`
5. `Num16RaceDetectorDemo`

This sequence builds from control flow to cancellation, then production patterns.
