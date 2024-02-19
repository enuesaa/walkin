package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/buildreq"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)


func PromptReq(repos repository.Repos, method string, url string) (invoke.Invocation, error) {
	builder := buildreq.New(repos, method, url)

	fmt.Printf("***\n")
	if builder.IsUrlEmpty() {
		if err := builder.AskUrl(); err != nil {
			return builder.Invocation, err
		}
	}
	fmt.Printf("* %s\n", builder.Endpoint())
	fmt.Printf("*\n")
	fmt.Printf("* [Headers]\n")

	for {
		if err := builder.AskHeader(); err != nil {
			if err == buildreq.SKIP_HEADER {
				break
			}
			return builder.Invocation, err
		}
		lastHeader := builder.GetLastHeader()
		fmt.Printf("* %s: %s\n", lastHeader.Key, lastHeader.Value)
	}
	if method == "POST" || method == "PUT" {
		fmt.Printf("*\n")

		if err := builder.AskBody(); err != nil {
			return builder.Invocation, err
		}
		fmt.Printf("* [Body]\n")
		fmt.Printf("* %s\n", builder.Invocation.RequestBody)
	}
	fmt.Printf("***\n")

	return builder.Invocation, nil
}