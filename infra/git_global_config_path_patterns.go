package infra

import (
	"github.com/Lickability/git-remind/domain"
	"github.com/Lickability/git-remind/infra/git"
	"strings"
)

var GitGlobalConfigGetPathPatterns domain.GetPathPatterns = func() (pathPatterns []string, err error) {
	paths, err := git.GetGlobalConfig("remind.paths")
	if err != nil {
		return
	}
	pathPatterns = strings.Split(paths, ",")
	return
}
