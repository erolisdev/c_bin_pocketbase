/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_2239120814")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE UNIQUE INDEX `idx_ekmLgtvf9F` ON `printers` (`printer_id`)",
      "CREATE INDEX `idx_5u96QSVlIP` ON `printers` (`printer_ip`)"
    ]
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_2239120814")

  // update collection data
  unmarshal({
    "indexes": []
  }, collection)

  return app.save(collection)
})
