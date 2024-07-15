package apijson

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func P[T any](v T) *T { return &v }

type TfsdkStructs struct {
	BoolValue     types.Bool           `tfsdk:"bool_value" json:"bool_value"`
	StringValue   types.String         `tfsdk:"string_value" json:"string_value"`
	Data          *EmbeddedTfsdkStruct `tfsdk:"data" json:"data"`
	FloatValue    types.Float64        `tfsdk:"float_value" json:"float_value"`
	OptionalArray *[]types.String      `tfsdk:"optional_array" json:"optional_array"`
}

type EmbeddedTfsdkStruct struct {
	EmbeddedString types.String `tfsdk:"embedded_string" json:"embedded_string"`
	EmbeddedInt    types.Int64  `tfsdk:"embedded_int" json:"embedded_int"`
}

type Primitives struct {
	A bool    `json:"a"`
	B int     `json:"b"`
	C uint    `json:"c"`
	D float64 `json:"d"`
	E float32 `json:"e"`
	F []int   `json:"f"`
}

type PrimitivePointers struct {
	A *bool    `json:"a"`
	B *int     `json:"b"`
	C *uint    `json:"c"`
	D *float64 `json:"d"`
	E *float32 `json:"e"`
	F *[]int   `json:"f"`
}

type Slices struct {
	Slice []Primitives `json:"slices"`
}

type DateTime struct {
	Date     time.Time `json:"date" format:"date"`
	DateTime time.Time `json:"date-time" format:"date-time"`
}

type DateTimeCustom struct {
	DateCustom     timetypes.RFC3339 `json:"date" format:"date"`
	DateTimeCustom timetypes.RFC3339 `json:"date-time" format:"date-time"`
}

type AdditionalProperties struct {
	A      bool                   `json:"a"`
	Extras map[string]interface{} `json:"-,extras"`
}

type TypedAdditionalProperties struct {
	A      bool           `json:"a"`
	Extras map[string]int `json:"-,extras"`
}

type EmbeddedStructs struct {
	AdditionalProperties
	A      *int                   `json:"number2"`
	Extras map[string]interface{} `json:"-,extras"`
}

type Recursive struct {
	Name  string     `json:"name"`
	Child *Recursive `json:"child"`
}

type UnknownStruct struct {
	Unknown interface{} `json:"unknown"`
}

type UnionStruct struct {
	Union Union `json:"union" format:"date"`
}

type Union interface {
	union()
}

type Inline struct {
	InlineField Primitives `json:"-,inline"`
}

type InlineArray struct {
	InlineField []string `json:"-,inline"`
}

func init() {
	RegisterUnion(reflect.TypeOf((*Union)(nil)).Elem(), "type",
		UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionTime{}),
		},
		UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionInteger(0)),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeA",
			Type:               reflect.TypeOf(UnionStructA{}),
		},
		UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "typeB",
			Type:               reflect.TypeOf(UnionStructB{}),
		},
	)
}

type UnionInteger int64

func (UnionInteger) union() {}

type UnionStructA struct {
	Type string `json:"type"`
	A    string `json:"a"`
	B    string `json:"b"`
}

func (UnionStructA) union() {}

type UnionStructB struct {
	Type string `json:"type"`
	A    string `json:"a"`
}

func (UnionStructB) union() {}

type UnionTime time.Time

func (UnionTime) union() {}

type ResultEnvelope struct {
	Result RecordsModel `json:"result"`
}

type RecordsModel struct {
	A types.String `tfsdk:"a" json:"a"`
	B types.String `tfsdk:"b" json:"b"`
	C types.String `tfsdk:"c" json:"c,computed"`
}

func DropDiagnostic[resType interface{}](res resType, diags diag.Diagnostics) resType {
	for _, d := range diags {
		panic(fmt.Sprintf("%s: %s", d.Summary(), d.Detail()))
	}
	return res
}

type JsonModel struct {
	Arr  jsontypes.Normalized            `tfsdk:"arr" json:"arr"`
	Bol  jsontypes.Normalized            `tfsdk:"bol" json:"bol"`
	Map  jsontypes.Normalized            `tfsdk:"map" json:"map"`
	Nil  jsontypes.Normalized            `tfsdk:"nil" json:"nil"`
	Num  jsontypes.Normalized            `tfsdk:"num" json:"num"`
	Str  jsontypes.Normalized            `tfsdk:"str" json:"str"`
	Arr2 []jsontypes.Normalized          `tfsdk:"arr2" json:"arr2"`
	Map2 map[string]jsontypes.Normalized `tfsdk:"map2" json:"map2"`
}

