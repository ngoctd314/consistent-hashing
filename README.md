# Consistent Hashing

To achieve horizontal scaling, it is important to distribute requests/data efficiently and evenly aross server. Consistent hashing is a commonly used technique to achieve this goal.

## The rehashing problem

If you have n cache servers, a common way to balance the load is use the following hash method:

serverIndex = hash(key) % N, where N is the size of the server pool

This approach works well when the size of the server pool is fixed, and the data distribution is even. However, problems arise when new servers are added, or existing servers are removed.

## Consistent hashing

Consistent hashing is a special kind of hashing such that when a hash table is re-sized and consistent hashing is used, only k/n keys need to be remapped on average, where k is the number of keys and n is the number of slots. In contrast, in most traditional hash tables, a change in the number of array slots causes nearly all keys to be remapped.

### Hash space and hash ring

### Hash servers

Using the same hash function f, we map servers based on server IP or name onto the ring

### Hash keys

One thing worth mentioning is that hash function used here is different from the one in "the rehashing problem" and there is no modular operation.

### Server lookup

To determine which server a key is stored on, we go clockwise from the key position on the ring until a server is found.

### Add a server

Using the logic described above, adding a new server will only request redistribution of a fraction of keys

### Remove a server

When a server is removed, only a small fraction of keys request redistribution with consistent hashing.

### Two issues in the basic approach