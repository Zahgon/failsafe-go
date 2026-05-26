package policytesting

import (
	"github.com/failsafe-go/failsafe-go/cachepolicy"
)

type TestCache[R any] struct {
	Cache map[string]R
}

func (c *TestCache[R]) Get(key string) (R, bool) { _ = "STUB: not implemented"; return *new(R), false }

func (c *TestCache[R]) Set(key string, value R) { _ = "STUB: not implemented"; return }

func NewCache[R any]() (map[string]R, cachepolicy.Cache[R]) {
	_ = "STUB: not implemented"
	return nil, nil
}
