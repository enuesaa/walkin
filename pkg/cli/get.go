package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateGetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "make a get request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			verbose, _ := cmd.Flags().GetBool("verbose")

			invocation := invoke.NewInvocation("GET", url)
			if err := usecase.PromptReq(repos, &invocation); err != nil {
				return err
			}
			fmt.Printf("%+v", invocation)
			if err := usecase.Invoke(repos, &invocation, !verbose); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("url", "", "url")
	cmd.Flags().BoolP("verbose", "v", false, "verbose")

	return cmd
}
