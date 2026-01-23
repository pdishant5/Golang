# Go Language Learning – Pointers & Structs

Today, I focused on understanding **pointers and memory behavior**, along with **structs as core data types**, including their memory layout, value vs reference semantics, and best practices.

---

## 📌 Topics Covered

### 1. Pointers in Go

I studied pointers to understand how Go handles memory, references, and data modification across function boundaries.

#### Concepts Learned
- What pointers are and how they store memory addresses  
- Declaring and initializing pointers  
- Using the address-of (`&`) and dereference (`*`) operators  
- Passing pointers to functions  
- Modifying underlying data using pointers  
- Difference between pointer value and the value being pointed to  
- Nil pointers and their behavior  

#### Key Takeaways
- Go is always pass-by-value, including pointers  
- Passing a pointer allows modification of the underlying data  
- Reassigning a pointer inside a function does not affect the caller  
- Pointers help avoid copying large data structures  

---

### 2. Value vs Reference Behavior

I explored how Go handles value and reference semantics across different data types.

#### Concepts Learned
- Value behavior of basic types and structs  
- Reference-like behavior of slices, maps, and pointers  
- Copy semantics when passing values to functions  
- Shared underlying data in slices and maps  
- Common pitfalls when assuming reference behavior  

#### Key Takeaways
- Structs and arrays are value types  
- Slices and maps share underlying data structures  
- Understanding data layout is essential to avoid unintended side effects  

---

### 3. Structs in Go

I studied structs as the primary way to group related data in Go and explored their behavior in depth.

#### Concepts Learned
- Defining and using structs  
- Struct initialization techniques  
- Named vs positional field initialization  
- Zero values of struct fields  
- Copy behavior of structs  
- Passing structs by value vs by pointer  
- Pointer receivers vs value receivers in methods  

#### Key Takeaways
- Structs are value types and are copied on assignment  
- Pointer receivers should be used when modifying struct fields  
- Field naming improves readability and safety  
- Structs provide explicit and predictable data modeling  

---

### 4. Struct Memory Layout and Alignment

I explored how Go lays out struct fields in memory and how alignment and padding affect memory usage.

#### Concepts Learned
- Memory alignment rules in Go  
- Automatic padding added by the compiler  
- Effect of field order on struct size  
- Measuring struct size using `unsafe.Sizeof`  
- Inspecting field offsets using `unsafe.Offsetof`  

#### Key Takeaways
- Field ordering directly impacts memory footprint  
- Proper ordering can significantly reduce memory usage  
- Memory layout awareness is important for performance-critical code  

---

## 🧠 Overall Learnings

- Go emphasizes explicit control over memory and data behavior  
- Understanding pointers is key to writing efficient and correct Go code  
- Structs are simple but powerful building blocks for data modeling  
- Memory layout and alignment have real performance implications  

---

## 📅 Progress Status

✔ Pointers and memory behavior  
✔ Value vs reference semantics  
✔ Structs and struct initialization  
✔ Struct memory layout and alignment  

---

## 🚀 Next Planned Topics

- Methods on structs and receiver behavior  
- Interfaces and method sets  
- Polymorphism using interfaces  
- Generics with structs and functions  

---

