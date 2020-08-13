
# Concurrency (multitasking)

Doing a lot of things at once. 
Ex. One secretary is responsible (doing) multiple things: Answering on mails, Receiving the phone calls, doing a report for management. Secretary can do a multiple things, but only ONE AT THE TIME. => I/O bound problem
Ex. Cooking a breakfest

> Solving problems: IO-bound, CPU-bound

> I/O Bound

Will your code be "waiting" for something, such as data from a database?
- If your answer is "yes", then your work is **I/O-bound**.

> CPU bound

Will your code be performing an expensive computation?
- If you answered "yes", then your work is **CPU-bound**.

## Cooperative concurrency (multitasking) = NONE
**DOESN'T EXIST IN GOLANG**

## Pre-emptive concurrency (multitasking) = GOROUTINES

- Go routines
- Channels


# Pararellism  = MULTIPROCESSING

Doing a lot of things simultaneously (like — doing numerical computations). Ex. Five accountants are calculating yearly summary for five companies. => CPU bound problem 

> Solving problems: CPU-bound problems

Each `process` is starting their own Python interpreter -> consuming more resources than Concurrency
