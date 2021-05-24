package sku

type ArgSyncStock struct {
	IdempotentId   string `json:"idempotent_id"`
	Incremental    string `json:"incremental"`
	OutWarehouseId string `json:"out_warehouse_id"`
	SkuId          string `json:"sku_id"`
	StepStockNum   string `json:"step_stock_num"`
	StockNum       string `json:"stock_num"`
	SupplierId     string `json:"supplier_id"`
}
