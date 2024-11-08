package zingmp3

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/bahamas0x00/zmp3/pkg/common"
	"github.com/cheggaaa/pb/v3"
)

func Download(obj *DownloadObject, downloadFolder string) error {
	if obj == nil {
		return common.InvalidDownloadObject
	}
	resp, err := http.Get(obj.DownloadUrl)
	if err != nil {
		return err
	}

	length := int(resp.ContentLength)
	defer resp.Body.Close()

	fileName := fmt.Sprintf("%s-%s.%s", obj.Title, obj.Artist, obj.Type)

	out, err := os.Create(path.Join(downloadFolder, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	bar := pb.New(length)
	bar.Set("prefix", fmt.Sprintf("[%s]", fileName))
	bar.SetCurrent(0)
	bar.SetWidth(80)
	bar.Start()

	rd := bar.NewProxyReader(resp.Body)
	_, err = io.Copy(out, rd)
	if err != nil {
		return err
	}

	bar.Finish()
	return nil
}
