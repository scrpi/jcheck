package jcheck

type rule struct {
	pattern string
	checks  []CheckFunc
}

func newRule(pattern string, checks ...CheckFunc) *rule {
	r := &rule{
		pattern: pattern,
		checks:  checks,
	}

	return r
}

func (r *rule) check(n *node) []*Result {
	var results []*Result

	if match(r.pattern, n.path) {
		for _, chk := range r.checks {
			msg, ok := chk(n)

			// Special checks
			if msg == permittedFuncMsg {
				// Propagate 'permitted' up the tree.
				p := n
				for p != nil {
					p.hasPermitRule = true
					p = p.parent
				}
				// 'Permitted()' is a pseudo-check. It's purpose is to set the
				// 'hasPermitRule' flag on a node. The actual permit check is
				// performed in the calling function.
				continue
			}

			if !ok {
				results = append(results, &Result{
					Path:    n.String(),
					Pattern: r.pattern,
					Pass:    ok,
					FailMsg: msg,
				})
			}
		}
	}

	return results
}
