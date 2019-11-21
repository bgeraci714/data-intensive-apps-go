# Designing Data Intensive Applications using Go
This is a repo used for both learning Go and exploring concepts in the book: Designing Data Intensive Applications. 

## Inverted Indexer 
There are three different impelementations for the inverted indexer. `inverted_indexer_main` contains the main function, and the `indexer` package has the three implementations: 
1. Single Threaded 
2. Multiple Reader threads, Single Writer thread
3. Multiple Reader threads, Multiple Writer threads, also includes Sorting threads

## Red-Black Tree Implementation
Basic red black tree implementation. Not explicitly threadsafe. 

## TODO 
1. ~~Build in memory key-value server using the red black tree implementation for store's structure.~~  
2. Add logging in advance of writing to the store in order to provide crash/fault tolerance. 
3. Develop routine that saves the current data store's state into a segment file which can then be incorporated into database queries. 
4. Develop merge and compression routine for the segment files. 
5. Allow for rollbacks of the database using the log