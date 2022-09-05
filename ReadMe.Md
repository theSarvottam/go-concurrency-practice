Process - 
Thread
CK10 Problem
Context Switching

Why Concurrency is hard ?
    sharing of memory
    try to access same area of memory
    data race condition and outcome will be undeterministic

    solution
        memory acess synchroization
        putting a lock
            lock force execute sequencilly
            reduce parally
            Deadlock
Goroutines
    how go implement concurrency
        CSP (Communicating sequential Processes) - 3 core ideas
            (1) Each process built for sequential execution
            (2) Data is communicated between processes (No shared memory)
            (3) Scale by adding more of same.
    GO Concurrency tool set
        goroutine
        channel
        select
        sync package
    GOroutine
        main go routine
    sync wait group
    Closures
        moves local variable of a function to heap from Stack, if used in diff go routine.  
        GO routine works on current value of variable, if we want to use specific vlue then pass as an argument to go routine function.
        eg : for loop variable.
    Go scheduler
        part of go runtimme
            M:N scheduler
            run in user space
        Work Stealing
        Channel
            Normal chanel (Unbuffer channel)
            Buffer Channel
            range over the channel
            Channel Direction - (to increase type safty of Program)
                Only send channel
                Only reciece channel

