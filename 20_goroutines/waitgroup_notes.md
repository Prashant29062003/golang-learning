**TL;DR Summary:**

Your code is accurate, production-ready, and idiomatic Go. You cleanly eliminated the `time.Sleep` anti-pattern by using `sync.WaitGroup`, and passing `w *sync.WaitGroup` as a **pointer** correctly prevents severe concurrency bugs.

---

## Why This Implementation Is Excellent

* **Passing `sync.WaitGroup` by Pointer (`*sync.WaitGroup`):**
* In Go, a `sync.WaitGroup` contains internal state (an atomic counter and a mutex).
* If you pass `sync.WaitGroup` by value (`w sync.WaitGroup`), Go creates a duplicate copy of the struct. Calling `w.Done()` on a copy would decrement the copy's counter, while `main()`'s `wg.Wait()` would block forever waiting on the original, leading to a **deadlock**.
* Passing `&wg` ensures all goroutines mutate the exact same counter instance in memory.


* **Using `defer w.Done()`:**
* Using `defer` guarantees that `w.Done()` is invoked when `task` returns, even if the function encounters an early `return` or a runtime panic.


* **Deterministic & Resource-Efficient Termination:**
* `wg.Wait()` blocks `main()` only for as long as needed. The moment the 1,000th goroutine completes, `main()` resumes immediately without wasting idle time.



---

## Output

Running this program prints all 1,000 task numbers (in non-deterministic scheduler order) followed by:

```text
Doing task:  ...
All tasks completed!

```