package stand

import (
	"testing"
)

func TestLibreOffice(t *testing.T) {
	result, err := ConvertWithLibreOffice("C:\\Users\\zhang\\.openclaw\\workspace\\板报稿区知识竞赛成绩简报.docx")
	if err != nil {
		t.Errorf("Error converting file: %v", err)
	}
	t.Logf("Conversion result: %s", result)
}
