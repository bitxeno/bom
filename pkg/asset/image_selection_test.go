package asset

import (
	"image"
	"testing"
)

func TestPickBestImageCandidatePrefersLargest(t *testing.T) {
	candidates := []imageCandidate{
		newImageCandidate("Icon-Small@2x.png", 58, 58, 1, 2),
		newImageCandidate("Icon-Small-60@3x.png", 180, 180, 1, 3),
		newImageCandidate("1024x1024.png", 1024, 1024, 6, 1),
	}

	best, err := pickBestImageCandidate(candidates, ImageOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if best == nil {
		t.Fatal("expected a best image candidate")
	}
	if best.renditionName != "1024x1024.png" {
		t.Fatalf("unexpected best candidate: got %s", best.renditionName)
	}
}

func TestPickBestImageCandidateFiltersByIdiomAndScale(t *testing.T) {
	candidates := []imageCandidate{
		newImageCandidate("Icon-Small-40@2x.png", 80, 80, 1, 2),
		newImageCandidate("Icon-Small-60@3x.png", 180, 180, 1, 3),
		newImageCandidate("Icon-83.5@2x.png", 167, 167, 2, 2),
	}

	best, err := pickBestImageCandidate(candidates, ImageOptions{Idiom: "pad", Scale: 2})
	if err != nil {
		t.Fatal(err)
	}
	if best == nil {
		t.Fatal("expected a filtered image candidate")
	}
	if best.renditionName != "Icon-83.5@2x.png" {
		t.Fatalf("unexpected filtered candidate: got %s", best.renditionName)
	}
}

func TestPickBestImageCandidateRejectsUnknownIdiom(t *testing.T) {
	_, err := pickBestImageCandidate([]imageCandidate{newImageCandidate("Icon.png", 60, 60, 1, 2)}, ImageOptions{Idiom: "desktop"})
	if err == nil {
		t.Fatal("expected an error for unknown idiom")
	}
}

func newImageCandidate(name string, width, height int, idiom, scale uint16) imageCandidate {
	return imageCandidate{
		renditionName: name,
		image:         image.NewRGBA(image.Rect(0, 0, width, height)),
		attrs: RenditionAttrs{
			kRenditionAttributeType_Idiom: uint16hex(idiom),
			kRenditionAttributeType_Scale: uint16hex(scale),
		},
	}
}