package requests

type Item struct {
	RailwayLine					int		`json:"RailwayLine"`
	TrainOperationVersion		int		`json:"TrainOperationVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	PlannedTrainOperationID		int		`json:"PlannedTrainOperationID"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
