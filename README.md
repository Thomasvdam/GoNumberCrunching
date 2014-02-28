# numberCruncher

A simple Go approach to solving the number cruncher pilot case for the Heuristics course from the Minor Programmeren at the UvA. Still under construction.

## Known Issues

1. Currently the BigSqrt function automatically floors the result. This might turn out to be undesirable.

2. The base is there for a concurrent approach, but proper set management is not yet in place.

3. A lot of code might be wrong because I only just picked up Go.

4. It's just not complete yet, mainly due to the lack of set management.

5. It needs 2 channels, one for the big.Ints and one for the smallers numbers that can be represented as floats, since math.Sqrt doesn't floor automatically. This means that I also need to implement a function that converts new whole floats to big.Ints and passes those to the relevant channel.