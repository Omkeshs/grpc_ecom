package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func GetQueryIDs(r *http.Request, queryParamName string) ([]int32, error) {

	var IDs []int32
	params := r.URL.Query()
	IDsStr := params.Get(queryParamName)

	IDList := strings.Split(IDsStr, ",")
	for _, param := range IDList {
		param = strings.TrimSpace(param)
		if param != "" {
			id, err := strconv.Atoi(param)
			if err != nil {
				return IDs, fmt.Errorf("invalid query params : %s", queryParamName)
			}
			IDs = append(IDs, int32(id))
		}
	}

	return IDs, nil
}

func GetPathID(r *http.Request, param string) (int32, error) {
	vars := mux.Vars(r)

	idStr, exists := vars[param]
	if !exists {
		err := errors.New(fmt.Sprintf("key not found %s", param))
		err = errors.Wrapf(err, "Message=Failed to get value for key ID")
		return 0, err
	}

	id, conversionError := strconv.ParseInt(idStr, 10, 32)
	if conversionError != nil || id <= 0 {
		return 0, errors.New(fmt.Sprintf("invalid %s", param))
	}

	return int32(id), nil
}
