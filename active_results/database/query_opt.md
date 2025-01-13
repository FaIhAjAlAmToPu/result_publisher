We want to return results fast
**Even data distribution:** every node should hold roughly the same amount of data for that we need suitable partition key.<br>
**To minimize the number of partitions accessed in a read query:** To make reads faster, we’d ideally have all the data required in a read query stored in a single Table. Although it’s fine to duplicate data across tables, in terms of performance, it’s better if the data needed for a read query is in one table.<br>
**NOT Minimizing the number of writes:** “cheap.” Reads, while still very fast, are usually more expensive than writes and are harder to fine-tune. We’d usually be ready to increase the number of writes to increase read efficiency. Keep in mind that the number of tables also affects consistency.
**NOT Avoiding data duplication:** get efficient reads. to avoid duplication in some cases using Secondary Indexes.

#### Primary Key to include more than one column
* first part of the Primary Key is called the Partition Key (pet_chip_id in the above example) and the second part is called the Clustering Key (time).
* every query must include all columns defined in the partition key.

### choosing compaction
https://opensource.docs.scylladb.com/stable/architecture/compaction/compaction-strategies.html#id1