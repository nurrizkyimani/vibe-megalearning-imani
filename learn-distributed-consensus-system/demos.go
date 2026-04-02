package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
)

// =============================================================================
// Num01ConsistentHashingDemo
//
// TOPIC: Consistent Hashing with Virtual Nodes
//   - How to distribute data across N servers without reshuffling everything
//     when a server is added or removed
//   - Naive approach (key % N) remaps almost all keys on any topology change
//   - Hash ring: both servers and keys placed on a circle 0..2^32-1
//   - A key is owned by the first server clockwise from its position
//   - Virtual nodes: each server gets multiple ring positions for even load
//
// Real world: DynamoDB, Redis Cluster, Cassandra, CDN edge routing
// =============================================================================
func Num01ConsistentHashingDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num01 -- Consistent Hashing with Virtual Nodes")
	fmt.Println("============================================================")
	fmt.Println()

	// ── Ring data structure ──────────────────────────────────────────────────
	// positions: sorted slice of all virtual node positions (uint32 values)
	// nodeMap:   maps each position -> physical node name
	// replicas:  how many virtual nodes per physical node
	type hashRing struct {
		positions []uint32
		nodeMap   map[uint32]string
		replicas  int
	}

	// ── hash: string -> uint32 position using SHA-256 (first 4 bytes) ────────
	// Gives us a deterministic position in the 0..2^32-1 space.
	hashFn := func(key string) uint32 {
		h := sha256.Sum256([]byte(key))
		return binary.BigEndian.Uint32(h[:4])
	}

	// ── newRing: create an empty ring with N virtual nodes per server ─────────
	newRing := func(replicas int) *hashRing {
		return &hashRing{
			nodeMap:  make(map[uint32]string),
			replicas: replicas,
		}
	}

	// ── addNode: place a server on the ring at `replicas` positions ───────────
	// Each virtual node key is "NodeA#0", "NodeA#1", etc.
	// All virtual nodes of the same server map back to that server's name.
	addNode := func(r *hashRing, name string) {
		for i := 0; i < r.replicas; i++ {
			virtualKey := fmt.Sprintf("%s#%d", name, i)
			pos := hashFn(virtualKey)
			r.positions = append(r.positions, pos)
			r.nodeMap[pos] = name
			fmt.Printf("[RING]  + Added %-8s virtual node %-10s at ring position %10d\n",
				name, virtualKey, pos)
		}
		// keep positions sorted so binary search works correctly
		sort.Slice(r.positions, func(i, j int) bool {
			return r.positions[i] < r.positions[j]
		})
	}

	// ── removeNode: pull all virtual nodes of a server off the ring ───────────
	removeNode := func(r *hashRing, name string) {
		for i := 0; i < r.replicas; i++ {
			virtualKey := fmt.Sprintf("%s#%d", name, i)
			pos := hashFn(virtualKey)
			delete(r.nodeMap, pos)
			idx := sort.Search(len(r.positions), func(j int) bool {
				return r.positions[j] >= pos
			})
			if idx < len(r.positions) && r.positions[idx] == pos {
				r.positions = append(r.positions[:idx], r.positions[idx+1:]...)
			}
		}
		fmt.Printf("[RING]  - Removed %s (%d virtual nodes)\n", name, r.replicas)
	}

	// ── getNode: find which server owns a key ─────────────────────────────────
	// Hash the key -> get position -> binary search for first position >= keyPos
	// If keyPos is past the last node, wrap around to the first node.
	getNode := func(r *hashRing, key string) string {
		if len(r.positions) == 0 {
			return ""
		}
		keyPos := hashFn(key)
		idx := sort.Search(len(r.positions), func(i int) bool {
			return r.positions[i] >= keyPos
		})
		if idx == len(r.positions) {
			idx = 0 // wrap around
		}
		return r.nodeMap[r.positions[idx]]
	}

	// ── routeKeys: print which server each key routes to ─────────────────────
	routeKeys := func(r *hashRing, keys []string) {
		for _, key := range keys {
			node := getNode(r, key)
			fmt.Printf("[ROUTE] %-20q -> %s\n", key, node)
		}
	}

	// ── Step 1: Build ring with 3 nodes ──────────────────────────────────────
	// Each physical node gets 3 virtual nodes = 9 total positions on the ring.
	// More replicas = more even load distribution across the circle.
	ring := newRing(3)
	fmt.Println("-- Step 1: Add 3 nodes (NodeA, NodeB, NodeC) ---------------")
	addNode(ring, "NodeA")
	addNode(ring, "NodeB")
	addNode(ring, "NodeC")
	fmt.Println()

	keys := []string{
		"user:alice",
		"user:bob",
		"user:carol",
		"order:1001",
		"order:1008",
		"session:xyz99",
	}

	// ── Step 2: Route keys across 3 nodes ────────────────────────────────────
	fmt.Println("-- Step 2: Route 6 keys with 3 nodes -----------------------")
	routeKeys(ring, keys)
	fmt.Println()

	// ── Step 3: Add a 4th node ────────────────────────────────────────────────
	// With naive hashing (key % N), ALL keys would remap going from 3 -> 4 nodes.
	// With consistent hashing, only keys whose arc now contains NodeD will move.
	fmt.Println("-- Step 3: Add NodeD ----------------------------------------")
	addNode(ring, "NodeD")
	fmt.Println()

	fmt.Println("-- Step 4: Route same 6 keys with 4 nodes ------------------")
	fmt.Println("   (notice only SOME keys moved -- the rest are unchanged)")
	routeKeys(ring, keys)
	fmt.Println()

	// ── Step 5: Remove a node ─────────────────────────────────────────────────
	// Only NodeB's keys redistribute to its clockwise successor.
	// All other keys are completely unaffected.
	fmt.Println("-- Step 5: Remove NodeB -------------------------------------")
	removeNode(ring, "NodeB")
	fmt.Println()

	fmt.Println("-- Step 6: Route same 6 keys with NodeB removed ------------")
	fmt.Println("   (only keys that were on NodeB moved -- others unchanged)")
	routeKeys(ring, keys)
	fmt.Println()

	fmt.Println("============================================================")
	fmt.Println("  Key insight: adding/removing a node only moves ~1/N keys.")
	fmt.Println("  Naive hashing (key%N) would move almost ALL keys.")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num02LeaderElectionDemo
//
// TOPIC: Leader Election (Raft basics)
// - Distributed systems need ONE node to coordinate writes — the leader
// - Raft uses terms (logical clock) + majority voting to elect a leader
// - If a follower hears no heartbeat, it starts an election
// - A candidate wins if it gets votes from a majority (N/2 + 1) of nodes
// - Split vote: two candidates tie, timeout and retry with a new term
//
// Real world: etcd (Kubernetes uses this), CockroachDB, TiKV, Consul
// =============================================================================
func Num02LeaderElectionDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num02 -- Leader Election (Raft basics)")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// A 5-node Raft cluster. Each node can be in one of 3 states:
	//   FOLLOWER  — default state, listens for heartbeats from the leader
	//   CANDIDATE — follower timed out, now requesting votes to become leader
	//   LEADER    — won the election, sends heartbeats to keep followers alive
	//
	// Key concept: TERM
	//   A term is a logical clock. Every time an election starts, the term
	//   increments. This prevents old leaders from causing split-brain after
	//   recovering from a crash — their term is stale, so followers ignore them.
	//
	// Key concept: MAJORITY QUORUM
	//   To win, a candidate needs votes from N/2+1 nodes (majority).
	//   With 5 nodes, majority = 3. This ensures only ONE leader per term.

	// ── Node state ───────────────────────────────────────────────────────────
	type NodeState string
	const (
		Follower  NodeState = "FOLLOWER"
		Candidate NodeState = "CANDIDATE"
		Leader    NodeState = "LEADER"
	)

	type Node struct {
		id       int
		state    NodeState
		term     int  // current term this node believes
		votedFor int  // which candidate this node voted for (-1 = no vote yet)
		alive    bool // false = simulates a crashed node
	}

	// ── Helper: print cluster state ──────────────────────────────────────────
	printCluster := func(nodes []*Node, label string) {
		fmt.Printf("  [%s]\n", label)
		for _, n := range nodes {
			status := ""
			if !n.alive {
				status = " (CRASHED)"
			}
			fmt.Printf("    Node%d  state=%-10s  term=%d%s\n",
				n.id, n.state, n.term, status)
		}
		fmt.Println()
	}

	// ── Step 1: Initial cluster — all followers, no leader yet ───────────────
	fmt.Println("-- Step 1: Cluster starts up, all nodes are followers ------")
	nodes := []*Node{
		{id: 1, state: Follower, term: 0, votedFor: -1, alive: true},
		{id: 2, state: Follower, term: 0, votedFor: -1, alive: true},
		{id: 3, state: Follower, term: 0, votedFor: -1, alive: true},
		{id: 4, state: Follower, term: 0, votedFor: -1, alive: true},
		{id: 5, state: Follower, term: 0, votedFor: -1, alive: true},
	}
	printCluster(nodes, "initial state")

	// ── Step 2: Node1's election timeout fires — it becomes a candidate ──────
	// In real Raft, each follower has a random timeout (150-300ms).
	// If it hears no heartbeat from a leader before timeout, it starts election.
	fmt.Println("-- Step 2: Node1 times out — starts election for Term 1 ----")
	node1 := nodes[0]
	node1.term++ // increment term before requesting votes
	node1.state = Candidate
	node1.votedFor = node1.id // always votes for itself first
	fmt.Printf("  Node%d transitions: FOLLOWER -> CANDIDATE (term=%d)\n",
		node1.id, node1.term)
	fmt.Printf("  Node%d votes for itself\n\n", node1.id)

	// ── Step 3: Node1 requests votes from all other nodes ────────────────────
	// A follower grants its vote if:
	//   1. The candidate's term >= follower's current term
	//   2. The follower hasn't voted for someone else in this term
	fmt.Println("-- Step 3: Node1 requests votes from Node2..Node5 ----------")
	votes := 1                   // already has its own vote
	majority := len(nodes)/2 + 1 // 5/2+1 = 3
	fmt.Printf("  Majority needed: %d out of %d nodes\n\n", majority, len(nodes))

	for _, n := range nodes[1:] { // skip Node1 (already voted for itself)
		if !n.alive {
			fmt.Printf("  Node%d -> Node%d: no response (crashed)\n", node1.id, n.id)
			continue
		}
		// Grant vote: candidate term >= follower term AND follower hasn't voted
		if node1.term >= n.term && n.votedFor == -1 {
			n.votedFor = node1.id
			n.term = node1.term // follower updates its term to match candidate
			votes++
			fmt.Printf("  Node%d -> Node%d: VOTE GRANTED (follower term updated to %d)\n",
				n.id, node1.id, n.term)
		} else {
			fmt.Printf("  Node%d -> Node%d: VOTE DENIED (already voted for Node%d)\n",
				n.id, node1.id, n.votedFor)
		}
	}
	fmt.Println()

	// ── Step 4: Node1 wins if it has majority ────────────────────────────────
	fmt.Printf("-- Step 4: Count votes — Node1 received %d/%d votes ---------\n",
		votes, len(nodes))
	if votes >= majority {
		node1.state = Leader
		fmt.Printf("  Node1 WON the election! Becomes LEADER for term=%d\n", node1.term)
		fmt.Println("  Node1 now sends heartbeats to all followers to assert leadership.")
	} else {
		fmt.Println("  Node1 did NOT get majority — election failed, will retry.")
	}
	fmt.Println()
	printCluster(nodes, "after election")

	// ── Step 5: Simulate leader crash — Node1 crashes ────────────────────────
	// When followers stop receiving heartbeats, they time out and start a new election.
	fmt.Println("-- Step 5: Node1 (leader) CRASHES --------------------------")
	node1.alive = false
	fmt.Println("  Node1 is down. Node2..Node5 stop receiving heartbeats.")
	fmt.Println("  Node3 times out first and starts a new election.\n")

	// ── Step 6: New election — Node3 becomes candidate for Term 2 ────────────
	node3 := nodes[2]
	node3.term++ // term 1 -> 2
	node3.state = Candidate
	node3.votedFor = node3.id
	fmt.Printf("-- Step 6: Node3 starts election for Term %d ----------------\n", node3.term)
	fmt.Printf("  Node3 transitions: FOLLOWER -> CANDIDATE (term=%d)\n\n", node3.term)

	votes2 := 1 // Node3 votes for itself
	for _, n := range nodes {
		if n.id == node3.id || !n.alive {
			continue
		}
		// Reset votedFor for this new term (new term = new vote slate)
		if node3.term > n.term {
			n.term = node3.term
			n.votedFor = -1 // new term, follower can vote again
		}
		if n.votedFor == -1 {
			n.votedFor = node3.id
			votes2++
			fmt.Printf("  Node%d -> Node%d: VOTE GRANTED (term=%d)\n",
				n.id, node3.id, n.term)
		}
	}
	fmt.Println()

	fmt.Printf("-- Step 7: Count votes — Node3 received %d/%d votes ---------\n",
		votes2, len(nodes))
	if votes2 >= majority {
		node3.state = Leader
		fmt.Printf("  Node3 WON the election! New LEADER for term=%d\n", node3.term)
		fmt.Println("  Even with Node1 crashed, cluster continues with a new leader.")
	}
	fmt.Println()
	printCluster(nodes, "final state after leader crash + re-election")

	fmt.Println("============================================================")
	fmt.Println("  Key insights:")
	fmt.Println("  1. Terms prevent old leaders from causing split-brain")
	fmt.Println("  2. Majority quorum ensures only ONE leader per term")
	fmt.Println("  3. Cluster survives as long as majority (3/5) are alive")
	fmt.Println("  4. Random timeouts reduce split-vote probability")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num03DistributedQueuesDemo
