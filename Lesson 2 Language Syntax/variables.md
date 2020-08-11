# Variables

* Type is life, determines the cost
* Size
* Representation
* int is based on architecture
  * Look at the address size
    * 32
    * 64
  * int will give us mechanical sympathies unless you have a specific reason to use a precision based int
    * e.g int64
* One of the warts in Go
  * TOO MANY WAYS TO DECLARE VARIABLES
* If you want a 0 value initialization for a variable, you want to use var. It works 100% of the time. This helps from a readability perspective
* When you create a string using var, it creates an empty string
* string variables are based architecture
* They are 2 words long
  * 4 and 4 bytes for 32 bit (8 bytes)
  * 8 and 8 bytes for 64 bit (16 bytes)