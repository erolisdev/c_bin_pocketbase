/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_2854637623")

  // update field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "number306617826",
    "max": null,
    "min": 0,
    "name": "category_id",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_2854637623")

  // update field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "number306617826",
    "max": null,
    "min": 0,
    "name": "category_id",
    "onlyInt": false,
    "presentable": false,
    "required": true,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
})