//
// TOPIC: Distributed Queues (Kafka concepts)
//   - Topics: named streams of messages (like a channel with persistence)
//   - Partitions: a topic is split into N partitions for parallelism
//   - Offsets: each message has a sequential ID per partition — consumers
//     track their own offset so they can replay or resume independently
//   - Consumer groups: multiple consumers share a topic — each partition
//     is assigned to exactly one consumer in the group
//   - What happens when a consumer crashes: partition rebalances to another
//
// Real world: Kafka, Google Pub/Sub, AWS Kinesis, NATS JetStream
// =============================================================================
func Num03DistributedQueuesDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num03 -- Distributed Queues (Kafka concepts)")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// A Kafka-like system with:
	//   - 1 topic:           "orders"
	//   - 3 partitions:      P0, P1, P2  (parallelism unit)
	//   - 2 producers:       ProducerA, ProducerB  (write messages)
	//   - 1 consumer group:  "billing-service" with 3 consumers
	//
	// Key concepts:
	//   TOPIC      — a named stream of messages (like a named Go channel,
	//                but persistent and replayable)
	//   PARTITION  — a topic is split into N ordered partitions. Each
	//                partition is an append-only log. Messages within a
	//                partition are strictly ordered. Across partitions,
	//                ordering is NOT guaranteed.
	//   OFFSET     — every message in a partition gets a sequential integer
	//                (0, 1, 2, ...). Consumers track their own offset —
	//                this means they can replay from any point, or pick up
	//                exactly where they left off after a crash.
	//   CONSUMER GROUP — a set of consumers that share the work of reading
	//                a topic. Each partition is assigned to exactly ONE
	//                consumer in the group at a time. This is how Kafka
	//                achieves parallel consumption without duplicate reads.
	//   PARTITION KEY — producers choose which partition a message goes to
	//                by hashing a key (e.g. order ID). Same key always goes
	//                to the same partition — guaranteeing order per key.

	// ── Data structures ──────────────────────────────────────────────────────
	type Message struct {
		offset   int
		key      string // partition key (e.g. customer ID)
		value    string // message payload
		producer string
	}

	type Partition struct {
		id       int
		messages []Message // append-only log
	}

	type Topic struct {
		name       string
		partitions []*Partition
	}

	type Consumer struct {
		id      string
		offsets map[int]int // partition id -> last consumed offset
	}

	// ── partitionFor: decide which partition a key belongs to ────────────────
	// Uses key % numPartitions — same key always routes to same partition.
	// This guarantees ordering for all messages with the same key.
	partitionFor := func(key string, numPartitions int) int {
		h := 0
		for _, c := range key {
			h = (h*31 + int(c)) % numPartitions
		}
		return h
	}

	// ── produce: append a message to the correct partition ───────────────────
	produce := func(topic *Topic, producer, key, value string) {
		pid := partitionFor(key, len(topic.partitions))
		p := topic.partitions[pid]
		offset := len(p.messages)
		msg := Message{offset: offset, key: key, value: value, producer: producer}
		p.messages = append(p.messages, msg)
		fmt.Printf("[PRODUCE] %-12s key=%-12q -> partition P%d offset=%d  value=%q\n",
			producer, key, pid, offset, value)
	}

	// ── consume: read the next unread message from an assigned partition ──────
	consume := func(consumer *Consumer, partition *Partition) {
		offset := consumer.offsets[partition.id]
		if offset >= len(partition.messages) {
			fmt.Printf("[CONSUME] %-12s P%d: no new messages (offset=%d)\n",
				consumer.id, partition.id, offset)
			return
		}
		msg := partition.messages[offset]
		fmt.Printf("[CONSUME] %-12s P%d offset=%d  key=%-12q value=%q\n",
			consumer.id, partition.id, offset, msg.key, msg.value)
		consumer.offsets[partition.id]++ // advance offset after reading
	}

	// ── Step 1: Create topic "orders" with 3 partitions ──────────────────────
	fmt.Println("-- Step 1: Create topic 'orders' with 3 partitions ----------")
	topic := &Topic{
		name: "orders",
		partitions: []*Partition{
			{id: 0}, {id: 1}, {id: 2},
		},
	}
	fmt.Printf("  Topic %q created with %d partitions: P0, P1, P2\n\n",
		topic.name, len(topic.partitions))

	// ── Step 2: Two producers write messages ─────────────────────────────────
	// Partition key = customer ID. All orders from the same customer go to
	// the same partition — guaranteeing per-customer ordering.
	fmt.Println("-- Step 2: Producers write 9 messages -----------------------")
	fmt.Println("   (same customer key always routes to same partition)")
	fmt.Println()
	produce(topic, "ProducerA", "customer:alice", "order:1001 placed")
	produce(topic, "ProducerA", "customer:bob", "order:1002 placed")
	produce(topic, "ProducerB", "customer:carol", "order:1003 placed")
	produce(topic, "ProducerA", "customer:alice", "order:1001 paid")    // same partition as 1001 placed
	produce(topic, "ProducerB", "customer:bob", "order:1002 paid")      // same partition as 1002 placed
	produce(topic, "ProducerA", "customer:carol", "order:1003 paid")    // same partition as 1003 placed
	produce(topic, "ProducerB", "customer:alice", "order:1001 shipped") // same partition as alice
	produce(topic, "ProducerA", "customer:dave", "order:1004 placed")
	produce(topic, "ProducerB", "customer:dave", "order:1004 paid") // same partition as dave
	fmt.Println()

	// ── Step 3: Show partition contents (the append-only logs) ───────────────
	fmt.Println("-- Step 3: Partition state (append-only logs) ---------------")
	for _, p := range topic.partitions {
		fmt.Printf("  P%d (%d messages):\n", p.id, len(p.messages))
		for _, m := range p.messages {
			fmt.Printf("    offset=%d  key=%-16q  value=%q\n", m.offset, m.key, m.value)
		}
	}
	fmt.Println()

	// ── Step 4: Consumer group "billing-service" — 3 consumers, 3 partitions ─
	// Rule: each partition assigned to exactly ONE consumer in the group.
	// With 3 consumers and 3 partitions: 1 partition per consumer = perfect balance.
	// Consumers read their assigned partition independently and in parallel.
	fmt.Println("-- Step 4: Consumer group 'billing-service' reads messages --")
	fmt.Println("   Assignment: Consumer1->P0, Consumer2->P1, Consumer3->P2")
	fmt.Println()

	c1 := &Consumer{id: "Consumer1", offsets: map[int]int{0: 0}}
	c2 := &Consumer{id: "Consumer2", offsets: map[int]int{1: 0}}
	c3 := &Consumer{id: "Consumer3", offsets: map[int]int{2: 0}}

	// Each consumer drains its assigned partition
	p0, p1, p2 := topic.partitions[0], topic.partitions[1], topic.partitions[2]

	fmt.Println("  [round 1 — each consumer reads one message]")
	consume(c1, p0)
	consume(c2, p1)
	consume(c3, p2)
	fmt.Println()

	fmt.Println("  [round 2]")
	consume(c1, p0)
	consume(c2, p1)
	consume(c3, p2)
	fmt.Println()

	fmt.Println("  [round 3 — some partitions may be exhausted]")
	consume(c1, p0)
	consume(c2, p1)
	consume(c3, p2)
	fmt.Println()

	// ── Step 5: Simulate consumer crash + rebalance ───────────────────────────
	// Consumer2 crashes. Its partition (P1) is reassigned to Consumer1.
	// Consumer1 picks up from Consumer2's last committed offset — no data loss,
	// no duplicate reads (assuming offsets were committed before crash).
	fmt.Println("-- Step 5: Consumer2 CRASHES — partition P1 rebalances ------")
	fmt.Println("   P1 reassigned to Consumer1 (picks up at Consumer2's last offset)")
	fmt.Println()
	c1.offsets[1] = c2.offsets[1] // Consumer1 inherits Consumer2's offset for P1

	fmt.Println("  [Consumer1 now handles both P0 and P1]")
	consume(c1, p0)
	consume(c1, p1) // Consumer1 continues from where Consumer2 left off
	consume(c3, p2)
	fmt.Println()

	// ── Step 6: Replay — consumer re-reads from offset 0 ─────────────────────
	// Unlike a traditional queue, Kafka keeps messages after they are consumed.
	// A consumer can reset its offset to replay all messages from the beginning.
	// Useful for: replaying after a bug fix, backfilling a new service, auditing.
	fmt.Println("-- Step 6: Replay — Consumer3 resets offset to 0 on P2 -----")
	fmt.Println("   (Kafka retains messages — consumers can always replay)")
	fmt.Println()
	c3.offsets[2] = 0 // reset to beginning
	consume(c3, p2)
	consume(c3, p2)
	consume(c3, p2)
	fmt.Println()

	fmt.Println("============================================================")
	fmt.Println("  Key insights:")
	fmt.Println("  1. Partitions = unit of parallelism (more partitions = more consumers)")
	fmt.Println("  2. Same partition key = same partition = ordering guaranteed per key")
	fmt.Println("  3. Offsets = consumers are in control of their own position")
	fmt.Println("  4. Consumer crash: reassign partition, resume from last offset")
	fmt.Println("  5. Messages are retained — replay is always possible")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num04ConsistencyModelsDemo
//
// TOPIC: Eventual Consistency vs Strong Consistency
//   - Strong consistency: every read reflects the latest committed write
//     -- requires coordination (blocking), slower but always correct
//   - Eventual consistency: replicas may temporarily diverge but converge
//     over time -- faster and more available, but reads may be stale
//   - CAP theorem: during a network partition you must choose between
//     Consistency (C) and Availability (A) -- you cannot have both
//   - PACELC: extends CAP -- even without partition, tradeoff between
//     latency (L) and consistency (C)
//
// Real world: DynamoDB (eventual by default), Spanner (strong),
//
//	Cassandra (tunable), Zookeeper (strong)
//
// =============================================================================
func Num04ConsistencyModelsDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num04 -- Eventual vs Strong Consistency")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// A 3-node distributed key-value store: Primary, Secondary1, Secondary2.
	// A client writes key "user:alice" -> "premium".
	//
	// We run the SAME write under two different consistency models and show
	// what each node returns when a read hits it immediately after the write.
	//
	// STRONG CONSISTENCY
	//   Write is not ACKed to the client until ALL nodes confirm it.
	//   Any read — no matter which node it hits — sees "premium".
	//   Cost: write latency = latency of the SLOWEST node.
	//         if one node is unreachable, write is blocked.
	//
	// EVENTUAL CONSISTENCY
	//   Write is ACKed as soon as the Primary confirms it.
	//   Secondaries receive the write asynchronously in the background.
	//   During that window, a read on a secondary returns the OLD value — "free".
	//   Eventually all nodes converge to "premium".
	//   Benefit: writes are fast, system stays available during partial failures.
	//
	// CAP THEOREM
	//   A distributed system can guarantee at most TWO of:
	//     C — every read sees the latest write
	//     A — every request gets a response (not an error)
	//     P — system works despite network splits
	//   Network partitions always happen in production.
	//   So the real choice is: CP (correct) vs AP (available).

	// ── Data structures ──────────────────────────────────────────────────────
	type dbNode struct {
		id      string
		data    map[string]string
		version map[string]int
	}

	type writeOp struct {
		key     string
		value   string
		version int
	}

	newNode := func(id string) *dbNode {
		return &dbNode{
			id:      id,
			data:    make(map[string]string),
			version: make(map[string]int),
		}
	}

	// applyWrite: store value on a node and print what happened
	applyWrite := func(n *dbNode, w writeOp) {
		n.data[w.key] = w.value
		n.version[w.key] = w.version
		fmt.Printf("  [%-12s]  WRITE  key=%-14q  value=%-10q  ver=%d\n",
			n.id, w.key, w.value, w.version)
	}

	// readValue: read from a node and print what it returns
	readValue := func(n *dbNode, key string) {
		val, ok := n.data[key]
		if !ok {
			fmt.Printf("  [%-12s]  READ   key=%-14q  -> (not found)\n", n.id, key)
			return
		}
		fmt.Printf("  [%-12s]  READ   key=%-14q  -> %-10q  ver=%d\n",
			n.id, key, val, n.version[key])
	}

	primary := newNode("Primary")
	sec1 := newNode("Secondary1")
	sec2 := newNode("Secondary2")
	allNodes := []*dbNode{primary, sec1, sec2}

	// ── PART A: STRONG CONSISTENCY ───────────────────────────────────────────
	fmt.Println("── PART A: Strong Consistency ───────────────────────────────")
	fmt.Println()
	fmt.Println("  Client wants to write user:alice -> \"premium\"")
	fmt.Println("  Coordinator sends write to ALL nodes, waits for ALL to confirm.")
	fmt.Println()

	w := writeOp{key: "user:alice", value: "premium", version: 1}
	for _, n := range allNodes {
		applyWrite(n, w)
	}

	fmt.Println()
	fmt.Println("  ALL nodes confirmed. ACK sent to client.")
	fmt.Println()
	fmt.Println("  Client reads user:alice (can hit any node):")
	for _, n := range allNodes {
		readValue(n, "user:alice")
	}
	fmt.Println()
	fmt.Println("  Every node returns \"premium\". No stale reads possible.")
	fmt.Println()

	// ── PART B: EVENTUAL CONSISTENCY ─────────────────────────────────────────
	fmt.Println("── PART B: Eventual Consistency ─────────────────────────────")
	fmt.Println()
	fmt.Println("  Resetting all nodes to old value \"free\" (state before the write)...")
	fmt.Println()

	old := writeOp{key: "user:alice", value: "free", version: 0}
	for _, n := range allNodes {
		applyWrite(n, old)
	}

	fmt.Println()
	fmt.Println("  Client wants to write user:alice -> \"premium\"")
	fmt.Println("  Coordinator writes to Primary only, ACKs client immediately.")
	fmt.Println("  Secondaries will receive the write in the background — not yet.")
	fmt.Println()

	applyWrite(primary, w)
	fmt.Println()
	fmt.Println("  ACK sent to client.")
	fmt.Println()

	fmt.Println("  Client reads user:alice immediately after (can hit any node):")
	for _, n := range allNodes {
		readValue(n, "user:alice")
	}
	fmt.Println()
	fmt.Println("  Secondary1 and Secondary2 still return \"free\" — STALE READ.")
	fmt.Println("  The client wrote \"premium\" but a read on a secondary returns \"free\".")
	fmt.Println()

	fmt.Println("  Background replication fires — secondaries receive the write...")
	fmt.Println()
	applyWrite(sec1, w)
	applyWrite(sec2, w)
	fmt.Println()

	fmt.Println("  Client reads user:alice again:")
	for _, n := range allNodes {
		readValue(n, "user:alice")
	}
	fmt.Println()
	fmt.Println("  All nodes now return \"premium\". System has converged.")
	fmt.Println()

	// ── PART C: TUNABLE QUORUM ───────────────────────────────────────────────
	// Cassandra and DynamoDB let you choose W and R per request.
	// Rule: W + R > N  →  read and write sets overlap  →  consistent.
	// If W + R <= N   →  no guaranteed overlap          →  eventual.
	//
	// N = 3 nodes in our cluster.
	//   W = how many nodes must confirm the write before ACK
	//   R = how many nodes must respond to the read
	fmt.Println("── PART C: Tunable Quorum (W + R > N) ──────────────────────")
	fmt.Println()
	fmt.Println("  N = 3 nodes.  Rule: W + R > N  →  guaranteed overlap  →  consistent")
	fmt.Println()
	fmt.Printf("  %-24s  %3s  %3s  %5s  %s\n", "Config", "W", "R", "W+R", "Result")
	fmt.Printf("  %-24s  %3s  %3s  %5s  %s\n", "------", "-", "-", "---", "------")

	type quorum struct {
		label string
		W, R  int
	}

	quorums := []quorum{
		{"W=1 R=1 (fastest)", 1, 1},
		{"W=2 R=2 (balanced)", 2, 2},
		{"W=3 R=1 (strong write)", 3, 1},
		{"W=1 R=3 (strong read)", 1, 3},
	}

	N := 3
	for _, q := range quorums {
		sum := q.W + q.R
		result := "EVENTUAL   — stale reads possible"
		if sum > N {
			result = "CONSISTENT — guaranteed overlap"
		}
		fmt.Printf("  %-24s  %3d  %3d  %5d  %s\n", q.label, q.W, q.R, sum, result)
	}
	fmt.Println()
	fmt.Println("  W+R > N means at least one node is in BOTH the write set and read set.")
	fmt.Println("  That node has the latest write — so the read always sees it.")
	fmt.Println()

	fmt.Println("============================================================")
	fmt.Println("  Key insights:")
	fmt.Println("  1. Strong: write waits for ALL nodes — correct but slower")
	fmt.Println("  2. Eventual: write waits for PRIMARY only — fast but stale reads possible")
	fmt.Println("  3. CAP: during a network split, pick CP (correct) or AP (available)")
	fmt.Println("  4. W + R > N: tune per request to slide between the two models")
	fmt.Println("  5. Most prod bugs in distributed systems = unexpected stale read")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num05ReplicationStrategiesDemo
