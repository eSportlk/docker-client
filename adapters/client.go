package adapters

import "github.com/docker/docker/client"

func GetClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		return nil
	}
	return cli
}
