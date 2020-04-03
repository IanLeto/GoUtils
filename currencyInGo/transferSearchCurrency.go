package currencyInGo

import (
	"context"
	"sync"
	"time"
	"transfer/esb"
	"transfer/logging"
)

type CallBackFuncType func(task Task)

type Task struct {
	BizID int // 业务id
	//Topo  []CCSearchBizInstTopoResponseInfo // topo 信息（暂时以int 代替）

	page  int       // 已经查询了多少条
	Start int       // 当次需要开始查询的内容， added by jackyliang 用于替换page，字段更加明确
	Limit int       // 防止特殊情况 的限制，暂时没用
	count int       // 这个业务一共多少台主机
	time  time.Time //  超时 还没想好怎么写
}

type TaskManager struct {
	JobQueue     chan Task
	tokenBucket  chan interface{}
	maxWorker    int
	ctx          context.Context
	cancel       context.CancelFunc
	wg           sync.WaitGroup
	isRunning    bool
	taskList     []Task
	jobWg        sync.WaitGroup
	callBackFunc CallBackFuncType
}

func NewTaskManage(ctx context.Context, maxWorker int, callBackFunc CallBackFuncType, task []Task) (*TaskManager, error) {
	var (
		logger     = logging.GetStdLogger()
		newCtx     context.Context
		cancelFunc context.CancelFunc
	)
	if maxWorker > esb.MaxWorkerConfig || maxWorker <= 0 {
		logger.Infof("maxWorker->[%d] is less than 0 or more than max, max->[%d] will set.", maxWorker, esb.MaxWorkerConfig)
		maxWorker = esb.MaxWorkerConfig
	}
	newCtx, cancelFunc = context.WithCancel(ctx)

	return &TaskManager{
		ctx:          newCtx,
		cancel:       cancelFunc,
		callBackFunc: callBackFunc,
		maxWorker:    maxWorker,
		taskList:     task,
	}, nil
}

func (m *TaskManager) Start() error {
	var logger = logging.GetStdLogger()
	if m.isRunning {

	}
	m.init()
	go func() {
		defer func() {
			m.isRunning = false
		}()
		for {
			select {
			case task := <-m.JobQueue:
				logger.Tracef("job 执行前，要先获取token")
				<-m.tokenBucket
				m.wg.Add(1)
				go func(taskInfo Task) {
					defer func() {
						m.wg.Done()
						m.tokenBucket <- struct{}{}
						m.jobWg.Done()
					}()
				}(task)
			case <-m.ctx.Done():
				return

			}
		}
	}()
	go func() {
		defer func() {
			m.jobWg.Done()

		}()
		for i := 0; i < len(m.taskList); {
			select {
			case m.JobQueue <- m.taskList[i]:

				m.jobWg.Add(1)
				i++
			case <-m.ctx.Done():

			}
		}
	}()
	return nil
}

func (m *TaskManager) init() {
	//var logger = logging.GetStdLogger()
	if m.isRunning {
		panic(1)
	}
	m.tokenBucket = make(chan interface{}, m.maxWorker)
	m.JobQueue = make(chan Task, m.maxWorker)
	for i := 0; i < m.maxWorker; i++ {
		m.tokenBucket <- struct{}{}
	}
	m.jobWg = sync.WaitGroup{}
	return
}
