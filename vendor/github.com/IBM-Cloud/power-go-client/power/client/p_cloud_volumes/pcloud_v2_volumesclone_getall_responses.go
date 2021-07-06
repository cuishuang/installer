// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumescloneGetallReader is a Reader for the PcloudV2VolumescloneGetall structure.
type PcloudV2VolumescloneGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudV2VolumescloneGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudV2VolumescloneGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudV2VolumescloneGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudV2VolumescloneGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudV2VolumescloneGetallOK creates a PcloudV2VolumescloneGetallOK with default headers values
func NewPcloudV2VolumescloneGetallOK() *PcloudV2VolumescloneGetallOK {
	return &PcloudV2VolumescloneGetallOK{}
}

/*PcloudV2VolumescloneGetallOK handles this case with default header values.

OK
*/
type PcloudV2VolumescloneGetallOK struct {
	Payload *models.VolumesClones
}

func (o *PcloudV2VolumescloneGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallBadRequest creates a PcloudV2VolumescloneGetallBadRequest with default headers values
func NewPcloudV2VolumescloneGetallBadRequest() *PcloudV2VolumescloneGetallBadRequest {
	return &PcloudV2VolumescloneGetallBadRequest{}
}

/*PcloudV2VolumescloneGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudV2VolumescloneGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallNotFound creates a PcloudV2VolumescloneGetallNotFound with default headers values
func NewPcloudV2VolumescloneGetallNotFound() *PcloudV2VolumescloneGetallNotFound {
	return &PcloudV2VolumescloneGetallNotFound{}
}

/*PcloudV2VolumescloneGetallNotFound handles this case with default header values.

Not Found
*/
type PcloudV2VolumescloneGetallNotFound struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallInternalServerError creates a PcloudV2VolumescloneGetallInternalServerError with default headers values
func NewPcloudV2VolumescloneGetallInternalServerError() *PcloudV2VolumescloneGetallInternalServerError {
	return &PcloudV2VolumescloneGetallInternalServerError{}
}

/*PcloudV2VolumescloneGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudV2VolumescloneGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
