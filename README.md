# numberCruncher

A simple Go approach to solving the number cruncher pilot case for the Heuristics course from the Minor Programmeren at the UvA. Still under construction.

At the moment it is able to find at least 50 solutions (52 if you include 1 and 2) and display the operator path.

## To Do

1. Implement the Sqrt in such a way that it uses the SqrtBig if the number is too big, but otherwise use Sqrt. This will require the use of two channels. Select statement might be helpful here.

2. Clear up the formatting for the operator paths, these are not very clear at the moment.

## Known Issues

1. Currently the BigSqrt function automatically floors the result. This might turn out to be undesirable.

2. I do not completely understand how the concurrency works. In this implementation I think it fires of factorials in a seperate 'thread' but blocks at the Sqrts. Why this is better than making both non-blocking is something I do not understand. T_T

3. A lot of code might be wrong because I only just picked up Go.
