/*
Copyright © 2021 Mazedur Rahman <mazedur.rahman.litn@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/yamaszone/json2env/internal"
)

var cfgFile string
var exportable bool
var input string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "json2env",
	Short:   "Export JSON file content to the environment",
	Aliases: []string{"j2e"},
	Long: `Given a JSON file input.json:
{
	"ENV_VAR1": "value1",
	"ENV_VAR2": "value2"
}

Example #1: Set as variables in your shell

	$ eval $(json2env -i input.json)
	$ echo $ENV_VAR1
	value1
	$ echo $ENV_VAR2
	value2

Example #2: Export as environment variables in your shell

	$ eval $(json2env -i input.json -e)
	$ echo $ENV_VAR1
	value1
	$ echo $ENV_VAR2
	value2

Example #3: Get output in .env format

	$ json2env -i input.json
	ENV_VAR1=value1
	ENV_VAR2=value2

Example #4: Get output to source in your shell

	$ json2env -i input.json -e
	export ENV_VAR1=value1
	export ENV_VAR2=value2
`,
	RunE: toEnv,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err, "OK")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Input JSON file name")
	rootCmd.MarkFlagRequired("input")
	rootCmd.PersistentFlags().BoolVarP(&exportable, "exportable", "e", false, "Print output as 'export ENV_VAR=value'")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".json2env" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".json2env")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func toEnv(cmd *cobra.Command, args []string) error {

	f, err := cmd.Flags().GetString("input")
	if err != nil {
		return err
	}

	exportable, err = cmd.Flags().GetBool("exportable")
	if err != nil {
		return err
	}

	err = json2env.ToEnv(f, exportable)
	if err != nil {
		return err
	}

	return nil
}
