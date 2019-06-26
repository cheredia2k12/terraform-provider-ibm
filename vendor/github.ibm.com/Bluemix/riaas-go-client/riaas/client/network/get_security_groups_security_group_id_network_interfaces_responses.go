// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSecurityGroupsSecurityGroupIDNetworkInterfacesReader is a Reader for the GetSecurityGroupsSecurityGroupIDNetworkInterfaces structure.
type GetSecurityGroupsSecurityGroupIDNetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesOK creates a GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK with default headers values
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesOK() *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK {
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK{}
}

/*GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK struct {
	Payload *models.GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody
}

func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/network_interfaces][%d] getSecurityGroupsSecurityGroupIdNetworkInterfacesOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError creates a GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError with default headers values
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError() *GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError {
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError{}
}

/*GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError handles this case with default header values.

error
*/
type GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security_groups/{security_group_id}/network_interfaces][%d] getSecurityGroupsSecurityGroupIdNetworkInterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}