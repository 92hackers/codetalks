/**

context

*/

package utils

import (
	"context"
	"time"
)

func WithTimeoutCtxSeconds(seconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
}

func WithTimeoutCtxMilliSeconds(milliSeconds int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(milliSeconds)*time.Millisecond)
}