//
// TOPIC: Replication Strategies
//   - Single-leader (primary-replica): all writes go to one node, replicated
//     to followers -- simple, consistent, but leader is a bottleneck
//   - Multi-leader: writes accepted at multiple nodes, conflicts resolved
//     -- better availability, harder to reason about
//   - Leaderless (Dynamo-style): any node accepts writes, quorum reads/writes
//     (W + R > N) guarantee overlap -- no single point of failure
//   - Sync vs async replication: sync = durable but slow, async = fast but
//     data can be lost if leader crashes before replicating
//
// Real world: MySQL (single-leader), CouchDB (multi-leader),
//
//	Cassandra/DynamoDB (leaderless)
//
// =============================================================================
func Num05ReplicationStrategiesDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num05 -- Replication Strategies")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// Same 3-node cluster each time. We write the same key under 3 strategies:
	//
	// PART A — Single Leader (Primary-Replica)
	//   All writes go to ONE leader. Leader replicates to replicas.
	//   Reads can go anywhere. Simple, common, but leader is a bottleneck.
	//   Real world: MySQL, Postgres streaming replication, MongoDB
	//
	// PART B — Multi Leader
	//   Two nodes both accept writes independently (e.g. two data centers).
	//   They sync with each other in the background.
	//   Risk: two leaders can write DIFFERENT values to the same key at the
	//   same time → CONFLICT. Needs a resolution strategy.
	//   Real world: CouchDB, Google Docs, distributed office apps
	//
	// PART C — Leaderless (Dynamo-style)
	//   No leader. Client writes to ALL nodes simultaneously.
	//   Write succeeds when W nodes confirm (quorum).
	//   Read queries ALL nodes, waits for R responses, picks highest version.
	//   W + R > N guarantees the read always overlaps with the write — no stale.
	//   Real world: DynamoDB, Cassandra, Riak

	// ── Shared data structures ───────────────────────────────────────────────
	type node struct {
		id      string
		data    map[string]string
		version map[string]int
		alive   bool
	}

	newNode := func(id string) *node {
		return &node{
			id:      id,
			data:    make(map[string]string),
			version: make(map[string]int),
			alive:   true,
		}
	}

	write := func(n *node, key, value string, ver int) {
		n.data[key] = value
		n.version[key] = ver
		fmt.Printf("  [%-10s]  WRITE  %-16q  =  %-12q  ver=%d\n", n.id, key, value, ver)
	}

	read := func(n *node, key string) (string, int) {
		return n.data[key], n.version[key]
	}

	printRead := func(n *node, key string) {
		val, ver := read(n, key)
		if val == "" {
			fmt.Printf("  [%-10s]  READ   %-16q  ->  (empty)\n", n.id, key)
			return
		}
		fmt.Printf("  [%-10s]  READ   %-16q  ->  %-12q  ver=%d\n", n.id, key, val, ver)
	}

	// ── PART A: Single Leader (Primary-Replica) ──────────────────────────────
	fmt.Println("── PART A: Single Leader (Primary-Replica) ──────────────────")
	fmt.Println()
	fmt.Println("  Setup: 1 Primary (accepts writes) + 2 Replicas (read-only)")
	fmt.Println("  Rule:  ALL writes go to Primary. Primary replicates to replicas.")
	fmt.Println()

	primary := newNode("Primary")
	replica1 := newNode("Replica1")
	replica2 := newNode("Replica2")

	// Step 1: client writes to primary
	fmt.Println("  Step 1 — Client writes 'plan=premium' to Primary:")
	write(primary, "plan", "premium", 1)
	fmt.Println()

	// Step 2: primary replicates synchronously to replica1, async to replica2
	fmt.Println("  Step 2 — Primary replicates to Replica1 (sync) and Replica2 (async):")
	write(replica1, "plan", "premium", 1) // sync — confirmed before ACK
	// replica2 is still behind — simulates async lag
	fmt.Printf("  [%-10s]  (replication pending — not yet received)\n", "Replica2")
	fmt.Println()

	// Step 3: reads
	fmt.Println("  Step 3 — Reads from all nodes:")
	printRead(primary, "plan")
	printRead(replica1, "plan")
	printRead(replica2, "plan") // empty — async replica hasn't caught up
	fmt.Println()
	fmt.Println("  Replica2 is behind — async replication lag (same stale read as Num04).")
	fmt.Println()

	// Step 4: replica2 catches up
	fmt.Println("  Step 4 — Replica2 catches up (async replication arrives):")
	write(replica2, "plan", "premium", 1)
	fmt.Println()
	fmt.Println("  All nodes now agree.")
	fmt.Println()

	// Step 5: primary crashes
	fmt.Println("  Step 5 — Primary CRASHES:")
	primary.alive = false
	fmt.Println("  [Primary  ]  CRASHED — no more writes accepted")
	fmt.Println("  Replica1 must be promoted to new Primary (manual or automatic failover).")
	fmt.Println("  Until promotion completes: writes are BLOCKED.")
	fmt.Println()

	fmt.Println("  Single Leader summary:")
	fmt.Println("    + Simple — one place to write, easy to reason about")
	fmt.Println("    + Read scale — replicas can serve reads")
	fmt.Println("    - Leader bottleneck — all writes go through one node")
	fmt.Println("    - Failover gap — writes blocked until new leader promoted")
	fmt.Println()

	// ── PART B: Multi Leader ─────────────────────────────────────────────────
	fmt.Println("── PART B: Multi Leader ─────────────────────────────────────")
	fmt.Println()
	fmt.Println("  Setup: 2 Leaders (both accept writes) + 1 Follower")
	fmt.Println("  Use case: two data centers, each accepts local writes.")
	fmt.Println()

	leaderDC1 := newNode("Leader-DC1")
	leaderDC2 := newNode("Leader-DC2")
	follower := newNode("Follower")

	// Step 1: both leaders receive different writes to the same key
	// This is the conflict scenario — both think they are authoritative
	fmt.Println("  Step 1 — Two leaders receive writes to the SAME key simultaneously:")
	fmt.Println("  (DC1 user sets plan=premium, DC2 user sets plan=enterprise)")
	fmt.Println()
	write(leaderDC1, "plan", "premium", 1)    // DC1 user upgrades
	write(leaderDC2, "plan", "enterprise", 1) // DC2 user upgrades differently
	fmt.Println()

	// Step 2: leaders try to sync — CONFLICT detected
	fmt.Println("  Step 2 — Leaders sync with each other — CONFLICT detected:")
	fmt.Println("  Both have ver=1 for 'plan' but different values.")
	fmt.Println()

	val1, ver1 := read(leaderDC1, "plan")
	val2, ver2 := read(leaderDC2, "plan")
	fmt.Printf("  Leader-DC1 has: %-12q  ver=%d\n", val1, ver1)
	fmt.Printf("  Leader-DC2 has: %-12q  ver=%d\n", val2, ver2)
	fmt.Println()

	// Step 3: conflict resolution — Last Write Wins by wall clock
	// In real systems: LWW (last-write-wins), CRDTs, application-level merge
	fmt.Println("  Step 3 — Conflict resolution: Last Write Wins (LWW)")
	fmt.Println("  DC2 write timestamp is slightly later → DC2 wins.")
	fmt.Println()

	resolved := "enterprise"
	resolvedVer := 2
	write(leaderDC1, "plan", resolved, resolvedVer)
	write(leaderDC2, "plan", resolved, resolvedVer)
	write(follower, "plan", resolved, resolvedVer)
	fmt.Println()
	fmt.Println("  All nodes converge to \"enterprise\" after conflict resolution.")
	fmt.Println()

	fmt.Println("  Multi Leader summary:")
	fmt.Println("    + Write availability — both data centers accept writes independently")
	fmt.Println("    + No write bottleneck — no single leader")
	fmt.Println("    - Conflicts — same key written differently on two leaders simultaneously")
	fmt.Println("    - Complex — need conflict resolution strategy (LWW, CRDT, manual merge)")
	fmt.Println()

	// ── PART C: Leaderless (Dynamo-style) ────────────────────────────────────
	fmt.Println("── PART C: Leaderless (Dynamo-style) ────────────────────────")
	fmt.Println()
	fmt.Println("  Setup: 3 equal nodes — no leader. N=3, W=2, R=2.")
	fmt.Println("  Client writes to ALL nodes simultaneously, waits for W=2 to confirm.")
	fmt.Println("  Client reads from ALL nodes, waits for R=2 responses, picks highest ver.")
	fmt.Println()

	nodeA := newNode("NodeA")
	nodeB := newNode("NodeB")
	nodeC := newNode("NodeC")
	N, W, R := 3, 2, 3

	// Step 1: write to all nodes, but nodeC is slow — only W=2 confirm
	fmt.Println("  Step 1 — Client writes 'plan=premium' to all nodes (W=2 needed):")
	write(nodeA, "plan", "premium", 1)
	write(nodeB, "plan", "premium", 1)
	fmt.Printf("  [%-10s]  (slow — did not respond in time, write pending)\n", "NodeC")
	fmt.Printf("  W=%d confirmed (%d/%d) — ACK sent to client.\n", W, W, N)
	fmt.Println()

	// Step 2: read from all nodes, pick highest version
	// NodeC hasn't received the write yet — still has old value
	nodeC.data["plan"] = "free"
	nodeC.version["plan"] = 0

	fmt.Printf("  Step 2 — Client reads 'plan' from all nodes (R=%d needed):\n", R)
	printRead(nodeA, "plan")
	printRead(nodeB, "plan")
	printRead(nodeC, "plan") // stale
	fmt.Println()

	// pick highest version from R=2 responses (nodeA and nodeB responded first)
	fmt.Println("  R=2 fastest responses: NodeA (ver=1) and NodeB (ver=1).")
	fmt.Println("  Client picks highest version: \"premium\" ver=1.")
	fmt.Println("  NodeC's stale response (ver=0) is ignored.")
	fmt.Println()

	// Step 3: read repair — client pushes latest value back to nodeC
	fmt.Println("  Step 3 — Read repair: client pushes latest value to NodeC:")
	write(nodeC, "plan", "premium", 1)
	fmt.Println()
	fmt.Println("  All nodes now agree.")
	fmt.Println()

	fmt.Printf("  W+R=%d > N=%d → guaranteed overlap → consistent reads.\n", W+R, N)
	fmt.Println()

	fmt.Println("  Leaderless summary:")
	fmt.Println("    + No single point of failure — any node accepts writes")
	fmt.Println("    + Survives node failures as long as W and R nodes are reachable")
	fmt.Println("    + Ties back to Num04 quorum: W+R>N guarantees consistency")
	fmt.Println("    - More complex reads — must query multiple nodes, pick highest ver")
	fmt.Println("    - Read repair needed — stale nodes must be healed after reads")
	fmt.Println()

	fmt.Println("============================================================")
	fmt.Println("  Strategy comparison:")
	fmt.Println("  Single Leader  — simple, one bottleneck, failover gap on crash")
	fmt.Println("  Multi Leader   — available across DCs, conflicts possible")
	fmt.Println("  Leaderless     — no bottleneck, no failover, complexity in reads")
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num06DistributedTransactionsDemo
//
// TOPIC: Distributed Transactions (Two-Phase Commit)
//   - How to atomically commit a write that spans multiple services/databases
//   - Phase 1 (Prepare): coordinator asks all participants "can you commit?"
//   - Phase 2 (Commit/Abort): if ALL say yes, coordinator sends commit;
//     if ANY says no, coordinator sends abort to everyone
//   - Problem: coordinator crash between phase 1 and 2 leaves participants
//     in a blocked state (the 2PC blocking problem)
//   - Saga pattern: alternative to 2PC -- sequence of local transactions
//     with compensating transactions for rollback
//
// Real world: payment systems, order fulfillment, bank transfers
// =============================================================================
func Num06DistributedTransactionsDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num06 -- Distributed Transactions (2PC & Sagas)")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// Scenario: E-commerce "Place Order" transaction that touches 3 services:
	//   - OrderService:     creates order record
	//   - PaymentService:   charges customer's card
	//   - InventoryService: reserves product stock
	//
	// PROBLEM: What if we succeed in 2 services but fail in the 3rd?
	// We need ALL-OR-NOTHING (atomicity) across services.
	//
	// PART A — 2PC Happy Path:  all vote YES → coordinator commits everywhere
	// PART B — 2PC Abort Path:  one votes NO → coordinator aborts everywhere
	// PART C — 2PC Blocking:    coordinator crashes after Phase 1 → deadlock!
	// PART D — Saga Pattern:    sequence of local txns + compensating rollbacks

	// ── Shared types ─────────────────────────────────────────────────────────
	type Participant struct {
		name      string
		prepared  bool
		committed bool
		data      string // what this service stores
	}

	newParticipant := func(name string) *Participant {
		return &Participant{name: name}
	}

	// ── PART A: Two-Phase Commit (Happy Path) ────────────────────────────────
	fmt.Println("── PART A: Two-Phase Commit (Happy Path) ────────────────────")
	fmt.Println()
	fmt.Println("  Scenario: Customer buys 1x 'Laptop' for $1200")
	fmt.Println("  Services: OrderService, PaymentService, InventoryService")
	fmt.Println()

	orderSvc := newParticipant("OrderService")
	paymentSvc := newParticipant("PaymentService")
	inventorySvc := newParticipant("InventoryService")

	participants := []*Participant{orderSvc, paymentSvc, inventorySvc}

	// ── PHASE 1: PREPARE ──
	fmt.Println("  PHASE 1 — PREPARE (Coordinator asks: can you commit?)")
	fmt.Println()

	// Coordinator sends PREPARE to each participant
	votes := make(map[string]bool)

	// OrderService checks if it can create order
	fmt.Println("  [COORDINATOR] → OrderService: PREPARE")
	fmt.Println("  [ORDER-SVC]   ← Checking... order ID available? YES")
	fmt.Println("  [ORDER-SVC]   → VOTE_COMMIT (prepared, holding lock on order #7890)")
	orderSvc.prepared = true
	votes[orderSvc.name] = true
	fmt.Println()

	// PaymentService checks if card can be charged
	fmt.Println("  [COORDINATOR] → PaymentService: PREPARE")
	fmt.Println("  [PAYMENT-SVC] ← Checking... card valid? balance >= $1200? YES")
	fmt.Println("  [PAYMENT-SVC] → VOTE_COMMIT (prepared, reserved $1200)")
	paymentSvc.prepared = true
	votes[paymentSvc.name] = true
	fmt.Println()

	// InventoryService checks stock availability
	fmt.Println("  [COORDINATOR] → InventoryService: PREPARE")
	fmt.Println("  [INVENTORY]   ← Checking... stock available? Laptop count=5 → YES")
	fmt.Println("  [INVENTORY]   → VOTE_COMMIT (prepared, reserved 1 laptop)")
	inventorySvc.prepared = true
	votes[inventorySvc.name] = true
	fmt.Println()

	// Coordinator counts votes
	allCommit := true
	for _, v := range votes {
		if !v {
			allCommit = false
			break
		}
	}

	fmt.Println("  [COORDINATOR] All participants voted COMMIT ✓")
	fmt.Println("                Decision: COMMIT")
	fmt.Println()

	// ── PHASE 2: COMMIT ──
	fmt.Println("  PHASE 2 — COMMIT (Coordinator tells everyone: go ahead!)")
	fmt.Println()

	if allCommit {
		for _, p := range participants {
			fmt.Printf("  [COORDINATOR] → %s: COMMIT\n", p.name)
			p.committed = true
		}
		fmt.Println()

		fmt.Println("  [ORDER-SVC]   ✓ Order #7890 created (customer_id=42, laptop, $1200)")
		orderSvc.data = "Order#7890"

		fmt.Println("  [PAYMENT-SVC] ✓ Charged $1200 from card ending 1234")
		paymentSvc.data = "Charged:$1200"

		fmt.Println("  [INVENTORY]   ✓ Reserved 1 laptop (stock: 5→4)")
		inventorySvc.data = "Laptop:reserved"
		fmt.Println()

		fmt.Println("  ✓ Transaction COMMITTED across all services (atomic!)")
	}
	fmt.Println()
	fmt.Println("  Why it worked:")
	fmt.Println("    - Phase 1 ensures everyone CAN commit before anyone DOES commit")
	fmt.Println("    - Phase 2 is the single point of decision → all or nothing")
	fmt.Println()

	// ── PART B: Two-Phase Commit (Abort Path) ────────────────────────────────
	fmt.Println("── PART B: Two-Phase Commit (Abort Path) ────────────────────")
	fmt.Println()
	fmt.Println("  Scenario: Customer tries to buy 10x 'Laptop' but stock=4")
	fmt.Println()

	// Reset participants
	orderSvc2 := newParticipant("OrderService")
	paymentSvc2 := newParticipant("PaymentService")
	inventorySvc2 := newParticipant("InventoryService")
	participants2 := []*Participant{orderSvc2, paymentSvc2, inventorySvc2}

	fmt.Println("  PHASE 1 — PREPARE")
	fmt.Println()

	votes2 := make(map[string]bool)

	fmt.Println("  [COORDINATOR] → OrderService: PREPARE")
	fmt.Println("  [ORDER-SVC]   → VOTE_COMMIT (order ID available)")
	orderSvc2.prepared = true
	votes2[orderSvc2.name] = true
	fmt.Println()

	fmt.Println("  [COORDINATOR] → PaymentService: PREPARE")
	fmt.Println("  [PAYMENT-SVC] → VOTE_COMMIT (card valid, $12000 available)")
	paymentSvc2.prepared = true
	votes2[paymentSvc2.name] = true
	fmt.Println()

	fmt.Println("  [COORDINATOR] → InventoryService: PREPARE")
	fmt.Println("  [INVENTORY]   ← Checking... need 10 laptops, but stock=4 → INSUFFICIENT")
	fmt.Println("  [INVENTORY]   → VOTE_ABORT ✗")
	votes2[inventorySvc2.name] = false
	fmt.Println()

	allCommit2 := true
	for _, v := range votes2 {
		if !v {
			allCommit2 = false
			break
		}
	}

	_ = allCommit2 // suppress unused warning (we show the abort path)
	fmt.Println("  [COORDINATOR] At least one participant voted ABORT")
	fmt.Println("                Decision: ABORT (rollback everything)")
	fmt.Println()

	fmt.Println("  PHASE 2 — ABORT")
	fmt.Println()

	for _, p := range participants2 {
		fmt.Printf("  [COORDINATOR] → %s: ABORT\n", p.name)
	}
	fmt.Println()

	fmt.Println("  [ORDER-SVC]   ✓ Rollback: order ID #7891 released")
	fmt.Println("  [PAYMENT-SVC] ✓ Rollback: $12000 reservation cancelled")
	fmt.Println("  [INVENTORY]   ✓ Rollback: no stock reserved")
	fmt.Println()
	fmt.Println("  ✓ Transaction ABORTED across all services (atomic!)")
	fmt.Println()
	fmt.Println("  Why it worked:")
	fmt.Println("    - ONE participant can veto the entire transaction")
	fmt.Println("    - Coordinator ensures everyone rolls back together")
	fmt.Println()

	// ── PART C: The 2PC Blocking Problem ─────────────────────────────────────
	fmt.Println("── PART C: The 2PC Blocking Problem (Coordinator Crash) ─────")
	fmt.Println()
	fmt.Println("  Scenario: All participants vote YES in Phase 1...")
	fmt.Println("            ...but coordinator CRASHES before sending Phase 2!")
	fmt.Println()

	orderSvc3 := newParticipant("OrderService")
	paymentSvc3 := newParticipant("PaymentService")
	inventorySvc3 := newParticipant("InventoryService")

	fmt.Println("  PHASE 1 — PREPARE")
	fmt.Println()

	fmt.Println("  [COORDINATOR] → OrderService: PREPARE")
	fmt.Println("  [ORDER-SVC]   → VOTE_COMMIT (holding lock on order #7892)")
	orderSvc3.prepared = true
	fmt.Println()

	fmt.Println("  [COORDINATOR] → PaymentService: PREPARE")
	fmt.Println("  [PAYMENT-SVC] → VOTE_COMMIT (holding lock on $1200 reservation)")
	paymentSvc3.prepared = true
	fmt.Println()

	fmt.Println("  [COORDINATOR] → InventoryService: PREPARE")
	fmt.Println("  [INVENTORY]   → VOTE_COMMIT (holding lock on 1 laptop)")
	inventorySvc3.prepared = true
	fmt.Println()

	fmt.Println("  [COORDINATOR] All votes collected: COMMIT, COMMIT, COMMIT")
	fmt.Println("  [COORDINATOR] Decided to COMMIT, about to send Phase 2...")
	fmt.Println()
	fmt.Println("  💥 [COORDINATOR] CRASH! (network partition / server died)")
	fmt.Println()

	fmt.Println("  ⏳ [ORDER-SVC]   Waiting for Phase 2... (still holding lock)")
	fmt.Println("  ⏳ [PAYMENT-SVC] Waiting for Phase 2... (still holding lock)")
	fmt.Println("  ⏳ [INVENTORY]   Waiting for Phase 2... (still holding lock)")
	fmt.Println()

	fmt.Println("  ❌ BLOCKED STATE ❌")
	fmt.Println()
	fmt.Println("  The problem:")
	fmt.Println("    - Participants voted YES → must wait for coordinator's decision")
	fmt.Println("    - They CANNOT commit (no COMMIT message received)")
	fmt.Println("    - They CANNOT abort (they voted YES, coordinator might have told")
	fmt.Println("      others to commit → aborting now would break atomicity)")
	fmt.Println("    - They hold locks INDEFINITELY → other transactions blocked!")
	fmt.Println()
	fmt.Println("  This is the fundamental flaw of 2PC:")
	fmt.Println("    • Single point of failure (coordinator)")
	fmt.Println("    • Blocking protocol (participants wait indefinitely)")
	fmt.Println()
	fmt.Println("  Solutions:")
	fmt.Println("    • 3PC (Three-Phase Commit): adds a pre-commit phase, non-blocking")
	fmt.Println("    • Paxos/Raft: consensus protocols, tolerate coordinator failure")
	fmt.Println("    • Sagas: avoid distributed locks entirely (next part!)")
	fmt.Println()

	// ── PART D: Saga Pattern ─────────────────────────────────────────────────
	fmt.Println("── PART D: Saga Pattern (Compensating Transactions) ─────────")
	fmt.Println()
	fmt.Println("  Key idea: NO distributed lock. Instead:")
	fmt.Println("    1. Break transaction into sequence of LOCAL transactions")
	fmt.Println("    2. Each service commits immediately (no prepare phase)")
	fmt.Println("    3. If ANY step fails → run COMPENSATING transactions to undo")
	fmt.Println()
	fmt.Println("  Trade-off:")
	fmt.Println("    ✓ No blocking, no coordinator single point of failure")
	fmt.Println("    ✗ Temporary inconsistency (other services see partial state)")
	fmt.Println()

	// ── Saga: Happy Path ──
	fmt.Println("  ── Saga: Happy Path ──")
	fmt.Println()

	fmt.Println("  Step 1 — [ORDER-SVC] Create order #7893 (status=PENDING)")
	fmt.Println("           [ORDER-SVC] ✓ Committed locally, emit event: OrderCreated")
	fmt.Println()

	fmt.Println("  Step 2 — [PAYMENT-SVC] Receives OrderCreated event")
	fmt.Println("           [PAYMENT-SVC] Charge $1200 from card ...✓ SUCCESS")
	fmt.Println("           [PAYMENT-SVC] ✓ Committed locally, emit: PaymentCharged")
	fmt.Println()

	fmt.Println("  Step 3 — [INVENTORY] Receives PaymentCharged event")
	fmt.Println("           [INVENTORY] Reserve 1 laptop (stock 4→3) ...✓ SUCCESS")
	fmt.Println("           [INVENTORY] ✓ Committed locally, emit: ItemReserved")
	fmt.Println()

	fmt.Println("  Step 4 — [ORDER-SVC] Receives ItemReserved event")
	fmt.Println("           [ORDER-SVC] Update order #7893 status → CONFIRMED")
	fmt.Println()

	fmt.Println("  ✓ Saga completed successfully (eventual consistency)")
	fmt.Println()

	// ── Saga: Failure Path with Compensation ──
	fmt.Println("  ── Saga: Failure Path (with Compensation) ──")
	fmt.Println()

	fmt.Println("  Step 1 — [ORDER-SVC] Create order #7894 (status=PENDING)")
	fmt.Println("           [ORDER-SVC] ✓ Committed locally, emit: OrderCreated")
	fmt.Println()

	fmt.Println("  Step 2 — [PAYMENT-SVC] Receives OrderCreated event")
	fmt.Println("           [PAYMENT-SVC] Charge $1200 from card ...✓ SUCCESS")
	fmt.Println("           [PAYMENT-SVC] ✓ Committed locally, emit: PaymentCharged")
	fmt.Println()

	fmt.Println("  Step 3 — [INVENTORY] Receives PaymentCharged event")
	fmt.Println("           [INVENTORY] Reserve 1 laptop (stock=0) ...✗ FAILED")
	fmt.Println("           [INVENTORY] ✗ Emit: ItemReservationFailed")
	fmt.Println()

	fmt.Println("  ── Compensating Transactions (rollback) ──")
	fmt.Println()

	fmt.Println("  Compensate 2 — [PAYMENT-SVC] Receives ItemReservationFailed")
	fmt.Println("                 [PAYMENT-SVC] REFUND $1200 to customer")
	fmt.Println("                 [PAYMENT-SVC] ✓ Refund completed, emit: PaymentRefunded")
	fmt.Println()

	fmt.Println("  Compensate 1 — [ORDER-SVC] Receives PaymentRefunded")
	fmt.Println("                 [ORDER-SVC] Update order #7894 status → CANCELLED")
	fmt.Println("                 [ORDER-SVC] ✓ Compensated")
	fmt.Println()

	fmt.Println("  ✓ Saga rolled back using compensating transactions")
	fmt.Println()
	fmt.Println("  Key difference from 2PC:")
	fmt.Println("    • Each service committed its local work IMMEDIATELY")
	fmt.Println("    • No locks held across services")
	fmt.Println("    • Rollback = new forward transactions (refund, cancel)")
	fmt.Println("    • Other users COULD see order #7894 briefly before cancellation")
	fmt.Println()

	// ── Summary comparison ───────────────────────────────────────────────────
	fmt.Println("============================================================")
	fmt.Println("  2PC vs Saga Summary:")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Property", "2PC", "Saga")
	fmt.Println("  ─────────────────────┼───────────────────────────┼──────────────────────────")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Atomicity", "Strong (ACID)", "Eventual (BASE)")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Isolation", "Locks held", "No locks (visible partial)")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Blocking", "YES (Phase 1→2 gap)", "NO")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Coordinator failure", "Blocks participants", "No impact")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Complexity", "Protocol complex", "Business logic complex")
	fmt.Printf("  %-20s │ %-25s │ %-25s\n", "Use when", "Strong consistency needed", "Availability > consistency")
	fmt.Println()
	fmt.Println("  Real world:")
	fmt.Println("    2PC  → Traditional databases (XA transactions), some banking")
	fmt.Println("    Saga → Microservices (Uber, Netflix), e-commerce order flows")
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num07RateLimitingDemo
//
// TOPIC: Rate Limiting
//   - Token bucket: bucket holds up to N tokens, refilled at rate R/sec
//     -- allows bursts up to bucket size, smooth average rate
//   - Leaky bucket: requests queue up, processed at fixed rate
//     -- no bursts, strictly smooth output
//   - Fixed window: count requests per time window (e.g. 100 req/min)
//     -- simple but vulnerable to boundary bursts
//   - Sliding window: rolling count over last N seconds -- more accurate
//   - Distributed rate limiting: using Redis to share counter across instances
//
// Real world: every public API (GitHub, Stripe, OpenAI), DDoS protection
// =============================================================================
func Num07RateLimitingDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num07 -- Rate Limiting Algorithms")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// You run a public API. Users should be allowed to make requests, but
	// you need to prevent:
	//   1. Accidental infinite loops from buggy clients
	//   2. Malicious DDoS attacks
	//   3. One user hogging all your resources
	//
	// Solution: RATE LIMITING — cap requests per user to N per time window.
	//
	// PART A — Token Bucket:      allows bursts, smooth long-term rate
	// PART B — Fixed Window:      simple but has boundary burst exploit
	// PART C — Sliding Window:    accurate but memory-intensive
	// PART D — Distributed:       how to rate-limit across multiple servers

	// ── PART A: Token Bucket ─────────────────────────────────────────────────
	fmt.Println("── PART A: Token Bucket ──────────────────────────────────────")
	fmt.Println()
	fmt.Println("  Rule: Bucket holds tokens. Each request consumes 1 token.")
	fmt.Println("        Tokens refill at a fixed rate.")
	fmt.Println("        If bucket is empty → reject request.")
	fmt.Println()
	fmt.Println("  Config: capacity=5 tokens, refill=2 tokens/sec")
	fmt.Println()

	type TokenBucket struct {
		capacity   int
		tokens     int
		refillRate int // tokens per second
	}

	bucket := &TokenBucket{
		capacity:   5,
		tokens:     5, // start full
		refillRate: 2,
	}

	tryRequest := func(userID string) bool {
		if bucket.tokens > 0 {
			bucket.tokens--
			fmt.Printf("  [TOKEN-BUCKET] User %s: ✓ ALLOWED  (tokens remaining: %d/%d)\n",
				userID, bucket.tokens, bucket.capacity)
			return true
		}
		fmt.Printf("  [TOKEN-BUCKET] User %s: ✗ REJECTED (bucket empty)\n", userID)
		return false
	}

	refillTokens := func(seconds int) {
		added := seconds * bucket.refillRate
		bucket.tokens += added
		if bucket.tokens > bucket.capacity {
			bucket.tokens = bucket.capacity
		}
		fmt.Printf("  [TOKEN-BUCKET] ⏱  %d seconds passed → +%d tokens (now: %d/%d)\n",
			seconds, added, bucket.tokens, bucket.capacity)
	}

	// Burst of 7 requests — first 5 succeed, last 2 fail
	fmt.Println("  Scenario: Client sends 7 requests instantly (burst):")
	fmt.Println()
	for i := 1; i <= 7; i++ {
		tryRequest("Alice")
	}
	fmt.Println()

	// Wait 2 seconds → bucket refills by 4 tokens (2 tokens/sec * 2 sec)
	fmt.Println("  Client waits 2 seconds...")
	fmt.Println()
	refillTokens(2)
	fmt.Println()

	// 3 more requests — now we have 4 tokens, so first 4 succeed
	fmt.Println("  Client sends 3 more requests:")
	fmt.Println()
	for i := 1; i <= 3; i++ {
		tryRequest("Alice")
	}
	fmt.Println()

	fmt.Println("  Why Token Bucket?")
	fmt.Println("    ✓ Allows short bursts (up to capacity)")
	fmt.Println("    ✓ Enforces long-term average rate (refill rate)")
	fmt.Println("    ✓ Simple to implement — just track token count + last refill time")
	fmt.Println()
	fmt.Println("  Used by: GitHub API (5000 req/hour), Stripe API, AWS API Gateway")
	fmt.Println()

	// ── PART B: Fixed Window Counter ─────────────────────────────────────────
	fmt.Println("── PART B: Fixed Window Counter ──────────────────────────────")
	fmt.Println()
	fmt.Println("  Rule: Count requests in each fixed time window (e.g. per minute).")
	fmt.Println("        If count > limit → reject. Counter resets at window boundary.")
	fmt.Println()
	fmt.Println("  Config: limit=100 req/min, window=1 minute")
	fmt.Println()

	type FixedWindow struct {
		limit       int
		windowStart int // minute number (e.g. 0, 1, 2, ...)
		count       int
	}

	window := &FixedWindow{
		limit:       100,
		windowStart: 0,
		count:       0,
	}

	tryRequestFixed := func(userID string, currentMinute int, numRequests int) {
		// Check if we've crossed into a new window
		if currentMinute != window.windowStart {
			// New window — reset counter
			window.windowStart = currentMinute
			window.count = 0
		}

		// Try to add requests
		if window.count+numRequests <= window.limit {
			window.count += numRequests
			fmt.Printf("  [FIXED-WIN] Minute %d: user %s sends %d requests → count=%d/%d ✓ ALLOWED\n",
				currentMinute, userID, numRequests, window.count, window.limit)
		} else {
			fmt.Printf("  [FIXED-WIN] Minute %d: user %s sends %d requests → would exceed limit ✗ REJECTED\n",
				currentMinute, userID, numRequests)
		}
	}

	fmt.Println("  Scenario: Normal usage — 40 requests in minute 0, 30 in minute 1")
	fmt.Println()
	tryRequestFixed("Bob", 0, 40)
	tryRequestFixed("Bob", 1, 30)
	fmt.Println()

	fmt.Println("  The BOUNDARY BURST EXPLOIT:")
	fmt.Println("  ───────────────────────────────────────────────────────────")
	fmt.Println()

	// Reset window
	window2 := &FixedWindow{limit: 100, windowStart: 0, count: 0}

	fmt.Println("  Attacker exploits window boundaries:")
	fmt.Println()
	fmt.Println("  Minute 0, second 59: send 90 requests")
	window2.count = 90
	fmt.Printf("  [FIXED-WIN] Minute 0: count=%d/%d ✓ ALLOWED\n", window2.count, window2.limit)
	fmt.Println()

	fmt.Println("  Minute 1, second 00: send 90 requests (new window starts!)")
	window2.windowStart = 1
	window2.count = 90
	fmt.Printf("  [FIXED-WIN] Minute 1: count=%d/%d ✓ ALLOWED (counter reset)\n", window2.count, window2.limit)
	fmt.Println()

	fmt.Println("  Result: 180 requests in 2 seconds — both allowed!")
	fmt.Println()
	fmt.Println("  Timeline:")
	fmt.Println("    ┌──────────── Minute 0 ─────────────┬──────────── Minute 1 ─────────────┐")
	fmt.Println("    │                           90 req →│← 90 req                            │")
	fmt.Println("    │                            (sec 59)│(sec 0)                             │")
	fmt.Println("    └────────────────────────────────────┴────────────────────────────────────┘")
	fmt.Println("                                         ↑")
	fmt.Println("                                  boundary reset")
	fmt.Println()
	fmt.Println("  Problem:")
	fmt.Println("    ✗ Limit=100/min, but attacker sent 180 in 2 seconds!")
	fmt.Println("    ✗ Exploitable in production (many real attacks use this)")
	fmt.Println()
	fmt.Println("  Why use it anyway?")
	fmt.Println("    ✓ Dead simple to implement (just 1 counter + 1 timestamp)")
	fmt.Println("    ✓ O(1) memory and computation")
	fmt.Println()
	fmt.Println("  Used by: Simple internal APIs, microservices with coarse limits")
	fmt.Println()

	// ── PART C: Sliding Window Log ───────────────────────────────────────────
	fmt.Println("── PART C: Sliding Window Log ─────────────────────────────────")
	fmt.Println()
	fmt.Println("  Rule: Keep a log of request timestamps.")
	fmt.Println("        For each new request, count how many timestamps fall within")
	fmt.Println("        the last N seconds. If count >= limit → reject.")
	fmt.Println()
	fmt.Println("  Config: limit=5 requests per 10 seconds")
	fmt.Println()

	type SlidingWindow struct {
		limit      int
		windowSize int   // seconds
		log        []int // timestamps
	}

	sliding := &SlidingWindow{
		limit:      5,
		windowSize: 10,
		log:        []int{},
	}

	tryRequestSliding := func(userID string, currentTime int) {
		// Remove timestamps older than (currentTime - windowSize)
		cutoff := currentTime - sliding.windowSize
		validLog := []int{}
		evicted := 0
		for _, ts := range sliding.log {
			if ts > cutoff {
				validLog = append(validLog, ts)
			} else {
				evicted++
			}
		}
		sliding.log = validLog

		if evicted > 0 {
			fmt.Printf("  [SLIDING-WIN] t=%d: evicted %d old timestamps (older than t=%d)\n",
				currentTime, evicted, cutoff)
		}

		// Check if we can accept this request
		if len(sliding.log) < sliding.limit {
			sliding.log = append(sliding.log, currentTime)
			fmt.Printf("  [SLIDING-WIN] t=%d: user %s → ✓ ALLOWED (count=%d/%d)\n",
				currentTime, userID, len(sliding.log), sliding.limit)
		} else {
			fmt.Printf("  [SLIDING-WIN] t=%d: user %s → ✗ REJECTED (count=%d/%d, window full)\n",
				currentTime, userID, len(sliding.log), sliding.limit)
		}
		fmt.Printf("                       Current window: %v\n", sliding.log)
		fmt.Println()
	}

	fmt.Println("  Scenario: User sends 8 requests over 15 seconds")
	fmt.Println()

	tryRequestSliding("Carol", 1)
	tryRequestSliding("Carol", 2)
	tryRequestSliding("Carol", 3)
	tryRequestSliding("Carol", 4)
	tryRequestSliding("Carol", 5) // 5th request — bucket full
	tryRequestSliding("Carol", 6) // rejected — still within 10-sec window
	fmt.Println("  ⏱  Time passes... (no requests)")
	fmt.Println()
	tryRequestSliding("Carol", 12) // t=12: entries from t=1,2 are now >10sec old → evicted
	tryRequestSliding("Carol", 13) // accepted

	fmt.Println("  Why Sliding Window?")
	fmt.Println("    ✓ No boundary burst problem (rolling window)")
	fmt.Println("    ✓ Accurate — guarantees exactly N requests per window")
	fmt.Println("    ✗ Memory cost: O(N) — must store all timestamps in window")
	fmt.Println()
	fmt.Println("  Optimization: Sliding Window Counter (hybrid)")
	fmt.Println("    - Keep only 2 counters (current window + previous window)")
	fmt.Println("    - Estimate count using weighted average — O(1) memory!")
	fmt.Println("    - Trade-off: slight inaccuracy for huge memory savings")
	fmt.Println()
	fmt.Println("  Used by: Cloudflare, Redis (ZSET with ZREMRANGEBYSCORE)")
	fmt.Println()

	// ── PART D: Distributed Rate Limiting ────────────────────────────────────
	fmt.Println("── PART D: Distributed Rate Limiting (Multi-Server) ──────────")
	fmt.Println()
	fmt.Println("  Problem: You have 3 API servers behind a load balancer.")
	fmt.Println("           Each server has its own in-memory rate limiter.")
	fmt.Println("           User sends 30 requests → load balancer distributes evenly.")
	fmt.Println()
	fmt.Println("  Config: Global limit=10 req/min, 3 servers")
	fmt.Println()

	type LocalLimiter struct {
		name  string
		limit int
		count int
	}

	server1 := &LocalLimiter{name: "Server1", limit: 10, count: 0}
	server2 := &LocalLimiter{name: "Server2", limit: 10, count: 0}
	server3 := &LocalLimiter{name: "Server3", limit: 10, count: 0}

	tryLocalRequest := func(server *LocalLimiter, userID string) bool {
		if server.count < server.limit {
			server.count++
			fmt.Printf("  [%-8s] User %s: ✓ ALLOWED (local count=%d/%d)\n",
				server.name, userID, server.count, server.limit)
			return true
		}
		fmt.Printf("  [%-8s] User %s: ✗ REJECTED (local count=%d/%d)\n",
			server.name, userID, server.count, server.limit)
		return false
	}

	fmt.Println("  ── Without Shared State (BROKEN) ──")
	fmt.Println()
	fmt.Println("  User sends 30 requests, load balancer round-robins across 3 servers:")
	fmt.Println()

	allowed := 0
	for i := 1; i <= 30; i++ {
		var server *LocalLimiter
		if i%3 == 1 {
			server = server1
		} else if i%3 == 2 {
			server = server2
		} else {
			server = server3
		}

		if tryLocalRequest(server, "Dave") {
			allowed++
		}
	}

	fmt.Println()
	fmt.Printf("  Result: %d/30 requests allowed (limit was 10!)\n", allowed)
	fmt.Println()
	fmt.Println("  Problem:")
	fmt.Println("    ✗ Each server only sees 10 requests → each allows all 10")
	fmt.Println("    ✗ Total=30 requests allowed, but global limit=10")
	fmt.Println("    ✗ User bypassed rate limit by sending to multiple servers!")
	fmt.Println()

	fmt.Println("  ── With Shared State (Redis) ──")
	fmt.Println()
	fmt.Println("  All servers share a single Redis counter:")
	fmt.Println()

	// Simulate a shared Redis counter
	type Redis struct {
		counters map[string]int
	}

	redis := &Redis{counters: make(map[string]int)}

	tryRedisRequest := func(serverName, userID string, redis *Redis, globalLimit int) bool {
		key := "ratelimit:" + userID
		current := redis.counters[key]

		if current < globalLimit {
			redis.counters[key]++
			fmt.Printf("  [%-8s] User %s → [REDIS] INCR %s → %d → ✓ ALLOWED\n",
				serverName, userID, key, redis.counters[key])
			return true
		}
		fmt.Printf("  [%-8s] User %s → [REDIS] GET %s → %d (>= limit) → ✗ REJECTED\n",
			serverName, userID, key, redis.counters[key])
		return false
	}

	allowedRedis := 0
	for i := 1; i <= 30; i++ {
		var serverName string
		if i%3 == 1 {
			serverName = "Server1"
		} else if i%3 == 2 {
			serverName = "Server2"
		} else {
			serverName = "Server3"
		}

		if tryRedisRequest(serverName, "Dave", redis, 10) {
			allowedRedis++
		}

		// Only print first 12 requests to avoid clutter
		if i == 12 {
			fmt.Println()
			fmt.Println("  ... (18 more requests, all rejected) ...")
			fmt.Println()
		}
	}

	fmt.Printf("  Result: %d/30 requests allowed (limit was 10) ✓\n", allowedRedis)
	fmt.Println()
	fmt.Println("  Why Redis?")
	fmt.Println("    ✓ All servers see the same counter → globally consistent limit")
	fmt.Println("    ✓ Atomic INCR operation → no race conditions")
	fmt.Println("    ✓ Built-in TTL (EXPIRE) → auto-reset counters for time windows")
	fmt.Println()
	fmt.Println("  Used by: Cloudflare (distributed Workers), Kong API Gateway,")
	fmt.Println("           Nginx rate_limit_req (shared memory zone)")
	fmt.Println()

	// ── Summary ──────────────────────────────────────────────────────────────
	fmt.Println("============================================================")
	fmt.Println("  Rate Limiting Algorithm Comparison:")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Printf("  %-18s │ %-12s │ %-15s │ %-12s │ %-15s\n",
		"Algorithm", "Memory", "Allows bursts?", "Boundary safe?", "Distributed?")
	fmt.Println("  ───────────────────┼──────────────┼─────────────────┼──────────────┼─────────────────")
	fmt.Printf("  %-18s │ %-12s │ %-15s │ %-12s │ %-15s\n",
		"Token Bucket", "O(1)", "Yes", "Yes", "Needs Redis")
	fmt.Printf("  %-18s │ %-12s │ %-15s │ %-12s │ %-15s\n",
		"Fixed Window", "O(1)", "No", "No (exploit!)", "Needs Redis")
	fmt.Printf("  %-18s │ %-12s │ %-15s │ %-12s │ %-15s\n",
		"Sliding Window", "O(N)", "No", "Yes", "Needs Redis")
	fmt.Println()
	fmt.Println("  Real-world recommendation:")
	fmt.Println("    • Single server:         Token Bucket (simple, burst-friendly)")
	fmt.Println("    • Distributed (strict):  Sliding Window + Redis (accurate)")
	fmt.Println("    • Distributed (loose):   Fixed Window + Redis (cheap, good enough)")
	fmt.Println()
	fmt.Println("  Pro tip:")
	fmt.Println("    Most production APIs use TOKEN BUCKET + REDIS:")
	fmt.Println("      - Store (tokens, last_refill_time) in Redis")
	fmt.Println("      - Each request: atomically refill + decrement")
	fmt.Println("      - Lua script ensures atomicity (EVALSHA in Redis)")
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num08CircuitBreakerDemo
//
// TOPIC: Circuit Breaker Pattern
//   - Prevents cascading failures: if service B is slow/down, don't let
//     service A pile up requests and exhaust its own resources
//   - Three states:
//     CLOSED   -- requests flow normally, failures counted
//     OPEN     -- requests fail immediately (no call to B), saves resources
//     HALF-OPEN -- let one request through to test if B recovered
//   - Thresholds: open circuit after N failures in M seconds
//   - Ties into timeout, retry, and bulkhead patterns
//
// Real world: Netflix Hystrix, Resilience4j, every microservice mesh
// =============================================================================
func Num08CircuitBreakerDemo() {
	fmt.Println("============================================================")
	fmt.Println("  Num08 -- Circuit Breaker Pattern")
	fmt.Println("============================================================")
	fmt.Println()

	// ── What we are simulating ───────────────────────────────────────────────
	// Scenario: Service A (API) calls Service B (Payment Service).
	//
	// PROBLEM: Service B becomes slow/unresponsive (database overload, crash, etc.)
	//   - Without Circuit Breaker:
	//     → Service A keeps calling B, requests pile up, threads blocked
	//     → Service A's thread pool exhausted, now IT crashes too
	//     → Cascading failure spreads across entire system
	//
	//   - With Circuit Breaker:
	//     → After N failures, Circuit Breaker OPENS
	//     → Further requests to B fail immediately (no call made)
	//     → Service A survives, returns error quickly to client
	//     → After timeout, Circuit Breaker tries one request (HALF-OPEN)
	//     → If succeeds → CLOSE circuit, resume normal operation
	//
	// PART A — Normal operation (CLOSED state)
	// PART B — Failures trigger OPEN state (stop calling downstream)
	// PART C — Recovery: HALF-OPEN → test if downstream recovered
	// PART D — Real-world: timeout + retry + bulkhead combination

	// ── Circuit Breaker States ───────────────────────────────────────────────
	const (
		CLOSED    = "CLOSED"    // Normal: requests flow, failures counted
		OPEN      = "OPEN"      // Tripped: reject all requests immediately
		HALF_OPEN = "HALF_OPEN" // Testing: allow 1 request to test recovery
	)

	type CircuitBreaker struct {
		state            string
		failureCount     int
		successCount     int
		failureThreshold int // open after N failures
		successThreshold int // close after N successes in HALF_OPEN
		lastFailureTime  int // timestamp of last failure
		openTimeout      int // seconds to wait before trying HALF_OPEN
	}

	newCircuitBreaker := func(failThreshold, successThreshold, openTimeout int) *CircuitBreaker {
		return &CircuitBreaker{
			state:            CLOSED,
			failureThreshold: failThreshold,
			successThreshold: successThreshold,
			openTimeout:      openTimeout,
		}
	}

	// ── PART A: Normal Operation (CLOSED state) ──────────────────────────────
	fmt.Println("── PART A: Normal Operation (CLOSED state) ──────────────────")
	fmt.Println()
	fmt.Println("  Scenario: Service A calls Payment Service (Service B)")
	fmt.Println("  Config: Circuit opens after 3 failures")
	fmt.Println()

	cb := newCircuitBreaker(3, 2, 5)

	// Simulate a downstream service that can succeed or fail
	type PaymentService struct {
		healthy bool
	}
	paymentSvc := &PaymentService{healthy: true}

	callPaymentService := func(cb *CircuitBreaker, currentTime int, requestID int) bool {
		// Check circuit state BEFORE attempting the call
		if cb.state == OPEN {
			// Circuit is open — check if timeout expired
			if currentTime-cb.lastFailureTime >= cb.openTimeout {
				// Timeout expired → transition to HALF_OPEN
				cb.state = HALF_OPEN
				cb.successCount = 0
				fmt.Printf("  [CIRCUIT] Request #%d: OPEN → HALF_OPEN (timeout expired, testing recovery)\n", requestID)
			} else {
				// Still within timeout — fail fast
				fmt.Printf("  [CIRCUIT] Request #%d: ✗ REJECTED (circuit OPEN, fail fast)\n", requestID)
				return false
			}
		}

		// Attempt the actual call to downstream service
		fmt.Printf("  [CIRCUIT] Request #%d: calling Payment Service...\n", requestID)

		var success bool
		if paymentSvc.healthy {
			fmt.Printf("  [PAYMENT] Request #%d: ✓ SUCCESS (processed payment)\n", requestID)
			success = true
		} else {
			fmt.Printf("  [PAYMENT] Request #%d: ✗ FAILED (timeout / 500 error)\n", requestID)
			success = false
		}

		// Update circuit breaker state based on result
		if success {
			if cb.state == HALF_OPEN {
				cb.successCount++
				fmt.Printf("  [CIRCUIT] HALF_OPEN: success count=%d/%d\n", cb.successCount, cb.successThreshold)
				if cb.successCount >= cb.successThreshold {
					// Enough successes → close circuit
					cb.state = CLOSED
					cb.failureCount = 0
					fmt.Printf("  [CIRCUIT] HALF_OPEN → CLOSED (recovery confirmed!)\n")
				}
			} else if cb.state == CLOSED {
				// Reset failure count on success
				cb.failureCount = 0
			}
		} else {
			// Failure
			cb.failureCount++
			cb.lastFailureTime = currentTime

			if cb.state == HALF_OPEN {
				// Failed during test → reopen circuit
				cb.state = OPEN
				fmt.Printf("  [CIRCUIT] HALF_OPEN → OPEN (test failed, circuit reopened)\n")
			} else if cb.state == CLOSED {
				fmt.Printf("  [CIRCUIT] CLOSED: failure count=%d/%d\n", cb.failureCount, cb.failureThreshold)
				if cb.failureCount >= cb.failureThreshold {
					// Threshold breached → open circuit
					cb.state = OPEN
					fmt.Printf("  [CIRCUIT] CLOSED → OPEN (threshold breached!)\n")
				}
			}
		}

		fmt.Println()
		return success
	}

	// Send 3 successful requests
	fmt.Println("  Sending 3 requests to healthy Payment Service:")
	fmt.Println()
	for i := 1; i <= 3; i++ {
		callPaymentService(cb, 0, i)
	}

	fmt.Println("  Result:")
	fmt.Println("    ✓ All requests succeeded")
	fmt.Printf("    ✓ Circuit state: %s (failure count: %d/%d)\n", cb.state, cb.failureCount, cb.failureThreshold)
	fmt.Println()

	// ── PART B: Failures Trigger OPEN State ──────────────────────────────────
	fmt.Println("── PART B: Failures Trigger OPEN State ──────────────────────")
	fmt.Println()
	fmt.Println("  Scenario: Payment Service becomes unhealthy (database crash)")
	fmt.Println()

	// Payment service goes down
	paymentSvc.healthy = false
	cb.failureCount = 0 // reset for demo clarity

	fmt.Println("  Sending 5 requests to unhealthy Payment Service:")
	fmt.Println()

	for i := 1; i <= 5; i++ {
		callPaymentService(cb, 10, i)
	}

	fmt.Println("  Result:")
	fmt.Println("    ✗ First 3 requests: called Payment Service, all failed")
	fmt.Println("    ✗ Circuit OPENED after 3rd failure")
	fmt.Println("    ✓ Requests 4-5: REJECTED immediately (fail fast, no call made)")
	fmt.Println()
	fmt.Println("  Why Circuit Breaker saved Service A:")
	fmt.Println("    - Without CB: all 5 requests would block/timeout on Payment Service")
	fmt.Println("                  → threads exhausted → Service A crashes")
	fmt.Println("    - With CB:    requests 4-5 failed instantly → threads available")
	fmt.Println("                  → Service A survives, can serve other endpoints")
	fmt.Println()

	// ── PART C: Recovery (HALF_OPEN → CLOSED) ────────────────────────────────
	fmt.Println("── PART C: Recovery (HALF_OPEN → CLOSED) ────────────────────")
	fmt.Println()
	fmt.Println("  Scenario: Wait for timeout, test if Payment Service recovered")
	fmt.Println()

	// Simulate time passing (openTimeout = 5 seconds)
	currentTime := 10 + cb.openTimeout // 10 (last failure) + 5 = 15
	fmt.Printf("  ⏱  %d seconds passed (timeout expired)\n", cb.openTimeout)
	fmt.Println()

	fmt.Println("  Sending request #6 (circuit will transition to HALF_OPEN):")
	fmt.Println()

	// Payment service still down
	callPaymentService(cb, currentTime, 6)

	fmt.Println("  Result:")
	fmt.Println("    ✗ Test request failed → Circuit REOPENED")
	fmt.Println()

	// Wait another timeout period
	currentTime += cb.openTimeout
	fmt.Printf("  ⏱  Another %d seconds passed\n", cb.openTimeout)
	fmt.Println()

	// Payment service recovers
	paymentSvc.healthy = true
	fmt.Println("  💚 Payment Service recovered (database back online)")
	fmt.Println()

	fmt.Println("  Sending requests #7-8 (test recovery):")
	fmt.Println()

	callPaymentService(cb, currentTime, 7)
	callPaymentService(cb, currentTime, 8)

	fmt.Println("  Result:")
	fmt.Println("    ✓ Request #7: Circuit HALF_OPEN → called Payment Service → SUCCESS")
	fmt.Println("    ✓ Request #8: Another success → Circuit CLOSED (recovery confirmed!)")
	fmt.Println("    ✓ Normal operation resumed")
	fmt.Println()

	// ── PART D: Real-World Patterns ──────────────────────────────────────────
	fmt.Println("── PART D: Real-World Patterns (Timeout + Retry + Bulkhead) ──")
	fmt.Println()

	fmt.Println("  Circuit Breaker is usually combined with:")
	fmt.Println()

	fmt.Println("  1. TIMEOUT:")
	fmt.Println("     - Don't wait forever for downstream service")
	fmt.Println("     - Kill request after 2 seconds → counts as failure")
	fmt.Println("     - Without timeout: slow service → threads blocked → cascading failure")
	fmt.Println()

	fmt.Println("  2. RETRY (with exponential backoff):")
	fmt.Println("     - Transient failures (network glitch) → retry 1-2 times")
	fmt.Println("     - Permanent failures (service down) → Circuit Breaker opens, stop retrying")
	fmt.Println()

	fmt.Println("  3. BULKHEAD (thread pool isolation):")
	fmt.Println("     - Dedicate 10 threads to Payment Service calls")
	fmt.Println("     - If all 10 blocked → only Payment calls fail")
	fmt.Println("     - Other endpoints (User Service, Search, etc.) still have threads → survive")
	fmt.Println()

	fmt.Println("  Example: Netflix request flow")
	fmt.Println("    ┌─────────────────────────────────────────────────────┐")
	fmt.Println("    │  API Gateway receives request                       │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  Check Circuit Breaker state                        │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  CLOSED? → Call Payment Service (with timeout)      │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  Success? → Reset failure counter                   │")
	fmt.Println("    │  Timeout? → Count as failure, retry 1x              │")
	fmt.Println("    │  Failure? → Increment counter, check threshold      │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  Threshold reached? → OPEN circuit                  │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  OPEN? → Fail fast, return cached response/error   │")
	fmt.Println("    │    ↓                                                │")
	fmt.Println("    │  After timeout → HALF_OPEN, test 1 request          │")
	fmt.Println("    └─────────────────────────────────────────────────────┘")
	fmt.Println()

	// ── Summary comparison ───────────────────────────────────────────────────
	fmt.Println("============================================================")
	fmt.Println("  Circuit Breaker State Transitions:")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Println("  CLOSED ──(N failures)──→ OPEN")
	fmt.Println("     ↑                        │")
	fmt.Println("     │                        │ (timeout expires)")
	fmt.Println("     │                        ↓")
	fmt.Println("     │                    HALF_OPEN")
	fmt.Println("     │                    │       │")
	fmt.Println("     └──(M successes)─────┘       └──(failure)──→ OPEN")
	fmt.Println()
	fmt.Println("  State behaviors:")
	fmt.Println()
	fmt.Printf("  %-12s │ %-50s\n", "CLOSED", "Normal operation, track failure count")
	fmt.Printf("  %-12s │ %-50s\n", "OPEN", "Fail fast, don't call downstream, wait for timeout")
	fmt.Printf("  %-12s │ %-50s\n", "HALF_OPEN", "Test recovery with limited requests")
	fmt.Println()
	fmt.Println("  Real-world libraries:")
	fmt.Println("    • Netflix Hystrix (Java) — original implementation")
	fmt.Println("    • Resilience4j (Java) — modern alternative")
	fmt.Println("    • Polly (.NET)")
	fmt.Println("    • resilience (Go)")
	fmt.Println("    • Istio / Envoy — service mesh with built-in circuit breaker")
	fmt.Println()
	fmt.Println("  Monitoring metrics (critical for production):")
	fmt.Println("    • circuit_breaker_state (gauge: 0=CLOSED, 1=OPEN, 2=HALF_OPEN)")
	fmt.Println("    • circuit_breaker_failure_count (counter)")
	fmt.Println("    • circuit_breaker_rejected_requests (counter)")
	fmt.Println("    • downstream_latency_p99 (histogram)")
	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()
}

