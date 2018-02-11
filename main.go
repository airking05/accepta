package accepta

import (
	apiclient "github.com/airking05/accepta/client"
	"github.com/airking05/accepta/client/container"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"context"
	"fmt"
	"github.com/airking05/accepta/wrapper"
)

func NewDefault() *wrapper.Client {
	transport := httptransport.New("0.0.0.0:2375", "",[]string{"http"})
	return &wrapper.Client{
		apiclient.New(transport, strfmt.Default),
	}
}

func main() {
	client := NewDefault()
	containerList,err := client.ListContainers(container.NewContainerListParamsWithContext(context.Background()))
	if err != nil {
		panic(err)
	}
	for _,c := range containerList {
		fmt.Println(c.ID)
		fmt.Println(c.Status)
		fmt.Println(c.Command)
	}
}