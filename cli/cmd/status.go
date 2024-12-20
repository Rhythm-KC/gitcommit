package cmd

import (
	"fmt"
	"os"

	"github.com/Rhythm-KC/gitcommit/cli/core"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
    Use: "status",
    Short: "Get the status of your git environmemt with each entitiy being "+
            "assigned a tag. That can be use for commit",
    Long: `For example:
    if Git status output is:
    
        dirA/fileA,
        dirB/fileB,
        dirC/fileC,
    gitcommit status output is
       1 dirA/fileA,
       2 dirB/fileB,
       3 dirC/fileC,
    where 1, 2, 3 are unique Id's to each file `,
    Run: runGitCommitStatus ,
} 
func init(){
    rootCmd.AddCommand(statusCmd)
}


func runGitCommitStatus(cmd *cobra.Command, args []string){
    working_dir, _ := os.Getwd();
    files, err := gitcommit.GetFilesAvailableForCommit(working_dir)
    if err != nil{
        fmt.Fprintln(os.Stderr, err.Error());
        return;
    }
    var output string = "Files to be Commited\n";
    for index, val := range files{
        output += fmt.Sprintf("[%d] %s\n", index, val)
    }
    println(output)

}
