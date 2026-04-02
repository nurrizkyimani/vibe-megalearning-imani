// ============================================================
// learn-hpc-rust — demos.rs
// All 12 HPC topic demos live here.
// Each function is a stub until we implement it together.
// ============================================================
#![allow(dead_code)]

// ------------------------------------------------------------
// Num01 — Memory Layout & Alignment
// Core: Stack vs heap, struct padding, cache-friendly layout
// Interview: "Our ML feature vector processing is slow — where do you look first?"
// Real world: C++/Rust hot path, ML inference
// ------------------------------------------------------------
pub fn num01_memory_layout_demo() {
    // ================================================================
    // TOPIC: Memory Layout & Alignment
    //
    // The CPU does NOT read memory one byte at a time. It reads in
    // chunks called "words" (8 bytes on 64-bit CPUs). To make that
    // work cleanly, the CPU requires each value to sit at an address
    // that is a multiple of its size. This is called ALIGNMENT.
    //
    //   u8  → must start at any address       (align = 1)
    //   u16 → must start at multiple of 2     (align = 2)
    //   u32 → must start at multiple of 4     (align = 4)
    //   u64 → must start at multiple of 8     (align = 8)
    //
    // When you put fields in a struct, Rust (like C) inserts invisible
    // "padding" bytes between fields to satisfy alignment. This means
    // field ORDER affects the total size of the struct.
    // ================================================================

    // ----------------------------------------------------------------
    // PART 1: Padding demo — same fields, different order, different size
    // ----------------------------------------------------------------

    // BAD layout — fields are in "random" order
    // Rust inserts padding to keep each field aligned:
    //
    //   [u8:  1 byte ] [padding: 7 bytes] [u64: 8 bytes] [u32: 4 bytes] [padding: 4 bytes]
    //   total = 24 bytes
    #[repr(C)] // repr(C) = use C-style layout (no Rust reordering), makes padding visible
    struct BadLayout {
        a: u8,  // 1 byte,  align=1
        b: u64, // 8 bytes, align=8 → needs 7 bytes of padding before it
        c: u32, // 4 bytes, align=4 → needs 4 bytes of padding after to reach align of struct
    }

    // GOOD layout — fields sorted largest → smallest
    // No padding needed between fields:
    //
    //   [u64: 8 bytes] [u32: 4 bytes] [u8: 1 byte] [padding: 3 bytes]
    //   total = 16 bytes
    #[repr(C)]
    struct GoodLayout {
        b: u64, // 8 bytes first
        c: u32, // 4 bytes next
        a: u8,  // 1 byte last — only 3 bytes padding at the end (struct align = 8)
    }

    println!("=== Num01: Memory Layout & Alignment ===\n");
    println!("--- Part 1: Struct Padding ---");
    println!(
        "BadLayout  (u8, u64, u32) size = {} bytes",
        std::mem::size_of::<BadLayout>()
    );
    println!(
        "GoodLayout (u64, u32, u8) size = {} bytes",
        std::mem::size_of::<GoodLayout>()
    );
    println!(
        "Same 3 fields, different order — {} bytes saved\n",
        std::mem::size_of::<BadLayout>() - std::mem::size_of::<GoodLayout>()
    );

    // ----------------------------------------------------------------
    // PART 2: Stack vs Heap
    //
    // STACK: fixed-size, allocated at compile time, extremely fast.
    //        Lives in the function's stack frame. Freed when function returns.
    //
    // HEAP:  dynamic size, allocated at runtime via the allocator (malloc).
    //        Lives until you explicitly free it (in Rust: when the owner drops).
    //        Slower to allocate, but can hold arbitrary sizes.
    // ----------------------------------------------------------------

    println!("--- Part 2: Stack vs Heap ---");

    // Stack allocation — size must be known at compile time
    let stack_array: [i32; 4] = [1, 2, 3, 4];
    println!(
        "Stack array [i32; 4]: size={} bytes, address={:p}",
        std::mem::size_of_val(&stack_array),
        &stack_array
    );

    // Heap allocation — Vec grows dynamically at runtime
    let heap_vec: Vec<i32> = vec![1, 2, 3, 4];
    println!(
        "Heap Vec<i32>:        size of Vec struct={} bytes (ptr+len+cap on stack), data on heap at {:p}\n",
        std::mem::size_of_val(&heap_vec),
        heap_vec.as_ptr()
    );

    // ----------------------------------------------------------------
    // PART 3: Cache-friendly (AoS vs SoA) layout
    //
    // Imagine you have 1 million particles and you only need to update
    // their positions each frame — you don't touch velocity or mass.
    //
    // AoS (Array of Structs): all fields of one particle are contiguous.
    //   [x,y,z, vx,vy,vz, mass] [x,y,z, vx,vy,vz, mass] ...
    //   When you loop over positions, you load velocity+mass into cache
    //   even though you never use them. Cache is wasted → cache misses.
    //
    // SoA (Struct of Arrays): each field is its own array.
    //   [x,x,x,...] [y,y,y,...] [z,z,z,...] [vx,vx,...] ...
    //   When you loop over positions, you load only x,y,z. Cache is
    //   fully used → much faster for SIMD and sequential access.
    // ----------------------------------------------------------------

    println!("--- Part 3: AoS vs SoA layout ---");

    // AoS — Array of Structs (cache-unfriendly for partial field access)
    struct ParticleAoS {
        x: f32,
        y: f32,
        z: f32, // position — what we want
        vx: f32,
        vy: f32,
        vz: f32,   // velocity — irrelevant for position loop
        mass: f32, // mass     — irrelevant for position loop
    }

    // SoA — Struct of Arrays (cache-friendly for partial field access)
    struct ParticlesSoA {
        x: Vec<f32>,
        y: Vec<f32>,
        z: Vec<f32>, // positions packed together
        vx: Vec<f32>,
        vy: Vec<f32>,
        vz: Vec<f32>,   // velocities packed together
        mass: Vec<f32>, // masses packed together
    }

    let n = 4;

    let aos: Vec<ParticleAoS> = (0..n)
        .map(|i| ParticleAoS {
            x: i as f32,
            y: i as f32,
            z: i as f32,
            vx: 0.0,
            vy: 0.0,
            vz: 0.0,
            mass: 1.0,
        })
        .collect();

    let soa = ParticlesSoA {
        x: (0..n).map(|i| i as f32).collect(),
        y: (0..n).map(|i| i as f32).collect(),
        z: (0..n).map(|i| i as f32).collect(),
        vx: vec![0.0; n],
        vy: vec![0.0; n],
        vz: vec![0.0; n],
        mass: vec![1.0; n],
    };

    println!(
        "AoS — size of one ParticleAoS struct = {} bytes",
        std::mem::size_of::<ParticleAoS>()
    );
    println!(
        "      to read {} positions, we touch {} bytes (includes unused vx/vy/vz/mass)",
        n,
        n * std::mem::size_of::<ParticleAoS>()
    );

    println!(
        "SoA — size of x array alone           = {} bytes",
        n * std::mem::size_of::<f32>()
    );
    println!(
        "      to read {} positions, we touch {} bytes (only x+y+z arrays)",
        n,
        3 * n * std::mem::size_of::<f32>()
    );

    // suppress unused warnings
    let _ = &aos;
    let _ = &soa;

    println!("\n=== Num01 done ===");
}

