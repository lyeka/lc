#!/bin/bash

# 指定目录和Markdown文件名
target_directory="doc/content"
markdown_file="README.md"
count=0

# 清空已存在的Markdown文件内容
> "$markdown_file"

# 递归遍历目录及其子目录下的所有文件名，并写入Markdown文件
echo "## 题目索引" >> "$markdown_file"
echo "" >> "$markdown_file"
# 定义递归函数
list_files() {
    local current_directory="$1"
    local indent="$2"

    for file in "$current_directory"/*; do
        if [ -f "$file" ]; then
            filename=$(basename "$file" .md)  # 移除 .md 后缀
            prefix=$(echo "$filename" | grep -Eo '^[0-9]+')  # 提取数字前缀
            echo "$indent- [$filename]($file)/[Code]($prefix/$filename.go)" >> "$markdown_file"
            count=$((count + 1))
        elif [ -d "$file" ]; then
            dirname=$(basename "$file")
            list_files "$file" "$indent  "
        fi
    done
}

list_files "$target_directory" ""
echo "" >> "$markdown_file"
echo "Total: $count of 150" >> "$markdown_file"

echo "Markdown文件已生成：$markdown_file"
