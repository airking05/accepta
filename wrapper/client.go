package wrapper

import (
	apiclient "github.com/airking05/go-docker-api-client/client"
	"github.com/airking05/go-docker-api-client/client/container"
	"github.com/airking05/go-docker-api-client/models"
	"io"
)

type Client struct{
	*apiclient.Core
}
func(c *Client) ListContainers(params *container.ContainerListParams) ([]*models.ContainerSummaryItems,error){
	resp,err:= c.Container.ContainerList(params)
	if err != nil {
		return resp.Payload,err
	}
	return resp.Payload,nil
}

func(c *Client) UpdateContainer(params *container.ContainerUpdateParams) (error) {
	_,err := c.Container.ContainerUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RenameContainer(params *container.ContainerRenameParams) (error){
	_, err := c.Container.ContainerRename(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) InspectContainer(params *container.ContainerInspectParams) (*models.ContainerInspectOKBody, error) {
	resp, err := c.Container.ContainerInspect(params)
	if err != nil {
		return resp.Payload,err
	}
	return resp.Payload,nil
}

func (c *Client) ContainerChanges(params *container.ContainerChangesParams) (*models.ContainerChangesOKBody, error) {
	resp,err := c.Container.ContainerChanges(params)
	if err != nil {
		return &resp.Payload,err
	}
	return &resp.Payload,nil
}

func (c *Client)CreateContainer(params *container.ContainerCreateParams) (*models.ContainerCreateCreatedBody,error) {
	resp, err := c.Container.ContainerCreate(params)
	if err != nil {
		return resp.Payload,err
	}
	return resp.Payload,nil
}

func (c *Client)StartConContainer(params *container.ContainerStartParams) (error) {
	_, err := c.Container.ContainerStart(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StopContainer(params *container.ContainerStopParams) (error) {
	_,err := c.Container.ContainerStop(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RestartContainer(params *container.ContainerRestartParams) (error) {
	_,err := c.Container.ContainerRestart(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PauseContainer(params *container.ContainerPauseParams) (error) {
	_,err := c.Container.ContainerPause(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UnpauseContainer(params *container.ContainerUnpauseParams) (error) {
	_,err := c.Container.ContainerUnpause(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) TopContainer(params *container.ContainerTopParams) (*models.ContainerTopOKBody,error) {
	resp,err := c.Container.ContainerTop(params)
	if err != nil {
		return resp.Payload,err
	}
	return resp.Payload,nil
}

func (c *Client) Stats(params *container.ContainerStatsParams) (error) {
	_,err := c.Container.ContainerStats(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) KillContainer(params *container.ContainerKillParams) (error) {
	_,err := c.Container.ContainerKill(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerDelete(params *container.ContainerDeleteParams) (error) {
	_,err := c.Container.ContainerDelete(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerWait(params *container.ContainerWaitParams) (error) {
	_,err := c.Container.ContainerWait(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerAttach(params *container.ContainerAttachParams) (error) {
	_,err := c.Container.ContainerAttach(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CoentainerAttachWebsocket(params *container.ContainerAttachWebsocketParams) (*container.ContainerAttachWebsocketOK,error) {
	resp,err := c.Container.ContainerAttachWebsocket(params)
	if err != nil {
		return resp,err
	}
	return resp,nil
}

func (c *Client) Logs(params *container.ContainerLogsParams, w io.Writer) (error) {
	_,err := c.Container.ContainerLogs(params,w)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerResize(params *container.ContainerResizeParams) (error) {
	_,err := c.Container.ContainerResize(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerExport(params *container.ContainerExportParams) (error) {
	_,err := c.Container.ContainerExport(params)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ContainerPrune(params *container.ContainerPruneParams) (*models.ContainerPruneOKBody,error) {
	resp,err := c.Container.ContainerPrune(params)
	if err != nil {
		return resp.Payload,err
	}
	return resp.Payload,nil
}

