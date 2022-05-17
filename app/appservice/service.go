package appservice

import (
	"github.com/Lickability/git-remind/app/cli/cliglobalopts"
	"github.com/Lickability/git-remind/domain"
	"github.com/Lickability/git-remind/infra"
)

var GetPathPatterns = domain.NewGetPathPatterns(
	domain.MultipleGetPathPatterns(
		cliglobalopts.GetPathPatterns,
		infra.GitGlobalConfigGetPathPatterns,
	),
)
var GetRepos = domain.NewGetRepos(GetPathPatterns, infra.FilesystemRepos)
var GetRepoStatuses = domain.NewGetRepoStatuses(GetRepos, infra.GetGitStatus)
var NotifyRepoStatues = domain.NewNotifyRepoStatuses(GetRepoStatuses, infra.BeeepRepoStatusNotifier)
