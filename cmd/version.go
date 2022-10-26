/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Branch    = "main"
	Author    = "uscan team"
	Email     = "<uscanteam@163.com>"
	Date      = "2022-10-25"
	Commit    = "821288f"
	GoVersion = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version info",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("branch: ", Branch)
		fmt.Println("author: ", Author)
		fmt.Println("email: ", Email)
		fmt.Println("date: ", Date)
		fmt.Println("git commit: ", Commit)
		fmt.Println("go verion: ", GoVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd represents the version command
