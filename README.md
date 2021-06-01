# Go Sorting Algorithm

In this repository you will find five basic sorting algorithms written in go language, the algorithms are:
bubble sort, insertion sort, selection sort, merge sort and quicksort.

You will find a main function in which you can call any of these algorithms, they will be between to timer 
variables that at the end of the function prints the time taken to execute. 

### Explanation of the algorithms: 

(Source Wikipedia)

#### Bubble sort:
The bubble sort compares adjacent elements and swaps them if they are in the wrong order. The pass through the 
list is repeated until the list is sorted. The algorithm

![Alt Text](https://upload.wikimedia.org/wikipedia/commons/c/c8/Bubble-sort-example-300px.gif)

#### Insertion sort:
Insertion sort iterates, consuming one input element each repetition, and grows a sorted output list. 
At each iteration, insertion sort removes one element from the input data, finds the location it 
belongs within the sorted list, and inserts it there. It repeats until no input elements remain.

![Alt Text](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

#### Selection sort:

The algorithm divides the input list into two parts: a sorted sublist of items which is built up from left to right at the front (left) of 
the list and a sublist of the remaining unsorted items that occupy the rest of the list. 

Find the smallest item in the list, trade it with the first then find the next minimum in the rest of the list and swap it with the second.

![Alt Text](https://upload.wikimedia.org/wikipedia/commons/9/94/Selection-Sort-Animation.gif)

#### Merge sort:
Conceptually, a merge sort works as follows:
1.	Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
2.	Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.

![Alt Text](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

#### Quick sort:
Quicksort is a divide-and-conquer algorithm. It works by selecting a 'pivot' element from the array and 
partitioning the other elements into two sub-arrays, according to whether they are less than or greater 
than the pivot. For this reason, it is sometimes called partition-exchange sort.[4] The sub-arrays are then 
sorted recursively. This can be done in-place, requiring small additional amounts of memory to perform the 
sorting.

![Alt Text](https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif)

