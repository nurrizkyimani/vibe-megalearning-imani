package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"syscall"
)

// =============================================================================
// Num01BuildYourOwnContainer
//
// TOPIC: Linux Primitives — Namespaces, Cgroups, Chroot, OverlayFS
//   - PID namespace: clone(CLONE_NEWPID) — first child gets PID 1 inside ns
//   - Net namespace: clone(CLONE_NEWNET) — own IP stack, veth pairs, iptables
//   - Cgroups v2: write cpu.max + memory.max under /sys/fs/cgroup/<group>/
//   - Chroot / pivot_root: syscall.Chroot(rootfs) — EPERM without CAP_SYS_CHROOT
//   - OverlayFS: lowerdir (read-only layers) + upperdir (writable) → merged
//   - mmap(MAP_ANON): anonymous memory mapping — how the allocator gets pages
//   - Pipe(2): kernel ring buffer — used for container stdio plumbing
//   - Rlimit: per-process resource limits (cgroup analog, no root needed)
//
// Real world: What `docker run` / `runc` does — one clone(2) call with all flags
// TikTok/GoTo: container runtime internals, Kata Containers, Firecracker MicroVMs
// =============================================================================
func Num01BuildYourOwnContainer() {
	fmt.Println("============================================================")
	fmt.Println("  Num01 -- Linux Primitives (Build Your Own Container)")
	fmt.Println("============================================================")
	fmt.Println()

	// ── Step 1: PID Namespace — who are you in the process tree? ─────────────
	// On Linux: clone(CLONE_NEWPID) gives the child PID 1 inside its namespace.
	// The same kernel thread is visible from the host as a high PID.
	// We use real syscall.Getpid/Getppid + spawn a real child to show the mapping.
	fmt.Println("── Step 1: PID Namespace ────────────────────────────────────")
	fmt.Println()

	myPID := syscall.Getpid()
	myPPID := syscall.Getppid()
	myUID := syscall.Getuid()
	myGID := syscall.Getgid()
	fmt.Printf("[HOST] this process:  PID=%d  PPID=%d  UID=%d  GID=%d\n", myPID, myPPID, myUID, myGID)

	// Spawn a real child process — kernel assigns it a real host PID.
	// On Linux with CLONE_NEWPID the child would see itself as PID 1.
	child := exec.Command("/bin/sh", "-c", "echo [child] my PID=$$; sleep 0.05")
	child.Stdout = os.Stdout
	if err := child.Start(); err != nil {
		fmt.Println("child start err:", err)
	} else {
		hostPID := child.Process.Pid
		fmt.Printf("[HOST] child spawned: host PID=%d\n", hostPID)
		fmt.Printf("[HOST] (on Linux with CLONE_NEWPID → child would see itself as PID 1)\n")

		// Signal 0: probe whether the process is alive — no actual signal sent
		if err := syscall.Kill(hostPID, 0); err == nil {
			fmt.Printf("[HOST] Kill(%d, SIG=0) → process alive (kernel confirmed)\n", hostPID)
		}
		child.Wait()
		// After exit, Kill with SIG=0 returns ESRCH — process no longer exists
		if err := syscall.Kill(hostPID, 0); err != nil {
			fmt.Printf("[HOST] Kill(%d, SIG=0) after exit → %v (ESRCH: gone)\n", hostPID, err)
		}
	}
	fmt.Println()
	fmt.Println("Key insight: host PID N == PID 1 inside the container namespace.")
	fmt.Println("             Container cannot see or signal PIDs outside its ns.")
	fmt.Println()

	// ── Step 2: Pipe — real kernel ring buffer, container stdio wiring ────────
	// containerd wires container stdout → a kernel pipe → your terminal.
	// syscall.Pipe allocates two real file descriptors in the kernel.
	// Write to fd[1] → kernel buffer → read from fd[0].
	fmt.Println("── Step 2: Pipe (container stdio wiring) ────────────────────")
	fmt.Println()

	var pipeFds [2]int
	if err := syscall.Pipe(pipeFds[:]); err != nil {
		fmt.Println("pipe err:", err)
	} else {
		rFd, wFd := pipeFds[0], pipeFds[1]
		fmt.Printf("[PIPE] kernel allocated fds: read=%d  write=%d\n", rFd, wFd)

		// Write a log line as if coming from a container process
		msg := []byte("[container stdout] server started on :8080\n")
		n, _ := syscall.Write(wFd, msg)
		fmt.Printf("[PIPE] container wrote %d bytes into kernel pipe buffer (fd=%d)\n", n, wFd)

		// Read it back — simulates the container runtime (containerd) draining stdout
		buf := make([]byte, 256)
		n, _ = syscall.Read(rFd, buf)
		fmt.Printf("[PIPE] runtime drained fd=%d: %s", rFd, buf[:n])

		syscall.Close(rFd)
		syscall.Close(wFd)
		fmt.Printf("[PIPE] fds %d and %d closed (kernel reclaimed)\n", rFd, wFd)
	}
	fmt.Println()
	fmt.Println("Key insight: every `docker logs` reads from a pipe or log file")
	fmt.Println("             backed by exactly this fd plumbing at the runtime level.")
	fmt.Println()

	// ── Step 3: Rlimit — enforce resource limits, no root needed ──────────────
	// Cgroups write limits into /sys/fs/cgroup/.../memory.max and cpu.max.
	// Rlimit is the per-process predecessor — same kernel enforcement mechanism.
	// RLIMIT_NOFILE = max open file descriptors (like ulimit -n in a Dockerfile).
	fmt.Println("── Step 3: Rlimit (resource limits — cgroup analog) ─────────")
	fmt.Println()

	resources := []struct {
		name     string
		resource int
	}{
		{"RLIMIT_AS (virtual mem)", syscall.RLIMIT_AS},
		{"RLIMIT_NOFILE (open fds)", syscall.RLIMIT_NOFILE},
		{"RLIMIT_CPU (CPU seconds)", syscall.RLIMIT_CPU},
	}
	for _, r := range resources {
		var lim syscall.Rlimit
		syscall.Getrlimit(r.resource, &lim)
		curStr, maxStr := fmt.Sprintf("%d", lim.Cur), fmt.Sprintf("%d", lim.Max)
		if lim.Cur == ^uint64(0) {
			curStr = "unlimited"
		}
		if lim.Max == ^uint64(0) {
			maxStr = "unlimited"
		}
		fmt.Printf("[RLIMIT] %-30s  cur=%-14s  max=%s\n", r.name, curStr, maxStr)
	}
	fmt.Println()

	// Tighten RLIMIT_NOFILE to 32 — simulates a container's fd limit
	var origLim syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &origLim)
	tightLim := syscall.Rlimit{Cur: 32, Max: origLim.Max}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &tightLim); err != nil {
		fmt.Printf("[RLIMIT] Setrlimit(RLIMIT_NOFILE, soft=32) err=%v\n", err)
	} else {
		fmt.Println("[RLIMIT] Setrlimit(RLIMIT_NOFILE, soft=32) applied — kernel now enforces this")
		// Try to open files until the kernel refuses with EMFILE
		fds := []int{}
		for i := 1; i <= 8; i++ {
			fd, err := syscall.Open("/dev/null", syscall.O_RDONLY, 0)
			if err != nil {
				fmt.Printf("[RLIMIT] Open #%d → %v (kernel: EMFILE — too many open files)\n", i, err)
				break
			}
			fds = append(fds, fd)
			fmt.Printf("[RLIMIT] Open #%d → fd=%d  OK\n", i, fd)
		}
		for _, fd := range fds {
			syscall.Close(fd)
		}
		syscall.Setrlimit(syscall.RLIMIT_NOFILE, &origLim)
		fmt.Println("[RLIMIT] restored original limit")
	}
	fmt.Println()
	fmt.Println("Key insight: cgroup memory.max and cpu.max are the same concept —")
	fmt.Println("             a number the kernel checks on every alloc/open/fork.")
	fmt.Println()

	// ── Step 4: mmap + runtime.ReadMemStats — real page allocation ────────────
	// malloc / make([]byte, N) ultimately calls mmap(MAP_ANON) or brk(2).
	// We call mmap directly to see what the kernel hands us.
	// runtime.ReadMemStats reads the Go allocator's live counters.
	// These are the exact numbers cgroup memory.max caps.
	fmt.Println("── Step 4: mmap + allocator stats (what memory.max actually caps) ──")
	fmt.Println()

	pageSize := syscall.Getpagesize()
	fmt.Printf("[MMAP] kernel page size: %d bytes (%d KiB)\n", pageSize, pageSize/1024)

	// mmap 4 pages directly — bypasses the Go allocator, goes straight to kernel
	mapSize := 4 * pageSize
	data, err := syscall.Mmap(
		-1, 0, mapSize,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
	)
	if err != nil {
		fmt.Println("mmap err:", err)
	} else {
		fmt.Printf("[MMAP] mmap(MAP_ANON, %d bytes = %d pages) → addr=%p\n",
			mapSize, mapSize/pageSize, &data[0])

		// Write to each page — triggers real page faults, kernel commits physical RAM
		for i := 0; i < len(data); i += pageSize {
			data[i] = byte(i / pageSize)
		}
		fmt.Printf("[MMAP] wrote to each page → forced %d page faults (kernel committed RAM)\n",
			mapSize/pageSize)
		fmt.Printf("[MMAP] readback: page[0]=%d  page[1]=%d  page[2]=%d  page[3]=%d\n",
			data[0], data[pageSize], data[2*pageSize], data[3*pageSize])

		syscall.Munmap(data)
		fmt.Printf("[MMAP] munmap → pages returned to kernel\n")
	}
	fmt.Println()

	// Read real allocator stats
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("[ALLOCATOR] runtime.ReadMemStats — live numbers from the Go allocator:\n")
	fmt.Printf("  HeapAlloc    = %8d bytes  (live heap objects right now)\n", ms.HeapAlloc)
	fmt.Printf("  HeapSys      = %8d bytes  (OS pages reserved for heap)\n", ms.HeapSys)
	fmt.Printf("  HeapInuse    = %8d bytes  (spans with at least one live object)\n", ms.HeapInuse)
	fmt.Printf("  HeapReleased = %8d bytes  (pages returned to OS)\n", ms.HeapReleased)
	fmt.Printf("  Mallocs      = %8d        (cumulative alloc calls)\n", ms.Mallocs)
	fmt.Printf("  Frees        = %8d        (cumulative free calls)\n", ms.Frees)
	fmt.Printf("  NumGC        = %8d        (GC cycles run so far)\n", ms.NumGC)
	fmt.Println()

	// Allocate a large slice to cause GC pressure, then force GC and show the drop
	big := make([]byte, 8*1024*1024) // 8 MiB
	big[0] = 1
	var msPeak runtime.MemStats
	runtime.ReadMemStats(&msPeak)
	fmt.Printf("[ALLOCATOR] after 8 MiB alloc  → HeapAlloc=%d bytes\n", msPeak.HeapAlloc)
	big = nil
	runtime.GC()
	var msPost runtime.MemStats
	runtime.ReadMemStats(&msPost)
	fmt.Printf("[ALLOCATOR] after runtime.GC()  → HeapAlloc=%d bytes  (GC #%d ran)\n",
		msPost.HeapAlloc, msPost.NumGC)
	fmt.Println()
	fmt.Println("Key insight: cgroup memory.max caps the total RSS across all mmap/brk.")
	fmt.Println("             OOM killer fires SIGKILL when the cgroup crosses the limit.")
	fmt.Println()

	// ── Step 5: CPU topology via sysctl — what cpu.max is carved from ─────────
	// containerd writes cpu.max = "<quota> <period>" into the cgroup.
	// We read the host's real CPU topology from the kernel via sysctl.
	fmt.Println("── Step 5: CPU topology (sysctl) + cpu.max calculation ──────")
	fmt.Println()

	physCPU, _ := syscall.SysctlUint32("hw.physicalcpu")
	logCPU, _ := syscall.SysctlUint32("hw.logicalcpu")
	hostname, _ := syscall.Sysctl("kern.hostname")
	osType, _ := syscall.Sysctl("kern.ostype")
	osRelease, _ := syscall.Sysctl("kern.osrelease")

	fmt.Printf("[SYSCTL] kern.hostname    = %s\n", hostname)
	fmt.Printf("[SYSCTL] kern.ostype      = %s\n", osType)
	fmt.Printf("[SYSCTL] kern.osrelease   = %s\n", osRelease)
	fmt.Printf("[SYSCTL] hw.physicalcpu   = %d\n", physCPU)
	fmt.Printf("[SYSCTL] hw.logicalcpu    = %d\n", logCPU)
	fmt.Printf("[RUNTIME] NumCPU()        = %d  (logical CPUs Go sees)\n", runtime.NumCPU())
	fmt.Printf("[RUNTIME] GOMAXPROCS(0)   = %d  (OS threads running Go code)\n", runtime.GOMAXPROCS(0))
	fmt.Println()

	// Show what containerd writes to cpu.max for common container sizes
	period := 100_000 // 100ms in microseconds — Linux CFS default
	type cpuSpec struct {
		label string
		vcpu  float64
	}
	specs := []cpuSpec{
		{"0.25 vCPU (small pod)", 0.25},
		{"0.50 vCPU (default)", 0.50},
		{"1.00 vCPU (standard)", 1.00},
		{"2.00 vCPU (large)", 2.00},
	}
	fmt.Printf("[CGROUP] cpu.max containerd writes (host has %d logical CPUs):\n", logCPU)
	for _, s := range specs {
		quota := int(s.vcpu * float64(period))
		pct := s.vcpu / float64(logCPU) * 100
		throttled := period - quota
		throttleNote := fmt.Sprintf("paused %dms / 100ms", throttled/1000)
		if throttled <= 0 {
			throttleNote = "no throttle (multi-core)"
		}
		fmt.Printf("  %-26s  cpu.max=%-16s  → %s\n",
			s.label,
			fmt.Sprintf("%d %d", quota, period),
			fmt.Sprintf("%.1f%% of host; %s", pct, throttleNote))
	}
	fmt.Println()
	fmt.Println("Key insight: kernel throttles the cgroup's threads when they exceed")
	fmt.Printf("             quota µs in each %dµs CFS period.\n", period)
	fmt.Println()

	// ── Step 6: Chroot — real syscall, real kernel response ───────────────────
	// On Linux, runc calls pivot_root(2) (safer — chroot can be escaped).
	// We call syscall.Chroot to show the actual kernel errno.
	fmt.Println("── Step 6: Chroot (real syscall → real errno) ───────────────")
	fmt.Println()

	fmt.Printf("[CHROOT] syscall.Chroot(%q)  uid=%d...\n", "/tmp", syscall.Getuid())
	err = syscall.Chroot("/tmp")
	if err != nil {
		fmt.Printf("[CHROOT] errno = %v\n", err)
		fmt.Println("[CHROOT] needs CAP_SYS_CHROOT (root or user namespace)")
		fmt.Println()
		fmt.Println("[CHROOT] what runc does as root inside a user namespace:")
		fmt.Println("  1. chdir(rootfs)")
		fmt.Println("  2. pivot_root(\".\", \"oldrootfs\")")
		fmt.Println("  3. umount2(\"oldrootfs\", MNT_DETACH)")
		fmt.Println("  4. rmdir(\"oldrootfs\")")
		fmt.Println("  → process's / now points at container rootfs only")
		fmt.Println("  → /home/alice on the host: ENOENT")
	} else {
		fmt.Println("[CHROOT] succeeded (running as root)")
	}
	fmt.Println()
	fmt.Println("Key insight: chroot is one syscall — updates fs->root in the kernel.")
	fmt.Println("             No copying, no VM overhead. Pure pointer swap.")
	fmt.Println()

	// ── Step 7: OverlayFS — exact kernel CoW algorithm in Go maps ────────────
	// We implement the overlay lookup the same way the kernel does it:
	// scan layers top-down (upper → lower1 → lower0), first match wins.
	// A write to a read-only file triggers copy-on-write into the upper layer.
	fmt.Println("── Step 7: OverlayFS (copy-on-write layer lookup) ───────────")
	fmt.Println()

	type layer struct {
		name     string
		readOnly bool
		files    map[string]string
	}
	layers := []*layer{
		{
			name:     "lower0: ubuntu:22.04",
			readOnly: true,
			files: map[string]string{
				"/bin/bash":       "bash 5.1",
				"/lib/libc.so":    "glibc 2.35",
				"/etc/os-release": "Ubuntu 22.04",
				"/etc/nginx.conf": "worker_processes 1;",
			},
		},
		{
			name:     "lower1: nginx:1.25",
			readOnly: true,
			files: map[string]string{
				"/usr/sbin/nginx": "nginx 1.25",
				"/etc/nginx.conf": "worker_processes 4;", // shadows lower0
			},
		},
		{
			name:     "upper: container writable",
			readOnly: false,
			files:    map[string]string{}, // starts empty — all writes land here
		},
	}
	upper := layers[len(layers)-1]

	// overlayLookup: kernel scans upper → lower1 → lower0, returns first match
	overlayLookup := func(path string) (string, string) {
		for i := len(layers) - 1; i >= 0; i-- {
			if v, ok := layers[i].files[path]; ok {
				return v, layers[i].name
			}
		}
		return "", "ENOENT"
	}

	// Collect all unique paths across all layers
	seen := map[string]bool{}
	for _, l := range layers {
		for f := range l.files {
			seen[f] = true
		}
	}
	allPaths := make([]string, 0, len(seen))
	for f := range seen {
		allPaths = append(allPaths, f)
	}
	sort.Strings(allPaths)

	fmt.Println("[OVERLAYFS] merged view before any writes:")
	fmt.Printf("  %-28s  %-28s  %s\n", "path", "content", "layer")
	fmt.Printf("  %-28s  %-28s  %s\n",
		strings.Repeat("-", 27), strings.Repeat("-", 27), strings.Repeat("-", 22))
	for _, p := range allPaths {
		content, src := overlayLookup(p)
		fmt.Printf("  %-28s  %-28s  %s\n", p, content, src)
	}
	fmt.Println()

	// CoW write: container edits /etc/nginx.conf
	target := "/etc/nginx.conf"
	_, srcLayer := overlayLookup(target)
	newVal := "worker_processes auto; # tuned"
	fmt.Printf("[OVERLAYFS] container: write(%q, %q)\n", target, newVal)
	fmt.Printf("  kernel step 1: lookup %q → found in %q (read-only)\n", target, srcLayer)
	fmt.Printf("  kernel step 2: copy-on-write → copy file to upper layer\n")
	fmt.Printf("  kernel step 3: apply write in upper layer\n")
	upper.files[target] = newVal
	fmt.Println()

	newFile := "/var/log/nginx/access.log"
	upper.files[newFile] = "GET / 200 0.002s"
	fmt.Printf("[OVERLAYFS] container: create(%q) → goes directly to upper (no CoW needed)\n\n", newFile)

	// Recompute merged view after writes
	seen = map[string]bool{}
	for _, l := range layers {
		for f := range l.files {
			seen[f] = true
		}
	}
	allPaths = make([]string, 0, len(seen))
	for f := range seen {
		allPaths = append(allPaths, f)
	}
	sort.Strings(allPaths)

	fmt.Println("[OVERLAYFS] merged view after writes:")
	fmt.Printf("  %-28s  %-32s  %s\n", "path", "content", "layer")
	fmt.Printf("  %-28s  %-32s  %s\n",
		strings.Repeat("-", 27), strings.Repeat("-", 31), strings.Repeat("-", 24))
	for _, p := range allPaths {
		content, src := overlayLookup(p)
		marker := ""
		if src == upper.name {
			marker = "  ← CoW/new"
		}
		fmt.Printf("  %-28s  %-32s  %s%s\n", p, content, src, marker)
	}
	fmt.Println()
	fmt.Println("Key insight: lower layers never change. 100 containers share 1 image")
	fmt.Println("             on disk. Each gets its own upper layer — only the delta.")
	fmt.Println()

	// ── Summary ──────────────────────────────────────────────────────────────
	fmt.Println("============================================================")
	fmt.Println("  Key Insights:")
	fmt.Println("  1. Container = 4 kernel primitives, NOT a VM:")
	fmt.Println("     Namespace  → clone(CLONE_NEWPID|CLONE_NEWNET|CLONE_NEWNS)")
	fmt.Println("     Cgroup     → write numbers to /sys/fs/cgroup/.../cpu.max")
	fmt.Println("     pivot_root → one syscall, kernel swaps fs->root pointer")
	fmt.Println("     OverlayFS  → lowerdir/upperdir, CoW on first write")
	fmt.Println("  2. Pipe(2) is how container stdout reaches your terminal")
	fmt.Println("  3. memory.max = the RSS cap the kernel checks on every mmap/brk")
	fmt.Println("  4. cpu.max = quota/period throttle, enforced per CFS period")
	fmt.Println("  5. containerd/runc run exactly this sequence on every pod start")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num02CNINetworkPlugin
