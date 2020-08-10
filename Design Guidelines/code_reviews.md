# Code Reviews

## Integrity

* Serious about reliability

### Driving Forces

* Is about allocation, read and write of memory being accurate, consistent and efficient
* Type system is CRITICAL, we need to have micro level of integrity
* It is about Data Transformation being accurate, consistent, and efficient. Writing less code and error handling is critical to make sure we have macro level integrity

## Write Less Code

* Average of 15 to 50 bugs per 1000 lines of code
* More places for bugs to hide with larger amounts of code

## Error Handling

* Should be treated as an exception and not part of the main code
* Most critical failures are due to error handling
* Most code is focused on failure
  * This is solved through writing good error handling

### Example

* 48 critical failures were found looking at a couple hundred bugs in Cassandra, HBase, HDFS, MapReduce, and Redis
* 92%: Failures from bad error handling
  * 35%: Incorrect handling
    * 25%: Simply ignoring an error
    * 8%: Catching the wrong exception
    * 2%: Incomplete TODOs
* 57%: System specific
  * 23%: Easily detectable
  * 34%: Complex Bugs
* 8%: Failures from latent human errors

## Readability

* Should be based on structure
* Readable code should have the desire to have all developers have an understanding of the Mental Model and can fix any bugs
* You need understand your team and find their strengths and weaknesses
* If you better at the project type you are working on, you need to write less clever code for your average coworkers
* If you are worse at the project type, you need to ask your coworkers to either help you get up to speed or to write less clever code
* Not hiding the cost of the code you are writing

## Simplicity

* Hide complexity
* Create encapsulations
* You refactor into simplicity
* Complexity sells, we need to stop this
* Encapsulations need to not hide cost, they need to hide complexity