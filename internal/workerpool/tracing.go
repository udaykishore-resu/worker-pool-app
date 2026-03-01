package workerpool

import (
	"context"

	"go.opentelemetry.io/otel"
)

func TraceJob(ctx context.Context, name string) (context.Context, func()) {
	tracer := otel.Tracer("workerpool")

	ctx, span := tracer.Start(ctx, name)

	return ctx, span.End
}
