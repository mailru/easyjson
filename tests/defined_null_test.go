package tests

import (
    "testing"
    "github.com/mailru/easyjson"
    "reflect"
)

func TestDefinedNull(t *testing.T) {
    cases := []struct{data string; res NullStringStruct}{
        {`{"NS": "ns", "VNS": "vns", "S": "s"}`, NullStringStruct{NullString{"ns", true, true}, VanillaNullString{"vns", true, true}, "s"}},
        {`{"NS": null, "VNS": null, "S": null}`, NullStringStruct{NullString{"", false, true}, VanillaNullString{"", false, true}, ""}},
        {`{"Unknown": "Value"}`, NullStringStruct{NullString{"", false, false}, VanillaNullString{"", false, false}, ""}},
    }
    for _, c := range cases {
        var res NullStringStruct
        if err := easyjson.Unmarshal([]byte(c.data), &res); err != nil {
            t.Errorf("Unexpected Unmarshal erorr: %v", err)
        }
        if !reflect.DeepEqual(res, c.res) {
            t.Errorf("Expected to unmarshal %+v, got %+v", c.res, res)
        }
    }
}


