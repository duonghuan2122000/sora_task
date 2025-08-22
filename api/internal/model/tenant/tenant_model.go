package tenantmodel

import (
	"context"
	basemodel "sorataskapi/internal/model/base"
)

// Dto req tạo tenant
type CreateTenantReqDto struct {
	Name string `json:"name"`
}

// Dto res tạo tenant
type CreateTenantResDto struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type SelectTenantReqDto struct {
	Id string `json:"id"`
}

func IntoContext(ctx context.Context, tenantId string) context.Context {
	return context.WithValue(ctx, basemodel.CurrentTenantIdKey, tenantId)
}

func FromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(basemodel.CurrentTenantIdKey)
	tenantId, ok := v.(string)
	return tenantId, ok
}
