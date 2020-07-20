from flask import Flask
import json
app = Flask(__name__)

@app.route('/')
def hello():
    return "hello, world"

@app.route('/text')
def text():
    t = {
        'id': 1,
        'str': "string",
        'list': [2, 3, 4],
        'dict': {"msg": "hello, world"}
    }
    return json.dumps(t)

if __name__ == "__main__":
    # 这种是不太推荐的启动方式，我这只是做演示用，官方启动方式参见：http://flask.pocoo.org/docs/0.12/quickstart/#a-minimal-application
    app.run(host="0.0.0.0", debug=True)
