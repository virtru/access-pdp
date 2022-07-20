package pdp_test

import (
	ctx "context"
	"fmt"

	"go.uber.org/zap"

	accesspdp "github.com/virtru/access-pdp"
	attrs "github.com/virtru/access-pdp/attributes"
)

//AnyOf tests
func Example() {
	zapLog, _ := zap.NewDevelopment()

	entityID := "4f6636ca-c60c-40d1-9f3f-015086303f74"
	attrAuthorities := []string{"https://example.org"}
	AttrDefinitions := []attrs.AttributeDefinition{
		{
			Authority: attrAuthorities[0],
			Name:      "MyAttr",
			Rule:      "anyOf",
			Order:     []string{"Value1", "Value2"},
		},
	}
	DataAttrs := []attrs.AttributeInstance{
		{
			Authority: attrAuthorities[0],
			Name:      AttrDefinitions[0].Name,
			Value:     AttrDefinitions[0].Order[1],
		},
		{
			Authority: attrAuthorities[0],
			Name:      AttrDefinitions[0].Name,
			Value:     AttrDefinitions[0].Order[0],
		},
	}
	EntityAttrs := map[string][]attrs.AttributeInstance{
		entityID: {
			{
				Authority: "https://example.org",
				Name:      "MyAttr",
				Value:     "Value2",
			},
			{
				Authority: "https://meep.org",
				Name:      "meep",
				Value:     "beepbeep",
			},
		},
	}
	accessPDP := accesspdp.NewAccessPDP(zapLog.Sugar())

	decisions, err := accessPDP.DetermineAccess(DataAttrs, EntityAttrs, AttrDefinitions, ctx.Background())
	if err != nil {
		zapLog.Error("Could not generate a decision!")
	}

	fmt.Printf("Decision result: %+v", decisions)
}
