# Concurrency with Shared Variables

## Race Conditions
A function is *concurrency-safe* if it continues to work correctly even when called concurrently. 

Avoid concurrent access by confining variables to a single goroutine or by maintaing a higher-level invariant of mutual exclusion. 

Reasons why a function might not work when called concurrently:
- deadlock
- livelock
- resource starvation

A *race condition* is a situation in which the program does not give the correct result for some interleavings of the operations of multiple goroutines. 

Even when a variable cannot be confined toa  single goroutine for its entire lifetime, confinement may still be a solution to the problem of concurrent access. 

4 ways to make variable access safe in shared-memory concurrency:

- *Confinement*: Don’t share the variable between threads. 
- *Immutability*: Make the shared data immutable. 
- *Threadsafe data type*: Encapsulate the shared data in an existing threadsafe data type that does the coordination for you.
- *Synchronization*: Use synchronization to keep the threads from accessing the variable at the same time. 

## Mutual Exclusion: sync.Mutex
A *binary semaphore* is a sempahore that counts only to 1.

Common pattern of mutual exclusion
```
var (
    sem = make(chan struct {}, 1)
    balance int
)

func Deposit(amount int) {
    sema <- struct{}{} //get token
    balance = balance + amount
    <- sema //release the token
}

func Balance() int {
    sema <- struct{}{} // get token
    b := balance
    <- sema // release token
    return b
}
```

Mutex updated version:

```
import "sync"

var (
    mu  sync.Mutex // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}
func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}

```
Variables guarded by mutexes are declared immediately after the declaration of the mutex itself by convention. 

Note: mutex locks are not re-entrant. You can lock a mutext that is already locked.

Encapsulation helps us maintain data structure invariants. 

When you use a mutex, make sure that both it and the variables guarded are not exported. 

## Read/Write Mutexes: sync.RWMutex

## Memory Synchronization 

## Lazy Initialization: sync.Once

## The Race Detector 

## Example: Concurrent Non-Blocking Cache

## Goroutines and Threads

## Growable Stacks

## Goroutine Scheduling

## GOMAXPROCS

## Goroutines Have No Identity 