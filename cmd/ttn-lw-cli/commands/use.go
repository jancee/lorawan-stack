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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	configFile = ".ttn-lw-cli.yml"
	useCommand = &cobra.Command{
		Use:              "use",
		Short:            "Use",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		RunE: func(cmd *cobra.Command, args []string) error {
			insecure, _ := cmd.Flags().GetBool("insecure")
			ca, _ := cmd.Flags().GetString("ca")
			host, _ := cmd.Flags().GetString("host")

			user, _ := cmd.Flags().GetBool("user")
			overwrite, _ := cmd.Flags().GetBool("overwrite")

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

			fname := configFile
			if user {
				dir, err := os.UserConfigDir()
				if err != nil {
					return err
				}
				fname = filepath.Join(dir, configFile)
			}

			_, err = os.Stat(fname)
			if err == nil && !overwrite {
				fmt.Printf("%s exists. Not overwriting.\n", fname)
				os.Exit(-1)
			}

			err = ioutil.WriteFile(fname, bytestr, 0644)
			if err == nil {
				fmt.Printf("Config file for %s written in %s\n", host, fname)
			}
			return err
		},
	}
)

func init() {
	useCommand.Flags().Bool("insecure", defaultInsecure, "Connect without TLS")
	useCommand.Flags().String("host", defaultClusterHost, "Server host name")
	useCommand.Flags().String("ca", "", "Certificate file to use")
	useCommand.Flags().Bool("user", false, "Write config file in user config directory")
	useCommand.Flags().Bool("overwrite", false, "Overwrite existing config files without confirmation")
	Root.AddCommand(useCommand)
}
