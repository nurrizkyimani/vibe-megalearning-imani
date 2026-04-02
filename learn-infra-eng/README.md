# M3/W3/D4 - Thu, 19 Mar 2026 (WIB)

## Cloud Native Platform Engineering — Full Curriculum

A structured learning path for Principal Platform Engineers at infrastructure-heavy firms like Platform9, HashiCorp, GoTo, Grab, and TikTok. Each topic is a runnable Go simulation focusing on Kubernetes Internals, Service Mesh, Multi-Cloud Architecture, and Building Internal Developer Platforms (IDP).

**Target audience:** Senior engineers who know how to **use** Kubernetes but need to learn how to **extend**, **secure**, and **architect** it at scale.

This curriculum is designed to make you the most sought-after infrastructure engineer at top tech companies. After mastering all 19 topics, you will command Principal-level compensation ($180K–$280K USD equivalent) and be able to design, build, and operate platforms that serve thousands of engineers.

---

## Curriculum Table

| # | Topic | Core Concept | Real World at TikTok/GoTo | Demo |
| --- | --- | --- | --- | --- |
| 01 | Linux Primitives | Namespaces (PID/Net), Cgroups, Chroot | Building a container runtime from scratch (like Docker) | Num01 ✅ |
| 02 | Container Networking (CNI) | Veth pairs, Bridges, Overlay Networks (VXLAN) | Debugging "CrashLoopBackOff" and pod-to-pod latency | Num02 🔲 |
| 03 | Kubernetes Architecture | API Server, Scheduler, Kubelet, Etcd | The Control Plane logic that manages 50+ clusters | Num03 🔲 |
| 04 | The Operator Pattern | CRDs, Controller Loops, Reconciliation | Building "Database-as-a-Service" on K8s (Crossplane) | Num04 🔲 |
| 05 | Service Mesh Internals | Sidecar Injection, Envoy Proxy, iptables capture | Rolling out Istio to 1,000+ services without downtime | Num05 🔲 |
| 06 | Workload Identity & Cert Lifecycle | SPIFFE/SPIRE, X.509 SVID rotation, trust domains | Securing internal traffic between Payment & User services | Num06 🔲 |
| 07 | GitOps Architecture | Drift Detection, Sync Loops, ApplicationSets | Managing 10k+ ArgoCD apps across multi-cloud | Num07 🔲 |
| 08 | Policy as Code (OPA) | Rego, Admission Controllers, Webhooks | Preventing devs from deploying insecure containers | Num08 🔲 |
| 09 | Ingress & API Gateway | L4 vs L7 load balancing, BGP, Anycast | Handling 100M requests/sec at the cluster edge | Num09 🔲 |
| 10 | Auto-Scaling Internals | HPA vs VPA, Custom Metrics, Cluster Autoscaler | Cost optimization: Packing bins tightly to save AWS bills | Num10 🔲 |
| 11 | Virtualization (KVM) | Hypervisors, VirtIO, Firecracker MicroVMs | Running "Bare Metal Kubernetes" (the profile's specific skill) | Num11 🔲 |
| 12 | Distributed Storage (CSI) | Persistent Volumes, Attach/Detach Controller | Managing stateful databases on dynamic pods | Num12 🔲 |
| 13 | Observability Stack | Metrics (Prometheus/Thanos), Logs (Loki), Traces (OpenTelemetry) | Aggregating metrics from 5,000 microservices | Num13 🔲 |
| 14 | Chaos Engineering | Fault Injection, Network Partitioning | Testing if the "Multi-Cluster Disaster Recovery" actually works | Num14 🔲 |
| 15 | IDP Architecture | Service Catalog, Scaffolding, Golden Paths (Backstage) | Building the "Internal Heroku" for developers | Num15 🔲 |
| 16 | eBPF & Cilium | BPF maps, kprobes, XDP, Cilium NetworkPolicy | Replacing iptables entirely; 0-overhead observability | Num16 🔲 |
| 17 | Multi-Cluster Networking | Karmada, Submariner, ClusterIP federation, VIP routing | Cross-region DR: "If us-east-1 dies, route to ap-southeast-1" | Num17 🔲 |
| 18 | Supply Chain Security | Cosign/Sigstore image signing, SBOM, SLSA provenance | ByteDance mandate: no unsigned image reaches prod | Num18 🔲 |
| 19 | FinOps Engineering | Karpenter node provisioning, Spot interruption, Chargeback | "Save $2M/month on AWS by packing bins smarter" | Num19 🔲 |

**Legend:** ✅ = Fully implemented | 🔲 = Skeleton stub (to be implemented)

**Cross-references for advanced topics:**

*   **Linux Primitives** → See Num01 (foundation for all container tech)
*   **CNI Networking** → See Num02 (also: eBPF in Num16)
*   **Kubernetes Control Plane** → See Num03 (also: Operators in Num04)
*   **Service Mesh** → See Num05 (also: mTLS in Num06)
*   **Workload Identity** → See Num06 (also: Zero Trust architecture)
*   **GitOps** → See Num07 (also: IDP in Num15)
*   **Policy Enforcement** → See Num08 (also: Supply Chain in Num18)
*   **Multi-Cluster** → See Num17 (also: Chaos Engineering in Num14)
*   **Cost Optimization** → See Num19 (also: Auto-Scaling in Num10)

---

## How to run

```
# run all topics (only uncommented ones in main.go will execute)
go run ./learn-infra-eng/
```

To run a specific topic, comment/uncomment the relevant line in `main.go`.

---

## Function signatures

```
func Num01BuildYourOwnContainer()      // ✅ fully implemented
func Num02CNINetworkPlugin()           // 🔲 stub — to be implemented
func Num03K8sSchedulerSimulator()      // 🔲 stub — to be implemented
func Num04CustomControllerDemo()       // 🔲 stub — to be implemented
func Num05SidecarProxyInjection()      // 🔲 stub — to be implemented
func Num06MTLSCertRotation()           // 🔲 stub — to be implemented
func Num07GitOpsSyncEngine()           // 🔲 stub — to be implemented
func Num08AdmissionWebhookDemo()       // 🔲 stub — to be implemented
func Num09LoadBalancerInternals()      // 🔲 stub — to be implemented
func Num10AutoScalingInternals()       // 🔲 stub — to be implemented
func Num11KVMVirtualization()          // 🔲 stub — to be implemented
func Num12CSIStorageDriver()           // 🔲 stub — to be implemented
func Num13ObservabilityPipeline()      // 🔲 stub — to be implemented
func Num14ChaosEngineering()           // 🔲 stub — to be implemented
func Num15IDPArchitecture()            // 🔲 stub — to be implemented
func Num16EBPFAndCilium()              // 🔲 stub — to be implemented
func Num17MultiClusterNetworking()     // 🔲 stub — to be implemented
func Num18SupplyChainSecurity()        // 🔲 stub — to be implemented
func Num19FinOpsEngineering()          // 🔲 stub — to be implemented
```

---

## Num01 — Linux Primitives (The Container)

### The Problem

You run `docker run nginx` and a web server appears. But what is a "container" actually?

Most engineers think containers are lightweight VMs. **Wrong.** A container is not a real kernel-level primitive. It's an illusion created by combining 4 existing Linux features.

Understanding these primitives is the difference between:

*   **Mid-level:** "I run `kubectl apply` and pods appear"
*   **Principal:** "I debugged a node where cgroup memory limits weren't being enforced due to a systemd slice misconfiguration"

### The Concept

A container is **4 Linux kernel features** working together:

**1\. Namespaces — Isolation**

Linux has 7 namespace types. Each isolates a different resource:

```
PID namespace   → process sees only its own children (appears as PID 1)
Net namespace   → isolated network stack (own IP, routes, iptables)
Mount namespace → isolated filesystem view
UTS namespace   → isolated hostname
IPC namespace   → isolated shared memory, message queues
User namespace  → isolated UID/GID mappings
Cgroup namespace → isolated cgroup view
```

Without namespaces, all processes share the same PID tree, network, and filesystem. With namespaces, each container thinks it's the only thing running.

**2\. Cgroups — Resource Limits**

Control Groups enforce resource quotas:

```
cpu.shares         → relative CPU weight (shares, not absolute cores)
memory.limit_in_bytes → hard RAM limit (OOM kill if exceeded)
blkio.throttle     → disk I/O limits
```

Without cgroups, one container can hog all 64 cores and 512 GB of RAM, starving others.

**3\. Chroot — Filesystem Jail**

Changes the root directory for a process. The container cannot see files outside its chroot jail.

```
HOST:        /home/user/secrets.txt  ← visible
CONTAINER:   /home/user/secrets.txt  ← no such file (different root)
```

**4\. Union Mounts (OverlayFS) — Layered Filesystem**

Docker images are **layers** stacked on top of each other. OverlayFS merges them into one view:

```
Layer 3 (writable)   → /var/lib/docker/overlay2/abc123/diff/
Layer 2 (read-only)  → ubuntu:22.04 + nginx
Layer 1 (read-only)  → ubuntu:22.04 base
```

When you write a file, it goes to Layer 3 (copy-on-write). Base layers stay unchanged, enabling layer reuse across containers.

### What the demo shows

```
Step 1: PID namespace — process appears as PID 1 inside container
Step 2: Net namespace — isolated network stack with veth pair
Step 3: Cgroups — enforce 512 MB memory limit (OOM kill if exceeded)
Step 4: Chroot — restrict filesystem root to a directory
Step 5: OverlayFS — merge 3 image layers into one unified view
```

### Key Insight

```
A container is NOT a VM. It shares the host kernel.
Container = Namespaces + Cgroups + Chroot + OverlayFS
`docker run` = syscalls to unshare(2), clone(2), setns(2), cgroupfs writes
containerd / CRI-O = the low-level runtime that orchestrates these syscalls
```

### Real World Usage

| System | Uses these primitives for |
| --- | --- |
| **Docker** | containerd → runc → syscalls (unshare, clone, setns) |
| **Kubernetes** | CRI (Container Runtime Interface) → containerd / CRI-O |
| **Firecracker (AWS Lambda)** | MicroVMs (KVM) instead of namespaces, but same isolation goals |
| **gVisor (Google)** | User-space kernel (Go) intercepts syscalls, no shared kernel |
| **Podman** | Daemonless Docker alternative, same underlying primitives |

### Connection to Other Topics

*   **Num02 (CNI):** The veth pair in Net namespace is how CNI plugins work
*   **Num05 (Service Mesh):** Istio injects an Envoy sidecar into the same Net namespace (shares pod IP)
*   **Num16 (eBPF):** eBPF attaches to kernel hooks (kprobes, XDP) — runs at lower level than namespaces

### Interview Tips

**"What is a container?"**

*   Start with: "A container is 4 Linux kernel features: namespaces, cgroups, chroot, and union mounts"
*   Explain one namespace type (PID or Net) in detail
*   Mention that containers share the host kernel (unlike VMs)

**"Why can't you run Windows containers on Linux?"**

*   Containers share the kernel. Linux kernel can't run Windows binaries.
*   You need Windows Server with Hyper-V isolation (which is actually a lightweight VM)

**"How does Docker limit a container to 2 GB RAM?"**

*   cgroup `memory.limit_in_bytes = 2147483648`
*   Kernel OOM killer terminates the process if exceeded
*   This is why you see "137" exit code (128 + 9 SIGKILL)

**"What happens when you** `**docker exec**` **into a running container?"**

*   `setns(2)` syscall to join the container's namespaces
*   New bash process spawned in the same PID, Net, Mount namespaces
*   The exec'd process sees the same filesystem and network as the container

---

## Num02 — Container Networking (CNI)

### The Problem

**To be implemented** — this section will cover veth pairs, Linux bridges, VXLAN overlay networks, and how CNI plugins (Calico, Flannel, Cilium) implement pod-to-pod networking.

---

## Num03 — Kubernetes Architecture

### The Problem

**To be implemented** — this section will cover the Kubernetes control plane: API Server (the source of truth), Scheduler (bin-packing algorithm), Kubelet (node agent), and etcd (distributed key-value store).

---

## Num04 — The Operator Pattern

### The Problem

**To be implemented** — this section will cover Custom Resource Definitions (CRDs), controller reconciliation loops, and how to build operators that extend Kubernetes (like Crossplane for multi-cloud infrastructure).

---

## Num05 — Service Mesh Internals

### The Problem

**To be implemented** — this section will cover sidecar injection (mutating webhooks), Envoy proxy configuration, iptables capture rules, and L7 traffic management.

---

## Num06 — Workload Identity & Cert Lifecycle

### The Problem

**To be implemented** — this section will cover SPIFFE/SPIRE architecture, X.509 SVID (identity certificates), automatic cert rotation, and trust domain federation for zero-trust networking.

---

## Num07 — GitOps Architecture

### The Problem

**To be implemented** — this section will cover drift detection (desired state vs actual state), sync loops, ApplicationSets for managing 10k+ apps, and how ArgoCD/Flux work under the hood.

---

## Num08 — Policy as Code (OPA)

### The Problem

**To be implemented** — this section will cover Rego policy language, admission webhooks (ValidatingWebhookConfiguration), and how to prevent insecure containers from being deployed.

---

## Num09 — Ingress & API Gateway

### The Problem

**To be implemented** — this section will cover L4 vs L7 load balancing, MetalLB BGP peering, Anycast routing, and handling 100M+ requests/sec at the cluster edge.

---

## Num10 — Auto-Scaling Internals

### The Problem

**To be implemented** — this section will cover Horizontal Pod Autoscaler (HPA) metrics pipeline, Vertical Pod Autoscaler (VPA) recommendations, custom metrics adapters, and Cluster Autoscaler node provisioning logic.

---

## Num11 — Virtualization (KVM)

### The Problem

**To be implemented** — this section will cover hypervisors (Type 1 vs Type 2), VirtIO paravirtualization, Firecracker MicroVMs, and running bare-metal Kubernetes on KVM.

---

## Num12 — Distributed Storage (CSI)

### The Problem

**To be implemented** — this section will cover Container Storage Interface (CSI) spec, Persistent Volume lifecycle, Attach/Detach Controller, and managing stateful databases (Postgres, Cassandra) on ephemeral pods.

---

## Num13 — Observability Stack

### The Problem

**To be implemented** — this section will cover the 3 pillars of observability: metrics (Prometheus scraping, Thanos long-term storage), logs (Loki pipeline, Fluentd), and traces (OpenTelemetry, Jaeger).

---

## Num14 — Chaos Engineering

### The Problem

**To be implemented** — this section will cover fault injection (Chaos Mesh, Litmus), network partitioning, pod kill experiments, and testing multi-cluster disaster recovery (DR) failover.

---

## Num15 — IDP Architecture

### The Problem

**To be implemented** — this section will cover Internal Developer Platforms (Backstage), service catalogs, scaffolding templates, golden paths, and building the "Internal Heroku" for your company.

---

## Num16 — eBPF & Cilium

### The Problem

**To be implemented** — this section will cover eBPF (extended Berkeley Packet Filter) programs, BPF maps, kprobes/tracepoints, XDP (eXpress Data Path), and how Cilium replaces iptables with eBPF for 10x performance.

---

## Num17 — Multi-Cluster Networking

### The Problem

**To be implemented** — this section will cover multi-cluster federation (Karmada), cross-cluster service discovery (Submariner), ClusterIP routing, and disaster recovery (DR) failover across regions.

---

## Num18 — Supply Chain Security

### The Problem

**To be implemented** — this section will cover image signing (Cosign/Sigstore), Software Bill of Materials (SBOM), SLSA provenance levels, Tekton Chains, and enforcing "no unsigned image reaches prod" policy at ByteDance scale.

---

## Num19 — FinOps Engineering

### The Problem

**To be implemented** — this section will cover Karpenter node provisioning, Spot instance interruption handling, chargeback/showback models, rightsizing APIs (Kubecost), and cost optimization strategies that save millions per month.

---

---

## What to Expect in a Principal Platform Engineer Interview

After mastering all 19 topics, here's what your interviews will look like at companies like TikTok, GoTo, Grab, HashiCorp, and Platform9.

### System Design Round (60–90 min)

**Typical prompt:**

> _"Design a platform that lets 500 engineering teams deploy microservices to Kubernetes across 3 clouds (AWS, GCP, Azure) without writing YAML."_

**What they're testing:**

Can you decompose this into layers?

*   **IDP layer** (Num15) — Backstage UI, service catalog, golden paths
*   **GitOps layer** (Num07) — ArgoCD ApplicationSets, drift detection
*   **Policy layer** (Num08) — OPA admission webhooks, image signing (Num18)
*   **Identity layer** (Num06) — SPIFFE/SPIRE for zero-trust mTLS
*   **Observability layer** (Num13) — unified metrics/logs/traces pipeline
*   **Cost layer** (Num19) — chargeback per team, Karpenter for Spot instances

Can you justify architectural choices with **tradeoffs**?

*   "ArgoCD vs Flux: ArgoCD has a better UI for 500 teams, Flux is lighter for 5 teams"
*   "Single large cluster vs multi-cluster: blast radius vs complexity"
*   "Istio vs Linkerd: Istio has more features, Linkerd is simpler to operate"

**The one answer that separates Principal from Senior:**

> "I'd start with a single shared IDP (Backstage) but **federate the GitOps layer** — each region has its own ArgoCD instance syncing from a global Git monorepo. Why? If the central ArgoCD goes down, all 3 regions stop deploying. Federated ArgoCD means us-east-1 dying doesn't block ap-southeast-1 deployments. Trade-off: now we need cross-region Git replication and consistent RBAC across 3 ArgoCD instances."

This answer shows you've operated systems at scale and debugged **availability failure modes**, not just happy paths.

### Deep-Dive / Debugging Round (60 min)

**Typical prompts:**

_"A pod is running but can't reach another pod in the same cluster. Walk me through your debugging steps."_

**Expected answer flow:**

*   Check CNI (Num02): `ip route`, `ip link`, veth pairs, bridge forwarding
*   Check NetworkPolicy (Num16): Cilium policy, iptables rules, eBPF programs
*   Check Service Mesh (Num05): Envoy sidecar logs, mTLS handshake failures
*   Check DNS: CoreDNS logs, `nslookup` inside pod
*   Check kube-proxy: iptables NAT rules for Service ClusterIP

_"The Kubernetes API server is responding slowly (5-second latency). How do you investigate?"_

**Expected answer flow:**

*   Check etcd (Num03): disk latency, compaction lag, raft leader election
*   Check admission webhooks (Num08): webhook timeout, external OPA service down
*   Check API server metrics (Num13): request latency histograms, queue depth
*   Check client rate limiting: too many LIST requests from controllers
*   Check network: DNS resolution for etcd endpoints, MTU issues

_"Pods are not being scheduled. What do you check?"_

**Expected answer flow:**

*   Check Scheduler logs (Num03): unschedulable reasons (taints, affinity, resources)
*   Check node resources: `kubectl describe node`, cgroup limits (Num01)
*   Check PVC binding: CSI driver issues (Num12), PV availability
*   Check Cluster Autoscaler (Num10): node scaling disabled, cloud quota exceeded

**The differentiator:** Principals don't just know the commands — they know **which layer failed** and **why** based on symptoms. They've debugged these failures at 3am in production.

### Coding / Take-Home Round (3–5 days)

**Typical prompt:**

> _"Write a Kubernetes Operator that watches a custom resource_ `_Database_` _and provisions a Postgres StatefulSet + Service + PVC automatically."_

**What they're testing:**

*   Can you write a controller reconciliation loop (Num04)?
*   Do you handle error cases (DB already exists, PVC provisioning fails)?
*   Do you implement finalizers (cleanup when `Database` is deleted)?
*   Do you write unit tests with fake clients?

**Sample code structure they expect:**

```
func (r *DatabaseReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
    // 1. Fetch the Database CR
    db := &v1alpha1.Database{}
    if err := r.Get(ctx, req.NamespacedName, db); err != nil {
        return reconcile.Result{}, client.IgnoreNotFound(err)
    }

    // 2. Check if StatefulSet exists
    sts := &appsv1.StatefulSet{}
    err := r.Get(ctx, types.NamespacedName{Name: db.Name, Namespace: db.Namespace}, sts)
    if errors.IsNotFound(err) {
        // 3. Create StatefulSet + Service + PVC
        sts := r.constructStatefulSet(db)
        if err := r.Create(ctx, sts); err != nil {
            return reconcile.Result{}, err
        }
        // Requeue to verify it's running
        return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
    }

    // 4. Update status
    db.Status.Phase = "Running"
    r.Status().Update(ctx, db)

    return reconcile.Result{}, nil
}
```

Topic 04 teaches this pattern. Without it, most candidates write an infinite loop and don't understand why the controller burns CPU.

### Compensation Signal

At firms like GoTo, Grab, TikTok, HashiCorp:

| Level | Knows | Compensation (USD equivalent) |
| --- | --- | --- |
| **Senior** | Uses Kubernetes | $120K–$180K |
| **Staff** | Extends Kubernetes (writes Operators) | $160K–$220K |
| **Principal** | Architectures Kubernetes (multi-cluster, eBPF, FinOps) | $200K–$300K |

**The eBPF differentiator (Num16):** Fewer than 5% of senior engineers can explain eBPF from first principles. If you can implement a toy eBPF program and explain why Cilium is 10x faster than iptables, you immediately jump to the top 5% of candidates. This alone is worth $20K–$50K in negotiation leverage.

### The One Answer That Proves You're Principal-Level

**Question:** _"How do you prevent one team from deploying a container that mines Bitcoin?"_

**Senior answer:** "Use OPA to block images from untrusted registries."

**Principal answer:**

> "Multi-layered defense. First, **supply chain** (Num18): enforce Sigstore image signing — only images signed by our CI/CD can deploy. Second, **admission policy** (Num08): OPA Rego rules block privileged containers, host network, and CAP\_SYS\_ADMIN. Third, **runtime** (Num01): cgroup cpu.shares limits prevent CPU hogging — even if a Bitcoin miner gets through, it can't monopolize the node. Fourth, **observability** (Num13): Prometheus alert on sustained high CPU usage by pod, OpenTelemetry traces to identify the call path. Fifth, **FinOps** (Num19): chargeback shows the offending team's cost spike — they get a bill.
> 
> Why all 5 layers? If one fails (e.g., someone bypasses OPA with a direct API call), the others catch it. This is **defense in depth**."

This answer shows you've designed **security architectures**, not just written policies. That's the difference between Staff and Principal.

---

---

## Why This Curriculum Makes You Hireable

### The Market Reality (2026)

Most "Kubernetes engineers" know:

*   `kubectl apply -f deployment.yaml`
*   Read logs with `kubectl logs`
*   Maybe helm install some charts

**They cannot:**

*   Explain why a pod can't reach another pod (CNI, eBPF, NetworkPolicy)
*   Write a Kubernetes Operator (CRDs, reconciliation loop)
*   Design multi-cluster disaster recovery (cross-region service mesh)
*   Reduce cloud costs by $2M/month (FinOps, Karpenter, Spot instances)

These 19 topics cover **exactly the gaps** between "I use Kubernetes" and "I build platforms on Kubernetes."

### The Interview Delta

After this curriculum, when an interviewer asks:

> _"Design a platform for 1,000 engineers"_

**Before:** "Uh... Kubernetes... and GitOps... and monitoring?"

**After:** "IDP (Num15) for self-service, GitOps (Num07) for deployment, OPA (Num08) + Sigstore (Num18) for security, SPIFFE (Num06) for zero-trust, Cilium (Num16) for 10x faster networking, Karpenter (Num19) for cost optimization, multi-cluster (Num17) for DR. Let me draw the architecture and explain the failure modes of each layer."

That's the difference between a $120K offer and a $250K offer.

### The Unique Value of Topics 16–19

These 4 topics (eBPF, Multi-Cluster, Supply Chain, FinOps) are **missing from 95% of Kubernetes learning resources**. They're also the topics that:

*   **eBPF (Num16):** Every CNI is moving to eBPF (Cilium, Calico eBPF mode). Knowing this in 2026 is like knowing Docker in 2016 — early-mover advantage.
*   **Multi-Cluster (Num17):** TikTok, Grab, GoTo all run 50+ clusters. Single-cluster knowledge doesn't scale.
*   **Supply Chain (Num18):** ByteDance (TikTok's parent) mandates Sigstore. This is now table stakes at top firms.
*   **FinOps (Num19):** Executives care about cost. Saving $2M/month gets you promoted faster than any technical feat.

These are **not toy topics**. They're the difference between Senior and Principal.

---

## License

This curriculum is open for educational use. Go build platforms. Go get Principal offers.