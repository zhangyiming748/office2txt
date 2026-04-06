package stand
import (
	"os"
	"github.com/nguyenthenguyen/docx"
)
func ReadDocx(filePath string) ( textBytes []byte, err error) {
	_, err = os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// 2. 打开docx并提取文本
	r, err := docx.ReadDocxFile(filePath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// 3. 获取纯文本
	doc := r.Editable()
	text := doc.GetContent() // 提取所有文本
	textBytes = []byte(text)

	return textBytes, nil
}
