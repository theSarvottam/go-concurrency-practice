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
