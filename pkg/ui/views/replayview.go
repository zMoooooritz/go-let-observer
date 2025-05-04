package views

import (
	"image/color"
	"os"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/zMoooooritz/go-let-observer/pkg/record"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type ReplayView struct {
	*BaseViewer

	replays        []string
	selectedReplay int
	visibleStart   int
	visibleCount   int
	errorMessage   string
}

func NewReplayView(bv *BaseViewer) *ReplayView {
	replayDir := util.Config.ReplaysDirectory
	replays := []string{}
	files, err := os.ReadDir(replayDir)

	if err == nil {
		for _, file := range files {
			if !file.IsDir() {
				replays = append(replays, file.Name())
			}
		}
	}

	rv := &ReplayView{
		BaseViewer:   bv,
		replays:      replays,
		visibleStart: 0,
		visibleCount: 5,
	}
	rv.backgroundImage = util.LoadGreeterImage()
	return rv
}

func (rv *ReplayView) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	if len(rv.replays) > 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyJ) {
			rv.selectedReplay = (rv.selectedReplay + 1) % len(rv.replays)
			if rv.selectedReplay >= rv.visibleStart+rv.visibleCount {
				rv.visibleStart++
			}
			if rv.selectedReplay == 0 {
				rv.visibleStart = 0
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyK) {
			rv.selectedReplay = (rv.selectedReplay - 1 + len(rv.replays)) % len(rv.replays)
			if rv.selectedReplay < rv.visibleStart {
				rv.visibleStart--
			}
			if rv.selectedReplay == len(rv.replays)-1 {
				rv.visibleStart = len(rv.replays) - rv.visibleCount
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			replay := rv.replays[rv.selectedReplay]
			bv := NewBaseViewer(rv.ctx)
			dataFetcher, err := record.NewMatchReplayer(path.Join(util.Config.ReplaysDirectory, replay))
			if err != nil {
				rv.errorMessage = "Error loading replay"
			}
			dataRecorder := record.NewNoRecorder()
			rv.ctx.TransitionTo(NewMapView(bv, dataFetcher, dataRecorder))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		bv := NewBaseViewer(rv.ctx)
		rv.ctx.TransitionTo(NewMainMenu(bv))
	}

	return nil
}

func (rv *ReplayView) Draw(screen *ebiten.Image) {
	rv.DrawBackground(screen)

	util.DrawScaledRect(screen, 0, 0, 1000, 400, shared.CLR_OVERLAY)

	util.DrawText(screen, "Select a Replay", 20, 40, shared.CLR_WHITE, util.Font.Title)

	if len(rv.replays) == 0 {
		util.DrawText(screen, "No replays found", 50, 100, color.Gray{Y: 200}, util.Font.Normal)
	} else {
		for i := rv.visibleStart; i < rv.visibleStart+rv.visibleCount && i < len(rv.replays); i++ {
			color := shared.CLR_WHITE
			if i == rv.selectedReplay {
				color = shared.CLR_SELECTED
			}
			util.DrawText(screen, rv.replays[i], 50, 100+(i-rv.visibleStart)*40, color, util.Font.Normal)
		}

		if rv.errorMessage != "" {
			util.DrawText(screen, rv.errorMessage, 50, 280, color.RGBA{255, 0, 0, 255}, util.Font.Normal)
		}

		util.DrawText(screen, "Use Arrow Keys to navigate, Enter to select", 50, 340, color.Gray{Y: 200}, util.Font.Normal)
	}
}
