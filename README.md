# Consistent Hashing

To achieve horizontal scaling, it is important to distribute requests/data efficiently and evenly aross server. Consistent hashing is a commonly used technique to achieve this goal.

## The rehashing problem

If you have n cache servers, a common way to balance the load is use the following hash method:

serverIndex = hash(key) % N, where N is the size of the server pool

## Consistent hashing

