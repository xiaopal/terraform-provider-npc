package npc

import "github.com/xiaopal/terraform-provider-npc/npc/api"

// Config 配置
type Config struct {
	Credentials *Credentials
	ApiEndpoint string
}

type Credentials api.ApiCredentials

func (config *Config) ApiClient() *api.ApiClient {
	return &api.ApiClient{
		Credentials: (*api.ApiCredentials)(config.Credentials),
		Endpoint:    config.ApiEndpoint,
	}
}
