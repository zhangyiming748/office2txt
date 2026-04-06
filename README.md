# office2txt

一个基于 LibreOffice 的命令行工具，用于将 Office 文档转换为文本格式。支持将 Word 文档转换为 TXT，Excel 表格转换为 CSV，所有输出文件均为 UTF-8 编码。

## 功能特性

- ✅ 支持 `.doc` 和 `.docx` 格式转换为 `.txt`（UTF-8 编码）
- ✅ 支持 `.xls` 和 `.xlsx` 格式转换为 `.csv`（UTF-8 编码）
- ✅ 自动检测并转换各种编码为 UTF-8
- ✅ 跨平台支持（Windows、Linux、macOS）
- ✅ 简洁的命令行界面
- ✅ 无需外部依赖（除 LibreOffice 外）

## 系统要求

### 必需软件

**LibreOffice** - 必须预先安装在系统中

#### Windows
- 下载地址：https://www.libreoffice.org/download/download/
- 默认安装路径：
  - `C:\Program Files\LibreOffice\program\soffice.exe`
  - `C:\Program Files (x86)\LibreOffice\program\soffice.exe`

#### Linux
```bash
# Ubuntu/Debian
sudo apt-get install libreoffice

# CentOS/RHEL/Fedora
sudo yum install libreoffice
```

#### macOS
- 下载地址：https://www.libreoffice.org/download/download/
- 或使用 Homebrew：
```bash
brew install --cask libreoffice
```

## 安装方法

### 方法一：下载预编译版本

从 [Releases](https://github.com/yourusername/office2txt/releases) 页面下载对应平台的二进制文件。

### 方法二：从源码编译

```bash
# 克隆仓库
git clone https://github.com/yourusername/office2txt.git
cd office2txt

# 构建
go build -o o2t

# 或者使用 GoReleaser
goreleaser build --snapshot --clean
```

## 使用方法

### 基本用法

```bash
# 转换 Word 文档
./o2t -i document.docx

# 转换 Excel 表格
./o2t -i spreadsheet.xlsx

# 使用长参数
./o2t --input document.docx
```

### 命令行参数

```
Usage:
  o2t [flags]

Flags:
  -h, --help           显示帮助信息
  -i, --input string   输入文件路径 (支持 .doc, .docx, .xls, .xlsx)
```

### 示例

```bash
# Windows
o2t.exe -i "C:\Documents\报告.docx"
o2t.exe -i "D:\Data\数据表.xlsx"

# Linux/macOS
./o2t -i /home/user/documents/report.docx
./o2t -i /home/user/data/spreadsheet.xlsx
```

## 输出说明

### Word 文档转换
- 输入：`.doc` 或 `.docx`
- 输出：同目录下的 `.txt` 文件（UTF-8 编码）
- 示例：`document.docx` → `document.txt`

### Excel 表格转换
- 输入：`.xls` 或 `.xlsx`
- 输出：同目录下的 `.csv` 文件（UTF-8 编码）
- 示例：`data.xlsx` → `data.csv`
- CSV 格式：逗号分隔，双引号包裹文本

## 注意事项

1. **LibreOffice 必须安装**：工具依赖系统安装的 LibreOffice 进行文件转换
2. **首次运行可能较慢**：LibreOffice 首次启动需要初始化，后续转换会更快
3. **文件权限**：确保对输入文件有读取权限，对输出目录有写入权限
4. **大文件处理**：转换大型文件可能需要较长时间，请耐心等待
5. **编码问题**：无论输入文件是什么编码，输出都是 UTF-8 编码
6. **不支持的文件类型**：如果输入不支持的文件类型，工具会静默退出

## 故障排除

### 找不到 LibreOffice

**Windows**: 确保 LibreOffice 安装在默认路径，或将 `soffice.exe` 添加到系统 PATH

**Linux/macOS**: 运行 `which soffice` 检查是否在 PATH 中

### 转换失败

1. 检查文件是否存在且可访问
2. 确认文件格式是否正确（`.doc`, `.docx`, `.xls`, `.xlsx`）
3. 查看错误信息进行排查

### 中文乱码

输出文件应该是 UTF-8 编码。如果使用文本编辑器打开时出现乱码，请确保编辑器使用 UTF-8 编码打开文件。

## 开发

### 项目结构

```
office2txt/
├── core/              # 核心转换逻辑
│   ├── libreoffice.go # LibreOffice 转换实现
│   ├── clean.go       # XML 清理工具
│   └── unit_test.go   # 单元测试
├── main.go            # CLI 入口
├── go.mod             # Go 模块定义
└── .goreleaser.yml    # GoReleaser 配置
```

### 构建发布版本

```bash
# 安装 GoReleaser
go install github.com/goreleaser/goreleaser/v2@latest

# 构建所有平台
goreleaser release --snapshot --clean
```

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 贡献

欢迎提交 Issue 和 Pull Request！

## 致谢

- [LibreOffice](https://www.libreoffice.org/) - 强大的开源办公套件
- [Cobra](https://github.com/spf13/cobra) - Go 语言的 CLI 库
- [GoReleaser](https://goreleaser.com/) - Go 项目的发布工具
