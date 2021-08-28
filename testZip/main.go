package testZip

import (
	"archive/zip"
	"io"
	"os"
)

func ToZip(src string, dest string) (err error) {
	sFile, err := os.Open(src)
	if err != nil {
		return err
	}

	os.RemoveAll(dest)
	dFile, err := os.Create(dest)
	defer dFile.Close()
	if err != nil {
		return err
	}

	w := zip.NewWriter(dFile)
	defer w.Close()

	err = compress(sFile, "", w)
	if err != nil {
		return err
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//func main() {
//	err := ToZip("pages", "./1.zip")
//	if err != nil {
//		fmt.Println(err)
//	}
//}
