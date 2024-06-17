package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-planned-train-operation-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-train-operation-deletes-rmq-kube/DPFM_API_Output_Formatter"

	"fmt"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.RailwayLine = %d", input.Header.RailwayLine)
	where = fmt.Sprintf("%s\nAND header.TrainOperationVersion = %d", where, input.Header.TrainOperationVersion)
	where = fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)
	where = fmt.Sprintf("%s\nAND header.PlannedTrainOperationID = %d", where, input.Header.PlannedTrainOperationID)
	rows, err := c.db.Query(
		`SELECT 
			header.Station, header.TrainOperationVersion, header.WeekdayType, header.PlannedTrainOperationID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_train_operation_header_data as header ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	where := fmt.Sprintf("WHERE item.RailwayLine = %d", input.Header.RailwayLine)
	where = fmt.Sprintf("%s\nAND item.TrainOperationVersion = %d", where, input.Header.TrainOperationVersion)
	where = fmt.Sprintf("%s\nAND item.WeekdayType = \"%s\"", where, input.Header.WeekdayType)
	where = fmt.Sprintf("%s\nAND item.PlannedTrainOperationID = %d", where, input.Header.PlannedTrainOperationID)
	where = fmt.Sprintf("%s\nAND item.RailwayLineStationID = %d", where, input.Header.Item.RailwayLineStationID)
	rows, err := c.db.Query(
		`SELECT 
			item.Station, item.TrainOperationVersion, item.WeekdayType, item.PlannedTrainOperationID, item.RailwayLineStationID
		FROM DataItemMastersAndTransactionsMysqlKube.data_item_planned_train_operation_item_data as item
		INNER JOIN DataItemMastersAndTransactionsMysqlKube.data_item_planned_train_operation_header_data as header
		ON header.RailwayLine = item.RailwayLine ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
