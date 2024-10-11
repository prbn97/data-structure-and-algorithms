> A study on data structure and algorithms by [prbn](https://github.com/prbn97/).
# A solid understanding of the relationship between data structures and algorithms.
    
**Algorithms** and **Data structures** are the tools for a programmer to create a _system_.

### First, you learn about [Big O notations.](https://www.bigocheatsheet.com/)

With this you can make good decisions about how to code, because it allows you to evaluate the efficiency of your code's algorithms, based on computational complexity.

## An efficient way of doing things...

algorithms and data structures are the tools for a programmer to create a system.

- data structures are used to organize information.
- algorithms are the instructions for processing this information.

A system can be made in "n" ways, one of them is _Object Oriented Programming_, using:

- classes / objects
- methods / inheritance

with P.O.O. we can...

1. Make data structures like _Stack and Linked List_ became classes.
2. Make Algorithms methods of these classes.
3. Make simpler or more widely used data structures and algorithms be used through inheritance.

```js
// Building a linked list to exemplify

class Node {
    constructor(value) {...}
}

class LinkedList {
    // algorithms
    constructor(value) {...}
    push(value) {...}
    unshift(value) {...}
    insert(index, value) {...}
    remove(index) {...}
    pop() {...}
    shift() {...}
}

// some of the things we can do by building linked list class.
let myLinkedList = new LinkedList(25)
myLinkedList.push(9)
myLinkedList.unshift(3)
myLinkedList.insert(1, 15)
myLinkedList.remove(1)
myLinkedList.pop()
myLinkedList.shift()

```