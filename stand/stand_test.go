package stand

import "testing"

func TestDocx(t *testing.T) {
	b,e:=ReadDocx("C:\\Users\\zhang\\.openclaw\\workspace\\板报稿区知识竞赛成绩简报.docx")
	t.Logf("b = %v\te = %v\n",string(b),e)
	s:=Clean(b)
	t.Logf("清洗后的文本 = %v\n",string(s)) // 打印清洗后的文本
}
