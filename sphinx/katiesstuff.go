package sphinx

import (
	"github.com/xlab/pocketsphinx-go/pocketsphinx"
)

func (d *Decoder) SetSearch(search string) {
	pocketsphinx.SetSearch(d.dec, search)
}
