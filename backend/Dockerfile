FROM python:3.9-alpine
WORKDIR /app
COPY ./requirements.txt .
RUN pip install -r requirements.txt
COPY ./src .
COPY .env .
CMD ["python", "main.py"]
EXPOSE 80
