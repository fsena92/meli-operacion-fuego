package cache

import (
    "github.com/patrickmn/go-cache"
    "time"
    "github.com/fsena92/meli-operacion-fuego/structs"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func SetCache(key string, object interface{}) bool {
    Cache.Set(key, object, cache.NoExpiration)
    return true
}

func GetCache(key string) (structs.SatelliteRequest, bool) {
    var object structs.SatelliteRequest
    var found bool
    data, found := Cache.Get(key)
    if found {
      object = data.(structs.SatelliteRequest)
    }
    return object, found
}

func FlushCache(){
	Cache.Flush()
}

func CountingItems() int {
	return Cache.ItemCount()
}