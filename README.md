# Designing Data Intensive Applications using Go
This is a repo used for both learning Go and exploring concepts in the book: Designing Data Intensive Applications. 

## Inverted Indexer 
There are three different impelementations for the inverted indexer. `inverted_indexer_main` contains the main function, and the `indexer` package has the three implementations: 
1. Single Threaded 
2. Multiple Reader threads, Single Writer thread
3. Multiple Reader threads, Multiple Writer threads, also includes Sorting threads
