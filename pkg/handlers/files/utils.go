package files

import "io"

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
