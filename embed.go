package neng

import "embed"

//go:embed res/adj
//go:embed res/adv
//go:embed res/noun
//go:embed res/verb
//go:embed res/adj.irr
//go:embed res/adj.suf
//go:embed res/adj.ncmp
//go:embed res/noun.irr
//go:embed res/noun.unc
//go:embed res/verb.irr
var efs embed.FS
