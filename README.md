# int-maths, Go mathematics for integers

My firs public module for Go. It' a work in progress!

## safe operations on integers

I like Go but it has some annoying flaws. One of them is the lack of overflow and underflow with simple operations on integers. Here is my version which use generics.

Here is a set of *pros* and *cons*:

| ***Pros***                                               | ***Cons***                                                        |
|----------------------------------------------------------|-------------------------------------------------------------------|
| One funtion for all integer types, thanks to *generics*. | We use the reflect package, which is not so good for performances |
| it's my first real project                               | it's my first real project                                        |

## the manual
### Safe operations
#### `SSub`
- **Note**: Not well tested
- **Syntax**:
```Go
func SSub[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64]
        (a V, b V) 
        (V, bool)
```
- **Explain**: Compute the diffrence betweeen two integers of the same size
- ***Exemple***:
```Go
var x uint16 = 60000
var y uint16 = 30000 
var delta, ok = SSub(x, y)
```
`ok` will be `true` and the result will be 30000

#### `SAdd`
- **Note**: Not well tested
- **Syntax**:
```Go
func SAdd[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64]
        (a V, b V) 
        (V, bool)
```
- **Explain**: Compute the sum of two integers of the same size
***Exemple***
```Go
var x uint16 = 60000
var y uint16 = 30000 
var delta, ok = SAdd(x, y)
```
`ok` will be `false` and the result will be `9000 mod 65536

#### `SMul`
- **Note**: Not well tested
- **Syntax**:
```Go
func SMul[V int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64]
        (a V, b V) 
        (V, bool)
```
- **Explain**: Compute the sum of two integers of the same size
- ***Exemple***:
```Go
var x uint16 = 60000
var y uint16 = 3 
var delta, ok = SMul(x, y)
```
`ok` will be `false` and the result will be `180000 mod 65536`