func time2time(t time.Time) timetypes.RFC3339 {
	return timetypes.NewRFC3339TimePointerValue(&t)
}

var tests = map[string]struct {
	buf string
	val interface{}
}{
	"true":               {"true", true},
	"false":              {"false", false},
	"int":                {"1", 1},
	"int_bigger":         {"12324", 12324},
	"int_string_coerce":  {`"65"`, 65},
	"int_boolean_coerce": {"true", 1},
	"int64":              {"1", int64(1)},
	"int64_huge":         {"123456789123456789", int64(123456789123456789)},
	"uint":               {"1", uint(1)},
	"uint_bigger":        {"12324", uint(12324)},
	"uint_coerce":        {`"65"`, uint(65)},
	"float_1.54":         {"1.54", float32(1.54)},
	"float_1.89":         {"1.89", float64(1.89)},
	"string":             {`"str"`, "str"},
	"string_int_coerce":  {`12`, "12"},
	"array_string":       {`["foo","bar"]`, []string{"foo", "bar"}},
	"array_int":          {`[1,2]`, []int{1, 2}},
	"array_int_coerce":   {`["1",2]`, []int{1, 2}},

	"ptr_true":               {"true", P(true)},
	"ptr_false":              {"false", P(false)},
	"ptr_int":                {"1", P(1)},
	"ptr_int_bigger":         {"12324", P(12324)},
	"ptr_int_string_coerce":  {`"65"`, P(65)},
	"ptr_int_boolean_coerce": {"true", P(1)},
	"ptr_int64":              {"1", P(int64(1))},
	"ptr_int64_huge":         {"123456789123456789", P(int64(123456789123456789))},
	"ptr_uint":               {"1", P(uint(1))},
	"ptr_uint_bigger":        {"12324", P(uint(12324))},
	"ptr_uint_coerce":        {`"65"`, P(uint(65))},
	"ptr_float_1.54":         {"1.54", P(float32(1.54))},
	"ptr_float_1.89":         {"1.89", P(float64(1.89))},

	"date_time":             {`"2007-03-01T13:00:00Z"`, time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC)},
	"date_time_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)},

	"date_time_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)},
	"date_time_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC)},
	"date_time_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60))},
	"date_time_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC)},

	"date_time_custom":             {`"2007-03-01T13:00:00Z"`, time2time(time.Date(2007, time.March, 1, 13, 0, 0, 0, time.UTC))},
	"date_time_custom_nano_coerce": {`"2007-03-01T13:03:05.123456789Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"date_time_custom_missing_t_coerce":              {`"2007-03-01 13:03:05Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_custom_missing_timezone_coerce":       {`"2007-03-01T13:03:05"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.UTC))},
	"date_time_custom_missing_timezone_colon_coerce": {`"2007-03-01T13:03:05+0100"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 0, time.FixedZone("", 60*60)))},
	"date_time_custom_nano_missing_t_coerce":         {`"2007-03-01 13:03:05.123456789Z"`, time2time(time.Date(2007, time.March, 1, 13, 3, 5, 123456789, time.UTC))},

	"map_string":    {`{"foo":"bar"}`, map[string]string{"foo": "bar"}},
	"map_interface": {`{"a":1,"b":"str","c":false}`, map[string]interface{}{"a": float64(1), "b": "str", "c": false}},

	"primitive_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}`,
		Primitives{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}},
	},

	"slices": {
		`{"slices":[{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4]}]}`,
		Slices{
			Slice: []Primitives{{A: false, B: 237628372683, C: uint(654), D: 9999.43, E: 43.76, F: []int{1, 2, 3, 4}}},
		},
	},

	"primitive_pointer_struct": {
		`{"a":false,"b":237628372683,"c":654,"d":9999.43,"e":43.76,"f":[1,2,3,4,5]}`,
		PrimitivePointers{
			A: P(false),
			B: P(237628372683),
			C: P(uint(654)),
			D: P(9999.43),
			E: P(float32(43.76)),
			F: &[]int{1, 2, 3, 4, 5},
		},
	},

	"datetime_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTime{
			Date:     time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
			DateTime: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
		},
	},

	"datetime_custom_struct": {
		`{"date":"2006-01-02","date-time":"2006-01-02T15:04:05Z"}`,
		DateTimeCustom{
			DateCustom:     time2time(time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)),
			DateTimeCustom: time2time(time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)),
		},
	},

	"additional_properties": {
		`{"a":true,"bar":"value","foo":true}`,
		AdditionalProperties{
			A: true,
			Extras: map[string]interface{}{
				"bar": "value",
				"foo": true,
			},
		},
	},

	"recursive_struct": {
		`{"child":{"name":"Alex"},"name":"Robert"}`,
		Recursive{Name: "Robert", Child: &Recursive{Name: "Alex"}},
	},

	"unknown_struct_number": {
		`{"unknown":12}`,
		UnknownStruct{
			Unknown: 12.,
		},
	},

	"unknown_struct_map": {
		`{"unknown":{"foo":"bar"}}`,
		UnknownStruct{
			Unknown: map[string]interface{}{
				"foo": "bar",
			},
		},
	},

	"union_integer": {
		`{"union":12}`,
		UnionStruct{
			Union: UnionInteger(12),
		},
	},

	"union_struct_discriminated_a": {
		`{"union":{"a":"foo","b":"bar","type":"typeA"}}`,
		UnionStruct{
			Union: UnionStructA{
				Type: "typeA",
				A:    "foo",
				B:    "bar",
			},
		},
	},

	"union_struct_discriminated_b": {
		`{"union":{"a":"foo","type":"typeB"}}`,
		UnionStruct{
			Union: UnionStructB{
				Type: "typeB",
				A:    "foo",
			},
		},
	},

	"union_struct_time": {
		`{"union":"2010-05-23"}`,
		UnionStruct{
			Union: UnionTime(time.Date(2010, 05, 23, 0, 0, 0, 0, time.UTC)),
		},
	},

	"tfsdk_null_string":  {"", types.StringNull()},
	"tfsdk_null_int":     {"", types.Int64Null()},
	"tfsdk_null_float":   {"", types.Float64Null()},
	"tfsdk_null_bool":    {"", types.BoolNull()},
	"tfsdk_null_dynamic": {"", types.DynamicNull()},

	"tfsdk_string":             {`"hey"`, types.StringValue("hey")},
	"tfsdk_true":               {"true", types.BoolValue(true)},
	"tfsdk_false":              {"false", types.BoolValue(false)},
	"tfsdk_int":                {"1", types.Int64Value(1)},
	"tfsdk_int_bigger":         {"12324", types.Int64Value(12324)},
	"tfsdk_int_string_coerce":  {`"65"`, types.Int64Value(65)},
	"tfsdk_int_boolean_coerce": {"true", types.BoolValue(true)},
	"tfsdk_float_1.54":         {"1.54", types.Float64Value(1.54)},
	"tfsdk_float_1.89":         {"1.89", types.Float64Value(1.89)},
	"tfsdk_array_ptr":          {"[\"hi\",null]", &[]types.String{types.StringValue("hi"), types.StringNull()}},
	"tfsdk_dynamic_string":     {`"hey"`, types.DynamicValue(types.StringValue("hey"))},
	"tfsdk_dynamic_int":        {"5", types.DynamicValue(types.Int64Value(5))},

	"embedded_tfsdk_struct": {
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
	},

	"json_struct": {
		`{"arr":[true,1,"one"],"arr2":[true,1,"one"],"bol":false,"map":{"nil":null,"bol":false,"str":"two"},"map2":{"bol":false,"nil":null,"str":"two"},"nil":null,"num":2,"str":"two"}`,
		JsonModel{
			Arr:  jsontypes.NewNormalizedValue(`[true,1,"one"]`),
			Bol:  jsontypes.NewNormalizedValue("false"),
			Map:  jsontypes.NewNormalizedValue(`{"nil":null,"bol":false,"str":"two"}`),
			Nil:  jsontypes.NewNormalizedValue("null"),
			Num:  jsontypes.NewNormalizedValue("2"),
			Str:  jsontypes.NewNormalizedValue(`"two"`),
			Arr2: []jsontypes.Normalized{jsontypes.NewNormalizedValue("true"), jsontypes.NewNormalizedValue("1"), jsontypes.NewNormalizedValue(`"one"`)},
			Map2: map[string]jsontypes.Normalized{"nil": jsontypes.NewNormalizedValue("null"), "bol": jsontypes.NewNormalizedValue("false"), "str": jsontypes.NewNormalizedValue(`"two"`)},
		},
	},

	"json_struct_nil1": {`{}`, JsonModel{Nil: jsontypes.NewNormalizedNull()}},
	"json_struct_nil2": {`{}`, JsonModel{}},
	"json_struct_nil3": {`{"nil":null}`, JsonModel{Nil: jsontypes.NewNormalizedValue("null")}},
}

