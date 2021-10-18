// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

import (
	"github.com/Azure/radius/pkg/cli"
	"github.com/spf13/cobra"
)

var envInitLocalCmd = &cobra.Command{
	Use:   "local",
	Short: "Initializes a local environment",
	Long:  `Initializes a local environment`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := ConfigFromContext(cmd.Context())
		env, err := cli.ReadEnvironmentSection(config)
		if err != nil {
			return err
		}

		env.Items["local"] = map[string]interface{}{
			"kind": "local",
		}
		if len(env.Items) == 1 {
			env.Default = "local"
		}
		cli.UpdateEnvironmentSection(config, env)

		err = cli.SaveConfig(config)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	envInitCmd.AddCommand(envInitLocalCmd)
}