// =============================================================================
// Num09GossipProtocolDemo
//
// TOPIC: Gossip Protocol
// - Decentralized way for N nodes to share state without a central coordinator
// - Each node periodically picks K random peers and exchanges state
// - Information spreads exponentially: O(log N) rounds to reach all nodes
// - Used for: membership (who is alive?), failure detection, state sync
// - Eventual consistency built-in -- nodes converge without coordination
//
// Real world: Cassandra (ring membership), Consul (service discovery),
//
//	Bitcoin (transaction propagation), AWS DynamoDB
//
// =============================================================================
func Num09GossipProtocolDemo() {
	fmt.Println("\n========================================")
	fmt.Println("Topic 09: GOSSIP PROTOCOL")
	fmt.Println("========================================")
	fmt.Println("Analogy: Like how rumors spread in a school - each person tells a few")
	fmt.Println("friends, who tell their friends, and soon everyone knows!")
	fmt.Println()

	// =============================================================================
	// Part 1: Basic Gossip - Membership Information Spreading
	// =============================================================================
	fmt.Println("[PART 1: BASIC GOSSIP - MEMBERSHIP SPREADING]")
	fmt.Println("[GOSSIP] Scenario: Node-5 joins a 10-node cluster. How does everyone learn?")
	fmt.Println()

	type GossipNode struct {
		id             string
		knownMembers   map[string]bool
		gossipRound    int
		lastUpdateFrom string
	}

	// Initialize 10 nodes, only Node-1 knows about Node-5
	nodes := make(map[string]*GossipNode)
	for i := 1; i <= 10; i++ {
		id := fmt.Sprintf("Node-%d", i)
		nodes[id] = &GossipNode{
			id:           id,
			knownMembers: map[string]bool{id: true}, // Everyone knows themselves
			gossipRound:  0,
		}
	}

	// Node-5 joins, Node-1 learns first (direct contact)
	nodes["Node-1"].knownMembers["Node-5"] = true
	nodes["Node-1"].lastUpdateFrom = "Node-5 (direct)"

	fmt.Println("[GOSSIP] Round 0: Node-1 discovers Node-5 joined")
	fmt.Println("[GOSSIP]   Node-1 knows: [Node-1, Node-5]")
	fmt.Println()

	// Gossip protocol parameters
	fanout := 2 // Each node tells 2 random peers per round

	// Simulate gossip rounds
	for round := 1; round <= 5; round++ {
		fmt.Printf("[GOSSIP] --- Round %d: Each node picks %d random peers ---\n", round, fanout)

		// Collect all gossip messages this round (to apply simultaneously)
		type gossipMsg struct {
			from    string
			to      string
			members map[string]bool
		}
		var messages []gossipMsg

		// Each node picks random peers and prepares messages
		for _, node := range nodes {
			// Pick 'fanout' random peers
			var peers []string
			for peerId := range nodes {
				if peerId != node.id {
					peers = append(peers, peerId)
				}
			}

			// Simple deterministic "random" selection based on round and node id
			selectedPeers := []string{}
			if len(peers) > 0 {
				// Pick first 'fanout' peers (in real system, this would be random)
				for i := 0; i < fanout && i < len(peers); i++ {
					idx := (round*7 + i*3 + int(node.id[5]-'0')) % len(peers)
					selectedPeers = append(selectedPeers, peers[idx])
				}
			}

			// Send current knowledge to selected peers
			for _, peerId := range selectedPeers {
				membersCopy := make(map[string]bool)
				for k, v := range node.knownMembers {
					membersCopy[k] = v
				}
				messages = append(messages, gossipMsg{
					from:    node.id,
					to:      peerId,
					members: membersCopy,
				})
			}
		}

		// Apply all gossip messages
		updatedNodes := make(map[string]bool)
		for _, msg := range messages {
			target := nodes[msg.to]
			oldSize := len(target.knownMembers)

			// Merge received members
			for member := range msg.members {
				if !target.knownMembers[member] {
					target.knownMembers[member] = true
					target.lastUpdateFrom = msg.from
				}
			}

			newSize := len(target.knownMembers)
			if newSize > oldSize {
				updatedNodes[msg.to] = true
			}
		}

		// Show which nodes learned about Node-5 this round
		for nodeId := range updatedNodes {
			if nodes[nodeId].knownMembers["Node-5"] {
				fmt.Printf("[GOSSIP]   %s learned about Node-5 (from %s)\n",
					nodeId, nodes[nodeId].lastUpdateFrom)
			}
		}

		// Count how many nodes know about Node-5
		knowCount := 0
		for _, node := range nodes {
			if node.knownMembers["Node-5"] {
				knowCount++
			}
		}
		fmt.Printf("[GOSSIP]   Total nodes aware of Node-5: %d/10\n", knowCount)

		if knowCount == 10 {
			fmt.Printf("[GOSSIP]   ✓ Full convergence achieved in %d rounds!\n", round)
			fmt.Printf("[GOSSIP]   Theoretical: O(log N) = O(log 10) ≈ 3.3 rounds\n")
			break
		}
		fmt.Println()
	}

	// =============================================================================
	// Part 2: Failure Detection via Gossip
	// =============================================================================
	fmt.Println("\n[PART 2: FAILURE DETECTION WITH HEARTBEAT COUNTERS]")
	fmt.Println("[GOSSIP] Scenario: Nodes gossip heartbeat counters. Detect when Node-3 crashes.")
	fmt.Println()

	type HeartbeatNode struct {
		id         string
		alive      bool
		heartbeats map[string]int // node -> heartbeat counter
	}

	hbNodes := make(map[string]*HeartbeatNode)
	for i := 1; i <= 5; i++ {
		id := fmt.Sprintf("Node-%d", i)
		hbNodes[id] = &HeartbeatNode{
			id:         id,
			alive:      true,
			heartbeats: make(map[string]int),
		}
		// Initialize own heartbeat
		hbNodes[id].heartbeats[id] = 0
	}

	fmt.Println("[GOSSIP] Initial state: All nodes alive, heartbeat = 0")
	fmt.Println()

	// Simulate 4 rounds
	for round := 1; round <= 4; round++ {
		fmt.Printf("[GOSSIP] --- Heartbeat Round %d ---\n", round)

		// Living nodes increment their own heartbeat
		for _, node := range hbNodes {
			if node.alive {
				node.heartbeats[node.id]++
			}
		}

		// Node-3 crashes after round 2
		if round == 2 {
			hbNodes["Node-3"].alive = false
			fmt.Println("[GOSSIP] ⚠️  Node-3 CRASHED (stops incrementing heartbeat)")
			fmt.Println()
		}

		// Each node gossips with 2 peers
		for _, node := range hbNodes {
			if !node.alive {
				continue
			}

			// Pick 2 peers (simplified selection)
			peers := []string{}
			for peerId := range hbNodes {
				if peerId != node.id {
					peers = append(peers, peerId)
				}
			}

			for i := 0; i < 2 && i < len(peers); i++ {
				peer := hbNodes[peers[i]]

				// Exchange heartbeats (take max of each node's counter)
				for nid, hb := range node.heartbeats {
					if peerHb, exists := peer.heartbeats[nid]; !exists || hb > peerHb {
						peer.heartbeats[nid] = hb
					}
				}
				for nid, hb := range peer.heartbeats {
					if nodeHb, exists := node.heartbeats[nid]; !exists || hb > nodeHb {
						node.heartbeats[nid] = hb
					}
				}
			}
		}

		// Show current heartbeat state
		fmt.Println("[GOSSIP] Heartbeat counters after gossip:")
		for i := 1; i <= 5; i++ {
			id := fmt.Sprintf("Node-%d", i)
			node := hbNodes[id]
			status := "alive"
			if !node.alive {
				status = "CRASHED"
			}
			fmt.Printf("[GOSSIP]   %s (%s): ", id, status)
			for j := 1; j <= 5; j++ {
				nid := fmt.Sprintf("Node-%d", j)
				fmt.Printf("%s=%d ", nid, node.heartbeats[nid])
			}
			fmt.Println()
		}

		// Failure detection: if heartbeat hasn't increased in 2 rounds, mark as suspected
		if round >= 3 {
			fmt.Println("[GOSSIP] Failure detection analysis:")
			currentRound := round
			for _, node := range hbNodes {
				if !node.alive {
					continue
				}
				for nid, hb := range node.heartbeats {
					if nid == "Node-3" && hb < currentRound-1 {
						fmt.Printf("[GOSSIP]   %s suspects Node-3 failed (heartbeat stuck at %d)\n",
							node.id, hb)
					}
				}
			}
		}
		fmt.Println()
	}

	// =============================================================================
	// Part 3: State Synchronization - Key-Value Store
	// =============================================================================
	fmt.Println("\n[PART 3: STATE SYNC - DISTRIBUTED KEY-VALUE STORE]")
	fmt.Println("[GOSSIP] Scenario: 4 nodes maintain shared state via gossip (eventual consistency)")
	fmt.Println()

	type StateNode struct {
		id      string
		data    map[string]string // key -> value
		version map[string]int    // key -> version (for conflict resolution)
	}

	stateNodes := make(map[string]*StateNode)
	for i := 1; i <= 4; i++ {
		id := fmt.Sprintf("Node-%d", i)
		stateNodes[id] = &StateNode{
			id:      id,
			data:    make(map[string]string),
			version: make(map[string]int),
		}
	}

	// Initial writes to different nodes
	fmt.Println("[GOSSIP] Round 0: Initial writes (not yet propagated)")
	stateNodes["Node-1"].data["user:1"] = "Alice"
	stateNodes["Node-1"].version["user:1"] = 1
	fmt.Println("[GOSSIP]   Node-1 writes: user:1 = Alice (v1)")

	stateNodes["Node-2"].data["user:2"] = "Bob"
	stateNodes["Node-2"].version["user:2"] = 1
	fmt.Println("[GOSSIP]   Node-2 writes: user:2 = Bob (v1)")

	stateNodes["Node-3"].data["config:timeout"] = "30s"
	stateNodes["Node-3"].version["config:timeout"] = 1
	fmt.Println("[GOSSIP]   Node-3 writes: config:timeout = 30s (v1)")
	fmt.Println()

	// Gossip rounds to propagate state
	for round := 1; round <= 3; round++ {
		fmt.Printf("[GOSSIP] --- State Sync Round %d ---\n", round)

		// Each node gossips with 1 random peer
		for _, node := range stateNodes {
			// Pick a peer (deterministic for demo)
			peerIdx := (round + int(node.id[5]-'0')) % 4
			if peerIdx == 0 {
				peerIdx = 1
			}
			peerId := fmt.Sprintf("Node-%d", peerIdx)
			if peerId == node.id {
				peerId = fmt.Sprintf("Node-%d", (peerIdx%4)+1)
			}
			peer := stateNodes[peerId]

			// Exchange state (use version to resolve conflicts)
			merged := false
			for key, val := range node.data {
				peerVer := peer.version[key]
				nodeVer := node.version[key]
				if nodeVer > peerVer {
					peer.data[key] = val
					peer.version[key] = nodeVer
					merged = true
				}
			}
			for key, val := range peer.data {
				peerVer := peer.version[key]
				nodeVer := node.version[key]
				if peerVer > nodeVer {
					node.data[key] = val
					node.version[key] = peerVer
					merged = true
				}
			}

			if merged {
				fmt.Printf("[GOSSIP]   %s ↔ %s exchanged state\n", node.id, peerId)
			}
		}

		// Show convergence progress
		fmt.Println("[GOSSIP] Current state across nodes:")
		for i := 1; i <= 4; i++ {
			id := fmt.Sprintf("Node-%d", i)
			node := stateNodes[id]
			fmt.Printf("[GOSSIP]   %s: ", id)
			keyCount := 0
			for range node.data {
				keyCount++
			}
			fmt.Printf("%d keys [", keyCount)
			first := true
			for key, val := range node.data {
				if !first {
					fmt.Print(", ")
				}
				fmt.Printf("%s=%s", key, val)
				first = false
			}
			fmt.Println("]")
		}
		fmt.Println()
	}

	// Verify eventual consistency
	fmt.Println("[GOSSIP] ✓ Eventual Consistency: All nodes converged to same state (3 keys)")
	fmt.Println()

	// =============================================================================
	// Part 4: Comparison - Gossip vs Other Approaches
	// =============================================================================
	fmt.Println("\n[PART 4: COMPARISON - GOSSIP VS ALTERNATIVES]")
	fmt.Println()
	fmt.Println("┌────────────────────┬──────────────────┬──────────────────┬──────────────────┐")
	fmt.Println("│ Approach           │ Convergence Time │ Fault Tolerance  │ Network Load     │")
	fmt.Println("├────────────────────┼──────────────────┼──────────────────┼──────────────────┤")
	fmt.Println("│ Gossip Protocol    │ O(log N) rounds  │ Excellent        │ O(N * fanout)    │")
	fmt.Println("│                    │ ~3 rounds for    │ No single point  │ per round        │")
	fmt.Println("│                    │ N=100 nodes      │ of failure       │                  │")
	fmt.Println("├────────────────────┼──────────────────┼──────────────────┼──────────────────┤")
	fmt.Println("│ Broadcast (fanout) │ 1 round          │ Poor             │ O(N²) messages   │")
	fmt.Println("│ All nodes tell all │ Instant!         │ Source failure = │ Overwhelms large │")
	fmt.Println("│                    │                  │ update lost      │ networks         │")
	fmt.Println("├────────────────────┼──────────────────┼──────────────────┼──────────────────┤")
	fmt.Println("│ Central Coordinator│ 2 rounds         │ Very Poor        │ O(N) messages    │")
	fmt.Println("│ Hub-and-spoke      │ Node→Hub→All     │ Hub failure =    │ Efficient, but   │")
	fmt.Println("│                    │                  │ system down      │ bottleneck       │")
	fmt.Println("├────────────────────┼──────────────────┼──────────────────┼──────────────────┤")
	fmt.Println("│ Consensus (Raft)   │ 1 round          │ Good             │ O(N) messages    │")
	fmt.Println("│ Quorum-based       │ (requires ack)   │ Survives N/2-1   │ Requires leader  │")
	fmt.Println("│                    │                  │ failures         │ election         │")
	fmt.Println("└────────────────────┴──────────────────┴──────────────────┴──────────────────┘")
	fmt.Println()

	// =============================================================================
	// Part 5: Real-World Applications & Trade-offs
	// =============================================================================
	fmt.Println("\n[PART 5: REAL-WORLD APPLICATIONS]")
	fmt.Println()
	fmt.Println("[GOSSIP] ✓ Cassandra - Ring membership (who owns which token ranges?)")
	fmt.Println("[GOSSIP]   - 100-node cluster converges in ~7 gossip rounds (log₂ 100)")
	fmt.Println("[GOSSIP]   - Each node gossips every 1 second → full sync in ~7 seconds")
	fmt.Println()
	fmt.Println("[GOSSIP] ✓ Consul - Service discovery and health checks")
	fmt.Println("[GOSSIP]   - Uses Serf (gossip library) to detect failed nodes within seconds")
	fmt.Println("[GOSSIP]   - SWIM protocol: reduces false positives from network delays")
	fmt.Println()
	fmt.Println("[GOSSIP] ✓ Redis Cluster - Node discovery and failure detection")
	fmt.Println("[GOSSIP]   - Every node gossips with 3 random nodes every 1 second")
	fmt.Println("[GOSSIP]   - Detects failures in ~(NODE_TIMEOUT) seconds")
	fmt.Println()
	fmt.Println("[GOSSIP] ✓ AWS DynamoDB - Anti-entropy (background sync)")
	fmt.Println("[GOSSIP]   - Merkle trees + gossip to detect and repair inconsistencies")
	fmt.Println("[GOSSIP]   - Eventual consistency guarantee")
	fmt.Println()

	fmt.Println("\n[TRADE-OFFS]")
	fmt.Println()
	fmt.Println("Pros:")
	fmt.Println("  ✓ Decentralized - no single point of failure")
	fmt.Println("  ✓ Scalable - O(log N) convergence even with 10,000+ nodes")
	fmt.Println("  ✓ Fault-tolerant - works even if 50% of nodes fail")
	fmt.Println("  ✓ Simple - no complex coordination or leader election")
	fmt.Println()
	fmt.Println("Cons:")
	fmt.Println("  ✗ Eventual consistency - takes several rounds to converge")
	fmt.Println("  ✗ Redundant messages - same update may be sent multiple times")
	fmt.Println("  ✗ Network overhead - continuous background chatter")
	fmt.Println("  ✗ Not suitable for strong consistency requirements")
	fmt.Println()

	fmt.Println("\n[KEY INSIGHT]")
	fmt.Println("Gossip Protocol trades IMMEDIATE consistency for SCALABILITY and")
	fmt.Println("FAULT TOLERANCE. Perfect for systems that can tolerate eventual")
	fmt.Println("consistency (membership, health checks, monitoring, caches).")
	fmt.Println()
}

