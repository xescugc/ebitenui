package widget

import (
	"image"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A Text lays out its glyphs once and draws the same ones every frame
// until its label changes: laying out every frame was most of a screen's
// text cost (worst on the web).
func TestText_GlyphsLaidOutOnce(t *testing.T) {
	for _, bb := range []bool{false, true} {
		tw := NewText(TextOpts.Text("Hello [color=ff0000]world[/color]\nsecond", loadFont(t), color.White), TextOpts.ProcessBBCode(bb))
		tw.Validate()
		tw.SetLocation(image.Rect(0, 0, 300, 100))
		screen := ebiten.NewImage(300, 100)
		tw.Render(screen)
		require.NotEmpty(t, tw.measurements.glyphs, "bbcode=%v", bb)
		first := &tw.measurements.glyphs[0][0].glyphs[0]
		tw.Render(screen)
		assert.Same(t, first, &tw.measurements.glyphs[0][0].glyphs[0], "same glyphs on the next frame (bbcode=%v)", bb)

		tw.Label = "changed"
		tw.Render(screen)
		assert.NotSame(t, first, &tw.measurements.glyphs[0][0].glyphs[0], "a new label lays out again (bbcode=%v)", bb)
	}
}