//
// TOPIC: Container Networking (CNI) — Veth pairs, Bridges, VXLAN
//   - stub — implement when teaching
//
// Real world: Debugging "CrashLoopBackOff" and pod-to-pod latency
// =============================================================================
func Num02CNINetworkPlugin() {
	// stub — implement when teaching
}

// =============================================================================
// Num03K8sSchedulerSimulator
//
// TOPIC: Kubernetes Architecture — API Server, Scheduler, Kubelet, Etcd
//   - stub — implement when teaching
//
// Real world: The Control Plane logic that manages 50+ clusters
// =============================================================================
func Num03K8sSchedulerSimulator() {
	// stub — implement when teaching
}

// =============================================================================
// Num04CustomControllerDemo
//
// TOPIC: The Operator Pattern — CRDs, Controller Loops, Reconciliation
//   - stub — implement when teaching
//
// Real world: Building "Database-as-a-Service" on K8s (Crossplane)
// =============================================================================
func Num04CustomControllerDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num05SidecarProxyInjection
//
// TOPIC: Service Mesh Internals — Sidecar Injection, Envoy Proxy, iptables
//   - stub — implement when teaching
//
// Real world: Rolling out Istio to 1,000+ services without downtime
// =============================================================================
func Num05SidecarProxyInjection() {
	// stub — implement when teaching
}

