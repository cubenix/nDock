FROM python:3.7.3-slim

WORKDIR /app

COPY requirements.txt .

RUN pip install --trusted-host pypi.python.org -r requirements.txt

COPY watchdock/. .

EXPOSE 5000

CMD ["python", "app.py"]