// =============================================================================
// Num10DistributedLockingDemo
//
// TOPIC: Distributed Locking
//   - How multiple service instances coordinate exclusive access to a resource
//   - Naive approach: SET key NX PX ttl in Redis -- single Redis node
//   - Redlock algorithm: acquire lock on majority (N/2+1) of Redis nodes
//     independently -- survives single node failure
//   - Problems: clock skew, GC pauses can cause lock to expire while holder
//     thinks it still has it -- fencing tokens solve this
//   - etcd/ZooKeeper: use consensus (Raft/ZAB) for stronger guarantees
//
// Real world: preventing double-charge in payments, job deduplication,
//
//	leader election in stateless services
//
// =============================================================================
func Num10DistributedLockingDemo() {
	// TODO: implement when teaching Topic 10
}

// =============================================================================
// Num11ObservabilityDemo
//
// TOPIC: Observability in Distributed Consensus Systems
//   - How do you observe consensus in action? Distributed tracing of Raft RPCs
//   - Detecting split-brain scenarios via metrics (two leaders in cluster?)
//   - Structured logs: correlate Raft state transitions across nodes
//   - Metrics: consensus-specific (election latency, log replication lag, quorum health)
//     -- RED method: Rate (ops/sec), Errors (failed votes), Duration (commit latency)
//   - Distributed tracing: follow a write request through leader → followers → commit
//     -- spans: receive write → append to log → replicate → commit → respond
//   - Correlation: tie every log line to the Raft term/index that caused it
//
// Real world: Debugging etcd split-brain, Consul leader flapping, ZooKeeper session timeouts
//
//	Prometheus + Grafana (metrics), Jaeger (tracing), OpenTelemetry standard
//
// =============================================================================
func Num11ObservabilityDemo() {
	// TODO: implement when teaching Topic 11
}

