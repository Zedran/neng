package neng

import "embed"

//go:embed res/adj
//go:embed res/adv
//go:embed res/noun
//go:embed res/verb
//go:embed res/adj.irr
//go:embed res/adj.nc
//go:embed res/noun.irr
//go:embed res/verb.irr
var efs embed.FS
