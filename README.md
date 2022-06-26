# int-maths, Go mathematics for integers

My firs public module for Go. ***It' a work in progress!*** Use it at your own risks!

## safe operations on integers

I like Go but it has some annoying flaws. One of them is the lack of overflow and underflow with simple operations on integers. Here is my version which use generics.

Here is a set of *pros* and *cons*:

| ***Pros***                                                | ***Cons***                                                          |
| --------------------------------------------------------- | ------------------------------------------------------------------- |
| it's my first real project in *Go*                        | it's my first real project in *Go*                                  |
| One function for all integer types, thanks to *generics*. | We use the `reflect` package, which is not so good for performances |
|                                                           | `SAdd` and `SSub` are mutually recursive                            |

### tests

|              | SSub          | SAdd          | SMul          |
| ------------ | ------------- | ------------- | ------------- |
| ***int8***   | ***partial*** | ***full***    | none          |
| ***int16***  | none          | none          | none          |
| ***int32***  | none          | none          | none          |
| ***int64***  | none          | *some*        | *some*        |
| ***uint8***  | ***partial*** | ***full***    | none          |
| ***uint16*** | none          | none          | none          |
| ***uint32*** | none          | none          | none          |
| ***uint64*** | none          | ***partial*** | ***partial*** |