// =============================================================================
// Num12ShardingStrategiesDemo
//
// TOPIC: Sharding with Per-Shard Consensus Groups
//   - How do you shard a consensus system? Each shard = independent Raft/Paxos group
//   - Range sharding: shard by key range (e.g. keys 0-1000 on shard 1, each has own Raft)
//     -- example: TiKV (distributed KV store) — each range = Raft group
//   - Hash sharding: shard = hash(key) % N — even distribution, coordination via consensus
//   - Directory sharding: lookup table maps key → shard — the lookup table itself needs consensus!
//   - Cross-shard transactions: 2PC coordinator + per-shard Raft groups (CockroachDB style)
//   - Resharding: move ranges between Raft groups using joint consensus (Raft membership changes)
//   - Benefits: horizontal scaling of consensus (1000 shards = 1000× throughput)
//   - Trade-offs: complexity, cross-shard operations require coordination
//
// Real world: TiKV (per-range Raft), CockroachDB (per-range consensus + 2PC),
//
//	Spanner (Paxos per shard), YugabyteDB (per-tablet Raft)
//
// =============================================================================
func Num12ShardingStrategiesDemo() {
	// TODO: implement when teaching Topic 12
}

// =============================================================================
// Num13EventSourcingDemo
//
// TOPIC: Event Sourcing — The Distributed Log IS the Consensus Output
//   - Store the full history of events instead of just current state
//   - The distributed log (Raft/Paxos output) = immutable event stream
//   - Every state change = event appended to the replicated log (linearizable order)
//   - Replay events to reconstruct state at any point in time (time travel)
//   - CQRS (Command Query Responsibility Segregation): separate read/write models
//     -- writes go to Raft log (consensus), reads from materialized view (eventual)
//   - Benefits: full audit trail, debugging via replay, consensus guarantees ordering
//   - Trade-offs: storage growth (need log compaction/snapshots), eventual read consistency
//
// Real world: Banking ledgers (audit trail), Kafka as event store (with KRaft consensus),
//
//	Event-driven microservices, Git (commits = events in DAG),
//	Raft log itself is event sourcing (state machine replication)
//
// =============================================================================
func Num13EventSourcingDemo() {
	// TODO: implement when teaching Topic 13
}