// ------------------------------------------------------------
// Num02 — Cache Lines & False Sharing
// Core: 64-byte cache lines, false sharing between threads
// Interview: "Our multi-threaded price engine slows down with more cores — why?"
// Real world: HFT engines, game servers
// ------------------------------------------------------------
pub fn num02_cache_lines_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num03 — SIMD & Vectorization
// Core: Process multiple data with one CPU instruction
// Interview: "How does numpy multiply a million floats faster than a for-loop?"
// Real world: NumPy, ML inference, audio/video
// ------------------------------------------------------------
pub fn num03_simd_vectorization_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num04 — Memory Allocators
// Core: Arena, pool, bump allocators vs system malloc
// Interview: "Our trading system has latency spikes every few seconds — diagnose it"
// Real world: Game engines, DB buffer pools
// ------------------------------------------------------------
pub fn num04_memory_allocators_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num05 — Lock-free Data Structures
// Core: Atomics, CAS, ABA problem
// Interview: "How do you build a queue thousands of threads push to without a mutex?"
// Real world: Disruptor (LMAX), Kafka internals
// ------------------------------------------------------------
pub fn num05_lock_free_data_structures_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num06 — CPU Branch Prediction
// Core: Pipeline stalls, branchless code
// Interview: "This sort is fast on random data but slow on sorted data — explain why"
// Real world: Hot loops, parsers, compilers
// ------------------------------------------------------------
pub fn num06_branch_prediction_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num07 — I/O Models
// Core: Blocking vs non-blocking vs async vs io_uring
// Interview: "How does Nginx serve 10k connections on one thread?"
// Real world: Nginx, Tokio, Node.js, epoll
// ------------------------------------------------------------
pub fn num07_io_models_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num08 — NUMA Architecture
// Core: Non-uniform memory access, CPU affinity
// Interview: "Our DB is slower on the 64-core server than the 16-core — why?"
// Real world: Multi-socket HPC, ML training
// ------------------------------------------------------------
pub fn num08_numa_architecture_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num09 — GPU Memory Hierarchy
// Core: Global/shared/register memory, coalesced access
// Interview: "Our CUDA kernel uses 80% GPU but throughput is low — what's wrong?"
// Real world: CUDA kernels, Triton, PyTorch
// ------------------------------------------------------------
pub fn num09_gpu_memory_hierarchy_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num10 — Linux Memory Management
// Core: Virtual memory, page faults, mmap, huge pages
// Interview: "Our model loading takes 8s on first request then is instant — explain"
// Real world: DB engines, ML frameworks
// ------------------------------------------------------------
pub fn num10_linux_memory_management_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num11 — Profiling & Flamegraphs
// Core: perf, flamegraph, identifying hot paths
// Interview: "P99 latency is 10x P50 in production — how do you find the cause?"
// Real world: Any production system
// ------------------------------------------------------------
pub fn num11_profiling_flamegraphs_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num12 — Kernel Bypass (io_uring/DPDK)
// Core: Skipping kernel for ultra-low latency I/O
// Interview: "How do HFT firms achieve sub-microsecond network latency?"
// Real world: HFT, network packet processing
// ------------------------------------------------------------
pub fn num12_kernel_bypass_demo() {
    // TODO: implement
}

