# base image - python runtime
FROM python:3.7.3-slim

# set the working directory in the container to /app
WORKDIR /app

COPY requirements.txt .

# execute everyone's favorite pip command, pip install -r
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# add the current directory to the container as /app
COPY watchdock/. .

# unblock port 80 for the Flask app to run on
EXPOSE 5000

# execute the Flask app
CMD ["python", "app.py"]