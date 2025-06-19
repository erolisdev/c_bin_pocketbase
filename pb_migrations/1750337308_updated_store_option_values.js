/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // update field
  collection.fields.addAt(13, new Field({
    "hidden": false,
    "id": "number1352515405",
    "max": null,
    "min": null,
    "name": "reset",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(15, new Field({
    "hidden": false,
    "id": "number2063623452",
    "max": null,
    "min": null,
    "name": "status",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(16, new Field({
    "hidden": false,
    "id": "number1476078466",
    "max": null,
    "min": null,
    "name": "not_recommanded",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(17, new Field({
    "hidden": false,
    "id": "number2128919991",
    "max": null,
    "min": null,
    "name": "image_type",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // update field
  collection.fields.addAt(13, new Field({
    "hidden": false,
    "id": "number1352515405",
    "max": null,
    "min": null,
    "name": "reset",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(15, new Field({
    "hidden": false,
    "id": "number2063623452",
    "max": null,
    "min": null,
    "name": "status",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(16, new Field({
    "hidden": false,
    "id": "number1476078466",
    "max": null,
    "min": null,
    "name": "not_recommanded",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // update field
  collection.fields.addAt(17, new Field({
    "hidden": false,
    "id": "number2128919991",
    "max": null,
    "min": null,
    "name": "image_type",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
})
