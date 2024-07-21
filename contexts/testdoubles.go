package contexts

import "testing"

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (spy SpyStore) assertWasCancelled() {
	spy.t.Helper()
	if !spy.cancelled {
		spy.t.Errorf("store was not told to canel")
	}
}

func (spy SpyStore) assertWasNotCancelled() {
	spy.t.Helper()

	if spy.cancelled {
		spy.t.Errorf("store was told to cancel")
	}
}
