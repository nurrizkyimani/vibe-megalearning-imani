# M3/W2/D1 - Sun, 08 Mar 2026 (13:50 WIB)

## HPC & Systems Programming in Rust

Learning high-performance computing and systems programming concepts from first principles, implemented in Rust, targeting senior SWE roles at Anthropic, OpenAI, and similar companies.

---

## How to Run

### 1\. Install Rust (if not already installed)

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source $HOME/.cargo/env
```

### 2\. Verify installation

```
rustc --version   # should print rustc 1.x.x
cargo --version   # should print cargo 1.x.x
```

### 3\. Run the project

```
cargo run --manifest-path learn-hpc-rust/Cargo.toml
```

### 4\. To run a specific topic

Uncomment the desired function call in `learn-hpc-rust/src/main.rs`, then run:

```
cargo run --manifest-path learn-hpc-rust/Cargo.toml
```

---

## How to Set Up This Folder From Scratch (Human Steps)

If you want to replicate this project structure yourself in any new Rust project, here are the exact commands:

### Step 1 — Create the folder structure

```
mkdir -p learn-hpc-rust/src
```

`-p` creates parent folders if they don't exist. `src/` is where all `.rs` files live.

### Step 2 — Create `Cargo.toml`

```
cat > learn-hpc-rust/Cargo.toml << 'EOF'
[package]
name = "learn-hpc-rust"
version = "0.1.0"
edition = "2021"
EOF
```

`Cargo.toml` is the manifest — like `package.json` in Node or `go.mod` in Go. No `[[bin]]` section needed; Cargo automatically uses `src/main.rs` as the entry point.

### Step 3 — Create `src/main.rs`

```
cat > learn-hpc-rust/src/main.rs << 'EOF'
mod demos;
#[allow(unused_imports)]
use demos::*;

fn main() {
    // num01_memory_layout_demo();
}
EOF
```

`mod demos;` tells Rust to load `src/demos.rs` as a module. `use demos::*;` brings all its public functions into scope.

### Step 4 — Create `src/demos.rs`

```
cat > learn-hpc-rust/src/demos.rs << 'EOF'
#![allow(dead_code)]

pub fn num01_memory_layout_demo() {
    // TODO
}
EOF
```

`pub fn` makes the function visible outside the module. `#![allow(dead_code)]` silences warnings for stubs not yet called.

### Step 5 — Verify it compiles

```
cargo build --manifest-path learn-hpc-rust/Cargo.toml
```

### Step 6 — Run it

```
cargo run --manifest-path learn-hpc-rust/Cargo.toml
```

Output: nothing (empty stub), no errors.

### Step 7 — To run a topic, uncomment its call in `main.rs`

```
// before
// num01_memory_layout_demo();

// after
num01_memory_layout_demo();
```

Then `cargo run` again.

---

### Where does `target/` come from?

You never create it. Cargo auto-generates it on first build:

```
learn-hpc-rust/
└── target/
    └── debug/
        └── learn-hpc-rust    ← the compiled binary
```

| Command | Profile | Output folder | Binary speed |
| --- | --- | --- | --- |
| `cargo build` | debug (default) | `target/debug/` | unoptimized, fast to compile |
| `cargo build --release` | release | `target/release/` | fully optimized, slow to compile |

During learning, always use debug (the default). Use `--release` only when benchmarking real performance numbers.

To wipe all build artifacts: `cargo clean`

---

### Shortcut: `cargo new` does steps 1–3 for you

If starting a fresh standalone project (not inside an existing repo):

```
cargo new learn-hpc-rust --vcs none   # --vcs none skips git init
cd learn-hpc-rust
cargo run
```

You still need to create `demos.rs` manually (step 4) and add `mod demos;` to `main.rs`.

---

## Curriculum

