package files

import (
	//"crypto/sha1"
	"errors"
	"fmt"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/handlers/fb"
	"gitlab.mcsolutions.ru/lib/common/config"
	"io"
	//"io/ioutil"
	"mime/multipart"
	//"os"
	//"path"
	//"strings"
)

var (
	MAX_UPLOAD_SIZE = config.GetEnvInt("MAX_UPLOAD_SIZE", 0)
	FILES_DIR       = config.GetEnv("FILES_DIR", "C:/Users/User/temp/")
)

type Files struct {
	fb *fb.Fb
}

func (f *Files) Upload(srcFile multipart.File, info *multipart.FileHeader) (string, error) {
	size, err := getSize(srcFile)
	if err != nil {
		return "", err
	}
	if MAX_UPLOAD_SIZE != 0 && size > int64(MAX_UPLOAD_SIZE) {
		return "", errors.New("uploaded file size exceeds the limit")
	}
	//body, err := ioutil.ReadAll(srcFile)
	//if err != nil {
	//	return "", err
	//}
	filename := info.Filename
	//if filename == "" {
	//	filename = fmt.Sprintf("%x", sha1.Sum(body))
	//}
	if f.fb == nil {
		f.fb = &fb.Fb{}
	}
	docName := "test2/" + filename
	path, err := f.fb.Upload(docName, srcFile)
	if err != nil {
		return "", err
	}
	users, err := f.fb.GetUsers()
	if err != nil {
		return "", err
	}
	fmt.Println(*users)
	return path, nil
	//dstPath := path.Join(FILES_DIR, filename)
	//dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	//if err != nil {
	//	return "", err
	//}
	//defer dstFile.Close()
	//written, err := dstFile.Write(body)
	//if err != nil {
	//	return "", err
	//}
	//if int64(written) != size {
	//	return "", errors.New("uploaded file size and written size differ")
	//}
	//uploadedURL := strings.TrimPrefix(dstPath, FILES_DIR)
	//if !strings.HasPrefix(uploadedURL, "/") {
	//	uploadedURL = "/" + uploadedURL
	//}
	//uploadedURL = "/files" + uploadedURL
	//return uploadedURL, nil
}

func getSize(content io.Seeker) (int64, error) {
	size, err := content.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	_, err = content.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return size, nil
}
