/*
Copyright © 2018-2022 blacktop

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/apex/log"
	"github.com/blacktop/go-macho"
	"github.com/blacktop/ipsw/pkg/ctf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	kernelcacheCmd.AddCommand(ctfdumpCmd)

	ctfdumpCmd.Flags().StringP("arch", "a", "", "Which architecture to use for fat/universal MachO")
	ctfdumpCmd.Flags().BoolP("pretty", "", false, "Pretty print JSON")
	ctfdumpCmd.Flags().BoolP("json", "j", false, "Output as JSON")
	viper.BindPFlag("kernel.ctfdump.arch", ctfdumpCmd.Flags().Lookup("arch"))
	viper.BindPFlag("kernel.ctfdump.pretty", ctfdumpCmd.Flags().Lookup("pretty"))
	viper.BindPFlag("kernel.ctfdump.json", ctfdumpCmd.Flags().Lookup("json"))
	ctfdumpCmd.MarkZshCompPositionalArgumentFile(1)
}

// ctfdumpCmd represents the ctfdump command
var ctfdumpCmd = &cobra.Command{
	Use:   "ctfdump",
	Short: "Dump CTF info",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		var m *macho.File

		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		// flags
		selectedArch := viper.GetString("kernel.ctfdump.arch")
		prettyJSON := viper.GetBool("kernel.ctfdump.pretty")
		outAsJSON := viper.GetBool("kernel.ctfdump.json")

		machoPath := filepath.Clean(args[0])

		if _, err := os.Stat(machoPath); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", machoPath)
		}

		// first check for fat file
		fat, err := macho.OpenFat(machoPath)
		if err != nil && err != macho.ErrNotFat {
			return err
		}

		if err == macho.ErrNotFat {
			m, err = macho.Open(machoPath)
			if err != nil {
				return err
			}
		} else {
			var options []string
			var shortOptions []string
			for _, arch := range fat.Arches {
				options = append(options, fmt.Sprintf("%s, %s", arch.CPU, arch.SubCPU.String(arch.CPU)))
				shortOptions = append(shortOptions, strings.ToLower(arch.SubCPU.String(arch.CPU)))
			}

			if len(selectedArch) > 0 {
				found := false
				for i, opt := range shortOptions {
					if strings.Contains(strings.ToLower(opt), strings.ToLower(selectedArch)) {
						m = fat.Arches[i].File
						found = true
						break
					}
				}
				if !found {
					return fmt.Errorf("--arch '%s' not found in: %s", selectedArch, strings.Join(shortOptions, ", "))
				}

			} else {
				choice := 0
				prompt := &survey.Select{
					Message: "Detected a universal MachO file, please select an architecture to analyze:",
					Options: options,
				}
				survey.AskOne(prompt, &choice)
				m = fat.Arches[choice].File
			}
		}

		c, err := ctf.Parse(m)
		if err != nil {
			return err
		}

		ids := make([]int, 0, len(c.Types))
		for id := range c.Types {
			ids = append(ids, id)
		}
		sort.Ints(ids)

		if len(args) > 1 {
			for _, id := range ids {
				if c.Types[id].Name() == args[1] {
					fmt.Println(c.Types[id])
					break
				}
			}
		} else { // DUMP
			if outAsJSON {
				var b []byte

				if prettyJSON {
					b, err = json.MarshalIndent(c, "", "    ")
					if err != nil {
						return fmt.Errorf("failed to marshal function as JSON: %v", err)
					}
				} else {
					b, err = json.Marshal(c)
					if err != nil {
						return fmt.Errorf("failed to marshal function as JSON: %v", err)
					}
				}

				cwd, _ := os.Getwd()
				log.Infof("Creating %s", filepath.Join(cwd, "ctfdump.json"))
				if err := ioutil.WriteFile("ctfdump.json", b, 0660); err != nil {
					return err
				}
			} else {
				fmt.Printf("- CTF Header -----------------------------------------------------------------\n\n")
				fmt.Println(c.Header)
				fmt.Printf("\n- Types ----------------------------------------------------------------------\n\n")
				for _, id := range ids {
					fmt.Println(c.Types[id].Dump())
				}
				fmt.Printf("\n- Data Objects ---------------------------------------------------------------\n\n")
				for _, g := range c.Globals {
					fmt.Println(g)
				}
				fmt.Printf("\n- Functions ------------------------------------------------------------------\n\n")
				for _, f := range c.Functions {
					fmt.Println(f)
				}
			}
		}

		return nil
	},
}
