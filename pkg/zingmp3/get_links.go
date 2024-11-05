package zingmp3

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bahamas0x00/zmp3/pkg/common"
)

func GetDownloadLinks(zingMp3Url string) (*SongInfo, error) {
	if !isValidRequestUrl(zingMp3Url) {
		return nil, common.InvalidZingMp3Url
	}

	httpClient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	httpRequest, err := http.NewRequest("GET", getUpStreamUrl(zingMp3Url), nil)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot query to upstream server with error code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	songInfo := &SongInfo{}

	err = json.Unmarshal(body, songInfo)
	if err != nil {
		return nil, err
	}

	return songInfo, nil
}

func isValidRequestUrl(rUrl string) bool {
	_, err := url.ParseRequestURI(rUrl)
	if err != nil {
		return false
	}
	if !strings.Contains(rUrl, Song) && !strings.Contains(rUrl, VideoClip) {
		return false
	}

	return true
}

func getUpStreamUrl(zingmp3Url string) string {
	return UpStream + zingmp3Url
}
