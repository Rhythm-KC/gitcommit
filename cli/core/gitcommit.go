package gitcommit

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/Rhythm-KC/gitcommit/cli/core/exceptions"
)

func GetFilesAvailableForCommit(rootDir string)([]string, error){
    if !isDirHere(rootDir){
        return nil, exceptions.NewGitException(
            fmt.Sprintf("Cound not find directory %s", rootDir),
            nil);
    }

    var baseGitPath string = path.Join(rootDir, ".git")
    if !isDirHere(baseGitPath){
        return nil, exceptions.NewGitException(
            fmt.Sprintf("Git hasen't been initlized at %s", rootDir),
            nil);
    }

    return exectueGitStatus();

}


func isDirHere(path string) bool {
    
    _, err := os.Stat(path);
    return err == nil;
}

func exectueGitStatus() ([]string, error) {
   exec.LookPath("git") 
   var cmd *exec.Cmd =  exec.Command("git", "status", "-s");
   var output []byte;
   var err error;
   output, err = cmd.Output();

   if err != nil{
       return nil, exceptions.NewGitException("Error when running git status -s",
                                              &err) 
   }

   outputStr := string(output);
   if len(outputStr) > 0 && outputStr[len(outputStr) - 1] == '\n'{
       outputStr = outputStr[0:len(outputStr)- 1]
   }

   return strings.Split(outputStr, "\n"), nil
}

func CommitGroup(rootDir string, tags[]uint, commitMsg string) error{

    gitPath := path.Join(rootDir,".git")
    if !isDirHere(gitPath){
        return exceptions.NewGitException(
            fmt.Sprintf("Git not initlized at %s", rootDir),
            nil);

    }

    files, err := exectueGitStatus()
    if err != nil{
        return err
    }
    if len(files) == 0{
        return exceptions.NewGitException("No Files to commit", nil);
    }

    validTags, err := fromTagsToFiles(&files, &tags);
    if err != nil{
        return err
    }

    err = executeGitAdd(validTags)
    if err!= nil{
        return err
    }

    err = executeGitCommit(commitMsg)
    if err != nil{ return err
    }

    return nil
}

func fromTagsToFiles(files *[]string, tags *[]uint) ([]string, error){
    validTags := []string{}
    for _, tag  := range *tags{
        if tag < 0 && int(tag) >= len(*tags){
            return nil, exceptions.NewGitException(
                fmt.Sprintf("Tag %d is not a valid tag.\n Commit Aborted!!", tag), 
                nil)
        }
        trim := strings.TrimSpace((*files)[tag])
        file:=strings.Split(trim, " ")

        validTags = append(validTags, file[len(file) -1])
    }
    return validTags, nil
}

func executeGitAdd(files []string) error{

    exec.LookPath("git")
    var cmd *exec.Cmd = exec.Command("git", append([]string{"add"}, files...)...)
    _, err := cmd.Output()
    if err != nil{
       return exceptions.NewGitException(
                           fmt.Sprintf("Error when running git add %v", files),
                           &err) 
    }
    return nil
}

func executeGitCommit(commitMsg string) error{

    exec.LookPath("git")
    var cmd *exec.Cmd = exec.Command("git", "commit", "-m", commitMsg)
    err := cmd.Run()
    if err != nil{
       return exceptions.NewGitException(
                fmt.Sprintf("Error when running git commit -m %s",commitMsg),
                &err) 
    }
    return nil
}




