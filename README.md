# int-maths, Go mathematics for integers

My firs public module for Go. ***It' a work in progress!*** Use it at your own risks!

## safe operations on integers

I like Go but it has some annoying flaws. One of them is the lack of overflow and underflow with simple operations on integers. Here is my version which use generics.

We have three function 

- `SSub` for subtracting two numbers; `SSub(a, b)` returns `a - b` and `true` if no over/underflow occurs, 

- `SAdd` for adding two numbers; `SAdd(a, b)` returns `a + b` and `true` if no over/underflow occurs,

- `SMul` for multiplying two numbers; `SMul(a, b)` returns `a * b` and `true` if no over/underflow occurs.

The two parameters must be of the same type and one of `int8` to `int64` or `uint8` to `uint64`.

Here is a set of *pros* and *cons*:

| ***Pros***                                                | ***Cons***                                                          |
| --------------------------------------------------------- | ------------------------------------------------------------------- |
| it's my first real project in *Go*                        | it's my first real project in *Go*                                  |
| One function for all integer types, thanks to *generics*. | We use the `reflect` package, which is not so good for performances |

### tests

|              | SSub       | SAdd          | SMul          |
| ------------ | ---------- | ------------- | ------------- |
| ***int8***   | ***full*** | ***full***    | none          |
| ***int16***  | ***full*** | ***full***    | none          |
| ***int32***  | none       | none          | none          |
| ***int64***  | none       | ***partial*** | ***partial*** |
| ***uint8***  | ***full*** | ***full***    | none          |
| ***uint16*** | ***full*** | ***full***    | none          |
| ***uint32*** | none       | none          | none          |
| ***uint64*** | none       | ***partial*** | ***partial*** |
