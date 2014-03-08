# numberCruncher

A simple Go approach to solving the number cruncher pilot case for the Heuristics course from the Minor Programmeren at the UvA. Still under construction.

At the moment it is able to find solutions for all integers in the range [1:100] and display the operator path. Biggest problem with this is that it is rather hard to check, most calculators can't handle this. :(

## To Do

1. Write tests for crunch package.

2. Write a checker that takes a operator path as input, preforms those operations, and checks them against the supposed output.

3. Repair FactorialSqrt (see issue 1)


## Known Issues

1. FactorialSqrt is currently not working, or at least not working as desired. Perhaps an idea is to add a third channel which receives numbers that are too big for factorial but may be doable through Roland's method.

2. I do not completely understand how the concurrency works. In this implementation I think it fires of factorials in a seperate 'thread' but blocks at the Sqrts. Why this is better than making both non-blocking is something I do not understand. T_T

3. A lot of code might be wrong because I only just picked up Go.
