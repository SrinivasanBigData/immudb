/*
Copyright 2019-2020 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	c "github.com/codenotary/immudb/cmd"
	"github.com/codenotary/immudb/cmd/immuadmin/commands"
	"github.com/spf13/cobra"
)

var o = &c.Options{}

func init() {
	cobra.OnInitialize(func() { o.InitConfig("immuadmin") })
}

func main() {
	cmd := &cobra.Command{
		Use:   "immuadmin",
		Short: "Admin client for the ImmuDB tamperproof database",
		Long: `Admin client for the ImmuDB tamperproof database.

Environment variables:
  IMMUADMIN_ADDRESS=127.0.0.1
  IMMUADMIN_PORT=3322
  IMMUADMIN_MTLS=true`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		DisableAutoGenTag: true,
	}
	commands.Init(cmd, o)
	if err := cmd.Execute(); err != nil {
		c.QuitToStdErr(err)
	}
}