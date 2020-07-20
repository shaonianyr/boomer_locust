FROM python:3.7-alpine
WORKDIR /app
RUN pip install flask -i https://pypi.tuna.tsinghua.edu.cn/simple
ADD hello.py .
CMD ["python3", "hello.py"]
