// Code generated by go-swagger; DO NOT EDIT.

package installations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/michalq/go-airly-api-client/models"
)

// NearestUsingGETReader is a Reader for the NearestUsingGET structure.
type NearestUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NearestUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNearestUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewNearestUsingGETOK creates a NearestUsingGETOK with default headers values
func NewNearestUsingGETOK() *NearestUsingGETOK {
	return &NearestUsingGETOK{}
}

/*NearestUsingGETOK handles this case with default header values.

List of Installations
*/
type NearestUsingGETOK struct {
	Payload []*models.Installation
}

func (o *NearestUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /v2/installations/nearest][%d] nearestUsingGETOK  %+v", 200, o.Payload)
}

func (o *NearestUsingGETOK) GetPayload() []*models.Installation {
	return o.Payload
}

func (o *NearestUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
