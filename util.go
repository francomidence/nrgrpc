package nrgrpc

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type ctxKeyClientSegment struct{}

func setSegment(ctx context.Context, seg *newrelic.Segment) context.Context {
	return context.WithValue(ctx, ctxKeyClientSegment{}, seg)
}

func getSegment(ctx context.Context) (st *newrelic.Segment, ok bool) {
	if v := ctx.Value(ctxKeyClientSegment{}); v != nil {
		st, ok = v.(*newrelic.Segment)
	}
	return
}
