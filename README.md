# marcel-lib-pack
A pack of my common utility functions

## Arrays
A package that do array operation similar to javascript array functions
### Reduce
The `Reduce` function is a generic utility function in Go that reduces a slice of values to a single result using a callback function.

Function signature:

`func Reduce[TSlice any, TResult any](s []TSlice, initialValue TResult, cb reduceFunc[TSlice, TResult]) TResult`

|Parameter    |Type     |Description                                                                                                                      |
|-------------|-----------------------------|-------------------------------------------------------------------------------------------------------------|
|s	          |[]TSlice	                    |A slice of type `TSlice` to be reduced.                                                                      |
|initialValue |[]TSlice	                    |The initial value of the result to start the reduction process.                                              |
|cb           |reduceFunc[TSlice, TResult]	|A callback function that defines how two values (an accumulator and the current element) combine.            |

|Generic Type|Description                                                         |
|------------|--------------------------------------------------------------------|
|TSlice      |Represents the type of the elements in the slice `s`.               |
|TResult     |Represents the type of the result produced by the reduction process.|

Example:
```go
s := []int{1, 2, 3, 4, 5}
total := arrays.Reduce(s, 0, func(total, currentValue int) int {
    return total + currentValue
})
```

### Map
The `Map` function is a generic utility function in Go that transforms a slice of values into another slice of values using a callback function.

Function signature: 

`func Map[TSlice any, TResult any](s []TSlice, cb mapFunc[TSlice, TResult]) []TResult`

|Parameter    |Type     |Description                                                                                                         |
|-------------|-----------------------------|------------------------------------------------------------------------------------------------|
|s	          |[]TSlice	                    |A slice of type TSlice that will be transformed.                                                |
|cb           |mapFunc[TSlice, TResult]	    |A callback function that defines how each element of the input slice is transformed.            |

|Generic Type|Description                                                         |
|------------|--------------------------------------------------------------------|
|TSlice      |Represents the type of the elements in the slice `s`.               |
|TResult     |Represents the type of the result produced by the reduction process.|

Example:
```go
s := []int{1, 2, 3, 4, 5}
result := arrays.Map(s, func(currentValue int) int {
    return currentValue * 2
})
```

### Filter
The `Filter` function is a generic utility function in Go that filters elements from a slice based on a provided condition defined in a callback function.

Function signature: 

`Filter[TSlice any](s []TSlice, cb func(currentValue TSlice) bool) []TSlice`

|Parameter    |Type                          |Description                                                                                     |
|-------------|------------------------------|------------------------------------------------------------------------------------------------|
|s	          |[]TSlice	                     |A slice of type TSlice to be filtered.                                                          |
|cb           |func(currentValue TSlice) bool|A callback function that determines if an element should be included in the resulting slice.    |

|Generic Type|Description                                                         |
|------------|--------------------------------------------------------------------|
|TSlice      |Represents the type of the elements in the slice `s`.               |

Example:
```go
s := []int{1, 2, 3, 4, 5}
result := arrays.Filter(s, func(currentValue int) bool {
    return currentValue%2 == 0
})
```

### Find
The `Find` function is a generic utility function in Go that searches through a slice and returns the first element that satisfies a condition defined in a callback function. If no element satisfies the condition, it returns `nil`.

Function signature: 

`func Find[TSlice any](s []TSlice, cb func(currentValue TSlice) bool) *TSlice`

|Parameter    |Type                          |Description                                                                                     |
|-------------|------------------------------|------------------------------------------------------------------------------------------------|
|s	          |[]TSlice	                     |A slice of type TSlice to be searched.                                                          |
|cb           |func(currentValue TSlice) bool|A callback function that determines whether an element satisfies the condition.                 |

|Generic Type|Description                                                         |
|------------|--------------------------------------------------------------------|
|TSlice      |Represents the type of the elements in the slice `s`.               |

Example:
```go
s := []int{1, 2, 3, 4, 5}
result := arrays.Find(s, func(currentValue int) bool {
    return currentValue%2 == 0
})
```

## Goment
A moment.js-like date string parser
```go
date, err := goment.New("2021-01-02", "YYYY-MM-DD")
// Handle error here

dateString := date.Formats("dddd, MMMM D YYYY") // Will return Saturday, January 2 2021
```

## try.Catch
An all panic recovery that will return error when recovered, shouldn't really use this though since you should always return errors as value and handle it.
```go
func throw() (err error) {
	defer try.Catch(&err)
	panic("error")
}

func main() {
	err := throw()

    if err != nil {
        log.Println(err.Error())
    }
}
```

## Gollama
A mini-SDK to communicate to [Ollama](https://ollama.com/) API. Currently can only do generate a chat completion.
```go
ollama := gollama.New(gollama.WithHost("http://localhost:11434")) // Put your ollama server here by default it will use http://localhost:11434

messages := []gollama.ChatMessage{}
messages = append(messages, gollama.ChatMessage{
    Role:    "system",
    Content: "Your system prompt here",
})
messages = append(messages, gollama.ChatMessage{
    Role:    "user",
    Content: "Content of the message to be sent to AI assistant",
})

res, err := ollama.Chat(gollama.ChatRequest{
    Model:    "llama3.2",
    Messages: messages,
})

log.Println(res.Message)
// result: {
//     Role: "assistant",
//     Content: "Message content of AI chat generation",
// }   
```