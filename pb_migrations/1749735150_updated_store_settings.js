/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_4242123432")

  // add field
  collection.fields.addAt(14, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text3923773533",
    "max": 0,
    "min": 0,
    "name": "ssl",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(15, new Field({
    "autogeneratePattern": "",
    "hidden": false,
    "id": "text4101391790",
    "max": 0,
    "min": 0,
    "name": "url",
    "pattern": "",
    "presentable": false,
    "primaryKey": false,
    "required": false,
    "system": false,
    "type": "text"
  }))

  // add field
  collection.fields.addAt(16, new Field({
    "hidden": false,
    "id": "number657989146",
    "max": 1,
    "min": 0,
    "name": "labelfastprint",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_4242123432")

  // remove field
  collection.fields.removeById("text3923773533")

  // remove field
  collection.fields.removeById("text4101391790")

  // remove field
  collection.fields.removeById("number657989146")

  return app.save(collection)
})
