/*
Copyright © 2019 mutl3y

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
package cmd

import (
	"fmt"
	"github.com/mutl3y/prtg_dns/sensor"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prtg_dns",
	Short: "simple dns resolve test for remote nodes",
	Long: `
simple dns resolve test for remote nodes using prtg

Examples:
	prtg_dns-windows-amd64.exe -a www.facebook.com,www.google.com -t 200ms
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		a, err := flags.GetStringSlice("addr")
		if err != nil {
			fmt.Println(err)
		}

		t, err := flags.GetDuration("timeout")
		if err != nil {
			fmt.Println(err)
		}
		err = sensor.PrtgLookup(a, t)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("addr", "a", []string{"www.google.com", "www.facebook.com"}, "up to 50 addresses")
	rootCmd.PersistentFlags().DurationP("timeout", "t", 500*time.Millisecond, "timeout string eg 500ms")
}