var decode_only_tests = map[string]struct {
	buf string
	val interface{}
}{
	"tfsdk_struct_decode": {
		`{"result":{"c":"7887590e1967befa70f48ffe9f61ce80","a":"88281d6015751d6172e7313b0c665b5e","extra":"property","another":2,"b":"http://example.com/example.html\t20"}`,
		ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}},
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		TfsdkStructs{
			BoolValue:   types.BoolNull(),
			StringValue: types.StringNull(),
			FloatValue:  types.Float64Null(),
		},
	},
}

var encode_only_tests = map[string]struct {
	buf string
	val interface{}
}{
	"tfsdk_struct_encode": {
		`{"result":{"a":"88281d6015751d6172e7313b0c665b5e","b":"http://example.com/example.html\t20"}}`,
		ResultEnvelope{RecordsModel{
			A: types.StringValue("88281d6015751d6172e7313b0c665b5e"),
			B: types.StringValue("http://example.com/example.html\t20"),
			C: types.StringValue("7887590e1967befa70f48ffe9f61ce80"),
		}},
	},

	"embedded_tfsdk_struct_nil": {
		`{}`,
		TfsdkStructs{
			BoolValue:   types.BoolNull(),
			StringValue: types.StringNull(),
			FloatValue:  types.Float64Null(),
		},
	},
}

