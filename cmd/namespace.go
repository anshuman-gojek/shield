package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/salt/printer"
	"github.com/odpf/shield/pkg/file"
	shieldv1beta1 "github.com/odpf/shield/proto/v1beta1"
	cli "github.com/spf13/cobra"
)

func NamespaceCommand(cliConfig *Config) *cli.Command {
	cmd := &cli.Command{
		Use:     "namespace",
		Aliases: []string{"namespaces"},
		Short:   "Manage namespaces",
		Long: heredoc.Doc(`
			Work with namespaces.
		`),
		Example: heredoc.Doc(`
			$ shield namespace create
			$ shield namespace edit
			$ shield namespace view
			$ shield namespace list
		`),
		Annotations: map[string]string{
			"group":  "core",
			"client": "true",
		},
	}

	cmd.AddCommand(createNamespaceCommand(cliConfig))
	cmd.AddCommand(editNamespaceCommand(cliConfig))
	cmd.AddCommand(viewNamespaceCommand(cliConfig))
	cmd.AddCommand(listNamespaceCommand(cliConfig))

	bindFlagsFromClientConfig(cmd)

	return cmd
}

func createNamespaceCommand(cliConfig *Config) *cli.Command {
	var filePath string

	cmd := &cli.Command{
		Use:   "create",
		Short: "Create a namespace",
		Args:  cli.NoArgs,
		Example: heredoc.Doc(`
			$ shield namespace create --file=<namespace-body>
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cli.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			var reqBody shieldv1beta1.NamespaceRequestBody
			if err := file.Parse(filePath, &reqBody); err != nil {
				return err
			}

			err := reqBody.ValidateAll()
			if err != nil {
				return err
			}

			client, cancel, err := createClient(cmd.Context(), cliConfig.Host)
			if err != nil {
				return err
			}
			defer cancel()

			res, err := client.CreateNamespace(cmd.Context(), &shieldv1beta1.CreateNamespaceRequest{
				Body: &reqBody,
			})
			if err != nil {
				return err
			}

			spinner.Stop()
			fmt.Printf("successfully created namespace %s with id %s\n", res.GetNamespace().GetName(), res.GetNamespace().GetId())
			return nil
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the namespace body file")
	cmd.MarkFlagRequired("file")

	return cmd
}

func editNamespaceCommand(cliConfig *Config) *cli.Command {
	var filePath string

	cmd := &cli.Command{
		Use:   "edit",
		Short: "Edit a namespace",
		Args:  cli.ExactArgs(1),
		Example: heredoc.Doc(`
			$ shield namespace edit <namespace-id> --file=<namespace-body>
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cli.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			var reqBody shieldv1beta1.NamespaceRequestBody
			if err := file.Parse(filePath, &reqBody); err != nil {
				return err
			}

			err := reqBody.ValidateAll()
			if err != nil {
				return err
			}

			client, cancel, err := createClient(cmd.Context(), cliConfig.Host)
			if err != nil {
				return err
			}
			defer cancel()

			namespaceID := args[0]
			res, err := client.UpdateNamespace(cmd.Context(), &shieldv1beta1.UpdateNamespaceRequest{
				Id:   namespaceID,
				Body: &reqBody,
			})
			if err != nil {
				return err
			}

			spinner.Stop()
			fmt.Printf("successfully edited namespace with id %s to id %s and name %s\n", namespaceID, res.GetNamespace().GetId(), res.GetNamespace().GetName())
			return nil
		},
	}

	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the namespace body file")
	cmd.MarkFlagRequired("file")

	return cmd
}

func viewNamespaceCommand(cliConfig *Config) *cli.Command {
	cmd := &cli.Command{
		Use:   "view",
		Short: "View a namespace",
		Args:  cli.ExactArgs(1),
		Example: heredoc.Doc(`
			$ shield namespace view <namespace-id>
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cli.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			client, cancel, err := createClient(cmd.Context(), cliConfig.Host)
			if err != nil {
				return err
			}
			defer cancel()

			namespaceID := args[0]
			res, err := client.GetNamespace(cmd.Context(), &shieldv1beta1.GetNamespaceRequest{
				Id: namespaceID,
			})
			if err != nil {
				return err
			}

			report := [][]string{}

			namespace := res.GetNamespace()

			spinner.Stop()

			report = append(report, []string{"ID", "NAME", "CREATED AT", "UPDATED AT"})
			report = append(report, []string{
				namespace.GetId(),
				namespace.GetName(),
				namespace.GetCreatedAt().AsTime().String(),
				namespace.GetUpdatedAt().AsTime().String(),
			})
			printer.Table(os.Stdout, report)

			spinner.Stop()

			return nil
		},
	}

	return cmd
}

func listNamespaceCommand(cliConfig *Config) *cli.Command {
	cmd := &cli.Command{
		Use:   "list",
		Short: "List all namespaces",
		Args:  cli.NoArgs,
		Example: heredoc.Doc(`
			$ shield namespace list
		`),
		Annotations: map[string]string{
			"group": "core",
		},
		RunE: func(cmd *cli.Command, args []string) error {
			spinner := printer.Spin("")
			defer spinner.Stop()

			client, cancel, err := createClient(cmd.Context(), cliConfig.Host)
			if err != nil {
				return err
			}
			defer cancel()

			res, err := client.ListNamespaces(cmd.Context(), &shieldv1beta1.ListNamespacesRequest{})
			if err != nil {
				return err
			}

			report := [][]string{}
			namespaces := res.GetNamespaces()

			spinner.Stop()

			fmt.Printf(" \nShowing %d namespaces\n \n", len(namespaces))

			report = append(report, []string{"ID", "NAME", "CREATED AT", "UPDATED AT"})
			for _, n := range namespaces {
				report = append(report, []string{
					n.GetId(),
					n.GetName(),
					n.GetCreatedAt().AsTime().String(),
					n.GetUpdatedAt().AsTime().String(),
				})
			}
			printer.Table(os.Stdout, report)

			return nil
		},
	}

	return cmd
}