// =============================================================================
// Num06MTLSCertRotation
//
// TOPIC: Workload Identity & Cert Lifecycle — SPIFFE/SPIRE, X.509 SVID rotation
//   - stub — implement when teaching
//
// Real world: Securing internal traffic between Payment & User services
// =============================================================================
func Num06MTLSCertRotation() {
	// stub — implement when teaching
}

// =============================================================================
// Num07GitOpsSyncEngine
//
// TOPIC: GitOps Architecture — Drift Detection, Sync Loops, ApplicationSets
//   - stub — implement when teaching
//
// Real world: Managing 10k+ ArgoCD apps across multi-cloud
// =============================================================================
func Num07GitOpsSyncEngine() {
	// stub — implement when teaching
}

// =============================================================================
// Num08AdmissionWebhookDemo
//
// TOPIC: Policy as Code (OPA) — Rego, Admission Controllers, Webhooks
//   - stub — implement when teaching
//
// Real world: Preventing devs from deploying insecure containers
// =============================================================================
func Num08AdmissionWebhookDemo() {
	// stub — implement when teaching
}

// =============================================================================
// Num09LoadBalancerInternals
//
// TOPIC: Ingress & API Gateway — L4 vs L7, BGP, Anycast
//   - stub — implement when teaching
//
// Real world: Handling 100M requests/sec at the cluster edge
// =============================================================================
func Num09LoadBalancerInternals() {
	// stub — implement when teaching
}

