/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_537519633")

  // update collection data
  unmarshal({
    "createRule": "",
    "updateRule": ""
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_537519633")

  // update collection data
  unmarshal({
    "createRule": "@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true",
    "updateRule": "@request.auth.id != '' && @request.auth.collectionName = 'users' &&  @request.auth.isManager = true"
  }, collection)

  return app.save(collection)
})
