# o.O

## Usage

    package main

    import "github.com/AndrewVos/o"

    type Something struct {
      Name string
      Value string
    }

    func main() {
      something := Something { Name: "Donald", Value: "Duck" }
      o.O(something)

      someString := "o.O is pretty awesome!"
      o.O(someString)
    }

Run this and you'll see something like this:

![output](https://github.com/AndrewVos/o/raw/master/output.png)
