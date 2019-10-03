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

### Selection sort

* With linked lists, your items can be anywhere in memory.
* Linked lists are great if you’re going to read all the items one at a time: you can read one item, follow the address to the next item, and so on. But if you’re going to keep jumping around, linked lists are terrible.
* Arrays are great if you want to read random elements, because you can look up any element in your array instantly. With a linked list, the elements aren’t next to each other, so you can’t instantly calculate the position of the fifth element in memory—you have to go to the first element to get the address to the second element, then go to the second element to get the address of the third element, and so on until you get to the fifth element.
* Lists are better if you want to insert elements into the middle.
* Arrays see a lot of use because they allow random access. There are two different types of access: *random access* and *sequential access*. Sequential access means reading the elements one by one, starting at the first element. Linked lists can only do sequential access. If you want to read the 10th element of a linked list, you have to read the first 9 elements and follow the links to the 10th element. Random access means you can jump directly to the 10th element. You’ll frequently hear me say that arrays are faster at reads. This is because they provide random access. A lot of use cases require random access, so arrays are used a lot.
* Linked lists are good for inserts/deletes, and arrays are good for random access.  

## Chapter 3

### Recursion
* Loops may achieve a performance gain for your program. Recursion may achieve a performance gain for your programmer. Choose which is more important in your situation!
* Using the stack is convenient, but there’s a cost: saving all that info can take up a lot of memory. Each of those function calls takes up some memory, and when your stack is too tall, that means your computer is saving information for many function calls. At that point, you have two options:
    * You can rewrite your code to use a loop instead.
* You can use something called tail recursion. That’s an advanced recursion topic that is out of the scope of this book. It’s also only supported by some languages, not all.
* Every recursive function has two cases: the base case and the recursive case.
* A stack has two operations: push and pop.
* All function calls go onto the call stack.
* The call stack can get very large, which takes up a lot of memory.


## Chapter 4

### Quicksort

* D&C works by breaking a problem down into smaller and smaller pieces. If you’re using D&C on a list, the base case is probably an empty array or an array with one element.
* If you’re implementing quicksort, choose a random element as the pivot. The average runtime of quicksort is O(n log n)!
* The constant in Big O notation can matter sometimes. That’s why quicksort is faster than merge sort.
* The constant almost never matters for simple search versus binary search, because O(log n) is so much faster than O(n) when your list gets big.


# Chapter 5

### Hash Tables

* Hash tables are a powerful data structure because they’re so fast and they let you model data in a different way. You might soon find that you’re using them all the time:
* You can make a hash table by combining a hash function with an array.
* Collisions are bad. You need a hash function that minimizes collisions.
* Hash tables have really fast search, insert, and delete.
* Hash tables are good for modeling relationships from one item to another item.
* Once your load factor is greater than .07, it’s time to resize your hash table.
* Hash tables are used for caching data (for example, with a web server).
* Hash tables are great for catching duplicates.

# Chapter 6

### Breatdth-first search

* Breadth-first search tells you if there’s a path from A to B.
* If there’s a path, breadth-first search will find the shortest path.
* If you have a problem like “find the shortest X,” try modeling your problem as a graph, and use breadth-first search to solve.
* A directed graph has arrows, and the relationship follows the direction of the arrow (rama -> adit means “rama owes adit money”).
* Undirected graphs don’t have arrows, and the relationship goes both ways (ross - rachel means “ross dated rachel and rachel dated ross”).
* Queues are FIFO (First In, First Out).
* Stacks are LIFO (Last In, First Out).
* You need to check people in the order they were added to the search list, so the search list needs to be a queue. Otherwise, you won’t get the shortest path.
* Once you check someone, make sure you don’t check them again. Otherwise, you might end up in an infinite loop.

# Chapter 7

## Dijkstra's algorithm

* A graph with weights is called a weighted graph. A graph without weights is called an unweighted graph.
* To calculate the shortest path in an unweighted graph, use breadth-first search. To calculate the shortest path in a weighted graph, use Dijkstra’s algorithm. Graphs can also have cycles. A cycle looks like this.
* Dijkstra’s algorithm only works with directed acyclic graphs, called DAGs for short.
* I hope this example showed you that the shortest path doesn’t have to be about physical distance. It can be *about minimizing something*. In this case, Rama wanted to minimize the amount of money he spent. Thanks, Dijkstra!
* You can’t use Dijkstra’s algorithm if you have negative-weight edges.
* If you want to find the shortest path in a graph that has negative-weight edges, there’s an algorithm for that! It’s called the Bellman-Ford algorithm.

# Chapter 8

## Greedy algorithms

* Usually there’s a very small difference between a problem that’s easy to solve and an NP-complete problem. For example, in the previous chapters, I talked a lot about shortest paths. You know how to calculate the shortest way to get from point A to point B.
* But if you want to find the shortest path that connects several points, that’s the traveling-salesperson problem, which is NP-complete. The short answer: there’s no easy way to tell if the problem you’re working on is NP-complete. Here are some giveaways:
* Your algorithm runs quickly with a handful of items but really slows down with more items.
* All combinations of X” usually point to an NP-complete problem.
* Do you have to calculate “every possible version” of X because you can’t break it down into smaller sub-problems? Might be NP-complete.
* If your problem involves a sequence (such as a sequence of cities, like traveling salesperson), and it’s hard to solve, it might be NP-complete.
* If your problem involves a set (like a set of radio stations) and it’s hard to solve, it might be NP-complete.