| # | Topic | Core Concept | Interview Question | Real World |
| --- | --- | --- | --- | --- |
| 01 | Memory Layout & Alignment | Stack vs heap, struct padding, cache-friendly layout | "Our ML feature vector processing is slow — where do you look first?" | C++/Rust hot path, ML inference |
| 02 | Cache Lines & False Sharing | 64-byte cache lines, false sharing between threads | "Our multi-threaded price engine slows down with more cores — why?" | HFT engines, game servers |
| 03 | SIMD & Vectorization | Process multiple data with one CPU instruction | "How does numpy multiply a million floats faster than a for-loop?" | NumPy, ML inference, audio/video |
| 04 | Memory Allocators | Arena, pool, bump allocators vs system malloc | "Our trading system has latency spikes every few seconds — diagnose it" | Game engines, DB buffer pools |
| 05 | Lock-free Data Structures | Atomics, CAS, ABA problem | "How do you build a queue thousands of threads push to without a mutex?" | Disruptor (LMAX), Kafka internals |
| 06 | CPU Branch Prediction | Pipeline stalls, branchless code | "This sort is fast on random data but slow on sorted data — explain why" | Hot loops, parsers, compilers |
| 07 | I/O Models | Blocking vs non-blocking vs async vs io\_uring | "How does Nginx serve 10k connections on one thread?" | Nginx, Tokio, Node.js, epoll |
| 08 | NUMA Architecture | Non-uniform memory access, CPU affinity | "Our DB is slower on the 64-core server than the 16-core — why?" | Multi-socket HPC, ML training |
| 09 | GPU Memory Hierarchy | Global/shared/register memory, coalesced access | "Our CUDA kernel uses 80% GPU but throughput is low — what's wrong?" | CUDA kernels, Triton, PyTorch |
| 10 | Linux Memory Management | Virtual memory, page faults, mmap, huge pages | "Our model loading takes 8s on first request then is instant — explain" | DB engines, ML frameworks |
| 11 | Profiling & Flamegraphs | perf, flamegraph, identifying hot paths | "P99 latency is 10x P50 in production — how do you find the cause?" | Any production system |
| 12 | Kernel Bypass (io\_uring/DPDK) | Skipping kernel for ultra-low latency I/O | "How do HFT firms achieve sub-microsecond network latency?" | HFT, network packet processing |

---

## Lessons

---

### Num01 — Memory Layout & Alignment

**Interview question:** "Our ML feature vector processing is slow — where do you look first?"

#### What we built

Three experiments in one function:
1. Same struct fields in different order → different sizes due to padding
2. Stack address vs heap address — where data actually lives
3. AoS vs SoA layout — cache-friendly data organisation

#### Output

```
=== Num01: Memory Layout & Alignment ===

--- Part 1: Struct Padding ---
BadLayout  (u8, u64, u32) size = 24 bytes
GoodLayout (u64, u32, u8) size = 16 bytes
Same 3 fields, different order — 8 bytes saved

--- Part 2: Stack vs Heap ---
Stack array [i32; 4]: size=16 bytes, address=0x16f0e2040
Heap Vec<i32>:        size of Vec struct=24 bytes (ptr+len+cap on stack), data on heap at 0x12ba04080

--- Part 3: AoS vs SoA layout ---
AoS — size of one ParticleAoS struct = 28 bytes
      to read 4 positions, we touch 112 bytes (includes unused vx/vy/vz/mass)
SoA — size of x array alone           = 16 bytes
      to read 4 positions, we touch 48 bytes (only x+y+z arrays)
```

#### Key lessons

**Alignment & padding**
The CPU reads memory in aligned chunks. A `u64` must start at an address divisible by 8. If the previous field left a gap, Rust inserts invisible padding bytes to satisfy that. Field order controls how much padding gets inserted.

```
BadLayout:  [ u8:1 ][ pad:7 ][ u64:8 ][ u32:4 ][ pad:4 ]  = 24 bytes
GoodLayout: [ u64:8 ][ u32:4 ][ u8:1 ][ pad:3 ]            = 16 bytes
```

Rule of thumb: sort fields **largest → smallest** to minimise padding.

**Stack vs Heap**
- Stack: fixed size known at compile time, lives in the function frame, freed automatically on return. Very fast.
- Heap: dynamic size, allocated at runtime by the allocator. The `Vec` struct itself (24 bytes) sits on the stack holding a pointer, length, and capacity — the actual data lives at a separate heap address.

**AoS vs SoA**
- AoS (Array of Structs): each particle is one struct with all fields packed. Iterating positions loads velocity + mass into cache even though you don't use them — wasted bandwidth.
- SoA (Struct of Arrays): each field is its own array. Iterating positions touches only x/y/z arrays — everything in cache is useful.

At 1 million particles: AoS wastes ~27 MB of cache per position loop. SoA uses ~12 MB. This is why ML inference engines, physics engines, and game engines all use SoA.