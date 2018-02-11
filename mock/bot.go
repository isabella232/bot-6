// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/andviro/middleware"
	"sync"
)

var (
	lockBotMockClose    sync.RWMutex
	lockBotMockCommand  sync.RWMutex
	lockBotMockContext  sync.RWMutex
	lockBotMockHears    sync.RWMutex
	lockBotMockMessage  sync.RWMutex
	lockBotMockNext     sync.RWMutex
	lockBotMockOn       sync.RWMutex
	lockBotMockReply    sync.RWMutex
	lockBotMockRun      sync.RWMutex
	lockBotMockText     sync.RWMutex
	lockBotMockUse      sync.RWMutex
	lockBotMockUserName sync.RWMutex
)

// BotMock is a mock implementation of Bot.
//
//     func TestSomethingThatUsesBot(t *testing.T) {
//
//         // make and configure a mocked Bot
//         mockedBot := &BotMock{
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             CommandFunc: func(in1 string) middleware.Predicate {
// 	               panic("TODO: mock out the Command method")
//             },
//             ContextFunc: func(in1 context.Context, in2 middleware.Handler) error {
// 	               panic("TODO: mock out the Context method")
//             },
//             HearsFunc: func(in1 string) middleware.Predicate {
// 	               panic("TODO: mock out the Hears method")
//             },
//             MessageFunc: func(in1 context.Context) bool {
// 	               panic("TODO: mock out the Message method")
//             },
//             NextFunc: func() bool {
// 	               panic("TODO: mock out the Next method")
//             },
//             OnFunc: func(in1 middleware.Predicate, in2 middleware.Handler)  {
// 	               panic("TODO: mock out the On method")
//             },
//             ReplyFunc: func(in1 context.Context, in2 interface{}) error {
// 	               panic("TODO: mock out the Reply method")
//             },
//             RunFunc: func() error {
// 	               panic("TODO: mock out the Run method")
//             },
//             TextFunc: func(in1 context.Context) string {
// 	               panic("TODO: mock out the Text method")
//             },
//             UseFunc: func(in1 ...middleware.Middleware)  {
// 	               panic("TODO: mock out the Use method")
//             },
//             UserNameFunc: func(in1 context.Context) string {
// 	               panic("TODO: mock out the UserName method")
//             },
//         }
//
//         // TODO: use mockedBot in code that requires Bot
//         //       and then make assertions.
//
//     }
type BotMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// CommandFunc mocks the Command method.
	CommandFunc func(in1 string) middleware.Predicate

	// ContextFunc mocks the Context method.
	ContextFunc func(in1 context.Context, in2 middleware.Handler) error

	// HearsFunc mocks the Hears method.
	HearsFunc func(in1 string) middleware.Predicate

	// MessageFunc mocks the Message method.
	MessageFunc func(in1 context.Context) bool

	// NextFunc mocks the Next method.
	NextFunc func() bool

	// OnFunc mocks the On method.
	OnFunc func(in1 middleware.Predicate, in2 middleware.Handler)

	// ReplyFunc mocks the Reply method.
	ReplyFunc func(in1 context.Context, in2 interface{}) error

	// RunFunc mocks the Run method.
	RunFunc func() error

	// TextFunc mocks the Text method.
	TextFunc func(in1 context.Context) string

	// UseFunc mocks the Use method.
	UseFunc func(in1 ...middleware.Middleware)

	// UserNameFunc mocks the UserName method.
	UserNameFunc func(in1 context.Context) string

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Command holds details about calls to the Command method.
		Command []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Context holds details about calls to the Context method.
		Context []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 middleware.Handler
		}
		// Hears holds details about calls to the Hears method.
		Hears []struct {
			// In1 is the in1 argument value.
			In1 string
		}
		// Message holds details about calls to the Message method.
		Message []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// Next holds details about calls to the Next method.
		Next []struct {
		}
		// On holds details about calls to the On method.
		On []struct {
			// In1 is the in1 argument value.
			In1 middleware.Predicate
			// In2 is the in2 argument value.
			In2 middleware.Handler
		}
		// Reply holds details about calls to the Reply method.
		Reply []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 interface{}
		}
		// Run holds details about calls to the Run method.
		Run []struct {
		}
		// Text holds details about calls to the Text method.
		Text []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
		// Use holds details about calls to the Use method.
		Use []struct {
			// In1 is the in1 argument value.
			In1 []middleware.Middleware
		}
		// UserName holds details about calls to the UserName method.
		UserName []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
	}
}

