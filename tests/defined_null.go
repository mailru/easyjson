package tests

import (
    "github.com/mailru/easyjson/jlexer"
    "bytes"
    "encoding/json"
)

//easyjson:json
type NullStringStruct struct {
    NS NullString
    VNS VanillaNullString
    S string
}

type NullString struct {
    V string
    Valid bool
    Set bool
}

func (s *NullString) UnmarshalEasyJSON(l *jlexer.Lexer) {
    s.Set = true
    if l.IsNull() {
        l.Skip()
        s.V, s.Valid = "", false
        return
    }
    s.V, s.Valid = l.String(), true
}

type VanillaNullString struct {
    V string
    Valid bool
    Set bool
}

func (s *VanillaNullString) UnmarshalJSON(v []byte) error {
    s.Set = true
    if bytes.Equal(v, []byte("null")) {
        s.V, s.Valid = "", false
        return nil
    }
    err := json.Unmarshal(v, &s.V)
    if err != nil {
        s.Valid = false
        return err
    }
    s.Valid = true
    return nil
}