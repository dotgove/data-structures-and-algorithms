# Notes

## Constants don't matter

| Don't    | Do     |
| -------- | ------ |
| O(2n)    | (O(n)  |
| O(500)   | O(1)   |
| O(13n^2) | O(n^2) |

## Smaller terms also don't matter

| Don't           | Do     |
| --------------- | ------ |
| O(n + 10)       | (O(n)  |
| O(1000n + 50)   | O(n)   |
| O(n^2 + 5n + 8) | O(n^2) |

# Big O Shorthands

-   Arithmetic operations are constant. The computer takes almost the same amount of time to add 2+2 or 2M + 2.
-   Variable assignment is also constant. It takes same amount of time to assign x a value of 2 or 2M
-   Accessing element in an array, slice or map by its index is constant.
-   In a loop, the complexity is the length of the loop _times_ the complexity of whatever happens inside the loop.

A couple of examples:

```js
function printAtLeast5(n) {
    for (let i = 1; i <= Math.max(5, n); i++) {
        console.log(i)
    }
}
```

The above function loops and prints over at least 5 time. However as the value of arg n increases significantly, It becomes clearer that the time complexity is O(n).

```js
function printAtMost5(n) {
    for (let i = 1; i <= Math.min(5, n); i++) {
        console.log(i)
    }
}
```

Despite the value of n, The above function loops and prints over at most 5 times. So its complexity is O(1)

## Space Complexity

### Auxillary Space Complexity

Sometimes the term auxillary space complexity is referred to the space required by the input algorithm, not including space taken up by the input arguments.

### Examples:

#### Sum of elements

```go
func sumOfElements(nums []int)(total int){
    for i:=0 ; i <=len(nums) ; i++{
        total+= nums[i]
    }
    return
}
```

In the above program no new variable is being created by the program. The Space complexity therefore is O(1)

#### Double Elements

```go
func doubleElements(nums []int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] += n * 2
	}
	return result
}
```

In the above code the space complexity is directly proportionate to the size of input. The complexity is O(n)
