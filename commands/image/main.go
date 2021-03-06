package image

import (
	"github.com/ItsClairton/Anny/base"
	"github.com/ItsClairton/Anny/utils/constants"
)

var Category = &base.Category{
	ID:       "image",
	Emote:    constants.PEPEART,
	Commands: []*base.Command{&CatCommand, &NekoCommand},
}
