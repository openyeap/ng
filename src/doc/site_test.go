package doc

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	site, _ := NewSite(".yaml")

	site.SiteMap["test"] = "test"
	fmt.Println("site map:", site.SiteMap)
	site.LoadLayout()
	site.WalkDoc(site.SiteMap["site.docs"])
}
