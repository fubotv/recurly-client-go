package recurly_test

import (
	"encoding/json"
	"testing"

	"github.com/fubotv/recurly-client-go"
)

func TestSubscriptionChangeCreate_AddOns(t *testing.T) {
	t.Parallel()

	t.Run("default case", func(t *testing.T) {
		t.Parallel()

		sampleCode := "sample-code"
		change := recurly.SubscriptionChangeCreate{AddOns: []recurly.SubscriptionAddOnUpdate{{Code: &sampleCode}}}

		jsonBytes, err := json.Marshal(change)
		if err != nil {
			t.Fatal(err)
		}

		expected := `{"add_ons":[{"code":"sample-code"}]}`
		if string(jsonBytes) != expected {
			t.Fatalf("unexpected JSON: %s", string(jsonBytes))
		}
	})

	t.Run("initialized but empty addons slice", func(t *testing.T) {
		t.Parallel()

		change := recurly.SubscriptionChangeCreate{AddOns: []recurly.SubscriptionAddOnUpdate{}}

		jsonBytes, err := json.Marshal(change)
		if err != nil {
			t.Fatal(err)
		}

		expected := `{"add_ons":[]}`
		if string(jsonBytes) != expected {
			t.Fatalf("unexpected JSON: %s", string(jsonBytes))
		}
	})

	t.Run("nil addons slice", func(t *testing.T) {
		t.Parallel()

		change := recurly.SubscriptionChangeCreate{AddOns: nil}

		jsonBytes, err := json.Marshal(change)
		if err != nil {
			t.Fatal(err)
		}

		expected := `{"add_ons":null}`
		if string(jsonBytes) != expected {
			t.Fatalf("unexpected JSON: %s", string(jsonBytes))
		}
	})
}
