# Goroutines and Channels 
2 styles of concurrent programming:
- communicating sequential processes
- shared memory multithreading

## Goroutines 
In Go, each concurrently executing activity is called a *goroutine*. 
In a *concurrent* program with two or more goroutines, calls to both functions can be active at the same time. 

## Channels
Channels are connections between goroutines. A *channel* is a communication mechanism that lets one goroutine send values to another goroutine. 

Each channel is a conduit for values of a particular type - called the channel's *element type*. 

creating a channel--> ch := make(chan int)

Channels can be compared using == and is true if both are references to the same channel data structure. 

2 communications:
- send ( ch <- x)
- receive (x = <- ch or <- ch)

Close - sets a flag indicating that no more values will be sent on this channel - subsequent attempts to send will panic. 

## Unbuffered Channels
Communication over an unbufferred channel causes the sending and receiving goroutines to *synchronize*. 

When a value is sent on an unbufferred channel, the reciept of the value happens before the reawakening of hte sending routine. 

2 important aspects to messages sent over channels:
- each message has a value
- existence/timing 

## Pipelines 
Pipelines are channels that are used to connect goroutines together so that the output of one is the input to another. 

It is only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent. A channel that is the garbage collector determines to be unreachable will have its resources reclaimed whether or not it is closed. 

## Unidirectional Channel Types 
The Go type system provides *unidirectional channel tpyes* that expose only one or the other of the send and receive operations. 

The type chan<-int is a send-only channel of int.
<-chan int is a receive-only channel of int. 

Conversions from bidirectional to unidirectional channel types are permitted in any assignment. However, there is no going back. Once you have a value of a unidirectional type such as chan<-int, there is no way to obtain from it a value of type chan int that referes to teh same channel structure. 

## Buffered Channels 
A bufferred channel has a queue of elements. The queue's max size is determined when created, by the capcity arg to make. 

Leaked goroutines are not automatically collected - it's important to make sure that goroutines termianted themselves when no longer needed. 

Unbuffered channels - give stronger synchronization guarantees because every send operation is synchronized with its corresponding receive.
With buffered channels these operations are decoupled. 

Failure to allocate sufficient buffer capacity would cause the program to deadlock. 
## Looping in Parallel 

## Multiplexing with select 

## Cancellation 

