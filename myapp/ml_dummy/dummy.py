# dummy_ml_service.py
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route("/generate", methods=["POST"])
def generate():
    data = request.get_json()
    text = data.get("text", "")
    summary = text[:75] + "..." if len(text) > 75 else text
    return jsonify({"summary": summary})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

