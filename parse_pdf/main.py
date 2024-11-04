import json
import os

import PyPDF2
from langchain_text_splitters import RecursiveCharacterTextSplitter

directory_path = "InnoglobalHack_2024"
output_file = "texts.json"

text_splitter = RecursiveCharacterTextSplitter(chunk_size=100, chunk_overlap=0)

with open(output_file, "w", encoding="utf-8") as out_file:
    all_texts = []
    for filename in os.listdir(directory_path):
        if filename.endswith(".pdf"):
            file_path = os.path.join(directory_path, filename)
            if not os.path.exists(file_path):
                print(f"Файл найден: {file_path}")
                continue

            print(f"Processing {file_path}")

            with open(file_path, "rb") as file:
                pdf_reader = PyPDF2.PdfReader(file)
                document_text = ""
                for page in pdf_reader.pages:
                    document_text += page.extract_text()

            texts = text_splitter.split_text(document_text)
            for t in texts:
                all_texts.append({"text": t})

    res = {"chunks": all_texts}

    json.dump(res, out_file, ensure_ascii=False, indent=4)

print(f"Тексты сохранены в файл {output_file}")
