# README.md

This repository serves for code examples and notes from book [A Common-Sense Guide to Data Structures and Algorithms](https://www.amazon.com/Common-Sense-Guide-Data-Structures-Algorithms/dp/1680502441).

## Contents
1. [Why Data Structures Matter](01_WhyDataStructuresMatter/README.md)
2. [Why Algorithms Matter](02_WhyAlgorithmsMatter/README.md)
3. [O Yes! Big O Notation](03_BigOh/README.md)
4. [Speeding Up Your Code with Big O](04_SpeedingUpYourCodeWithBigOh/README.md)
5. [Optimizing Code with and Without Big O](05_SpeedingUpYourCodeWithOrWithoutBigOh/README.md)
6. [Optimizing for Optimistic Scenarios](06_OptimizingForOptimisticScenarios/README.md)
7. [Big O in Everyday Code](07_BigOhInEverydayCode/README.md)
8. [Blazing Fast Lookup with Hash Tables](08_BlazingFastLookupWithHashTables/README.md)
9. [Crafting Elegant Code with Stacks and Queues Stacks](09_CraftingElegantCodeWithStacksAndQueues/README.md)
10. [Recursively Recurse with Recursion](10_RecursivelyRecurseWithRecursion/README.md)
11. [Learning to Write in Recursive](11_LearningToWriteInRecursive/README.md)
12. [Dynamic Programming](12_DynamicProgramming/README.md)
13. [Recursive Algorithms for Speed](13_RecursiveAlgorithmsForSpeed/README.md)
14. [Node-Based Data Structures](14_NodeBasedDataStructures/README.md)
15. [Speeding Up All the Things with Binary Search Trees](15_SpeedingUpAllTheThingsWithBST/README.md)
16. [Keeping Your Priorities Straight with Heaps](16_KeepingYourPrioritiesStraightWithHeaps/README.md)
17. [It Doesn’t Hurt to Trie](17_ItDoesntHurtToTrie/README.md)
18. [Connecting Everything with Graphs](18_ConnectingEverythingWithGraphs/README.md)
19. [Dealing with Space Constraints](19_DealingWithSpaceConstraints/README.md)
20. [Techniques for Code Optimization](20_TechniquesForCodeOptimization/README.md)

## Summary
- All in all, great book for starting with algorithms and data structures
- O(N) - constants are skipped, so even if we loop 10 times over an array, it is realistically O(10), but uses O(N) notation
- This way, we can make faster code without improving Big Oh (BubbleSort -> SelectionSort)
- Stacks and queues were introduced. Stack example was great - a very basic linting done via stacks. Queues had worse example
- First, sorts were reviewed to introduce these concepts. Done sorts are bubble, selection, insertion and quick
  - Note that quicksort is not O(log N) in space, it could fail on large numbers unless updated
- Then, hash tables were introduced. In JS, this is very obvious, as every object is basically a hash table
- Recursion showed quick sort and memoization (passing object with previous results down the line)
- Then, node based data structures came
- Node based data structures showed linked lists, which are great data structure for insertion/deletion at the end/start. They are great for queues for example
- Then, Binary Search Trees were shown. These are great if we can afford to keep our numbers sorted from start
- Heaps showed that we can work with an array as we would with a tree, but it has to be specific
  - Heaps must be a complete tree, and complete tree can be rewritten into an array
  - This allows us to be very fast
- Then, tries were shown. Yet another tree, but with multiple children. Autocomplete and autocorrect was done here
- Graphs showed that they are a superset of trees - all trees are graphs, but not all graphs are trees
- Graphs focus on relationship between data
- We've looked at weighted and unweighted graphs
  - Unweighted - followers on twitter, friends on facebook. How are data related to one another. Followers were directed (one way), friends were undirected (2-way)
  - Weighted were a different view - we added values to edges. Basically, this allowed us to create a simple algorithm for calculating plane prices between cities
- As part of graphs, we looked at interesting algorithms for traversing/finding closest node/pathfinding
  - Breadth-first search, depth-first search (Unweighted graphs)
  - Dijkstra (Weighted)
- Small chapter on space constraints was also part of this book. Here, we learned what space constraints are and how to deal with them
- Lastly, there were some techniques for code optimization. Basically showed various tasks that at first glance can be achieved in slow time, but they could be improved by
  - Finding patterns
  - Using different data structure
  - Find better algorithm

## To be done as a result of reading this book
- Implement own queue using doubly linked list, rather than currently implemented array
- Use such queue in BFS
- Review BFS/DFS/Dijkstra
- Apply in AoC2021 task