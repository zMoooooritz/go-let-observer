package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/views"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type UI struct {
	state shared.State
}

func NewUI(mode shared.PresentationMode) *UI {
	util.ScaleFactor = float32(util.Config.UIOptions.ScreenSize) / float32(shared.ROOT_SCALING_SIZE)

	ui := &UI{}
	bv := views.NewBaseViewer(ui)

	switch mode {
	case shared.MODE_NONE:
		ui.state = views.NewMainMenu(bv)
	case shared.MODE_VIEWER:
		ui.state, _ = views.CreateState(bv, shared.MODE_VIEWER, nil)
	case shared.MODE_RECORD:
		ui.state, _ = views.CreateState(bv, shared.MODE_RECORD, nil)
	case shared.MODE_REPLAY:
		ui.state = views.NewReplayView(bv)
	}

	return ui
}

func (ui *UI) TransitionTo(newState shared.State) {
	ui.state = newState
}

func (ui *UI) Update() error {
	return ui.state.Update()
}

func (ui *UI) Draw(screen *ebiten.Image) {
	ui.state.Draw(screen)
}

func (ui *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ui.state.Layout(outsideWidth, outsideHeight)
}
