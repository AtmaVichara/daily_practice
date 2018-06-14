# Change

Correctly determine the fewest number of coins to be given to a customer such
that the sum of the coins' value would equal the correct amount of change.

### Sources
This exercise was taken from the [exercism.io](http://exercism.io/exercises/ruby/change/readme)

### Status
So far only 4 tests passing.

### Reflection on Current Code: 06-12-2018
The presented solution is a greedy algorithm that takes the first occurrence of a solution, but doesn't present the right solutions for edge cases. If the first coin is able to take from the amount, the algorithm will produce a solution, but not the optimal one. For example: in the case of
```ruby
 Change.generate([1, 5, 10, 21, 25], 63)
```
the solution should return [21, 21, 21], but instead produces [1, 1, 1, 10, 25, 25], since 25 can cleanly be taken from 63.
