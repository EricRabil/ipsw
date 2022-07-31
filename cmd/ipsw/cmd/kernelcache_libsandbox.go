/*
Copyright © 2022 blacktop

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
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/apex/log"
	"github.com/blacktop/ipsw/pkg/dyld"
	"github.com/blacktop/ipsw/pkg/sandbox"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	kernelcacheCmd.AddCommand(libsandboxCmd)
	libsandboxCmd.Flags().BoolP("json", "j", false, "Output to stdout as JSON")
	libsandboxCmd.Flags().StringP("output", "o", "", "Folder to write JSON")
	libsandboxCmd.MarkZshCompPositionalArgumentFile(1, "dyld_shared_cache*")
}

// libsandboxCmd represents the libsandbox command
var libsandboxCmd = &cobra.Command{
	Use:           "libsandbox",
	Short:         "🚧 [WIP] Get libsandbox data",
	Args:          cobra.MinimumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	Hidden:        true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		// flags
		asJSON, _ := cmd.Flags().GetBool("json")
		output, _ := cmd.Flags().GetString("output")

		dscPath := filepath.Clean(args[0])

		fileInfo, err := os.Lstat(dscPath)
		if err != nil {
			return fmt.Errorf("file %s does not exist", dscPath)
		}

		// Check if file is a symlink
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			symlinkPath, err := os.Readlink(dscPath)
			if err != nil {
				return errors.Wrapf(err, "failed to read symlink %s", dscPath)
			}
			// TODO: this seems like it would break
			linkParent := filepath.Dir(dscPath)
			linkRoot := filepath.Dir(linkParent)

			dscPath = filepath.Join(linkRoot, symlinkPath)
		}

		f, err := dyld.Open(dscPath)
		if err != nil {
			return err
		}
		defer f.Close()

		fi, err := sandbox.GetFilterInfo(f)
		if err != nil {
			return err
		}

		mi, err := sandbox.GetModifierInfo(f)
		if err != nil {
			return err
		}

		oi, err := sandbox.GetOperationInfo(f)
		if err != nil {
			return err
		}

		if asJSON {
			dat, err := json.Marshal(sandbox.LibSandbox{
				Operations: oi,
				Filters:    fi,
				Modifiers:  mi,
			})
			if err != nil {
				return fmt.Errorf("failed to marshal json: %w", err)
			}

			if len(output) > 0 {
				fpath := filepath.Join(output, fmt.Sprintf("libsandbox_%s.gz", f.Headers[f.UUID].OsVersion.String()))
				log.Infof("Creating %s", fpath)
				var buf bytes.Buffer
				gz := gzip.NewWriter(&buf)
				gz.Write(dat)
				gz.Close()
				if err := ioutil.WriteFile(fpath, buf.Bytes(), 0755); err != nil {
					return err
				}
			} else {
				fmt.Println(string(dat))
			}
		} else {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
			fmt.Println("Filter Info")
			fmt.Println("===========")
			for _, f := range fi {
				if f.ArgumentType != sandbox.SB_VALUE_TYPE_NONE {
					fmt.Fprintf(w, "%d)\t%s\t%s\ttype: %d\tcost: %#x\flags: %#x\taliases: %d\n",
						f.ID,
						f.Name,
						f.Category,
						f.ArgumentType,
						f.CostFactor,
						f.Flags,
						len(f.Aliases))
					for _, a := range f.Aliases {
						if a.Unknown != 0 {
							fmt.Fprintf(w, "    - %d) %s\t%d\n", a.ID, a.Name, a.Unknown)
						} else {
							fmt.Fprintf(w, "    - %d) %s\n", a.ID, a.Name)
						}
					}
				}
			}
			w.Flush()

			fmt.Println()
			fmt.Println("Modifier Info")
			fmt.Println("=============")
			for _, modifier := range mi {
				if modifier.ID > 0 {
					fmt.Fprintf(w, "%02d) %-20s\taction: %d\ttype: %d\tflag: %-4d\tmask: %-4d\n",
						modifier.ID,
						modifier.Name,
						modifier.Action,
						modifier.Type,
						modifier.ActionFlag,
						modifier.ActionMask,
					)
					for _, a := range modifier.Aliases {
						if a.Unknown != 0 {
							fmt.Fprintf(w, "    - %d) %s\t%d\n", a.ID, a.Name, a.Unknown)
						} else {
							fmt.Fprintf(w, "    - %d) %s\n", a.ID, a.Name)
						}
					}
				}
			}
			w.Flush()

			fmt.Println()
			fmt.Println("Operation Names")
			fmt.Println("===============")
			for idx, o := range oi {
				if o.NodeType == sandbox.OPERATION_NODE_TYPE_TERMINAL {
					fmt.Fprintf(w, "%3d) %s\tTERMINAL\tjump_op: %d\taction: %d\tmsg_filt_ops: %v\n", idx, o.Name, o.JumpTargetOperation, o.Action, o.MsgFilterOps)
				} else {
					fmt.Fprintf(w, "%3d) %s\tNON_TERMINAL\tjump_op: %d\taction: %d\tmsg_filt_ops: %v\n", idx, o.Name, o.JumpTargetOperation, o.Action, o.MsgFilterOps)
				}
			}
			w.Flush()
		}

		return nil
	},
}
