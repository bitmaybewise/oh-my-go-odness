package file

import "testing"

func Test_file_Exists(t *testing.T) {
	t.Run("given a file exists", func(t *testing.T) {
		if !Exists("file.go") {
			t.Error("file expected to exist")
		}
	})

	t.Run("given a file does not exist", func(t *testing.T) {
		if Exists("666.go") {
			t.Error("file expected not to exist")
		}
	})
}