func merge[T interface{}](test_array ...map[string]T) map[string]T {
	out := make(map[string]T)
	for _, tests := range test_array {
		for name, t := range tests {
			out[name] = t
		}
	}
	return out
}

func TestDecode(t *testing.T) {
	for name, test := range merge(tests, decode_only_tests) {
		t.Run(name, func(t *testing.T) {
			result := reflect.New(reflect.TypeOf(test.val))
			if err := Unmarshal([]byte(test.buf), result.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", result, err)
			}
			if !reflect.DeepEqual(result.Elem().Interface(), test.val) {
				t.Fatalf("expected '%s' to deserialize to \n%#v\nbut got\n%#v", test.buf, test.val, result.Elem().Interface())
			}
		})
	}
}

func TestEncode(t *testing.T) {
	for name, test := range merge(tests, encode_only_tests) {
		if strings.HasSuffix(name, "_coerce") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			raw, err := Marshal(test.val)
			if err != nil {
				t.Fatalf("serialization of %v failed with error %v", test.val, err)
			}
			if string(raw) != test.buf {
				t.Fatalf("expected %+#v to serialize to \n%s\n but got \n%s\n", test.val, test.buf, string(raw))
			}
		})
	}
}

var updateTests = map[string]struct {
	state    interface{}
	plan     interface{}
	expected string
}{
	"true":           {true, true, "true"},
	"terraform_true": {types.BoolValue(true), types.BoolValue(true), "true"},

	"null to true":   {types.BoolNull(), types.BoolValue(true), "true"},
	"false to true":  {types.BoolValue(false), types.BoolValue(true), "true"},
	"unset bool":     {types.BoolValue(false), types.BoolNull(), "null"},
	"omit null bool": {types.BoolNull(), types.BoolNull(), ""},

	"string set":       {types.StringNull(), types.StringValue("two"), "\"two\""},
	"string update":    {types.StringValue("one"), types.StringValue("two"), "\"two\""},
	"unset string":     {types.StringValue("hey"), types.StringNull(), "null"},
	"omit null string": {types.StringNull(), types.StringNull(), ""},

	"int set":       {types.Int64Null(), types.Int64Value(42), "42"},
	"int update":    {types.Int64Value(42), types.Int64Value(43), "43"},
	"unset int":     {types.Int64Value(42), types.Int64Null(), "null"},
	"omit null int": {types.Int64Null(), types.Int64Null(), ""},

	"dynamic omit null":                     {types.DynamicNull(), types.DynamicNull(), ""},
	"dynamic omit underlying null state":    {types.DynamicValue(types.Int64Null()), types.DynamicNull(), ""},
	"dynamic omit underlying null plan":     {types.DynamicNull(), types.DynamicValue(types.Int64Null()), ""},
	"dynamic omit unknown":                  {types.DynamicUnknown(), types.DynamicUnknown(), ""},
	"dynamic omit underlying unknown state": {types.DynamicValue(types.Int64Unknown()), types.DynamicUnknown(), ""},
	"dynamic omit underlying unknown plan":  {types.DynamicUnknown(), types.DynamicValue(types.Int64Unknown()), ""},
	"dynamic unset null":                    {types.DynamicValue(types.Int64Value(4)), types.DynamicNull(), "null"},
	"dynamic int set":                       {types.DynamicNull(), types.DynamicValue(types.Int64Value(5)), "5"},
	"dynamic int update":                    {types.DynamicValue(types.Int64Value(4)), types.DynamicValue(types.Int64Value(5)), "5"},

	"set struct fields": {
		TfsdkStructs{},
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		`{"bool_value":true,"data":{"embedded_int":17,"embedded_string":"embedded_string_value"},"float_value":3.14,"optional_array":["hi","there"],"string_value":"string_value"}`,
	},

	"update some struct fields": {
		TfsdkStructs{
			BoolValue:   types.BoolValue(true),
			StringValue: types.StringValue("string_value"),
			FloatValue:  types.Float64Value(3.14),
		},
		TfsdkStructs{
			BoolValue:   types.BoolValue(false),
			StringValue: types.StringValue("another_string"),
			FloatValue:  types.Float64Value(1.14),
		},
		`{"bool_value":false,"float_value":1.14,"string_value":"another_string"}`,
	},

	"unset nested struct fields": {
		TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Value(17),
			},
		},
		TfsdkStructs{
			OptionalArray: &[]types.String{types.StringValue("hi")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedInt: types.Int64Null(),
			},
		},
		`{"data":{"embedded_int":null},"optional_array":["hi"]}`,
	},

	"unset struct fields": {
		TfsdkStructs{
			BoolValue:     types.BoolValue(true),
			StringValue:   types.StringValue("string_value"),
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
			Data: &EmbeddedTfsdkStruct{
				EmbeddedString: types.StringValue("embedded_string_value"),
				EmbeddedInt:    types.Int64Value(17),
			},
		},
		TfsdkStructs{},
		`{"bool_value":null,"data":null,"float_value":null,"optional_array":null,"string_value":null}`,
	},

	"set empty array": {
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{types.StringValue("hi"), types.StringValue("there")},
		},
		TfsdkStructs{
			FloatValue:    types.Float64Value(3.14),
			OptionalArray: &[]types.String{},
		},
		`{"float_value":3.14,"optional_array":[]}`,
	},
}

