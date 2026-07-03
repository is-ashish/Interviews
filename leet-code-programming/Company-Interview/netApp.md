1 - Create LRU Cache -: 

Design a data structure that follows the constraints of a Least Recently Used Cache.

Implement the LRUCache class:
(1) LRUCache(int capacity) Initialize the LRU cache with positive size capacity.

(2) int get(int key) Return the value of the key if the key exists, otherwise return -1.

(3) void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. 

If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.


LRUCache(3)

put(1,1)

put(2,2)

put(3,3)

put(2,4) -----------------------------

put(4,4) ------------------------------------::  (2,4), (3,3), (4,4)

 

get(1) = ? 1

get(2) = ? 4


===============


2nd round -: 

(1) 
Create a job schedules

// Given a list of tasks with dependencies, execute them with maximum parallelism.
// Input:
//   A → []        (no deps)

//   B → [A]       (depends on A)

//   C → [A]       (depends on A)

//   D → [B, C]    (depends on B and C)

// Output: Execute in order respecting deps, maximize parallelism

(2)  SQL Query -: 

table -: products and orders


table orders -- referecne product, dates, 


product -: id, name, 


highest sales product last 1 years - quality


select count(product.name)
from product
inner join orders
on product.id = orders.product
group by product.name
where 




3 .  What are the machanisma to communicate between microservices.


(4)


// Given a list of tasks with dependencies, execute them with maximum parallelism.
// Input:
//   A → []        (no deps)

//   B → [A]       (depends on A)

//   C → [A]       (depends on A)

//   D → [B, C]    (depends on B and C)

// Output: Execute in order respecting deps, maximize parallelism
