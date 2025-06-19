/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_4242123432")

  // update field
  collection.fields.addAt(17, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text4145232419",
    "max": 0,
    "min": 0,
    "name": "day_reference",
    "pattern": "^([01]\\d|2[0-3]):([0-5]\\d)$",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_4242123432")

  // update field
  collection.fields.addAt(17, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text4145232419",
    "max": 0,
    "min": 0,
    "name": "day_reference",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  return app.save(collection)
})
