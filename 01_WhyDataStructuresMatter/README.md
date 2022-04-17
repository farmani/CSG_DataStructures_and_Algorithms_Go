# Chapter 1

| Data Structure | Arrays | Sets    | 
|----------------|--------|---------|
| Reading        | O(1)   | O(1)    | 
| Searching      | O(n)   | O(n)    |
| Insertion      | O(n+1) | O(2n+1) |
|  Deletion      | O(n)   | O(n)    |

## Why Data Structures Matter
- One important measure is code maintainability. Maintainability of code involves aspects such as the readability, organization, and modularity of one’s code.
- This book is about writing efficient code. Having the ability to write code that runs quickly is an important aspect of becoming a better software developer.
- The first step in writing fast code is to understand what data structures are and how different data structures can affect the speed of our code.
- In fact, even the most complex pieces of data usually break down into a bunch of numbers and strings.
- Data structures refer to how data is organized. You’re going to learn how the same data can be organized in a variety of ways.
- Depending on how you choose to organize your data, your program may run faster or slower by orders of magnitude. And if you’re building a program that needs to deal with lots of data, or a web app used by thousands of people simultaneously, the data structures you select may affect whether your software runs at all, or simply conks out because it can’t handle the load.

### The Array: The Foundational Data Structure
- *Read:* Reading refers to looking something up at a particular spot within the data structure.
- *Search:* Searching refers to looking for a particular value within a data structure.
- *Insert:* Insertion refers to adding a new value to our data structure.
- *Delete:* Deletion refers to removing a value from our data structure.

- when we measure how “fast” an operation takes, we do not refer to how fast the operation takes in terms of pure time, but instead in how many steps it takes.
- Measuring the speed of an operation in terms of time is undependable, since the time will always change depending on the hardware it is run on.
- Measuring the speed of an operation is also known as measuring its time complexity. Throughout this book, I’ll use the terms speed, time complexity, efficiency, performance, and runtime interchangeably. They all refer to the number of steps a given operation takes.

#### Reading
- A computer can read from an array in just one step. This is because the computer has the ability to jump to any particular index in the array and peer inside.
  1. A computer can jump to any memory address in one step.
  2. Whenever a computer allocates an array, it also makes note at which memory address the array begins.
- A computer can also find the value at any index by performing simple addition.
- Naturally, an operation that takes just one step is the fastest type of operation.

#### Searching
- The computer checks each cell one at a time—is known as linear search.
- For N cells in an array, linear search would take a maximum of N steps.

#### Insertion
- The efficiency of inserting a new piece of data into an array depends on where within the array you’re inserting it.
- Another fact about computers: when allocating an array, the computer always keeps track of the array’s size.
- Once the computer calculates which memory address to insert the new value into, it can do so in one step.
- Inserting a new piece of data at the beginning or in the middle of an array is a different story. In these cases, we need to shift pieces of data to make room for what we’re inserting, leading to additional steps.
- This is because when inserting at the beginning of the array, we have to move all the other values one cell to the right.
- We can say that insertion in a worst-case scenario can take N + 1 steps for an array containing N elements. This is because we need to shift all N elements over, and then finally execute the actual insertion step.

#### Deletion
- While the actual deletion of "cucumbers" technically took just one step, we now have a problem: we have an empty cell sitting smack in the middle of our array. An array is not effective when there are gaps in the middle of it, so to resolve this issue, we need to shift "dates" and "elderberries" to the left.
- We can say then, that for an array containing N elements, the maximum number of steps that deletion would take is N steps.

### Sets
- A set is a data structure that does not allow duplicate values to be contained within it.
- I’ll talk about an array-based set. This set is just like an array—it is a simple list of values. The only difference between this set and a classic array is that the set never allows duplicate values to be inserted into it.
- Sets are useful when you need to ensure that you don’t have duplicate data.
- Insertion, however, is where arrays and sets diverge. Let’s first explore inserting a value at the end of a set
- With a set, however, the computer first needs to determine that this value doesn’t already exist in this set
- In the worst-case scenario, where we’re inserting a value at the beginning of a set, the computer needs to search N cells to ensure that the set doesn’t already contain that value, another N steps to shift all the data to the right, and another final step to insert the new value. That’s a total of 2N + 1 steps. Contrast this to insertion into the beginning of a regular array, which only takes N + 1 steps.
