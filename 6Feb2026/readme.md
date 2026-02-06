# Go Language Learning – Atomic Operations, Context & HTTP Basics

Today, I focused on understanding **low-level synchronization using atomic operations**, **context-based cancellation and request scoping**, and began exploring the **`net/http` package**, with emphasis on internal behavior and real-world usage.

---

## 📌 Topics Covered

### 1. `sync/atomic` Package

I studied atomic operations as Go’s lowest-level synchronization primitives and explored how they interact with the CPU and Go memory model.

- Meaning of atomic operations and why they are needed  
- Atomicity at CPU and memory-model level  
- Lock-free synchronization using atomic primitives  
- Old atomic API vs new typed atomic API (`atomic.Int64`, etc.)  
- Core atomic operations:
  - Load
  - Store
  - Add
  - Swap
  - Compare-And-Swap (CAS)  
- Memory ordering and happens-before guarantees  
- Use of atomics for counters, flags, and simple shared state  
- Limitations and pitfalls of atomic operations  

#### Key Takeaways
- Atomic operations are fast and lock-free but low-level  
- Atomics are suitable only for simple shared variables  
- Mixing atomic and non-atomic access causes data races  
- Complex invariants require mutexes, not atomics  

---

### 2. `context.Context` Package

I explored context as Go’s standard mechanism for cancellation, deadlines, and request-scoped data propagation.

- Purpose of `context` in concurrent systems  
- `context.Context` interface and its core methods  
- Root contexts: `Background` and `TODO`  
- Derived contexts:
  - `WithCancel`
  - `WithTimeout`
  - `WithDeadline`
  - `WithValue`  
- Cancellation propagation through context trees  
- `Done()` channel as a cancellation signal  
- Meaning and usage of `Err()` and `Deadline()`  
- Proper handling of timeouts and cancellation  
- Internal tree-based structure and child tracking  
- Interaction with goroutines and scheduler  

#### Key Takeaways
- Contexts are immutable and form a cancellation tree  
- Cancellation flows downward, never upward  
- Contexts must always be passed explicitly  
- Contexts prevent goroutine leaks in long-running systems  

---

### 3. `net/http` Package (Introduction)

I started exploring Go’s standard HTTP package to understand how networking integrates with concurrency and context.

- Overview of the `net/http` package  
- HTTP client vs HTTP server concepts  
- Request–response lifecycle basics  
- Use of contexts in HTTP requests  
- Importance of timeouts and cancellation in network calls  

---

## 🧠 Overall Learnings

- Atomic operations provide lock-free synchronization for simple state  
- Context is essential for safe cancellation and request scoping  
- Runtime and memory-model understanding improves concurrency correctness  
- Go’s standard library is deeply integrated with concurrency primitives  

---
