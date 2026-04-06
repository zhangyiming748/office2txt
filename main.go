package main

import (
	"fmt"
	"os"

	"office2txt/core"

	"github.com/spf13/cobra"
)

/*
main函数 作为一个cobra实现的cli命令行工具 主命令为 o2t 参数为 -i --input 对应函数ConvertWithLibreOffice(filePath string)
*/

func main() {
	var inputPath string

	var rootCmd = &cobra.Command{
		Use:   "o2t",
		Short: "Office to Text - 将 Office 文件转换为文本格式",
		Long:  "Office to Text 是一个使用 LibreOffice 将 Office 文档（doc, docx, xls, xlsx）转换为文本格式（txt, csv）的工具。",
		Run: func(cmd *cobra.Command, args []string) {
			if inputPath == "" {
				fmt.Println("错误: 请使用 -i 或 --input 指定输入文件路径")
				cmd.Help()
				os.Exit(1)
			}

			// 检查文件是否存在
			if _, err := os.Stat(inputPath); os.IsNotExist(err) {
				fmt.Printf("错误: 文件不存在: %s\n", inputPath)
				os.Exit(1)
			}

			fmt.Printf("正在转换文件: %s\n", inputPath)
			outputPath, err := core.ConvertWithLibreOffice(inputPath)
			if err != nil {
				fmt.Printf("转换失败: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("转换成功!\n输出文件: %s\n", outputPath)
		},
	}

	// 添加命令行参数
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "输入文件路径 (支持 .doc, .docx, .xls, .xlsx)")
	rootCmd.MarkFlagRequired("input")

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
