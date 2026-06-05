package assets

import (
	"embed"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

//go:embed fonts/* icon/* image/* roles/* spawns/* tacmaps/*
var Assets embed.FS

func ToFileName(mapIdentifier hll.MapIdentifier) string {
	switch mapIdentifier {
	case hll.MAP_STMEREEGLISE:
		return "sme"
	case hll.MAP_STMARIEDUMONT:
		return "smdm"
	case hll.MAP_UTAHBEACH:
		return "utah"
	case hll.MAP_OMAHABEACH:
		return "omaha"
	case hll.MAP_PURPLEHEARTLANE:
		return "phl"
	case hll.MAP_CARENTAN:
		return "carentan"
	case hll.MAP_HURTGENFOREST:
		return "hurtgen"
	case hll.MAP_HILL400:
		return "hill400"
	case hll.MAP_FOY:
		return "foy"
	case hll.MAP_KURSK:
		return "kursk"
	case hll.MAP_SMOLENSK:
		return "smolensk"
	case hll.MAP_STALINGRAD:
		return "stalingrad"
	case hll.MAP_REMAGEN:
		return "remagen"
	case hll.MAP_KHARKOV:
		return "kharkov"
	case hll.MAP_DRIEL:
		return "driel"
	case hll.MAP_ELALAMEIN:
		return "elalamein"
	case hll.MAP_MORTAIN:
		return "mortain"
	case hll.MAP_ELSENBORNRIDGE:
		return "elsenborn"
	case hll.MAP_TOBRUK:
		return "tobruk"
	case hll.MAP_JUNOBEACH:
		return "juno"
	default:
		return ""
	}
}
