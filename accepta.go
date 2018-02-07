package main

import (
	apiclient "github.com/airking05/accepta/client"
	"github.com/airking05/accepta/client/container"
	httptransport "github.com/go-openapi/runtime/client"
	"fmt"
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/davecgh/go-spew/spew"
	"net/http/httptest"
)

func main() {
	w := httptest.NewRecorder()
	transport := httptransport.New("0.0.0.0:2375", "",[]string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	resp,err:=client.Container.ContainerLogs(container.NewContainerLogsParamsWithContext(context.Background()),w)
	if err != nil {
		fmt.Println(err)
		return
	}
	spew.Dump(resp)

}