/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // update field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "number3646409222",
    "max": null,
    "min": 0,
    "name": "option_value_id",
    "onlyInt": true,
    "presentable": false,
    "required": true,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // update field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "number3646409222",
    "max": null,
    "min": 0,
    "name": "option_value_id",
    "onlyInt": false,
    "presentable": false,
    "required": true,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
})
