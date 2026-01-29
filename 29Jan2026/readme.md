# Go Language Learning – Map Implementations & Concurrency

Today, I focused on understanding **hash table–based map implementations in Go**, their **runtime internals**, and **concurrent map designs**, along with performance trade-offs and real-world use cases.

---

## 📌 Topics Covered

### 1. Hash Table Implementation in Go (Classic Maps)

- Fundamentals of hash tables and collision handling  
- Classic Go map design using buckets and overflow buckets  
- Role of `hmap` and `bmap` runtime structures  
- Bucket layout, tophash metadata, and key/value storage  
- Lookup, insertion, deletion, and incremental rehashing  
- Load factor thresholds and map growth behavior  

---

### 2. Swiss Table Implementation (Modern Go Maps)

- Swiss-table-inspired map design (Go 1.24+)  
- Control bytes and hash fingerprints for fast lookups  
- Probing-based collision handling without overflow buckets  
- Improved cache locality and reduced pointer chasing  
- Performance and memory improvements over classic maps  
- Comparison with bucket-based hash tables  

---

### 3. `sync.Map` Internals

- Purpose and design goals of `sync.Map`  
- Read-only and dirty map separation  
- Lock-free reads using atomic operations  
- Mutex-protected writes and delete operations  
- Promotion of dirty map to read-only map  
- Suitable workloads and limitations of `sync.Map`  

---

### 4. Concurrent Maps (CMap / Sharded Maps)

- Sharded map design for concurrent access  
- Hash-based shard selection  
- Per-shard locking using mutexes or RWMutex  
- Read, write, and delete operations across shards  
- Strong consistency guarantees  
- Comparison with `sync.Map` and `map + mutex`  

---

## 🧠 Key Learnings

- Go maps are runtime-managed hash tables with evolving internal designs  
- Swiss tables improve performance via cache-friendly data layouts  
- `sync.Map` is optimized for read-heavy concurrent workloads  
- Sharded concurrent maps scale well for balanced read/write access  
- Choosing the right map implementation depends on access patterns  

---
