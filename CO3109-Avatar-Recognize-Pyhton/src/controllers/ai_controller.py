from flask import Blueprint, jsonify, request
from services.ai_service import AIService  # Import AI service

ai_blueprint = Blueprint("ai", __name__)
ai_service = AIService()  # Initialize AI service

@ai_blueprint.route("/save/<string:shopID>/<string:employeeID>", methods=["POST"])
def save(shopID, employeeID):
    # print(f"Received request to save image for shopId: {shopID}, employeeID: {employeeID}")
    if len(request.files) == 0:
        return jsonify({"error": "Image file is required"}), 400
    
    image_file = request.files["image"]
    
    try:
        # Can return error or success
        result = ai_service.save(shopID, employeeID, image_file)
        # else:
        #     result = ai_service.saveBatch(shopId, employeeID, images)
    except Exception as e:
        print(f"Error: {e}")
        return jsonify({"error": "Server-side error! Please check AI Backend console."}), 500
        
    return jsonify({"message": "Image saved successfully"}), 200
@ai_blueprint.route("/find/<string:shopID>", methods=["POST"])
def find(shopID):
    print(f"Received request to find employee in shop: {shopID}")
    if len(request.files) == 0:
        print("Error: No image file provided")
        return jsonify({"error": "Image file is required"}), 400
    image = request.files["image"]
    
    try:
        print(f"Processing image for shop: {shopID}")
        result, status_code = ai_service.find(shopID, image)
        print(f"Find operation result: {result}, status: {status_code}")
        
        response = jsonify(result)
        return response, status_code
    except Exception as e:
        print(f"Error during find operation: {e}")
        error_response = jsonify({"error": "Server-side error! Please check AI Backend console."})
        return error_response, 500
