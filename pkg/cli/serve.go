package cli

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve",
		Run: func(cmd *cobra.Command, args []string) {	
			app := fiber.New()

			app.Get("/*", func(c *fiber.Ctx) error {
				path := c.Path() // like `/`
				path = filepath.Join("./", path, "index.html")
				fmt.Printf("%s", path)

				f, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				fileExt := filepath.Ext(path)
				mimeType := mime.TypeByExtension(fileExt)
				c.Set(fiber.HeaderContentType, mimeType)

				return c.SendString(string(f))
			})
			app.Listen(":3000")
		},
	}

	return cmd
}
