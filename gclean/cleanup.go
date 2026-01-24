package gclean

type IDeferCleanup interface {
	Add(f func())
	Cleanup()
}

func NewDeferCleanup() IDeferCleanup {
	return &deferCleanup{}
}

type deferCleanup struct {
	cleanups []func()
}

func (d *deferCleanup) Add(f func()) {
	d.cleanups = append(d.cleanups, f)
}

func (d *deferCleanup) Cleanup() {
	for i := len(d.cleanups) - 1; i >= 0; i-- {
		d.cleanups[i]()
	}
}
