package views

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type MainMenu struct {
	*BaseViewer

	selectedOption int
}

func NewMainMenu(bv *BaseViewer) *MainMenu {
	mm := &MainMenu{
		BaseViewer:     bv,
		selectedOption: 0,
	}
	mm.backgroundImage = util.LoadGreeterImage()
	return mm
}

func (mm *MainMenu) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	options := mm.getMenuOptions()

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		for {
			mm.selectedOption = (mm.selectedOption + 1) % len(options)
			if options[mm.selectedOption].Enabled {
				break
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyK) {
		for {
			mm.selectedOption = (mm.selectedOption - 1 + len(options)) % len(options)
			if options[mm.selectedOption].Enabled {
				break
			}
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		selectedOption := options[mm.selectedOption]
		if selectedOption.Enabled && selectedOption.Action != nil {
			return selectedOption.Action()
		}
	}

	return nil
}

func (mm *MainMenu) Draw(screen *ebiten.Image) {
	mm.DrawBackground(screen)

	util.DrawScaledRect(screen, 0, 0, 1000, 400, shared.CLR_OVERLAY)

	util.DrawText(screen, "HLL Observer", 20, 40, shared.CLR_WHITE, util.Font.Title)

	options := mm.getMenuOptions()
	for i, option := range options {
		var clr color.Color = shared.CLR_WHITE
		if !option.Enabled {
			clr = color.Gray{Y: 128}
		} else if i == mm.selectedOption {
			clr = shared.CLR_SELECTED
		}
		util.DrawText(screen, option.Label, 50, 100+i*40, clr, util.Font.Normal)
	}

	util.DrawText(screen, "Use Arrow Keys to navigate, Enter to select", 50, 300, color.Gray{Y: 200}, util.Font.Normal)
}

func (mm *MainMenu) getMenuOptions() []MenuOption {
	replayDirConfigured := util.Config.ReplaysDirectory != ""

	return []MenuOption{
		{
			Label:   "Start",
			Enabled: true,
			Action: func() error {
				bv := NewBaseViewer(mm.ctx)
				state, _ := CreateState(bv, shared.MODE_VIEWER, nil)
				mm.ctx.TransitionTo(state)
				return nil
			},
		},
		{
			Label:   "Record",
			Enabled: replayDirConfigured,
			Action: func() error {
				bv := NewBaseViewer(mm.ctx)
				state, _ := CreateState(bv, shared.MODE_RECORD, nil)
				mm.ctx.TransitionTo(state)
				return nil
			},
		},
		{
			Label:   "Replay",
			Enabled: replayDirConfigured,
			Action: func() error {
				bv := NewBaseViewer(mm.ctx)
				mm.ctx.TransitionTo(NewReplayView(bv))
				return nil
			},
		},
		{
			Label:   "Quit",
			Enabled: true,
			Action: func() error {
				return ebiten.Termination
			},
		},
	}
}

type MenuOption struct {
	Label   string
	Enabled bool
	Action  func() error
}
