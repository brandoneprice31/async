async - go package that makes concurrency easier
==================================================

------
Installation
================

To install assert, using `go get`:

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
Makes multiple concurrent functions calls and collects the errors.

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
