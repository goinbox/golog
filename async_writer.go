package golog

import "sync"

const (
	AsyncMsgKindWrite = 1
	AsyncMsgKindFlush = 2
	AsyncMsgKindFree  = 3
)

type asyncMsg struct {
	kind int
	msg  []byte
}

type asyncWriter struct {
	w Writer

	msgCh chan *asyncMsg
	wg    *sync.WaitGroup
}

func NewAsyncWriter(w Writer, queueSize int) *asyncWriter {
	a := &asyncWriter{
		w: w,

		msgCh: make(chan *asyncMsg, queueSize),
		wg:    new(sync.WaitGroup),
	}

	go a.asyncLogRoutine()
	a.wg.Add(1)

	return a
}

func (a *asyncWriter) asyncLogRoutine() {
	defer a.wg.Done()

	for {
		am := <-a.msgCh
		switch am.kind {
		case AsyncMsgKindWrite:
			_, _ = a.w.Write(am.msg)
		case AsyncMsgKindFlush:
			_ = a.w.Flush()
		case AsyncMsgKindFree:
			a.w.Free()
			return
		}
	}
}

func (a *asyncWriter) Write(p []byte) (int, error) {
	a.msgCh <- &asyncMsg{AsyncMsgKindWrite, p}

	return len(p), nil
}

func (a *asyncWriter) Flush() error {
	a.msgCh <- &asyncMsg{AsyncMsgKindFlush, nil}

	return nil
}

func (a *asyncWriter) Free() {
	a.msgCh <- &asyncMsg{AsyncMsgKindFree, nil}

	a.wg.Wait()
}
