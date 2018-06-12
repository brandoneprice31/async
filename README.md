async - go package that makes concurrency easier
==================================================

------
Installation
================

To install assert, use `go get`:

    go get github.com/brandoneprice31/async

Import the `async` package into your code like so:

```go
package main

import (
  "github.com/brandoneprice31/async"
)

func main() {
  async.Parallel()
}
```

------
Parallel
============

```go
func Parallel(...func() error) Errors
```

Makes multiple parallel functions calls and collects their returned errors safely.  Each argument is a function that returns an error.  The returned object contains all of the errors that were returned by these arguments.

Parallel doesn't guarantee the synchronization of shared memory between the functions being called.  Therefore all variables being shared by these functions need to make use of go sync primitives like channels or mutex locks.

```go
func main() {
  if err := UploadAndDownload(); err != nil {
    panic(err)
  }
}

func UploadAndDownload() error {
  api := Connect()

  var email string
  errs := async.Parallel(
    func() error {
      return api.Post(Request{
        Name: "Brandon",
        Email: "brandoneprice31@gmail.com",
      })
    },

    func() error {
      user, err = api.Get(Request{
        ID: 1234,
      })
      if err != nil {
        return err
      }

      email = user.Email
      return nil
    },
  )

  fmt.Println("email:", email)
  return errs.ToError()
}
```
