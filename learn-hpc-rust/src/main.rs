// ============================================================
// learn-hpc-rust — HPC & Systems Programming in Rust
// Dispatcher: uncomment a topic to run it.
// ============================================================

mod demos;
#[allow(unused_imports)]
use demos::*;

fn main() {
    // num01_memory_layout_demo();
    // num02_cache_lines_demo();
    // num03_simd_vectorization_demo();
    // num04_memory_allocators_demo();
    // num05_lock_free_data_structures_demo();
    // num06_branch_prediction_demo();
    // num07_io_models_demo();
    // num08_numa_architecture_demo();
    // num09_gpu_memory_hierarchy_demo();
    // num10_linux_memory_management_demo();
    // num11_profiling_flamegraphs_demo();
    // num12_kernel_bypass_demo();
    num13_columnar_db_demo();

    println!("learn-hpc-rust: uncomment a topic in main.rs to run it.");
}
