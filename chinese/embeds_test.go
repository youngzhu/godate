package chinese

import "testing"

func TestEmbedSanityCheck(t *testing.T) {
	// 列出所有嵌入文件
	entries, err := LocalResources.ReadDir("data/timor")
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) == 0 {
		t.Error("没有找到任何嵌入文件")
	}

	for _, entry := range entries {
		t.Logf("找到嵌入文件: %s", entry.Name())
	}
}
