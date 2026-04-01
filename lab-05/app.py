
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/predict')
def predict():
    x = int(request.args.get('x', 0))
    result = "Even" if x % 2 == 0 else "Odd"
    return jsonify({"input": x, "prediction": result})

app.run(host='0.0.0.0', port=5000)