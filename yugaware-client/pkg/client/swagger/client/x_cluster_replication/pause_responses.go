// Code generated by go-swagger; DO NOT EDIT.

package x_cluster_replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// PauseReader is a Reader for the Pause structure.
type PauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPauseOK creates a PauseOK with default headers values
func NewPauseOK() *PauseOK {
	return &PauseOK{}
}

/* PauseOK describes a response with status code 200, with default header values.

successful operation
*/
type PauseOK struct {
	Payload *models.YBPTask
}

func (o *PauseOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/customers/{cUUID}/universes/{uniUUID}/pause_xcluster_replication][%d] pauseOK  %+v", 200, o.Payload)
}
func (o *PauseOK) GetPayload() *models.YBPTask {
	return o.Payload
}

func (o *PauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