// Close calls CloseFunc.
func (mock *BotMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: BotMock.CloseFunc is nil but Bot.Close was just called")
	}
	callInfo := struct {
	}{}
	lockBotMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockBotMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedBot.CloseCalls())
func (mock *BotMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockBotMockClose.RLock()
	calls = mock.calls.Close
	lockBotMockClose.RUnlock()
	return calls
}

// Command calls CommandFunc.
func (mock *BotMock) Command(in1 string) middleware.Predicate {
	if mock.CommandFunc == nil {
		panic("moq: BotMock.CommandFunc is nil but Bot.Command was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockBotMockCommand.Lock()
	mock.calls.Command = append(mock.calls.Command, callInfo)
	lockBotMockCommand.Unlock()
	return mock.CommandFunc(in1)
}

// CommandCalls gets all the calls that were made to Command.
// Check the length with:
//     len(mockedBot.CommandCalls())
func (mock *BotMock) CommandCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockBotMockCommand.RLock()
	calls = mock.calls.Command
	lockBotMockCommand.RUnlock()
	return calls
}

// Context calls ContextFunc.
func (mock *BotMock) Context(in1 context.Context, in2 middleware.Handler) error {
	if mock.ContextFunc == nil {
		panic("moq: BotMock.ContextFunc is nil but Bot.Context was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 middleware.Handler
	}{
		In1: in1,
		In2: in2,
	}
	lockBotMockContext.Lock()
	mock.calls.Context = append(mock.calls.Context, callInfo)
	lockBotMockContext.Unlock()
	return mock.ContextFunc(in1, in2)
}

// ContextCalls gets all the calls that were made to Context.
// Check the length with:
//     len(mockedBot.ContextCalls())
func (mock *BotMock) ContextCalls() []struct {
	In1 context.Context
	In2 middleware.Handler
} {
	var calls []struct {
		In1 context.Context
		In2 middleware.Handler
	}
	lockBotMockContext.RLock()
	calls = mock.calls.Context
	lockBotMockContext.RUnlock()
	return calls
}

// Hears calls HearsFunc.
func (mock *BotMock) Hears(in1 string) middleware.Predicate {
	if mock.HearsFunc == nil {
		panic("moq: BotMock.HearsFunc is nil but Bot.Hears was just called")
	}
	callInfo := struct {
		In1 string
	}{
		In1: in1,
	}
	lockBotMockHears.Lock()
	mock.calls.Hears = append(mock.calls.Hears, callInfo)
	lockBotMockHears.Unlock()
	return mock.HearsFunc(in1)
}

// HearsCalls gets all the calls that were made to Hears.
// Check the length with:
//     len(mockedBot.HearsCalls())
func (mock *BotMock) HearsCalls() []struct {
	In1 string
} {
	var calls []struct {
		In1 string
	}
	lockBotMockHears.RLock()
	calls = mock.calls.Hears
	lockBotMockHears.RUnlock()
	return calls
}

// Message calls MessageFunc.
func (mock *BotMock) Message(in1 context.Context) bool {
	if mock.MessageFunc == nil {
		panic("moq: BotMock.MessageFunc is nil but Bot.Message was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	lockBotMockMessage.Lock()
	mock.calls.Message = append(mock.calls.Message, callInfo)
	lockBotMockMessage.Unlock()
	return mock.MessageFunc(in1)
}

// MessageCalls gets all the calls that were made to Message.
// Check the length with:
//     len(mockedBot.MessageCalls())
func (mock *BotMock) MessageCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	lockBotMockMessage.RLock()
	calls = mock.calls.Message
	lockBotMockMessage.RUnlock()
	return calls
}

// Next calls NextFunc.
func (mock *BotMock) Next() bool {
	if mock.NextFunc == nil {
		panic("moq: BotMock.NextFunc is nil but Bot.Next was just called")
	}
	callInfo := struct {
	}{}
	lockBotMockNext.Lock()
	mock.calls.Next = append(mock.calls.Next, callInfo)
	lockBotMockNext.Unlock()
	return mock.NextFunc()
}

// NextCalls gets all the calls that were made to Next.
// Check the length with:
//     len(mockedBot.NextCalls())
func (mock *BotMock) NextCalls() []struct {
} {
	var calls []struct {
	}
	lockBotMockNext.RLock()
	calls = mock.calls.Next
	lockBotMockNext.RUnlock()
	return calls
}

// On calls OnFunc.
func (mock *BotMock) On(in1 middleware.Predicate, in2 middleware.Handler) {
	if mock.OnFunc == nil {
		panic("moq: BotMock.OnFunc is nil but Bot.On was just called")
	}
	callInfo := struct {
		In1 middleware.Predicate
		In2 middleware.Handler
	}{
		In1: in1,
		In2: in2,
	}
	lockBotMockOn.Lock()
	mock.calls.On = append(mock.calls.On, callInfo)
	lockBotMockOn.Unlock()
	mock.OnFunc(in1, in2)
}

// OnCalls gets all the calls that were made to On.
// Check the length with:
//     len(mockedBot.OnCalls())
func (mock *BotMock) OnCalls() []struct {
	In1 middleware.Predicate
	In2 middleware.Handler
} {
	var calls []struct {
		In1 middleware.Predicate
		In2 middleware.Handler
	}
	lockBotMockOn.RLock()
	calls = mock.calls.On
	lockBotMockOn.RUnlock()
	return calls
}

// Reply calls ReplyFunc.
func (mock *BotMock) Reply(in1 context.Context, in2 interface{}) error {
	if mock.ReplyFunc == nil {
		panic("moq: BotMock.ReplyFunc is nil but Bot.Reply was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 interface{}
	}{
		In1: in1,
		In2: in2,
	}
	lockBotMockReply.Lock()
	mock.calls.Reply = append(mock.calls.Reply, callInfo)
	lockBotMockReply.Unlock()
	return mock.ReplyFunc(in1, in2)
}

// ReplyCalls gets all the calls that were made to Reply.
// Check the length with:
//     len(mockedBot.ReplyCalls())
func (mock *BotMock) ReplyCalls() []struct {
	In1 context.Context
	In2 interface{}
} {
	var calls []struct {
		In1 context.Context
		In2 interface{}
	}
	lockBotMockReply.RLock()
	calls = mock.calls.Reply
	lockBotMockReply.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *BotMock) Run() error {
	if mock.RunFunc == nil {
		panic("moq: BotMock.RunFunc is nil but Bot.Run was just called")
	}
	callInfo := struct {
	}{}
	lockBotMockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	lockBotMockRun.Unlock()
	return mock.RunFunc()
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//     len(mockedBot.RunCalls())
func (mock *BotMock) RunCalls() []struct {
} {
	var calls []struct {
	}
	lockBotMockRun.RLock()
	calls = mock.calls.Run
	lockBotMockRun.RUnlock()
	return calls
}

// Text calls TextFunc.
func (mock *BotMock) Text(in1 context.Context) string {
	if mock.TextFunc == nil {
		panic("moq: BotMock.TextFunc is nil but Bot.Text was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	lockBotMockText.Lock()
	mock.calls.Text = append(mock.calls.Text, callInfo)
	lockBotMockText.Unlock()
	return mock.TextFunc(in1)
}

// TextCalls gets all the calls that were made to Text.
// Check the length with:
//     len(mockedBot.TextCalls())
func (mock *BotMock) TextCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	lockBotMockText.RLock()
	calls = mock.calls.Text
	lockBotMockText.RUnlock()
	return calls
}

// Use calls UseFunc.
func (mock *BotMock) Use(in1 ...middleware.Middleware) {
	if mock.UseFunc == nil {
		panic("moq: BotMock.UseFunc is nil but Bot.Use was just called")
	}
	callInfo := struct {
		In1 []middleware.Middleware
	}{
		In1: in1,
	}
	lockBotMockUse.Lock()
	mock.calls.Use = append(mock.calls.Use, callInfo)
	lockBotMockUse.Unlock()
	mock.UseFunc(in1...)
}

// UseCalls gets all the calls that were made to Use.
// Check the length with:
//     len(mockedBot.UseCalls())
func (mock *BotMock) UseCalls() []struct {
	In1 []middleware.Middleware
} {
	var calls []struct {
		In1 []middleware.Middleware
	}
	lockBotMockUse.RLock()
	calls = mock.calls.Use
	lockBotMockUse.RUnlock()
	return calls
}

// UserName calls UserNameFunc.
func (mock *BotMock) UserName(in1 context.Context) string {
	if mock.UserNameFunc == nil {
		panic("moq: BotMock.UserNameFunc is nil but Bot.UserName was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	lockBotMockUserName.Lock()
	mock.calls.UserName = append(mock.calls.UserName, callInfo)
	lockBotMockUserName.Unlock()
	return mock.UserNameFunc(in1)
}

// UserNameCalls gets all the calls that were made to UserName.
// Check the length with:
//     len(mockedBot.UserNameCalls())
func (mock *BotMock) UserNameCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	lockBotMockUserName.RLock()
	calls = mock.calls.UserName
	lockBotMockUserName.RUnlock()
	return calls
}