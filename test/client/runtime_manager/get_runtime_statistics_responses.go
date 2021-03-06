// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// GetRuntimeStatisticsReader is a Reader for the GetRuntimeStatistics structure.
type GetRuntimeStatisticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeStatisticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRuntimeStatisticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuntimeStatisticsOK creates a GetRuntimeStatisticsOK with default headers values
func NewGetRuntimeStatisticsOK() *GetRuntimeStatisticsOK {
	return &GetRuntimeStatisticsOK{}
}

/*GetRuntimeStatisticsOK handles this case with default header values.

A successful response.
*/
type GetRuntimeStatisticsOK struct {
	Payload *models.OpenpitrixGetRuntimeStatisticsResponse
}

func (o *GetRuntimeStatisticsOK) Error() string {
	return fmt.Sprintf("[GET /v1/runtimes/statistics][%d] getRuntimeStatisticsOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeStatisticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixGetRuntimeStatisticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