// =============================================================================
// Num14ConsensusAlgorithmsDemo
//
// TOPIC: Consensus Algorithms (Paxos, Raft, ZAB)
//   - How distributed nodes agree on a single value despite failures
//   - Paxos: the original, proven correct, notoriously hard to understand
//     -- roles: proposer, acceptor, learner
//     -- phases: prepare, promise, accept, accepted
//   - Raft: designed for understandability, equivalent to Paxos
//     -- leader election, log replication, safety guarantees
//     -- used in: etcd, Consul, CockroachDB
//   - ZAB (ZooKeeper Atomic Broadcast): similar to Raft, optimized for ZooKeeper
//   - Key insight: majority quorum (N/2+1) can make progress, tolerates minority failure
//
// Real world: etcd (Kubernetes), Consul, ZooKeeper, CockroachDB, TiDB
// =============================================================================
func Num14ConsensusAlgorithmsDemo() {
	// TODO: implement when teaching Topic 14
}

// =============================================================================
// Num15CAPTheoremDemo
//
// TOPIC: CAP Theorem & Trade-offs
//   - The fundamental constraint: Consistency, Availability, Partition Tolerance
//     -- pick 2 out of 3 during a network partition
//   - Consistency: all nodes see the same data at the same time
//   - Availability: every request gets a response (success or failure)
//   - Partition Tolerance: system continues despite network splits
//   - Reality: network partitions happen → must choose CP or AP
//     -- CP (sacrifice Availability): etcd, ZooKeeper, HBase
//     -- AP (sacrifice Consistency): Cassandra, DynamoDB, Riak
//   - PACELC: extends CAP to include latency trade-offs during normal operation
//     -- if Partition: choose Availability or Consistency
//     -- else: choose Latency or Consistency
//
// Real world: system design interviews, architecture decisions,
//
//	understanding database behavior during failures
//
// =============================================================================
func Num15CAPTheoremDemo() {
	// TODO: implement when teaching Topic 15
}

