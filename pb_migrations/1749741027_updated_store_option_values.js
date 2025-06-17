/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // add field
  collection.fields.addAt(2, new Field({
    "hidden": false,
    "id": "number2814647663",
    "max": null,
    "min": null,
    "name": "option_id",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_417461541")

  // remove field
  collection.fields.removeById("number2814647663")

  return app.save(collection)
})
