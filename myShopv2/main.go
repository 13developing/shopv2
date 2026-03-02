package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "myShopv2/internal/packed"

	_ "myShopv2/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"myShopv2/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
