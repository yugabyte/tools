// Code generated by go-swagger; DO NOT EDIT.

package customer_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yugabyte/yb-tools/yugaware-client/pkg/client/swagger/models"
)

// DeleteCustomerConfigV2Reader is a Reader for the DeleteCustomerConfigV2 structure.
type DeleteCustomerConfigV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomerConfigV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomerConfigV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCustomerConfigV2OK creates a DeleteCustomerConfigV2OK with default headers values
func NewDeleteCustomerConfigV2OK() *DeleteCustomerConfigV2OK {
	return &DeleteCustomerConfigV2OK{}
}

/* DeleteCustomerConfigV2OK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteCustomerConfigV2OK struct {
	Payload *models.YBPTask
}

func (o *DeleteCustomerConfigV2OK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/customers/{cUUID}/configs/{configUUID}/delete][%d] deleteCustomerConfigV2OK  %+v", 200, o.Payload)
}
func (o *DeleteCustomerConfigV2OK) GetPayload() *models.YBPTask {
	return o.Payload
}

func (o *DeleteCustomerConfigV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.YBPTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
