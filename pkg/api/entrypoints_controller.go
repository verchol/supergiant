package api

import (
	"net/http"

	"github.com/supergiant/supergiant/pkg/core"
	"github.com/supergiant/supergiant/pkg/models"
)

func ListEntrypoints(core *core.Core, r *http.Request) (*Response, error) {
	return handleList(core, r, new(models.Entrypoint))
}

func CreateEntrypoint(core *core.Core, r *http.Request) (*Response, error) {
	item := new(models.Entrypoint)
	if err := decodeBodyInto(r, item); err != nil {
		return nil, err
	}
	if err := core.Entrypoints.Create(item); err != nil {
		return nil, err
	}
	return itemResponse(core, item, http.StatusCreated)
}

func UpdateEntrypoint(core *core.Core, r *http.Request) (*Response, error) {
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	item := new(models.Entrypoint)
	if err := decodeBodyInto(r, item); err != nil {
		return nil, err
	}
	if err := core.Entrypoints.Update(id, new(models.Entrypoint), item); err != nil {
		return nil, err
	}
	return itemResponse(core, item, http.StatusAccepted)
}

func GetEntrypoint(core *core.Core, r *http.Request) (*Response, error) {
	item := new(models.Entrypoint)
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	if err := core.Entrypoints.Get(id, item); err != nil {
		return nil, err
	}
	return itemResponse(core, item, http.StatusOK)
}

func DeleteEntrypoint(core *core.Core, r *http.Request) (*Response, error) {
	item := new(models.Entrypoint)
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	if err := core.Entrypoints.Delete(id, item).Async(); err != nil {
		return nil, err
	}
	return itemResponse(core, item, http.StatusAccepted)
}
