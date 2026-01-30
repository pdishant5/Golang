# Go Language Learning – Garbage Collection & Concurrency Basics

Today, I focused on gaining an **overview understanding of Go’s garbage collector** and **concurrency fundamentals**, including core primitives and commonly used concurrency patterns.

---

## 📌 Topics Covered

### 1. Garbage Collector (GC) in Go

I explored how Go manages memory automatically using its garbage collection mechanisms.

- Purpose of garbage collection in Go  
- Overview of the classic **mark-and-sweep** garbage collection algorithm  
- Understanding limitations of stop-the-world GC approaches  
- Introduction to Go’s modern **Green Tea (concurrent mark-sweep)** GC  
- High-level idea of concurrent marking and sweeping  
- How Go reduces pause times while maintaining memory safety  

---

### 2. Concurrency Basics in Go

I started learning Go’s concurrency model and its fundamental building blocks.

- Understanding concurrency vs parallelism (conceptual overview)  
- Goroutines as lightweight concurrent execution units  
- How goroutines differ from OS threads  
- Basic idea of goroutine scheduling by the Go runtime  

---

### 3. Channels and Communication

I explored channels as Go’s primary mechanism for communication between goroutines.

- Purpose of channels in concurrent programs  
- Sending and receiving data through channels  
- Blocking behavior of channel operations  
- Conceptual difference between buffered and unbuffered channels  
- Channels as a safer alternative to shared-memory concurrency  

---

### 4. `select` Statement and Concurrency Patterns

I studied the `select` statement and common patterns built on top of it.

- Purpose of `select` for handling multiple channel operations  
- Non-deterministic selection among ready channels  
- Using `select` to avoid blocking  
- Overview of the **for-select** concurrency pattern  
- Overview of the **done channel** pattern for signaling cancellation  

---

## 🧠 Key Learnings

- Go’s garbage collector evolved to minimize pause times and improve performance  
- Modern Go GC works concurrently with application code  
- Goroutines enable efficient concurrency without heavy OS threads  
- Channels promote safe communication between concurrent tasks  
- `select` enables coordination across multiple concurrent operations  

---
