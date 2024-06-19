package binglib_test

import (
	"testing"

	binglib "github.com/chengtx809/bing-lib"
)

const cookieImg = "Complete cookie"

func TestImage(t *testing.T) {
	i := binglib.NewImage(cookieImg)
	imgs, id, err := i.Image("猫")

	t.Log("id: ", id)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(imgs)
}
