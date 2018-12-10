package jsonUtil

import "fmt"
import "encoding/json"
import "reflect"

func JsonParse(p []byte) interface{}{
  var v interface{};
  err := json.Unmarshal(p, &v);
  if err != nil {
    return v
  }
  return nil
}

func FindInJson(v interface{}, key string)(interface{},string){
  m := v.(map[string]interface{});
  for k, val := range m {
    if reflect.TypeOf(val).Kind().String() == "map" {
      FindInJson(val, key);
    }
    if(k == key) {
      fmt.Println(val);
      switch reflect.TypeOf(val).Kind().String() {
        case "bool":
          return val, "boolean"
        case "float64":
          return val, "number"
        case "array":
          return val, "array"
        case "nil":
          return val, "nil"
        case "string":
          return val, "string"
        case "map":
          return val, "object"
        default:
          return val, "unknown"
      }
    }
  }
  return nil,"error"
}