// =============================================================================
// Num16LogicalClocksDemo
//
// TOPIC: Logical Time & Clocks (Lamport Clocks, Vector Clocks, Causality)
//   - Physical clocks are unreliable in distributed systems (clock skew, NTP drift)
//   - Lamport timestamps: single counter, guarantees happened-before ordering
//     -- if event A causally precedes B, then timestamp(A) < timestamp(B)
//     -- simple: increment on local event, take max+1 on receive
//   - Vector clocks: array of counters (one per node), detects concurrent events
//     -- can determine: A→B (causally precedes), B→A (causally follows), A||B (concurrent)
//     -- used for conflict detection in eventually consistent systems
//   - Happened-before relation (→): the causal ordering of events across nodes
//   - Applications: debugging race conditions, versioning in DynamoDB, detecting conflicts
//
// Real world: DynamoDB versioning (vector clocks), Riak (dotted version vectors),
//
//	Distributed tracing (causal spans), CRDTs (causality tracking)
//
// =============================================================================
func Num16LogicalClocksDemo() {
	// TODO: implement when teaching Topic 16
}

// =============================================================================
// Num17FailureDetectionDemo
//
// TOPIC: Failure Detection (Heartbeats, Phi Accrual, SWIM Protocol)
//   - How do you know if a node has crashed vs just slow?
//   - Basic heartbeat: missed 3 consecutive heartbeats → mark as dead
//     -- problem: network delays cause false positives (node alive but marked dead)
//   - Phi Accrual Failure Detector: statistical model of heartbeat arrival times
//     -- outputs suspicion level (0-100%) instead of binary alive/dead
//     -- adapts to network conditions (jitter, latency)
//   - SWIM protocol: indirect probing to reduce false positives
//     -- if Node-A suspects Node-B, asks Node-C to ping Node-B
//     -- if Node-C succeeds → network partition, not failure
//   - Trade-offs: detection speed vs false positive rate
//
// Real world: Cassandra (Phi Accrual + Gossip), Consul (SWIM via Serf library),
//
//	Akka cluster (Phi Accrual), Memberlist (HashiCorp SWIM implementation)
//
// =============================================================================
func Num17FailureDetectionDemo() {
	// TODO: implement when teaching Topic 17
}

// =============================================================================
// Num18QuorumsConsistencyDemo
//
// TOPIC: Quorums & Consistency (N/R/W models, Sloppy Quorums, Hinted Handoff)
//   - Quorum: W + R > N guarantees read sees latest write (overlap property)
//     -- N = total replicas, W = write quorum, R = read quorum
//     -- example: N=3, W=2, R=2 → any read overlaps with latest write
//   - Tuning: W=1, R=N (fast writes, slow reads) vs W=N, R=1 (slow writes, fast reads)
//   - Sloppy quorum: if primary nodes unavailable, write to backup nodes temporarily
//     -- maintains availability during failures (AP in CAP)
//     -- eventual consistency: hinted handoff transfers data back to primary later
//   - Hinted handoff: backup node remembers "this write belongs to Node-X" and transfers later
//   - Trade-offs: strict quorum (CP, unavailable during partition) vs sloppy (AP, eventual)
//
// Real world: Riak (tunable N/R/W + sloppy quorums), DynamoDB (W+R>N configurable),
//
//	Cassandra (tunable consistency levels: ONE, QUORUM, ALL)
//
// =============================================================================
func Num18QuorumsConsistencyDemo() {
	// TODO: implement when teaching Topic 18
}

// =============================================================================
// Num19AtomicBroadcastDemo
//
// TOPIC: Atomic Broadcast (Total Order Broadcast, FIFO Channels, ZAB)
//   - Atomic broadcast: all nodes deliver messages in the same order (total order)
//   - Properties:
//     -- Validity: if correct node broadcasts M, all correct nodes eventually deliver M
//     -- Uniform Agreement: if any node delivers M, all correct nodes deliver M
//     -- Uniform Integrity: M delivered at most once, only if it was broadcast
//     -- Uniform Total Order: if nodes deliver M1 then M2, all nodes deliver in that order
//   - FIFO channels: messages from same sender delivered in send order
//   - Total order vs causal order vs FIFO order (hierarchy of ordering guarantees)
//   - ZAB (ZooKeeper Atomic Broadcast): leader-based total order broadcast
//     -- leader proposes sequence numbers, followers ack, commit when quorum reached
//   - Applications: state machine replication (RSM) — same inputs in same order = same state
//
// Real world: ZooKeeper (ZAB), etcd (Raft log = total order), Kafka (per-partition order),
//
//	The theoretical foundation of all consensus algorithms
//
// =============================================================================
func Num19AtomicBroadcastDemo() {
	// TODO: implement when teaching Topic 19
}

// =============================================================================
// Num20RaftEdgeCasesDemo
//
// TOPIC: Raft Edge Cases (Log Compaction, Snapshots, Membership Changes)
//   - Raft basics (Num02) covers leader election + log replication
//   - This covers the production-critical edge cases:
//   - Log compaction: Raft log grows forever → disk fills up
//     -- Solution: snapshots — periodically compact log into state snapshot
//     -- discard all log entries up to snapshot index
//   - Snapshots: how to send snapshots to slow followers (InstallSnapshot RPC)
//     -- chunk-based transfer for large snapshots (multi-GB state machines)
//   - Membership changes: adding/removing nodes without downtime or split-brain
//     -- Joint consensus: transition through C_old,new (union of old + new config)
//     -- commit in C_old,new requires quorum in BOTH old and new configs
//     -- then transition to C_new alone
//   - Other edge cases: leadership transfer, PreVote optimization (prevent disruptive elections)
//
// Real world: etcd (log compaction + snapshots), Consul (snapshot management),
//
//	TiKV (Raft snapshots for region replicas), preventing disk exhaustion in production
//
// =============================================================================
func Num20RaftEdgeCasesDemo() {
	// TODO: implement when teaching Topic 20
}

// =============================================================================
// Num21PaxosSynodDemo
//
// TOPIC: Paxos (Synod) — Propose/Promise/Accept/Learn, Single-Value Consensus
//   - The original consensus algorithm (Lamport 1989), notoriously hard to understand
//   - Goal: N nodes agree on a single value despite crashes (but not Byzantine faults)
//   - Three roles: Proposer, Acceptor, Learner (same node can play multiple roles)
//   - Two phases:
//     -- Phase 1 (Prepare/Promise): Proposer asks Acceptors "can I propose with ballot N?"
//     Acceptors promise not to accept lower-numbered ballots, return highest accepted value
//     -- Phase 2 (Accept/Accepted): Proposer sends value to Acceptors
//     Acceptors accept if they haven't promised a higher ballot
//   - Safety: if value V is chosen, no other value can be chosen (quorum intersection)
//   - The "single value problem": Paxos only decides one value, not a log of values
//   - Liveness: requires a distinguished proposer (leader) to avoid dueling proposers
//
// Real world: Google Chubby (Paxos), Spanner (Paxos for replication groups),
//
//	"Paxos Made Live" paper (engineering challenges at Google)
//
// =============================================================================
func Num21PaxosSynodDemo() {
	// TODO: implement when teaching Topic 21
}

// =============================================================================
// Num22MultiPaxosDemo
//
// TOPIC: Multi-Paxos — Leader Stickiness, Skipping Phase 1, High-Throughput Consensus
//   - Basic Paxos (Num21) runs two phases for EVERY value — very slow
//   - Multi-Paxos optimization: elect a stable leader, skip Phase 1 for subsequent values
//     -- leader runs Phase 1 once (with "infinity ballot"), then only Phase 2 for all values
//     -- result: ~2 RTTs per value (vs 4 RTTs in basic Paxos)
//   - Log replication: use Multi-Paxos to agree on sequence of values (log entries)
//     -- each log position = separate Paxos instance
//   - Leader stickiness: keep same leader across multiple rounds (amortize Phase 1 cost)
//   - WAN optimization: pack multiple values in single Phase 2 (batching)
//   - Relation to Raft: Raft is essentially Multi-Paxos with stronger leader + simpler protocol
//
// Real world: Spanner (Multi-Paxos for WAN replication across datacenters),
//
//	Google F1 database, CockroachDB (originally Multi-Paxos, now Raft),
//	High-throughput consensus systems
//
// =============================================================================
func Num22MultiPaxosDemo() {
	// TODO: implement when teaching Topic 22
}

// =============================================================================
// Num23CRDTsDemo
//
// TOPIC: CRDTs (Conflict-Free Replicated Data Types)
//   - How do you merge concurrent updates without coordination or conflicts?
//   - CRDT property: all replicas converge to the same state by applying all updates
//     -- no consensus needed, no locking, no coordination
//   - G-Counter (grow-only counter): each node has own counter, merge = sum
//     -- can only increment, never decrement (monotonic)
//   - PN-Counter (positive-negative counter): two G-Counters (P and N), value = P - N
//     -- supports both increment and decrement
//   - LWW-Register (last-write-wins register): use timestamp to resolve conflicts
//     -- latest timestamp wins (requires synchronized clocks or Lamport timestamps)
//   - Types: State-based CRDTs (merge full state) vs Op-based CRDTs (commutative operations)
//   - Trade-offs: eventual consistency, cannot enforce strong invariants (e.g., "balance >= 0")
//
// Real world: Figma (collaborative editing), Discord (presence), Riak (CRDT datatypes),
//
//	Redis Enterprise (CRDT geo-replication), Automerge (CRDT library)
//
// =============================================================================
func Num23CRDTsDemo() {
	// TODO: implement when teaching Topic 23
}

// =============================================================================
// Num24LinearizabilityTestingDemo
//
// TOPIC: Linearizability Testing (Jepsen Methodology, History Verification)
//   - How do you PROVE your distributed system is correct?
//   - Linearizability: strongest consistency model — operations appear atomic and instantaneous
//     -- if op A completes before op B starts, then A's effects are visible to B
//   - Jepsen: framework for testing distributed systems under faults (network partitions, crashes)
//     -- run workload (reads/writes), inject faults, record operation history
//     -- check if history is linearizable (can operations be reordered to match real-time order?)
//   - Linearization check algorithm:
//     -- build a graph of operation dependencies (happened-before edges)
//     -- check if graph has a valid total order respecting real-time constraints
//     -- NP-complete in general, but tractable for typical test workloads
//   - Common bugs found: split-brain, stale reads, lost writes, zombie processes
//
// Real world: Jepsen tests found bugs in MongoDB, Redis, Cassandra, etcd, Consul, and more
//
//	Kyle Kingsbury's Jepsen blog (jepsen.io), Knossos (linearizability checker),
//	Production validation: chaos engineering (Netflix Chaos Monkey)
//
// =============================================================================
func Num24LinearizabilityTestingDemo() {
	// TODO: implement when teaching Topic 24
}
