import json
import os
import PyPDF2
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_text_splitters import SentenceTransformersTokenTextSplitter
from langchain_experimental.text_splitter import SemanticChunker
from langchain_huggingface import HuggingFaceEmbeddings


# Путь к директории с PDF файлами
directory_path = "InnoglobalHack_2024"
# Путь к файлу для сохранения результатов
output_file1 = "texts1.json"
output_file2 = "texts2.json"
output_file3 = "texts3.json"


#----------------------------------------------------------------------------------------
all_texts2=[]
print(1)
text_embedder = HuggingFaceEmbeddings(model_name='intfloat/multilingual-e5-large')

print(2)
text_splitter = SemanticChunker(text_embedder, breakpoint_threshold_type="percentile", breakpoint_threshold_amount=95)
print(3)
print(text_splitter)

with open(output_file3, "w", encoding="utf-8") as out_file:
    all_texts = []
    for filename in os.listdir(directory_path):
        if filename.endswith(".pdf"):
            file_path = os.path.join(directory_path, filename)
            if not os.path.exists(file_path):
                print(f"Файл найден: {file_path}")
                continue

            print(f"Processing {file_path}")

            # Чтение текста из PDF
            with open(file_path, "rb") as file:
                pdf_reader = PyPDF2.PdfReader(file)
                document_text = ""
                for page in pdf_reader.pages:
                    document_text += page.extract_text()
            print("Docs cobran")
            chunks = text_splitter.split_text(document_text)
            print("chunks")
            for t in chunks:
                all_texts2.append({"text": t})
            print(f"Всего чанков: {len(chunks)}")
            print("Первые N чанков:")
        break
    res2 = {"chunks": all_texts2}
    # Сохранение всех кусков текста в JSON файл
    json.dump(res2, out_file, ensure_ascii=False, indent=4)

    print(f"Тексты сохранены в файл {output_file3}")
