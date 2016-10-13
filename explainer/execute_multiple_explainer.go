package explainer

import (
	"sync"

	c "github.com/arrizalamin/codexplainer/context"
)

func ExecuteExplainers(ctx *c.Context) func(...c.ExplainerFunc) {
	return func(explainers ...c.ExplainerFunc) {
		for _, explainer := range explainers {
			explainer(ctx)
		}
	}
}

func Concurrent(explainers ...c.ExplainerFunc) c.ExplainerFunc {
	return func(ctx *c.Context) {
		var wg sync.WaitGroup
		for _, e := range explainers {
			wg.Add(1)
			go func(explainer c.ExplainerFunc) {
				defer wg.Done()
				explainer(ctx)
			}(e)
		}
		wg.Wait()
	}
}
