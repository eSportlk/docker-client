package main

// func main() {
// 	ctx := context.Background()
// 	cli, err := client.NewClientWithOpts(client.FromEnv)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cli.NegotiateAPIVersion(ctx)

// 	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	io.Copy(os.Stdout, reader)

// 	resp, err := cli.ContainerCreate(ctx, &container.Config{
// 		Image: "alpine",
// 		Cmd:   []string{"echo", "hello world"},
// 	}, nil, nil, "")
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
// 		panic(err)
// 	}

// 	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
// 	select {
// 	case err := <-errCh:
// 		if err != nil {
// 			panic(err)
// 		}
// 	case <-statusCh:
// 	}

// 	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
// 	if err != nil {
// 		panic(err)
// 	}

// 	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
// }