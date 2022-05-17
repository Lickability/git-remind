package infra

import (
	"github.com/gen2brain/beeep"
)

var BeeepRepoStatusNotifier = &beeepRepoStatusNotifier{}

type beeepRepoStatusNotifier struct {
}

func (*beeepRepoStatusNotifier) NotifyNeedToCommit(path string) (err error) {
	return exec.Command(osa, "-e", `display notification "`+path+`" with title "Remind to git commit" sound name "default"`).Run()
}

func (*beeepRepoStatusNotifier) NotifyNeedToPush(path string) (err error) {
	return exec.Command(osa, "-e", `display notification "`+path+`" with title "Remind to git push" sound name "default"`).Run()
}

func (*beeepRepoStatusNotifier) NotifyNeedToCommitAndPush(path string) (err error) {
	return exec.Command(osa, "-e", `display notification "`+path+`" with title "Remind to git commit/push" sound name "default"`).Run()
}
