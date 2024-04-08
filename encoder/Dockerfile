FROM python:3.11.7-slim
WORKDIR /app
COPY requirements.txt .

RUN pip install -r requirements.txt
RUN pip install huggingface_hub["cli"] 
RUN huggingface-cli download sentence-transformers/LaBSE

COPY . .

ENTRYPOINT python src/server.py
