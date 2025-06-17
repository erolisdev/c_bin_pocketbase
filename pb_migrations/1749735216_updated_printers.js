/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_2239120814")

  // add field
  collection.fields.addAt(13, new Field({
    "hidden": false,
    "id": "number1119987104",
    "max": 1,
    "min": 0,
    "name": "has_product_sort",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_2239120814")

  // remove field
  collection.fields.removeById("number1119987104")

  return app.save(collection)
})
