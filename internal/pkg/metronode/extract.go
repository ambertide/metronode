package metronode

import (
	common "ambertide/metronode/internal/pkg/metronode/common"
	referencetypes "ambertide/metronode/internal/pkg/metronode/reference_types"
	"errors"
)

func ExtractSystemFromString(systemName string, description string) (*common.Line, error) {
	extractors := map[string]func(string) *common.Line{
		"IZBAN": referencetypes.ExtractIzban,
	}

	extractor := extractors[systemName]

	if extractor != nil {
		return extractor(description), nil
	}

	return nil, errors.New("unable to extract this line")
}
