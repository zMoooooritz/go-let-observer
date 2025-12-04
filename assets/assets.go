package assets

import (
	"embed"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

//go:embed fonts/* roles/* spawns/* tacmaps/* image/*
var Assets embed.FS

func ToFileName(gameMap hll.Map) string {
	switch gameMap {
	case hll.MP_STMEREEGLISE:
		return "sme"
	case hll.MP_STMARIEDUMONT:
		return "smdm"
	case hll.MP_UTAHBEACH:
		return "utah"
	case hll.MP_OMAHABEACH:
		return "omaha"
	case hll.MP_PURPLEHEARTLANE:
		return "phl"
	case hll.MP_CARENTAN:
		return "carentan"
	case hll.MP_HURTGENFOREST:
		return "hurtgen"
	case hll.MP_HILL400:
		return "hill400"
	case hll.MP_FOY:
		return "foy"
	case hll.MP_KURSK:
		return "kursk"
	case hll.MP_SMOLENSK:
		return "smolensk"
	case hll.MP_STALINGRAD:
		return "stalingrad"
	case hll.MP_REMAGEN:
		return "remagen"
	case hll.MP_KHARKOV:
		return "kharkov"
	case hll.MP_DRIEL:
		return "driel"
	case hll.MP_ELALAMEIN:
		return "elalamein"
	case hll.MP_MORTAIN:
		return "mortain"
	case hll.MP_ELSENBORNRIDGE:
		return "elsenborn"
	case hll.MP_TOBRUK:
		return "tobruk"
	default:
		return ""
	}
}
