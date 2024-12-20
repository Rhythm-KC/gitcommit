package cmd

import (
	"fmt"
	"os"

	"github.com/Rhythm-KC/gitcommit/cli/core"
	"github.com/spf13/cobra"
)

var commitCmd= &cobra.Command{
    Use: "commit",
    Short: "Use this to add tags to a single commit.",
    Long: "long update",
    Run: runCommitCmd,
} 

var tags []uint;
var commitMessage string;

func init(){
    commitCmd.Flags().UintSliceVarP(&tags, "tags", "t", []uint{}, 
                                    "comma separated values for tags")

    commitCmd.Flags().StringVarP(&commitMessage,"message", "m", "", 
                                 "Message for the commit")

    commitCmd.MarkFlagRequired("message")
    commitCmd.MarkFlagRequired("tags")

    rootCmd.AddCommand(commitCmd)
}


func runCommitCmd(cmd *cobra.Command, args []string){

    working_dir, _ := os.Getwd();
    err := gitcommit.CommitGroup(working_dir, tags, commitMessage);

    if err != nil{
        fmt.Fprintln(os.Stderr, err.Error());
        return;
    }
}
