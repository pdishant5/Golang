# Go Language Learning – Concurrency & Runtime Internals

Today, I focused on **advanced concurrency concepts in Go**, including **goroutine lifecycle management**, **bounded concurrency using worker pools**, and **runtime-level object reuse**, with emphasis on how these mechanisms work under the hood.

---

## 📌 Topics Covered

### 1. Zombie Goroutines

I studied goroutine lifecycle behavior and how improper design can lead to goroutine leaks.

- Goroutines have no parent–child relationship  
- Goroutines do not terminate automatically when the creator exits  
- Common causes of zombie goroutines:
  - Blocked channel operations
  - Infinite loops
  - Missing cancellation signals  
- Difference between blocked and running zombie goroutines  
- Why Go does not support forcefully killing goroutines  
- Importance of explicit cancellation using channels or `context.Context`  
- Best practices to prevent goroutine leaks  

---

### 2. Worker Pool Pattern

I learned and implemented the worker pool pattern to achieve bounded and predictable concurrency.

- Purpose of worker pools for controlled concurrency  
- Fixed number of long-lived worker goroutines  
- Job distribution using channels  
- Sequential task execution per worker  
- Parallelism achieved across multiple workers  
- Backpressure through blocking sends on job channels  
- Comparison with goroutine-per-task model  
- Graceful shutdown using channel closing and `sync.WaitGroup`  
- Clean goroutine lifecycle management  

---

### 3. `sync.Pool` – Object Reuse and Runtime Interaction

I explored how `sync.Pool` enables temporary object reuse and how it interacts with the Go runtime and garbage collector.

- Purpose of `sync.Pool` to reduce allocations and GC pressure  
- `Get` and `Put` methods and their behavior  
- Per-P (processor-local) pool design  
- Private and shared pool entries  
- Victim cache and GC-driven eviction mechanism  
- Best-effort reuse with no lifetime guarantees  
- Appropriate use cases for `sync.Pool`  
- Common pitfalls and misuse scenarios  

---

## 🧠 Key Learnings

- Goroutines must be explicitly managed to avoid leaks and unintended execution  
- Worker pools provide bounded concurrency, backpressure, and system stability  
- Goroutine-per-task model is simple but unsafe at scale  
- `sync.Pool` optimizes performance but does not guarantee object reuse  
- Understanding runtime behavior is essential for writing scalable and safe concurrent Go programs  

---
