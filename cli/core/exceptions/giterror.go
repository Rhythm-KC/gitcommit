package exceptions

import (
	"fmt"

)

const git_error_perfix = "[GIT ERROR]:"

type Giterror struct{
    message string
    base *error
}

func (e Giterror) Error() string{
    return e.message;
}
func NewGitException(errMsg string, baseError *error) error{ 
    return &Giterror{
        message:fmt.Sprintf("%s %s", git_error_perfix, errMsg),
        base: baseError,
    }
}
