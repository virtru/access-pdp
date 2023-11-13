package attributes

import (
	"reflect"
	"testing"
)

func TestAttributeDefinition_GetAuthority(t *testing.T) {
	type fields struct {
		Authority string
		Name      string
		Rule      string
		State     string
		Order     []string
		GroupBy   *AttributeInstance
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive",
			fields: fields{
				Authority: "myauthority",
				Name:      "",
				Rule:      "",
				State:     "",
				Order:     nil,
				GroupBy:   nil,
			},
			want: "myauthority",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrdef := AttributeDefinition{
				Authority: tt.fields.Authority,
				Name:      tt.fields.Name,
				Rule:      tt.fields.Rule,
				State:     tt.fields.State,
				Order:     tt.fields.Order,
				GroupBy:   tt.fields.GroupBy,
			}
			if got := attrdef.GetAuthority(); got != tt.want {
				t.Errorf("GetAuthority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttributeDefinition_GetCanonicalName(t *testing.T) {
	type fields struct {
		Authority string
		Name      string
		Rule      string
		State     string
		Order     []string
		GroupBy   *AttributeInstance
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive",
			fields: fields{
				Authority: "a",
				Name:      "b",
				Rule:      "c",
				State:     "d",
				Order:     nil,
				GroupBy:   nil,
			},
			want: "a/attr/b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrdef := AttributeDefinition{
				Authority: tt.fields.Authority,
				Name:      tt.fields.Name,
				Rule:      tt.fields.Rule,
				State:     tt.fields.State,
				Order:     tt.fields.Order,
				GroupBy:   tt.fields.GroupBy,
			}
			if got := attrdef.GetCanonicalName(); got != tt.want {
				t.Errorf("GetCanonicalName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttributeInstance_GetAuthority(t *testing.T) {
	type fields struct {
		Authority string
		Name      string
		Value     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive",
			fields: fields{
				Authority: "a",
				Name:      "b",
				Value:     "c",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrdef := AttributeInstance{
				Authority: tt.fields.Authority,
				Name:      tt.fields.Name,
				Value:     tt.fields.Value,
			}
			if got := attrdef.GetAuthority(); got != tt.want {
				t.Errorf("GetAuthority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttributeInstance_GetCanonicalName(t *testing.T) {
	type fields struct {
		Authority string
		Name      string
		Value     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive",
			fields: fields{
				Authority: "a",
				Name:      "b",
				Value:     "c",
			},
			want: "a/attr/b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := AttributeInstance{
				Authority: tt.fields.Authority,
				Name:      tt.fields.Name,
				Value:     tt.fields.Value,
			}
			if got := attr.GetCanonicalName(); got != tt.want {
				t.Errorf("GetCanonicalName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttributeInstance_String(t *testing.T) {
	type fields struct {
		Authority string
		Name      string
		Value     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive",
			fields: fields{
				Authority: "a",
				Name:      "b",
				Value:     "c",
			},
			want: "a/attr/b/value/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := AttributeInstance{
				Authority: tt.fields.Authority,
				Name:      tt.fields.Name,
				Value:     tt.fields.Value,
			}
			if got := attr.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClusterByAuthority(t *testing.T) {
	type args[attrCluster Clusterable] struct {
		attrs []attrCluster
	}
	type testCase[attrCluster Clusterable] struct {
		name string
		args args[attrCluster]
		want map[string][]attrCluster
	}
	tests := []testCase[AttributeDefinition]{
		{
			name: "positive",
			args: args[AttributeDefinition]{},
			want: make(map[string][]AttributeDefinition),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClusterByAuthority(tt.args.attrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterByAuthority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClusterByCanonicalName(t *testing.T) {
	type args[attrCluster Clusterable] struct {
		attrs []attrCluster
	}
	type testCase[attrCluster Clusterable] struct {
		name string
		args args[attrCluster]
		want map[string][]attrCluster
	}
	tests := []testCase[AttributeDefinition]{
		{
			name: "positive",
			args: args[AttributeDefinition]{},
			want: make(map[string][]AttributeDefinition),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClusterByCanonicalName(tt.args.attrs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterByCanonicalName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInstanceFromParts(t *testing.T) {
	type args struct {
		namespace string
		name      string
		value     string
	}
	tests := []struct {
		name    string
		args    args
		want    AttributeInstance
		wantErr bool
	}{
		{
			name: "positive",
			args: args{
				namespace: "https://a",
				name:      "b",
				value:     "c",
			},
			want: AttributeInstance{
				Authority: "https://a",
				Name:      "b",
				Value:     "c",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstanceFromParts(tt.args.namespace, tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstanceFromParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInstanceFromParts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInstanceFromURI(t *testing.T) {
	type args struct {
		attributeURI string
	}
	tests := []struct {
		name    string
		args    args
		want    AttributeInstance
		wantErr bool
	}{
		{
			name: "positive",
			args: args{
				attributeURI: "https://a/attr/b/value/c",
			},
			want: AttributeInstance{
				Authority: "https://a",
				Name:      "b",
				Value:     "c",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstanceFromURI(tt.args.attributeURI)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstanceFromURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInstanceFromURI() got = %v, want %v", got, tt.want)
			}
		})
	}
}
