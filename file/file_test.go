package file

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	var datas = []struct {
		write []byte
		want  []byte
	}{
		{[]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus condimentum mi et mattis accumsan. Nunc feugiat ligula in purus auctor aliquam. Vestibulum tellus purus, fringilla at vulputate sit amet, consectetur eu quam. Nullam at commodo lorem, non auctor nisl. Integer efficitur tempus interdum. Donec sagittis aliquam sem non pellentesque. Nullam finibus est ipsum, eu iaculis nibh fermentum a. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam porta venenatis accumsan."),
			[]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus condimentum mi et mattis accumsan. Nunc feugiat ligula in purus auctor aliquam. Vestibulum tellus purus, fringilla at vulputate sit amet, consectetur eu quam. Nullam at commodo lorem, non auctor nisl. Integer efficitur tempus interdum. Donec sagittis aliquam sem non pellentesque. Nullam finibus est ipsum, eu iaculis nibh fermentum a. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam porta venenatis accumsan.")},
	}

	homeDir := HomeDir()
	for i, data := range datas {
		filePath := filepath.Join(homeDir, "test_data", fmt.Sprintf("s_%d", i))
		if err := WriteDataToFile(filePath, data.write); err != nil {
			t.Fatal(err)
		}

		out, err := FetchDataFromFile(filePath)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(out, data.want) {
			t.Logf("unexpect result, want %b, got %b", data.want, out)
			t.FailNow()
		}

	}

	os.RemoveAll(filepath.Join(homeDir, "test_data"))

}
