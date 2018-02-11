[![GoDoc](https://godoc.org/github.com/airking05/accepta?status.svg)](https://godoc.org/github.com/airking05/accepta)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/airking05/accepta/blob/master/LICENSE)

- Accepta is a smart API client for "Docker", implemented by golang.
- We can use this to newer API version(eg. 1.33) with swagger code gen.

# What is this :

- this is an API client for Docker.

# How to Install :

```bash
go get github.com/airking05/go-docker-api-client
```

# How to Use :

```Go
package main

import (
	"github.com/airking05/go-docker-api-client"
	"github.com/airking05/go-docker-api-client/client/container"
	"context"
	"fmt"
	)

func main() {
	client := accepta.NewDefault()
	
	// Container List
    containerList,err := client.ListContainers(container.NewContainerListParamsWithContext(context.Background()))
    
    if err != nil {
    	panic(err)
    }
    
    for _,c := range containerList {
   	    fmt.Println(c.ID)
		fmt.Println(c.Status)
		fmt.Println(c.Command)
    }}
    
```