// ------------------------------------------------------------
// Num13 — Columnar DB: AoS vs SoA applied to real database problem
//
// CONTEXT — the office problem:
//   You joined a team building an analytics database.
//   A customer runs: SELECT AVG(salary) FROM employees
//   The table has 1 million rows and 50 columns.
//   Your system takes 45 seconds. Postgres takes 2 seconds. Boss asks why.
//
// ROOT CAUSE — it's AoS vs SoA, the exact same insight as Num01 Part 3.
//
// ROW storage (AoS) — how Postgres and MySQL store data by default:
//   Each row is one contiguous block of all its columns.
//
//   Memory layout:
//   [id, name, age, salary, dept, ...47 more cols] ← row 0, all 50 cols
//   [id, name, age, salary, dept, ...47 more cols] ← row 1, all 50 cols
//   [id, name, age, salary, dept, ...47 more cols] ← row 2, all 50 cols
//   ...
//
//   To compute AVG(salary), you must scan ALL rows and in each row
//   skip past id, name, age to get to salary — then skip the other 46
//   columns you don't need. You touch every byte of the table even
//   though only 1/50th of it is salary data.
//   → 50x more data read than necessary → cache thrash → slow.
//
// COLUMNAR storage (SoA) — how ClickHouse, DuckDB, Parquet store data:
//   Each column is its own contiguous array.
//
//   Memory layout:
//   [id0, id1, id2, ...]           ← all ids packed
//   [name0, name1, name2, ...]     ← all names packed
//   [salary0, salary1, salary2, ...] ← all salaries packed  ← only this is read
//   [dept0, dept1, dept2, ...]     ← all depts packed
//   ...
//
//   To compute AVG(salary), you read ONLY the salary column.
//   You touch exactly 1/50th of the table. Everything in cache is useful.
//   → 50x less data read → fits in cache → fast.
//
// This is literally why columnar databases (ClickHouse, DuckDB, Snowflake,
// BigQuery, Parquet files) exist. The insight is pure SoA.
// ------------------------------------------------------------
pub fn num13_columnar_db_demo() {
    use std::time::Instant;

    println!("=== Num13: Columnar DB — AoS vs SoA in database context ===\n");

    // ----------------------------------------------------------------
    // Setup: 500_000 "employee" records, 6 columns each.
    // We will query: SELECT AVG(salary) FROM employees
    // Only 1 out of 6 columns is needed.
    // ----------------------------------------------------------------

    const N: usize = 500_000;

    // ----------------------------------------------------------------
    // ROW storage — AoS
    // Each employee is one struct. All columns are packed together.
    // This is how Postgres, MySQL store a row on disk and in memory.
    // ----------------------------------------------------------------
    struct EmployeeRow {
        id: u32,        // 4 bytes — not needed for AVG(salary)
        age: u8,        // 1 byte  — not needed
        department: u8, // 1 byte  — not needed
        _pad: [u8; 2],  // 2 bytes — padding so salary is 4-byte aligned
        salary: f32,    // 4 bytes — THIS is what we need
        score: f32,     // 4 bytes — not needed
    }
    // Total per row: 16 bytes. Only 4 bytes (salary) are useful for this query.
    // We read 16 bytes to get 4 useful bytes = 75% wasted bandwidth.

    // ----------------------------------------------------------------
    // COLUMNAR storage — SoA
    // Each column is its own Vec. This is how ClickHouse, DuckDB, Arrow store data.
    // ----------------------------------------------------------------
    struct EmployeeColumns {
        id: Vec<u32>,
        age: Vec<u8>,
        department: Vec<u8>,
        salary: Vec<f32>, // ← only this Vec gets touched for AVG(salary)
        score: Vec<f32>,
    }

    // --- populate row storage ---
    let rows: Vec<EmployeeRow> = (0..N)
        .map(|i| EmployeeRow {
            id: i as u32,
            age: (20 + i % 40) as u8,
            department: (i % 10) as u8,
            _pad: [0; 2],
            salary: 50_000.0 + (i % 100_000) as f32,
            score: (i % 100) as f32,
        })
        .collect();

    // --- populate columnar storage ---
    let cols = EmployeeColumns {
        id: (0..N).map(|i| i as u32).collect(),
        age: (0..N).map(|i| (20 + i % 40) as u8).collect(),
        department: (0..N).map(|i| (i % 10) as u8).collect(),
        salary: (0..N).map(|i| 50_000.0 + (i % 100_000) as f32).collect(),
        score: (0..N).map(|i| (i % 100) as f32).collect(),
    };

    // ----------------------------------------------------------------
    // QUERY: AVG(salary)
    //
    // Row scan — must touch every field of every row to get to salary.
    // The CPU loads 16-byte structs into cache lines. Per 64-byte cache
    // line we get 4 rows (64 / 16 = 4). But only 4 of 64 bytes are salary.
    // ----------------------------------------------------------------
    let t0 = Instant::now();
    let avg_row = {
        let mut sum = 0.0f64;
        for row in &rows {
            // We only use row.salary — but the CPU had to load the entire
            // 16-byte struct (id, age, department, _pad, salary, score)
            // into a cache line just to read this one field.
            sum += row.salary as f64;
        }
        sum / N as f64
    };
    let row_time = t0.elapsed();

    // ----------------------------------------------------------------
    // Columnar scan — touches ONLY the salary Vec.
    // salary is Vec<f32> = tightly packed f32 values.
    // One 64-byte cache line holds 16 salaries (64 / 4 = 16).
    // Every byte we load from cache is a salary we need.
    // Zero wasted bandwidth.
    // ----------------------------------------------------------------
    let t1 = Instant::now();
    let avg_col = {
        let mut sum = 0.0f64;
        for &s in &cols.salary {
            sum += s as f64;
        }
        sum / N as f64
    };
    let col_time = t1.elapsed();

    // ----------------------------------------------------------------
    // Results
    // ----------------------------------------------------------------
    println!("Dataset: {} employees, 6 columns, query = AVG(salary)\n", N);

    println!("Row layout (AoS — like Postgres):");
    println!(
        "  bytes per row    = {} (only 4 are salary)",
        std::mem::size_of::<EmployeeRow>()
    );
    println!(
        "  total data read  = {} KB",
        (N * std::mem::size_of::<EmployeeRow>()) / 1024
    );
    println!(
        "  useful data      = {} KB (salary only)",
        (N * std::mem::size_of::<f32>()) / 1024
    );
    println!(
        "  wasted bandwidth = {}%",
        100 - 100 * std::mem::size_of::<f32>() / std::mem::size_of::<EmployeeRow>()
    );
    println!("  AVG(salary)      = {:.2}", avg_row);
    println!("  time             = {:?}\n", row_time);

    println!("Columnar layout (SoA — like ClickHouse/DuckDB):");
    println!("  bytes per salary = 4 (f32, tightly packed)");
    println!(
        "  total data read  = {} KB (salary column only)",
        (N * std::mem::size_of::<f32>()) / 1024
    );
    println!("  wasted bandwidth = 0%");
    println!("  AVG(salary)      = {:.2}", avg_col);
    println!("  time             = {:?}\n", col_time);

    println!(
        "Columnar was {:.1}x faster for this query.",
        row_time.as_secs_f64() / col_time.as_secs_f64()
    );
    println!("The more columns the table has, the bigger this gap gets.");
    println!("A 50-column table = up to 50x more data read in row layout.");

    println!("\n=== Num13 done ===");
}
