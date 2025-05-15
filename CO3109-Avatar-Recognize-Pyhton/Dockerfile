FROM python:3.12.9-slim

WORKDIR /AIBackend

RUN apt-get update && apt-get install ffmpeg libsm6 libxext6  -y

COPY ./requirements.txt ./ 
RUN pip install --no-cache-dir -r requirements.txt

RUN pip3 install tensorflow[and-cuda]

COPY ./src ./
COPY ./.env ./

# Set Flask environment variables
ENV FLASK_APP=src/app.py
ENV FLASK_RUN_HOST=0.0.0.0
ENV FLASK_RUN_PORT=8000

# Expose the port
EXPOSE 8000

# Run the application
# CMD ["flask", "run", "--debug"]
