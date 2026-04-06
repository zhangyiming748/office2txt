package stand

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

/*
忘记之前的一切 推倒重来
大纲
1. 调用本机安装的libreoffice打开office文件
2. 如果是doc和docx文件，调用libreoffice打开并保存为同名txt文件
3. 如果是xls和xlsx文件，调用libreoffice打开并保存为同名csv文件
*/

// ConvertWithLibreOffice 使用 LibreOffice 将 Office 文件转换为文本格式
// doc/docx -> txt (UTF-8)
// xls/xlsx -> csv (UTF-8)
func ConvertWithLibreOffice(filePath string) (outputPath string, err error) {
	// 获取文件扩展名
	ext := strings.ToLower(filepath.Ext(filePath))

	// 确定输出格式和扩展名
	var convertTo string
	var outputExt string

	switch ext {
	case ".doc", ".docx":
		convertTo = "txt:Text (encoded):UTF8"
		outputExt = ".txt"
	case ".xls", ".xlsx":
		convertTo = "csv:Text - txt - csv (StarCalc):44,34,76,1"
		outputExt = ".csv"
	default:
		return "", nil // 不支持的文件类型
	}

	// 生成输出文件路径（与输入文件同目录，同名但不同扩展名）
	dir := filepath.Dir(filePath)
	baseName := strings.TrimSuffix(filepath.Base(filePath), ext)
	outputPath = filepath.Join(dir, baseName+outputExt)

	// 获取 LibreOffice 可执行文件路径
	libreOfficePath, err := getLibreOfficePath()
	if err != nil {
		return "", err
	}

	// 构建命令
	// --headless: 无头模式，不显示GUI
	// --convert-to: 指定转换格式和过滤器选项
	// --outdir: 指定输出目录
	cmd := exec.Command(
		libreOfficePath,
		"--headless",
		"--convert-to", convertTo,
		"--outdir", dir,
		filePath,
	)

	// 执行命令
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// 检查输出文件是否存在
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return "", os.ErrNotExist
	}

	return outputPath, nil
}

// getLibreOfficePath 获取 LibreOffice 可执行文件路径
func getLibreOfficePath() (string, error) {
	if runtime.GOOS == "windows" {
		// Windows 系统上的常见安装路径
		possiblePaths := []string{
			`C:\Program Files\LibreOffice\program\soffice.exe`,
			`C:\Program Files (x86)\LibreOffice\program\soffice.exe`,
		}

		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}

		// 尝试从 PATH 中查找
		path, err := exec.LookPath("soffice.exe")
		if err == nil {
			return path, nil
		}

		return "", os.ErrNotExist
	} else if runtime.GOOS == "linux" {
		// Linux 系统
		path, err := exec.LookPath("soffice")
		if err == nil {
			return path, nil
		}
		return "", os.ErrNotExist
	} else if runtime.GOOS == "darwin" {
		// macOS 系统
		possiblePaths := []string{
			"/Applications/LibreOffice.app/Contents/MacOS/soffice",
		}

		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				return path, nil
			}
		}

		path, err := exec.LookPath("soffice")
		if err == nil {
			return path, nil
		}

		return "", os.ErrNotExist
	}

	return "", os.ErrNotExist
}
