package lagou

import (
	"fmt"
	"testing"
)

func TestFetch_FetchURL(t *testing.T) {
	f := NewLagouFetch()
	content,err := f.FetchURL("https://www.lagou.com/jobs/4797912.html?show=c688715c31d241609f9a6cee4da9ef77")

	fmt.Println(string(content),err)
}
