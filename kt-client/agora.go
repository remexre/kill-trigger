package main

import (
	"github.com/PuerkitoBio/agora/compiler"
	"github.com/PuerkitoBio/agora/runtime"
	"github.com/PuerkitoBio/agora/runtime/stdlib"
)

func agora(m map[string]string) error {
	ctx := runtime.NewCtx(MapResolver{
		"code": m["code"],
	}, new(compiler.Compiler))

	ctx.RegisterNativeModule(new(stdlib.FilepathMod))
	ctx.RegisterNativeModule(new(stdlib.FmtMod))
	ctx.RegisterNativeModule(new(stdlib.MathMod))
	ctx.RegisterNativeModule(new(stdlib.OsMod))
	ctx.RegisterNativeModule(new(stdlib.StringsMod))
	ctx.RegisterNativeModule(new(stdlib.TimeMod))

	ctx.RegisterNativeModule(new(HTTPMod))

	mod, err := ctx.Load("code")
	if err != nil {
		return err
	}

	_, err = mod.Run(nil)
	return err
}
