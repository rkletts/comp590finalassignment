### Go Concurrent Pipeline 

### Description:
This Go program implements a 3-stage concurrent pipeline using goroutines and buffered channels to coordinate communication. The pipeline consists of:
- Stage 1: Two producers (one makes odd numbers, the other even)
- Stage 2: Two consumers that square the numbers
- Stage 3: A final filter that only prints values that are strictly increasing

I  used channels instead of shared memory to safely pass data between stages and avoid race conditions.

### Discussion Questions:
# What argument can you give that there will be no deadlock among these processes?

There shouldn’t be any deadlocks because every channel used in the program is buffered (with a capacity of 5), and all goroutines are constantly making progress. The producers write into inCh, and the consumers are actively reading from it. Similarly, the consumers write into outCh, and the final filter is always reading from outCh. Since there’s always someone on the other end of the channel (reading or writing), the system flows smoothly.

Also, I made sure that all goroutines exit properly once their job is done, by closing the channels when needed. That way, there isn't any goroutine stuck forever waiting to send or receive.

### We get "easy" fan-out and fan-in here with Go channels, because a channel is a data structure that is not bound to any single process. How would we handle this situation in Elixir, where the mailbox semantics in actors mean that our "channel" is always attached to a single process? How do we have 2 producers creating and distributing values to 2 consumers?

In Go, multiple goroutines can write to and read from the same channel easily, which makes fan out and fan in super straightforward.

In Elixir, though, each process has its own mailbox, which makes things trickier. To do fan in, you’d probably need an intermediary process that acts like a “hub”. The producers would each send their messages to this hub, and then the hub would be in charge of distributing those messages to the 2 consumer processes.

So instead of writing to a shared channel, producers send messages to the router process using send, and the router uses send again to forward the messages to whichever consumer it decides. You’d have to handle some kind of manual load-balancing (like alternating between consumers or checking if one is busy), which Go doesn’t require thanks to its flexible channels.
