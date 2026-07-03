In Go, channels can be buffered or unbuffered, and the difference lies in how they handle sending and receiving values.
---
🔑 Unbuffered Channels
Definition: Created without specifying capacity (make(chan int)).
Behavior:
A send (ch &lt;- value) blocks until another goroutine is ready to receive (&lt;-ch).
A receive (&lt;-ch) blocks until another goroutine sends.
Use Case: Ensures strict synchronization between sender and receiver.
Example:
ch := make(chan int) // unbuffered
go func() {
    ch <- 42 // blocks until main receives
}()
fmt.Println(<-ch) // prints 42

---
🔑 Buffered Channels
Definition: Created with a capacity (make(chan int, 3)).
Behavior:
A send does not block until the buffer is full.
A receive does not block until the buffer is empty.
Use Case: Allows decoupling of sender and receiver, useful for pipelines or when producers are faster than consumers.
Example:
ch := make(chan int, 2) // buffered with capacity 2
ch <- 1
ch <- 2
// ch <- 3 would block because buffer is full
fmt.Println(<-ch) // prints 1
fmt.Println(<-ch) // prints 2

---
Comparison Table
Feature	Unbuffered Channel	Buffered Channel
Creation	make(chan T)	make(chan T, capacity)
Send behavior	Blocks until receiver is ready	Blocks only if buffer is full
Receive behavior	Blocks until sender is ready	Blocks only if buffer is empty
Synchronization	Strict (forces rendezvous)	Looser (allows temporary storage)
Typical use case	Coordination, signaling	Pipelines, producer-consumer patterns
---
👉 In short:
Unbuffered channels = direct handoff, strict synchronization.
Buffered channels = temporary storage, more flexibility.
