// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	useCommand = &cobra.Command{
		Use:              "use",
		Short:            "Use",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		RunE: func(cmd *cobra.Command, args []string) error {
			insecure, _ := cmd.Flags().GetBool("insecure")
			ca, _ := cmd.Flags().GetString("ca")
			host, _ := cmd.Flags().GetString("host")

			// Build configuration
			cfg := BuildDefaultConfig(host, insecure)

			// Add CA certificate
			if !insecure && ca != "" {
				if abs, err := filepath.Abs(ca); err == nil {
					cfg.CA = abs
				}
			}

			bytestr, err := yaml.Marshal(cfg)
			if err != nil {
				return err
			}
			str := string(bytestr)

			// Print to stdout
			fmt.Println(str)

			return nil
		},
	}
)

func init() {
	useCommand.Flags().Bool("insecure", defaultInsecure, "Connect without TLS")
	useCommand.Flags().String("host", defaultClusterHost, "Server host name")
	useCommand.Flags().String("ca", "", "Certificate file to use")
	Root.AddCommand(useCommand)
}
