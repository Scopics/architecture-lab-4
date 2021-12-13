package engine

import "sync"

type commandQueue struct {
	sync.Mutex

	cmds      []Command
	noCommand bool
	pushed    chan struct{}
}

func (cq *commandQueue) push(cmd Command) {
	cq.Lock()
	defer cq.Unlock()

	cq.cmds = append(cq.cmds, cmd)

	if cq.noCommand {
		cq.noCommand = false
		cq.pushed <- struct{}{}
	}
}

func (cq *commandQueue) pull() Command {
	cq.Lock()
	defer cq.Unlock()
	if len(cq.cmds) == 0 {
		cq.noCommand = true
		<-cq.pushed
	}

	cmd := cq.cmds[0]
	cq.cmds[0] = nil
	cq.cmds = cq.cmds[1:]

	return cmd
}

type EventLoop struct {
	queue        *commandQueue
	awaitFinish  bool
	finishSignal chan struct{}
}

func (el *EventLoop) Start() {
	el.queue = &commandQueue{
		pushed: make(chan struct{}),
	}
	el.finishSignal = make(chan struct{})

	go func() {
		for !(el.awaitFinish && len(el.queue.cmds) == 0) {
			cmd := el.queue.pull()
			cmd.Execute(el)
		}
		el.finishSignal <- struct{}{}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	el.queue.push(cmd)
}

type cmdExecutor struct {
	executor func()
}

func (ce *cmdExecutor) Execute(h Handler) {
	ce.executor()
}

func (el *EventLoop) AwaitFinish() {
	finish := &cmdExecutor{func() { el.awaitFinish = false }}
	el.Post(finish)

	<-el.finishSignal
}
