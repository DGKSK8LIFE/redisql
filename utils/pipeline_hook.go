package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/DGKSK8LIFE/redisql/utils/logging"
)

type PipelineHook struct {
	//Log bool
}

func (p PipelineHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	//if p.Log {
		//fmt.Printf("Pipeline starting processing\n")
	//}
	return CTX, nil
}

func (p PipelineHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	logging.LogResultList(cmds, 3)
	return nil
}

// Required to implement a hook
func (p PipelineHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error){
	return CTX, nil
}

// Required to implement a hook
func (p PipelineHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	return nil
}