// =============================================================================
// Num10AutoScalingInternals
//
// TOPIC: Auto-Scaling Internals — HPA vs VPA, Custom Metrics, Cluster Autoscaler
//   - stub — implement when teaching
//
// Real world: Cost optimization: Packing bins tightly to save AWS bills
// =============================================================================
func Num10AutoScalingInternals() {
	// stub — implement when teaching
}

// =============================================================================
// Num11KVMVirtualization
//
// TOPIC: Virtualization (KVM) — Hypervisors, VirtIO, Firecracker MicroVMs
//   - stub — implement when teaching
//
// Real world: Running "Bare Metal Kubernetes" (the profile's specific skill)
// =============================================================================
func Num11KVMVirtualization() {
	// stub — implement when teaching
}

// =============================================================================
// Num12CSIStorageDriver
//
// TOPIC: Distributed Storage (CSI) — Persistent Volumes, Attach/Detach Controller
//   - stub — implement when teaching
//
// Real world: Managing stateful databases on dynamic pods
// =============================================================================
func Num12CSIStorageDriver() {
	// stub — implement when teaching
}

// =============================================================================
// Num13ObservabilityPipeline
//
// TOPIC: Observability Stack — Metrics (Prometheus/Thanos), Logs (Loki), Traces (OpenTelemetry)
//   - stub — implement when teaching
//
// Real world: Aggregating metrics from 5,000 microservices
// =============================================================================
func Num13ObservabilityPipeline() {
	// stub — implement when teaching
}