func TestEncodeForUpdate(t *testing.T) {
	for name, test := range updateTests {
		t.Run(name, func(t *testing.T) {
			raw, err := MarshalForUpdate(test.plan, test.state)
			if err != nil {
				t.Fatalf("serialization of %v, %v failed with error %v", test.plan, test.state, err)
			}
			if string(raw) != test.expected {
				t.Fatalf("expected %+#v, %+#v to serialize to \n%s\n but got \n%s\n", test.state, test.plan, test.expected, string(raw))
			}
		})
	}
}

var decode_from_value_tests = map[string]struct {
	buf      string
	starting interface{}
	expected interface{}
}{

	"tfsdk_dynamic_null": {
		`null`,
		types.DynamicNull(),
		types.DynamicNull(),
	},

	"tfsdk_dynamic_string_from_null": {
		`"hey"`,
		types.DynamicNull(),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_string_from_unknown": {
		`"hey"`,
		types.DynamicUnknown(),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_string_from_value": {
		`"hey"`,
		types.DynamicValue(types.StringValue("before_value")),
		types.DynamicValue(types.StringValue("hey")),
	},

	"tfsdk_dynamic_int_from_null": {
		`14`,
		types.DynamicNull(),
		types.DynamicValue(types.Int64Value(14)),
	},

	"tfsdk_dynamic_int_from_unknown": {
		`14`,
		types.DynamicUnknown(),
		types.DynamicValue(types.Int64Value(14)),
	},

	"tfsdk_dynamic_int_from_value": {
		`14`,
		types.DynamicValue(types.Int64Value(5)),
		types.DynamicValue(types.Int64Value(14)),
	},
}

func TestDecodeFromValue(t *testing.T) {
	for name, test := range decode_from_value_tests {
		t.Run(name, func(t *testing.T) {
			v := reflect.ValueOf(test.starting)
			starting := reflect.New(v.Type())
			starting.Elem().Set(v)

			if err := Unmarshal([]byte(test.buf), starting.Interface()); err != nil {
				t.Fatalf("deserialization of %v failed with error %v", test.buf, err)
			}
			if !reflect.DeepEqual(starting.Elem().Interface(), test.expected) {
				t.Fatalf("expected '%s' to deserialize to \n%#v\nbut got\n%#v", test.buf, test.expected, starting)
			}
		})
	}
}
