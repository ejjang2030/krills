# krills
Krills is Simple Programming Language.

## Getting Started
### git clone command
`git clone https://github.com/ejjang2030/krills.git`

Then access to the directory which at the git cloned the files with `cd` command in the terminal or cmd.<br>
Run this command `go run main.go` in the directory.

## Syntax

### Variable Types 
1. Integer
2. Boolean
3. String
4. Function
5. Array
6. HashMap

### Variable Definition and Initialization
```
>> let n1 = 5;
>> let n2 = 10;
>> n1 + n2
15
```

### Function Definition
```
let func1 = fun(x, y) {
    x + y;
  }
```
or
```
fun test(x) {
    return x;
}
```
**function identifier must not contain numbers**

## Future Goals to Implementation
### Loop
1. Condition Loop
```
    while: a.next() { 
        // statements
    }
```
2. for Loop
``` 
    for: (let i = 0; i < a.length; i++) {
        // statements 
    }
```
3. Nested for Loop in one Line
``` 
    for: (let i in range(1, 3)), (let j in range(1, 10)) {
        print("$i X $j = ${i * j}")
    }
```
This code is same to below code.
```
    for: (let i in range(1, 3)) {
        for: (let j in range(1, 10)) {
            print("$i X $j = ${i * j}")
        }
    }
```
4. Nested while Loop in one Line
```
    while: (!a.isEmpty()), (a.value != 0) {
        // statements
    }
```
This code is same to below code.
```
    while: (!a.isEmpty()) {
        while: (a.value != 0) {
            // statements
        }
    }
```
 

## If you want to contribute this Krills Project, then please don't get some hesitation to contact me. 
Here's my email: ejjang2030@gmail.com