// =============================================================================
// Num14ChaosEngineering
//
// TOPIC: Chaos Engineering — Fault Injection, Network Partitioning
//   - stub — implement when teaching
//
// Real world: Testing if the "Multi-Cluster Disaster Recovery" actually works
// =============================================================================
func Num14ChaosEngineering() {
	// stub — implement when teaching
}

// =============================================================================
// Num15IDPArchitecture
//
// TOPIC: IDP Architecture — Service Catalog, Scaffolding, Golden Paths
//   - stub — implement when teaching
//
// Real world: Building the "Internal Heroku" (Backstage) for developers
// =============================================================================
func Num15IDPArchitecture() {
	// stub — implement when teaching
}

// =============================================================================
// Num16EBPFAndCilium
//
// TOPIC: eBPF & Cilium — BPF maps, kprobes, XDP, Cilium NetworkPolicy
//   - stub — implement when teaching
//
// Real world: Replacing iptables entirely; 0-overhead observability
// =============================================================================
func Num16EBPFAndCilium() {
	// stub — implement when teaching
}

// =============================================================================
// Num17MultiClusterNetworking
//
// TOPIC: Multi-Cluster Networking — Karmada, Submariner, ClusterIP federation
//   - stub — implement when teaching
//
// Real world: Cross-region DR: "If us-east-1 dies, route to ap-southeast-1"
// =============================================================================
func Num17MultiClusterNetworking() {
	// stub — implement when teaching
}

// =============================================================================
// Num18SupplyChainSecurity
//
// TOPIC: Supply Chain Security — Cosign, Sigstore, SBOM, SLSA provenance
//   - stub — implement when teaching
//
// Real world: ByteDance mandate: no unsigned image reaches prod
// =============================================================================
func Num18SupplyChainSecurity() {
	// stub — implement when teaching
}

// =============================================================================
// Num19FinOpsEngineering
//
// TOPIC: FinOps Engineering — Karpenter, Spot interruption, Chargeback, Rightsizing
//   - stub — implement when teaching
//
// Real world: "Save $2M/month on AWS by packing bins smarter"
// =============================================================================
func Num19FinOpsEngineering() {
	// stub — implement when teaching
}
