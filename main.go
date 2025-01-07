package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
)

const usage string = "gh noreply-email [USER]"

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, usage)
		return
	}

	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	endpoint := "user"
	if len(args) == 1 {
		endpoint = "users/" + args[0]
	}
	response := struct {
		Id    int
		Login string
	}{}
	err = client.Get(endpoint, &response)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("%d+%s@noreply.users.github.com\n", response.Id, response.Login)
}
