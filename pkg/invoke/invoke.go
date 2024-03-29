package invoke

import (
	"fmt"
	"io"
	"net/http"
)

func Invoke(invocation *Invocation) error {
	req, err := http.NewRequest(invocation.Method, invocation.Url, nil)
	if err != nil {
		return err
	}

	for _, requestHeader := range invocation.RequestHeaders {
		req.Header.Add(requestHeader.Key, requestHeader.Value)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	invocation.Status = res.StatusCode
	for key, value := range res.Header {
		if len(value) == 0 {
			return fmt.Errorf("failed to map response header because there is no value supplied.")
		}
		invocation.ResponseHeaders = append(invocation.ResponseHeaders, Header{
			Key:   key,
			Value: value[0],
		})
	}

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	invocation.ResponseBody = string(resbody)

	return nil
}
