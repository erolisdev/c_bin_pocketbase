/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1988651086")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE UNIQUE INDEX idx_live_orders_order_number_date ON live_orders (order_number, date)",
      "CREATE INDEX `idx_order_status_id` ON `live_orders` (order_status_id)"
    ]
  }, collection)

  // remove field
  collection.fields.removeById("number504055712")

  // remove field
  collection.fields.removeById("number1577914996")

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1988651086")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE UNIQUE INDEX idx_live_orders_order_number_date ON live_orders (order_number, date)",
      "CREATE UNIQUE INDEX idx_remote_order_id_date ON live_orders (remote_order_id,date)",
      "CREATE INDEX `idx_order_status_id` ON `live_orders` (order_status_id)"
    ]
  }, collection)

  // add field
  collection.fields.addAt(2, new Field({
    "hidden": false,
    "id": "number504055712",
    "max": null,
    "min": null,
    "name": "local_order_id",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  // add field
  collection.fields.addAt(3, new Field({
    "hidden": false,
    "id": "number1577914996",
    "max": null,
    "min": null,
    "name": "remote_order_id",
    "onlyInt": false,
    "presentable": false,
    "required": false,
    "system": false,
    "type": "number"
  }))

  return app.save(collection)
})
