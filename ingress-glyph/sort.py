def sort_lines_in_file(input_file_path, output_file_path):
    try:
        with open(input_file_path, 'r', encoding='utf-8') as file:
            lines = file.readlines()

        result = []
        for line in lines:
            result.append(line.title())

        result.sort()
        with open(output_file_path, 'w', encoding='utf-8') as file:
            file.writelines(result)
        
    except Exception as e:
        print("处理文件时出错:", e)

# 使用示例
input_path = 'static/glyph.txt'  # 这里替换为你的输入文件路径
output_path = 'sorted_glyph.txt'  # 这里替换为你想要的输出文件路径
sort_lines_in_file(input_path, output_path)