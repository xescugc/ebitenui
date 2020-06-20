package widget

import (
	"image/color"
	"testing"

	"github.com/blizzy78/ebitenui/event"
	"github.com/matryer/is"
)

func TestTabBook_Tab_Initial(t *testing.T) {
	is := is.New(t)

	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.WithTabs(tab1, tab2),
		TabBookOpts.WithTabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			is.Fail() // event fired without previous action
		}))

	is.Equal(tb.Tab(), tab1)
}

func TestTabBook_SetTab(t *testing.T) {
	is := is.New(t)

	var eventArgs *TabBookTabSelectedEventArgs
	numEvents := 0

	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.WithTabs(tab1, tab2),
		TabBookOpts.WithTabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	tb.SetTab(tab2)
	event.FireDeferredEvents()

	is.Equal(tb.Tab(), tab2)
	is.Equal(eventArgs.Tab, tab2)
	is.Equal(eventArgs.PreviousTab, tab1)

	tb.SetTab(tab2)
	event.FireDeferredEvents()
	is.Equal(numEvents, 1)
}

func TestTabBook_TabSelectedEvent_User(t *testing.T) {
	is := is.New(t)

	var eventArgs *TabBookTabSelectedEventArgs
	numEvents := 0

	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.WithTabs(tab1, tab2),
		TabBookOpts.WithTabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	leftMouseButtonClick(tabBookButtons(tb)[1], t)

	is.Equal(tb.Tab(), tab2)
	is.Equal(eventArgs.Tab, tab2)
	is.Equal(eventArgs.PreviousTab, tab1)

	leftMouseButtonClick(tabBookButtons(tb)[1], t)
	is.Equal(numEvents, 1)
}

func newTabBook(t *testing.T, opts ...TabBookOpt) *TabBook {
	t.Helper()

	tb := NewTabBook(append(opts, []TabBookOpt{
		TabBookOpts.WithTabButtonImage(&ButtonImage{
			Idle: newNineSliceEmpty(t),
		}, &ButtonImage{
			Idle: newNineSliceEmpty(t),
		}),
		TabBookOpts.WithTabButtonText(loadFont(t), &ButtonTextColor{
			Idle:     color.Transparent,
			Disabled: color.Transparent,
		}),
	}...)...)

	event.FireDeferredEvents()
	render(tb, t)
	return tb
}

func tabBookButtons(t *TabBook) []*StateButton {
	buttons := []*StateButton{}
	for _, tab := range t.tabs {
		buttons = append(buttons, t.tabToButton[tab])
	}
	return buttons
}
