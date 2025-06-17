/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_3611327082")

  // add field
  collection.fields.addAt(19, new Field({
    "hidden": false,
    "id": "number2051285377",
    "max": null,
    "min": null,
    "name": "lbl_printer_id",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // add field
  collection.fields.addAt(20, new Field({
    "hidden": false,
    "id": "number1189890378",
    "max": null,
    "min": null,
    "name": "printer_id",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // add field
  collection.fields.addAt(21, new Field({
    "hidden": false,
    "id": "number1769643497",
    "max": null,
    "min": null,
    "name": "c_sort_order",
    "onlyInt": true,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // add field
  collection.fields.addAt(22, new Field({
    "hidden": false,
    "id": "number306617826",
    "max": null,
    "min": null,
    "name": "category_id",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_3611327082")

  // remove field
  collection.fields.removeById("number2051285377")

  // remove field
  collection.fields.removeById("number1189890378")

  // remove field
  collection.fields.removeById("number1769643497")

  // remove field
  collection.fields.removeById("number306617826")

  return app.save(collection)
})
