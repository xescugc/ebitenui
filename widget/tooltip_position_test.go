package widget

import (
	"image"
	"testing"

	"github.com/ebitenui/ebitenui/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matryer/is"
)

// A tooltip taller than the room both below and above the cursor must
// still stay inside the window (an item tooltip on an inventory's top row).
func TestToolTip_TallContentStaysOnScreen(t *testing.T) {
	is := is.New(t)
	input.Draw(ebiten.NewImage(1920, 1080))

	tt := &ToolTip{
		Position:                TOOLTIP_POS_CURSOR_FOLLOW,
		ContentOriginVertical:   TOOLTIP_ANCHOR_START,
		ContentOriginHorizontal: TOOLTIP_ANCHOR_START,
		Offset:                  image.Point{10, 10},
	}
	p := tt.processContentPosition(image.Point{1700, 404}, 400, 900, image.Rectangle{})

	is.True(p.Y >= 0)
	is.True(p.Y+900 <= 1080)
	is.True(p.X >= 0)
	is.True(p.X+400 <= 1920)
}
