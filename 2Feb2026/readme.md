# Go Language Learning – Scheduler, Channels & Synchronization

Today, I focused on understanding **Go runtime scheduling**, **channel internals**, and **synchronization primitives**, with emphasis on how concurrency is managed under the hood.

---

## 📌 Topics Covered

### 1. Go Runtime Scheduler (GMP Model)

I studied how Go schedules goroutines efficiently on OS threads.

- G–M–P (Goroutine–Machine–Processor) scheduling model  
- Role of `GOMAXPROCS` in controlling parallelism  
- Relationship between goroutines, OS threads, and processors  
- Goroutine creation and lightweight execution model  
- How blocking operations release threads and keep CPUs busy  
- Context switching between goroutines vs OS threads  

---

### 2. Channels – Internal Working

I explored how channels are implemented and managed by the Go runtime.

- Channels as runtime-managed synchronization structures  
- Internal channel structure (`hchan`) overview  
- Send and receive operation flow  
- Blocking and parking of goroutines on channels  
- Sender and receiver wait queues  
- Interaction between channels and the scheduler  

---

### 3. Buffered vs Unbuffered Channels

I understood behavioral and performance differences between channel types.

- Unbuffered channels as direct goroutine handoff  
- Buffered channels with internal circular buffer  
- Blocking behavior in send and receive operations  
- Impact of buffer size on concurrency and throughput  
- Appropriate use cases for each channel type  

---

### 4. `sync.WaitGroup` – Usage and Internals

I studied how `WaitGroup` coordinates goroutine completion.

- Purpose of `sync.WaitGroup` for goroutine synchronization  
- `Add`, `Done`, and `Wait` methods and their contracts  
- Correct usage patterns and common pitfalls  
- Internal counter and semaphore-based implementation  
- One-shot lifecycle and reuse considerations  
- Memory visibility guarantees provided by `WaitGroup`  

---

### 5. Basic Concurrency Patterns

I explored fundamental concurrency patterns using channels.

- Simple pipelining using channels  
- Coordinating goroutines with channel-based workflows  
- Combining channels with `WaitGroup` for synchronization  

---

## 🧠 Key Learnings

- Go scheduler efficiently multiplexes goroutines over OS threads  
- Channels block goroutines, not threads, enabling scalable concurrency  
- Buffered and unbuffered channels serve different coordination needs  
- `WaitGroup` is a lightweight and strict synchronization primitive  
- Understanding internals helps avoid subtle concurrency bugs  

---
