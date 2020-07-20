FROM python:3.7
RUN pip install locust prometheus-client -i https://pypi.tuna.tsinghua.edu.cn/simple
ADD prometheus_exporter.py .
EXPOSE 5557 5558 8089
CMD ["locust", "--master", "-f", "prometheus_exporter.py"]