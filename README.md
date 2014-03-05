# numberCruncher

A simple Go approach to solving the number cruncher pilot case for the Heuristics course from the Minor Programmeren at the UvA. Still under construction.

At the moment it is able to find solutions for all integers in the range [1:100] and display the operator path. Biggest problem with this is that it is rather hard to check, most calculators can't handle this. :(

## To Do

1. Implement the Sqrt in such a way that it uses the SqrtBig if the number is too big, but otherwise use Sqrt.

2. Repair FactorialSqrt (see issue 2)

## Known Issues

0. (Not really an issue) Floor is not displayed in the operator path. This was a conscious decision but it might be useful to put it in after all.

1. FactorialSqrt is currently not working, or at least not working as desired. Perhaps an idea is to add a third channel which receives numbers that are too big for factorial but may be doable through Roland's method.

2. I do not completely understand how the concurrency works. In this implementation I think it fires of factorials in a seperate 'thread' but blocks at the Sqrts. Why this is better than making both non-blocking is something I do not understand. T_T

3. A lot of code might be wrong because I only just picked up Go.
