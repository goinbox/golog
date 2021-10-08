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

func (w *asyncWriter) asyncLogRoutine() {
	defer w.wg.Done()

	for {
		am := <-w.msgCh
		switch am.kind {
		case AsyncMsgKindWrite:
			_, _ = w.w.Write(am.msg)
		case AsyncMsgKindFlush:
			_ = w.w.Flush()
		case AsyncMsgKindFree:
			w.w.Free()
			return
		}
	}
}

func (w *asyncWriter) Write(p []byte) (int, error) {
	w.msgCh <- &asyncMsg{AsyncMsgKindWrite, p}

	return len(p), nil
}

func (w *asyncWriter) Flush() error {
	w.msgCh <- &asyncMsg{AsyncMsgKindFlush, nil}

	return nil
}

func (w *asyncWriter) Free() {
	w.msgCh <- &asyncMsg{AsyncMsgKindFree, nil}

	w.wg.Wait()
}
