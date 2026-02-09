# Go Language Learning – Serialization, Testing & HTTP Backend Fundamentals

Today, I focused on understanding **data serialization using marshaling and unmarshaling**, explored Go’s **testing ecosystem in depth** (unit tests, benchmarks, lifecycle control, and fuzz testing), and studied the **`net/http` package** thoroughly by implementing a small, production-style HTTP backend.

---

## 📌 Topics Covered

### 1. Marshaling and Unmarshaling

I studied how Go converts in-memory data structures into external representations (and back), which is essential for APIs, persistence, and distributed systems.

- Meaning of marshaling and unmarshaling  
- Why serialization is required in backend systems  
- `json.Marshal` and `json.Unmarshal`  
- Streaming JSON handling with:
  - `json.NewEncoder`
  - `json.NewDecoder`  
- Role of reflection in encoding/decoding  
- Use of struct tags (`json:"field,omitempty"`)  
- Handling missing fields and zero values  
- Error handling during serialization  
- Performance and memory considerations  

#### Key Takeaways
- Marshaling converts Go values into transferable formats  
- Unmarshaling reconstructs Go values from serialized data  
- Streaming encoders are preferred in HTTP handlers  
- Struct tags define the external data contract  

---

### 2. Go Testing Package (`testing`)

I explored Go’s built-in testing framework in detail, covering correctness, performance, lifecycle control, and robustness testing.

---

#### 2.1 Unit Testing with `testing.T`

- Structure and conventions of Go test files (`*_test.go`)  
- Writing unit tests using `testing.T`  
- Difference between `t.Error` and `t.Fatal`  
- Subtests using `t.Run`  
- Parallel test execution and safety considerations  

##### Table-Driven Testing
- Motivation for table-driven tests  
- Separation of test data from test logic  
- Testing multiple scenarios and edge cases  
- Readability and scalability benefits  

**Key Takeaways**
- Unit tests validate correctness of isolated logic  
- Table-driven tests are idiomatic and scalable  
- Go testing relies on conventions over configuration  

---

#### 2.2 Benchmark Testing with `testing.B`

- Purpose of benchmark tests  
- Adaptive iteration control using `b.N`  
- Writing micro-benchmarks  
- Table-driven benchmarks for different input sizes  
- Measuring memory allocations with `-benchmem`  
- Understanding benchmark output (`ns/op`, `B/op`, `allocs/op`)  

**Key Takeaways**
- Benchmarks are for comparison, not absolute speed  
- Memory allocations are critical performance indicators  
- Setup code must be excluded using `b.ResetTimer`  

---

#### 2.3 Test Lifecycle Control with `testing.M`

- Purpose of `TestMain`  
- One-time setup and teardown at package level  
- Role of `m.Run()` and exit codes  
- Importance of `os.Exit(m.Run())`  
- Common pitfalls (using `defer` with `os.Exit`)  

**Key Takeaways**
- `TestMain` gives full control over test lifecycle  
- Cleanup must occur before calling `os.Exit`  
- Useful for shared resources and environment setup  

---

#### 2.4 Context Usage with `TestMain`

- Creating a root context in `TestMain`  
- Deriving per-test contexts using `context.WithCancel`  
- Avoiding shared cancellation across tests  
- Preventing goroutine leaks in tests  
- Aligning test patterns with production shutdown behavior  

**Key Takeaways**
- Shared contexts should not be cancelled inside tests  
- Each test must use its own derived context  
- Context cancellation ensures deterministic cleanup  

---

#### 2.5 Fuzz Testing with `testing.F`

- Purpose of fuzz testing and bug discovery  
- Difference between fuzz tests and unit tests  
- Writing fuzz tests using `testing.F`  
- Seeding inputs with `f.Add`  
- Input mutation and coverage-guided fuzzing  
- Understanding fuzz test output and execution statistics  
- Saving and replaying failing inputs  

**Key Takeaways**
- Fuzz testing discovers unexpected edge cases  
- Fuzzing is automated, non-interactive, and reproducible  
- Panics and crashes are primary fuzzing targets  

---

### 3. `net/http` Package (In Depth)

I explored Go’s standard HTTP package and implemented a small backend using production-oriented design principles.

- Core abstractions:
  - `http.Handler`
  - `http.HandlerFunc`
  - `http.Request`
  - `http.ResponseWriter`  
- Request–response lifecycle  
- One-goroutine-per-request concurrency model  
- `http.Server` configuration and timeouts  
- Routing using `http.ServeMux`  
- Query parameter extraction  
- Path parameters using `{id}` syntax (Go 1.22+)  
- Middleware implementation using handler wrapping  
- Logging middleware  
- Context propagation via `r.Context()`  
- Clean separation of router, handlers, and middleware  

#### Folder Structure Used
- `cmd/` for application entry point  
- `internal/` for application-specific logic  
- Clear separation of routing, handlers, and middleware  

---

## 🧠 Overall Learnings

- Serialization is fundamental to API communication  
- Go’s testing tools cover correctness, performance, lifecycle, and robustness  
- Proper use of `TestMain` and context prevents flaky tests and goroutine leaks  
- Fuzz testing helps uncover edge cases missed by traditional tests  
- `net/http` provides a powerful, concurrency-aware foundation for backend services  
- Clean structure and explicit patterns are key to production-ready Go code  

---
