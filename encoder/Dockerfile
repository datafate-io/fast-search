FROM python:3.11.7-slim
WORKDIR /app
RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT python src/server.py
