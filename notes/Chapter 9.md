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
sync.RWMutex is a multiple readers, single writer lock

Note: Only use an RWMutex when most of the goroutines that aquire the lock are readers and the lock is under contention (goroutines routinely have to wait to aquire it)

Downside: slower than a regular mutex for uncontended locks

## Memory Synchronization 
Writes to memory are buffered within each processor and flushed out to main memory only when necessary. Synchronization primitives (channel coms and mutex ops) cause the processor to flush out and commit all its accumulated writes.

Concurrency problems can be avoided by the consistent use of simple, established patterns. Where possible, confine variables to a single goroutine; for all other variables - use mutual exclusion. 

## Lazy Initialization: sync.Once
The tactic of delaying the creation of an object, the calculation of a value, or some other expensive process until the first time it is needed. 

sync.Once can be used for one-time initializations.
Conceptually: a Once consists of a mutex and a boolean variable recording if initialization has taken place or not. Mutex guards both boolean and client's data structures. 

sync.Once can be used to avoid sharing variables with other goroutines until they have been properly constructed. 

## The Race Detector 
Just adding the -race flag to go build/go run or go test causes the compiler to build a mod version of ur app/test with additiaonl instrumentation that records all accesses to shared variables that occurred during execution. 

The rac detector studies the stream of events searching for cases in which 1 goroutine reads/writes a shared variable that was most recently written by a diff goroutine witout synch operations then prints a report that identifies the problem areas. 

Con: can only detect data races that were actually executed. Cant catch that none will occur. 


## Growable Stacks
Threads has a fixed-size block off memory for its stack. Changing the fixed size can improve space efficiency and allow more threads to be created OR it can enable more deeply recursive functions - cannot do both.

Goroutines start with a small stack and can grow/shrink as needed. 

## Goroutine Scheduling
Threads are scheduled by the OS kernel. Slow.

Go runtime contains own scheduler that uses m:n scheduling because it multiplexes m goroutines on n OS threads. It is not invoked periodicly but implicitly by certain Go lang constructs. Rescheduling a goroutine is much cheaper than a thread. 


## GOMAXPROCS
GOMAXPROCS determines how many OS threads may be actively executing Go code simultaenously. 

## Goroutines Have No Identity 
A current thread has a distinct identity that can be easily obtained (integer/pointer)

Goroutines do not by design so as not to be abused like thread-local storage tends to be. 