# Grokking Algorithms

errata: <http://adit.io/errata.html>

## Chapter 1

### Binary search

* “With each step of binary search, you cut the number of words in half until you’re left with only one word.”
* Only works when the list is in sorted order;
* log2 n
* “That is, as the number of items increases, binary search takes a little more time to run. But simple search takes a lot more time to run. So as the list of numbers gets bigger, binary search suddenly becomes a lot faster than simple search. Bob thought binary search was 15 times faster than simple search, but that’s not correct. If the list has 1 billion items, it’s more like 33 million times faster. That’s why it’s not enough to know how long an algorithm takes to run—you need to know how the running time increases as the list size increases. That’s where Big O notation comes in.”
* “don’t grow at the same rate.”
* “Big O notation tells you how fast an algorithm is. For example, suppose you have a list of size n. Simple search needs to check each element, so it will take n operations.”
* “Big O notation is about the worst-case scenario”

## Chapter 2
