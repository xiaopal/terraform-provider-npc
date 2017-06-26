package api

import (
	"errors"
	"fmt"
)

type ApiNamespaceResult struct {
	Id         int64  `json:"id"`
	Name       string `json:"display_name"`
	UniqueName string `json:"namespace"`
	Deletable  bool   `json:"deletable"`
}

type apiListNamespacesResult struct {
	Namespaces []ApiNamespaceResult `json:"namespaces"`
}

func (c *ApiClient) ListNamespaces() ([]ApiNamespaceResult, error) {
	result := &apiListNamespacesResult{}
	return result.Namespaces, c.Get("/api/v1/namespaces", result)
}

type apiCreateNamespaceResult struct {
	Id int64 `json:"namespace_Id"`
}

func (c *ApiClient) CreateNamespace(name string) (int64, error) {
	result := &apiCreateNamespaceResult{}
	return result.Id,
		c.Post("/api/v1/namespaces",
			map[string]string{"name": name},
			result)
}

func (c *ApiClient) DeleteNamespace(id int64) error {
	return c.Do("DELETE", fmt.Sprintf("/api/v1/namespaces/%d", id), nil, nil)
}

func (c *ApiClient) LookupNamespace(name string) (int64, error) {
	namespaces, err := c.ListNamespaces()
	if err != nil {
		return 0, err
	}
	for _, ns := range namespaces {
		if ns.Name == name {
			return ns.Id, nil
		}
	}
	return 0, errors.New("Not found")
}

func (c *ApiClient) LookupOrCreateNamespace(name string) (int64, error) {
	ns, err := c.LookupNamespace(name)
	if err != nil {
		return c.CreateNamespace(name)
	}
	return ns, nil
}

func (c *ApiClient) DeleteNamespaceByName(name string) error {
	ns, err := c.LookupNamespace(name)
	if err != nil {
		return err
	}
	return c.DeleteNamespace(ns)
}
