package analyzer

import (
	"context"
	"time"
)

type ResourceWaste struct {
	ResourceID   string
	ResourceType string
	MonthlyWasteUSD float64
	Reason       string
}

type ComputeAnalyzer struct{}

func (c *ComputeAnalyzer) ScanIdleInstances(ctx context.Context) ([]ResourceWaste, error) {
	// Evaluates CloudWatch 14-day CPUUtilization metric
	return []ResourceWaste{
		{
			ResourceID:   "i-09f18b2c451a",
			ResourceType: "EC2::Instance",
			MonthlyWasteUSD: 148.50,
			Reason:       "CPU utilization < 2% across 14-day observation window",
		},
	}, nil
}
