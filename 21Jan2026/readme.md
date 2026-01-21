# Go Language Learning – Control Flow & Functions

Today, I focused on understanding **control flow constructs** and **functions**, including their behavior, syntax, edge cases, and best practices.

---

## 📌 Topics Covered

### 1. Loops in Go

I explored Go’s looping mechanism.

#### Concepts Learned
- Go uses a **single looping construct**: `for`
- Different forms of `for`:
  - Traditional `for` loop
  - `for` as a while loop
  - Infinite loops
  - Range-based loops
- Loop initialization, condition, and post statements
- Loop variable scope and lifetime
- `break`, `continue`, and labeled loops

#### Key Takeaways
- Go intentionally avoids `while` and `do-while`
- Loop variables are reused in each iteration

---

### 2. Switch-Case Statements

I studied Go’s `switch` statement in detail, including internal behavior and edge cases.

#### Concepts Learned
- Basic `switch` with expressions
- `switch` without expressions (acts like `if-else`)
- Multiple values in a single case
- Default case behavior
- Case matching order (top to bottom)
- Implicit `break` (no fallthrough by default)
- Explicit `fallthrough` and its rules
- Type switch basics

#### Important Observations
- The switch expression is evaluated **only once**
- `fallthrough` skips condition checking of the next case
- `fallthrough` must be the last statement in a case
- Type switches do not support `fallthrough`

---

### 3. Functions in Go (Covered Till Core Concepts)

I explored Go functions focusing on structure, arguments, and execution behavior.

#### Concepts Learned
- Basic function declaration and syntax
- Function signatures and types
- Multiple return values
- Named return values
- Parameter passing semantics (Go is **pass-by-value**)
- Argument evaluation order (left to right)
- Variadic functions (`...`)
- Passing slices to variadic parameters
- Functions as first-class values
- Anonymous functions
- Basic higher-order functions

#### Key Takeaways
- All arguments in Go are passed by value
- Slices, maps, and pointers appear mutable due to shared underlying data
- Variadic parameters are treated as slices inside functions
- Function arguments are evaluated before function execution
- Go avoids default parameters and function overloading for clarity

---

## 🧠 Overall Learnings

- Go emphasizes **simplicity, clarity, and explicit behavior**
- Control flow constructs are minimal but powerful
- Many common pitfalls come from misunderstanding:
  - Variable scope
  - Value vs reference semantics
  - Implicit behaviors (which Go intentionally avoids)

---

## 📅 Progress Status

✔ Loops
✔ Switch-case statements
✔ Functions (core concepts up to arguments and function values)

More advanced topics like closures, defer, panic/recover, and function internals will be covered next.

---

## 🚀 Next Planned Topics

- Closures and variable capture
- `defer` and execution order
- Error handling patterns
- Panic and recover
- Function internals and performance considerations

---

