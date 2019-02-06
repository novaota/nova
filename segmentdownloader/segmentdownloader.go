package segmentdownloader

import (
	"log"
	"path"
	"path/filepath"

	"nova/shared"
	"nova/updatepackage"
)

type DownloadChangedCallback interface {
	DownloadChanged(current int, max int)
}

func NewSegmentDownloader(signService updatepackage.SignService) *SegmentDownloader {
	return &SegmentDownloader{
		signService: signService,
	}
}

type SegmentDownloader struct {
	DownloadChangedCallback
	signService updatepackage.SignService
}

func (sd *SegmentDownloader) DownloadPackage(deployment updatepackage.UpdatePackageDeployment, destDir string) []updatepackage.Segment {
	segments := deployment.Segments

	max := len(segments)

	var result []updatepackage.Segment

	for i, segment := range segments {
		url := segment.Path
		_, filename := filepath.Split(url)
		dest := path.Join(destDir, filename)
		sd.DownloadFile(url, dest)

		ok, err := sd.signService.IsSigned(segment.Signature, dest)

		if !ok || err != nil {
			log.Fatalf("Signature for segment %v failed. Aborting", segment.Path)
			return nil
		}

		sd.OnDownloadChanged(i, max)

		downloadedSegment := updatepackage.Segment{
			Index:i,
			Path:dest,
			Signature:segment.Signature,
		}

		result = append(result, downloadedSegment)
	}
	return result
}

func (sd *SegmentDownloader) DownloadFile(src string, dest string) {
	shared.DefaultHTTPUtil.Download(src, dest)
}

func (sd *SegmentDownloader) OnDownloadChanged(current int, max int) {
	if sd.DownloadChangedCallback == nil {
		return
	}

	sd.DownloadChangedCallback.DownloadChanged(current, max)
}