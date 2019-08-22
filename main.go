package main

import (
	"fmt"
	"github.com/SurgicalSteel/sosfilter"
)

func main() {
	sos := generateMassiveSlice()
	filter := sosfilter.NewFilter(sos, []sosfilter.FilterRule{
		sosfilter.FilterRule{
			FilterFunc: sosfilter.StartsWith,
			Param:      "D",
		},
		sosfilter.FilterRule{
			FilterFunc: sosfilter.EndsWith,
			Param:      "se",
		},
	})
	filteredSOS, err := filter.Run()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(len(filteredSOS))
		fmt.Println(filteredSOS)
	}

}
