# o.O

## Demo

![output](https://github.com/AndrewVos/o/raw/master/o.gif)

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

## Install
    go get github.com/AndrewVos/o

## FAQ
- Q: How do I pronounce o.O?
- A: [Kind of like this](http://www.youtube.com/watch?v=140S4LkQxxA)
