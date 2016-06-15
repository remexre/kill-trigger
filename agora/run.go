package agora

import (
	"github.com/PuerkitoBio/agora/compiler"
	"github.com/PuerkitoBio/agora/runtime"
	"github.com/PuerkitoBio/agora/runtime/stdlib"
)

// Run runs agora code.
func Run(code string) (interface{}, error) {
	r := NewAggregateResolver(MapResolver{
		"code": code,
	})
	r.Add("pastebin", NewPastebinResolver())

	ctx := runtime.NewCtx(r, new(compiler.Compiler))

	ctx.RegisterNativeModule(new(stdlib.FilepathMod))
	ctx.RegisterNativeModule(new(stdlib.FmtMod))
	ctx.RegisterNativeModule(new(stdlib.MathMod))
	ctx.RegisterNativeModule(new(stdlib.OsMod))
	ctx.RegisterNativeModule(new(stdlib.StringsMod))
	ctx.RegisterNativeModule(new(stdlib.TimeMod))

	ctx.RegisterNativeModule(new(HTTPMod))

	mod, err := ctx.Load("code")
	if err != nil {
		return "", err
	}

	val, err := mod.Run(nil)
	if err != nil {
		return "", err
	}

	switch v := val.(type) {
	case runtime.String:
		return string(v), nil
	case runtime.Number:
		return float64(v), nil
	case runtime.Bool:
		return bool(v), nil
	default:
		return val, nil
	}
}
