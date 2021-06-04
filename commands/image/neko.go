package image

import (
	"math/rand"
	"strings"

	"github.com/ItsClairton/Anny/base"
	"github.com/ItsClairton/Anny/services/image"
	"github.com/ItsClairton/Anny/utils/Emotes"
)

var NekoCommand = base.Command{
	Name: "neko", Description: "Manda uma imagem aleatoria de uma neko",
	Handler: func(ctx *base.CommandContext) {

		gif := ctx.Args != nil && strings.HasPrefix(ctx.Args[0], "gif") || rand.Float32() < 0.2

		var url string
		var err error

		if gif {
			url, err = image.GetFromNekos("ngif")
		} else {
			if rand.Float32() < 0.5 {
				url, err = image.GetFromNekoBot("neko")
			} else {
				url, err = image.GetFromNekos("neko")
			}
		}

		if err != nil {
			ctx.Reply(Emotes.MIKU_CRY, "Um erro ocorreu ao procurar por uma neko, desculpa.")
		} else {
			ctx.Send(url)
		}
	},
}
