/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1317801859")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE INDEX `idx_ZAjuDkHDEP` ON `store_zones` (`city`)",
      "CREATE INDEX `idx_pZdmt7dQ9e` ON `store_zones` (`zone`)"
    ],
    "name": "store_zones"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1317801859")

  // update collection data
  unmarshal({
    "indexes": [
      "CREATE INDEX `idx_ZAjuDkHDEP` ON `zones` (`city`)",
      "CREATE INDEX `idx_pZdmt7dQ9e` ON `zones` (`zone`)"
    ],
    "name": "zones"
  }, collection)

  return app.save(collection)
})
