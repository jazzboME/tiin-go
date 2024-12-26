package tiingo

//
// import (
// 	"context"
// 	"testing"
// )
//
// func TestClient_FilteredSymbolList(t *testing.T) {
// 	t.Parallel()
//
// 	c, err := getClient()
// 	if err != nil {
// 		t.Fatalf("failed to get tiingo client: %s", err)
// 	}
//
// 	var count int
// 	symbols, err := c.FilteredSymbolList(context.Background(), func(asset SymbolItem) bool {
// 		if count >= 10 {
// 			return false
// 		}
// 		if asset.Exchange != "NYSE" {
// 			return false
// 		}
// 		if asset.AssetType != "Stock" {
// 			return false
// 		}
//
// 		count++
// 		return true
// 	})
// 	if err != nil {
// 		t.Fatalf("failed to get symbol list: %s", err)
// 	}
//
// 	if len(symbols) != 10 {
// 		t.Fatalf("symbols list length is wrong. expected=10, got=%d", len(symbols))
// 	}
//
// 	for _, s := range symbols {
// 		if s.AssetType != "Stock" {
// 			t.Fatalf("asset type wrong. expected=stock got=%s", s.AssetType)
// 		}
// 		if s.Exchange != "NYSE" {
// 			t.Fatalf("exchange wrong. expected=NYSE got=%s", s.Exchange)
// 		}
// 	}
// }
