# Insertion
-  Adds an element at the given index
- Two types of insertion are possible 
1. Insertion at the end of array
2. Insertion in the middle of the array


Write a program to insert an element in an array?
Consider
- Where to insert?
- What are the current elements in the array?
- What value you want to insert?
- How many elements are there in this array?

### Problem1 
- This problem will be implemented using python and go the file will be named as problem1 with their respective file extension.

Algorithm: Insertion of an element ITEM into Kth position in a Linear Array A with N elements.
Input:
A: Array A ;  N: Array A that has N elements in it
ITEM :item to insert ;   Pos: insert item at position Pos
Output:
Array A with item inserted at position index ‘Pos’ Algorithm
1. [Initialize Counter.] Set I = N.
2. Repeat Steps 3 and 4 while I >= Pos.
3. [Move Ith element downward] Set A[I+1] = A[I].
4. [Decrease Counter.] Set I = I-1. [End of Step 2 loop]
5. [Insert element.] Set A[Pos] = ITEM.
6. [Reset N] N = N+1.
7. Exit. 


### Problem2 
-  This problem will be implemented using python and go the file will be named as problem2 with their respective file extension.
Algorithm (for index starting from 0 and pos is also in the
form of index)
1. [Initialize Counter.] Set I = N-1. 2. Repeat Steps 3 and 4 while I >=Pos. 3. [Move Ith element downward] Set A[I+1] = A[I].
4. [Decrease Counter.] Set I = I-1. [End of Step 2 loop]
5. [Insert element.] Set A[Pos] = ITEM.
6. [Reset N] N = N+1.
7. Exit