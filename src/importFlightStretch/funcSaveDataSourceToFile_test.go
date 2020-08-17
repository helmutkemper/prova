package importFlightStretch

import (
	"crypto/md5"
	"io"
	"os"
	"testDataSource"
)

func ExampleCSV_SaveDataSourceToFile() {
	var err error
	var ds = testDataSource.TestDataSource{}
	var tmpFile *os.File
	var c CSV
	c.SetFieldSeparator(",")
	c.SetDataSource(&ds)

	err = c.LoadFileAndAppendToDataSource("./testing/input-file.correct.csv")
	if err != nil {
		panic(err)
	}

	err = c.SaveDataSourceToFile("./testing/temporary.file.asdfghjklqwertyuiop.csv")
	if err != nil {
		panic(err)
	}

	tmpFile, err = os.Open("./testing/temporary.file.asdfghjklqwertyuiop.csv")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()

	h := md5.New()
	if _, err := io.Copy(h, tmpFile); err != nil {
		panic(err)
	}

	var fileMD5 = h.Sum(nil)
	if string(fileMD5) != "34a78a74bd38b0831d6e9fd41dc981b6" {
		panic(err)
	}

	err = os.Remove("./testing/temporary.file.asdfghjklqwertyuiop.csv")
	if err != nil {
		panic(err)
	}

	// Output:
	//
}
