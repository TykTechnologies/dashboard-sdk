package assets

import (
	"context"
	"log"
	"net/http"

	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var sampleAsset = `{"data":{"info":{"title":"Our Sample OAS","version":"1.0.0"},"openapi":"3.0.3","paths":{"/anything":{"post":{"operationId":"anythingpost","responses":{"200":{"description":"post created"}}}}},"x-tyk-api-gateway":{"middleware":{"global":{"cache":{"cacheAllSafeRequests":true,"enabled":true,"timeout":5}}}}},"description":"My first template","id":"my-unique-template-id","kind":"oas-template","name":"sample asset"}`

func CreateAssets(ctx context.Context, client *dashboard.APIClient) (*dashboard.ApiResponse, error) {
	var assetData dashboard.NullableAddAssetRequest
	err := assetData.UnmarshalJSON([]byte(sampleAsset))
	if err != nil {
		return nil, err
	}
	apiResponse, resp, err := client.AssetsAPI.AddAsset(ctx).AddAssetRequest(*assetData.Get()).Execute()
	if err != nil {
		ExtractErrMessage(resp)
		return nil, err
	}
	log.Printf("Asset created successfully with ID: %s", apiResponse.GetID())
	log.Println("-------------------------------------------------------------")
	return apiResponse, nil
}

func GetAssetsByID(ctx context.Context, client *dashboard.APIClient, AssetId string) (*dashboard.Asset, error) {
	asset, resp, err := client.AssetsAPI.GetAsset(ctx, AssetId).Execute()
	if err != nil {
		ExtractErrMessage(resp)
		return nil, err
	}
	log.Printf("Asset fetched successfully with ID: %s and Name: %s", asset.GetId(), asset.GetName())
	log.Println("-------------------------------------------------------------")
	return asset, nil
}

func ExtractErrMessage(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		log.Println(resp.Body)
	}
}
