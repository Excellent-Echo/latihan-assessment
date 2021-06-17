package helper

type DeleteFormat struct {
	Data string `json:"data"`
}

func FormatDelete(data string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Data: data,
	}
	return deleteFormat
}
