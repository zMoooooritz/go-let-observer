package views

import (
	"github.com/zMoooooritz/go-let-loose/pkg/rcon"
	"github.com/zMoooooritz/go-let-observer/pkg/rcndata"
	"github.com/zMoooooritz/go-let-observer/pkg/record"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func CreateState(bv *BaseViewer, targetMode shared.PresentationMode, rconCreds *rcon.ServerConfig) (shared.State, error) {
	if rconCreds == nil {
		creds := util.Config.GetServerCredentials()
		rconCreds = &creds
	}
	rcn, rcnErr := rcon.NewRcon(*rconCreds, shared.RCON_WORKER_COUNT)
	if rcnErr == nil {
		dataFetcher := rcndata.NewRconDataFetcher(rcn)

		var dataRecorder record.DataRecorder
		if targetMode == shared.MODE_VIEWER {
			dataRecorder = record.NewNoRecorder()
		} else {
			currMap, _ := rcn.GetCurrentMap()
			dataRecorder, _ = record.NewMatchRecorder(util.Config.ReplaysDirectory, currMap)
		}

		return NewMapView(bv, dataFetcher, dataRecorder), nil
	}
	return NewLoginView(bv, targetMode), rcnErr
}
