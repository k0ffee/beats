// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package billing

import (
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2019-10-01/consumption"

	"github.com/shopspring/decimal"

	"github.com/k0ffee/beats/v7/metricbeat/mb"
	"github.com/elastic/elastic-agent-libs/mapstr"
)

func EventsMapping(subscriptionId string, results Usage) []mb.Event {
	var events []mb.Event
	if len(results.UsageDetails) > 0 {
		for _, usageDetail := range results.UsageDetails {
			event := mb.Event{
				ModuleFields: mapstr.M{
					"resource": mapstr.M{
						"type":  usageDetail.ConsumedService,
						"group": getResourceGroupFromId(*usageDetail.InstanceID),
						"name":  usageDetail.InstanceName,
					},
					"subscription_id": usageDetail.SubscriptionGUID,
				},
				MetricSetFields: mapstr.M{
					"pretax_cost":       usageDetail.PretaxCost,
					"department_name":   usageDetail.DepartmentName,
					"product":           usageDetail.Product,
					"usage_start":       usageDetail.UsageStart.ToTime(),
					"usage_end":         usageDetail.UsageEnd.ToTime(),
					"currency":          usageDetail.Currency,
					"billing_period_id": usageDetail.BillingPeriodID,
					"account_name":      usageDetail.AccountName,
				},
				Timestamp: time.Now().UTC(),
			}
			event.RootFields = mapstr.M{}
			event.RootFields.Put("cloud.provider", "azure")
			event.RootFields.Put("cloud.region", usageDetail.InstanceLocation)
			event.RootFields.Put("cloud.instance.name", usageDetail.InstanceName)
			event.RootFields.Put("cloud.instance.id", usageDetail.InstanceID)
			events = append(events, event)
		}
	}

	groupedCosts := make(map[*string][]consumption.Forecast)
	for _, forecast := range results.ForecastCosts {
		groupedCosts[forecast.UsageDate] = append(groupedCosts[forecast.UsageDate], forecast)
	}
	for _, forecast := range results.ActualCosts {
		groupedCosts[forecast.UsageDate] = append(groupedCosts[forecast.UsageDate], forecast)
	}
	for usageDate, items := range groupedCosts {
		var actualCost *decimal.Decimal
		var forecastCost *decimal.Decimal
		for _, item := range items {
			if item.ChargeType == consumption.ChargeTypeActual {
				actualCost = item.Charge
			} else {
				forecastCost = item.Charge
			}
		}
		parsedDate, err := time.Parse("2006-01-02", *usageDate)
		if err != nil {
			parsedDate = time.Now().UTC()
		}
		event := mb.Event{
			RootFields: mapstr.M{
				"cloud.provider": "azure",
			},
			ModuleFields: mapstr.M{
				"subscription_id": subscriptionId,
			},
			MetricSetFields: mapstr.M{
				"actual_cost":   actualCost,
				"forecast_cost": forecastCost,
				"usage_date":    parsedDate,
				"currency":      items[0].Currency,
			},
			Timestamp: time.Now().UTC(),
		}
		//event.ID = generateEventID(parsedDate)
		events = append(events, event)
	}
	return events
}

// getResourceGroupFromId maps resource group from resource ID
func getResourceGroupFromId(path string) string {
	params := strings.Split(path, "/")
	for i, param := range params {
		if param == "resourceGroups" {
			return fmt.Sprintf("%s", params[i+1])
		}
	}
	return ""
}
