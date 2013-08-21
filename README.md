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

Run that and you'll see something like this:

![output](https://github.com/AndrewVos/o/raw/master/output.png)

## FAQ
- Q: How do I pronounce o.O?
- A: [Kind of like this](http://www.youtube.com/watch?v=140S4LkQxxA)
