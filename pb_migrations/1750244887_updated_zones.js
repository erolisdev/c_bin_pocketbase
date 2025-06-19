/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1317801859")

  // add field
  collection.fields.addAt(1, new Field({
    "hidden": false,
    "id": "number2343330479",
    "max": null,
    "min": null,
    "name": "city_id",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1317801859")

  // remove field
  collection.fields.removeById("number2343330479")

  return app.save(collection)
})
