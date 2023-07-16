import pymongo


client = pymongo.MongoClient("localhost", 27017)
db = client.cooker

mexican_recipes = db.recipes.find({"tags": "mexican", "rating_avg": {"$gt": 3.5}})
for r in mexican_recipes:
    print(r["title"])
