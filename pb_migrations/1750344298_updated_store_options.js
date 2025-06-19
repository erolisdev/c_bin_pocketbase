/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_537519633")

  // update field
  collection.fields.addAt(11, new Field({
    "hidden": false,
    "id": "json3820684983",
    "maxSize": 0,
    "name": "descriptions",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "json"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_537519633")

  // update field
  collection.fields.addAt(11, new Field({
    "hidden": false,
    "id": "json3820684983",
    "maxSize": 0,
    "name": "option_desc",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "json"
  }))

  return app.save(collection)
})
