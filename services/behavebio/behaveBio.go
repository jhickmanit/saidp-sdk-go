package behavebio

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	sa "github.com/secureauthcorp/saidp-sdk-go"
)

/*
**********************************************************************
*   @author jhickman@secureauth.com
*
*  Copyright (c) 2017, SecureAuth
*  All rights reserved.
*
*    Redistribution and use in source and binary forms, with or without modification,
*    are permitted provided that the following conditions are met:
*
*    1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
*
*    2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer
*    in the documentation and/or other materials provided with the distribution.
*
*    3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived
*    from this software without specific prior written permission.
*
*    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
*    THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
*    CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
*    PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
*    LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
*    EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**********************************************************************
 */

const (
	jsEndpoint     = "/api/v1/behavebio/js"
	behaveEndpoint = "/api/v1/behavebio"
)

// Response :
//	Response struct that will be populated after the post request.
type Response struct {
	Source          string             `json:"src,omitempty"`
	BehaviorResults BehaviorBioResults `json:"BehaviorBioResults,omitempty"`
	Status          string             `json:"status"`
	Message         string             `json:"message"`
	RawJSON         string             `json:"-"`
	HTTPResponse    *http.Response     `json:"-"`
}

// Request :
//	Request struct to build required post parameters.
// Fields:
//	[Required] UserID: the username that you wish to submit a behaviobiometrics for.
//	BehaviorProfile: json string from the behavioBio javascript source.
//	HostAddress: the ip address of the user your are validating.
//	UserAgent: the user agent the user is using in the request.
//	FieldName: used for resetting the behavio profile.
//	FieldType: used for resetting the behavio profile.
//	DeviceType: used for the resetting the behavio profile.
type Request struct {
	UserID          string `json:"userId,omitempty"`
	BehaviorProfile string `json:"behaviorProfile,omitempty"`
	HostAddress     string `json:"hostAddress,omitempty"`
	UserAgent       string `json:"userAgent,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	FieldType       string `json:"fieldType,omitempty"`
	DeviceType      string `json:"deviceType,omitempty"`
}

// BehaviorBioResults :
//	Details for the behavior bio results
type BehaviorBioResults struct {
	TotalScore      float32   `json:"TotalScore,omitempty"`
	TotalConfidence float32   `json:"TotalConfidence,omitempty"`
	Device          string    `json:"Device,omitempty"`
	Results         []Results `json:"Results,omitempty"`
}

// Results :
//	Details for the behavior bio results
type Results struct {
	ControlID  string  `json:"ControlID,omitempty"`
	Score      float32 `json:"Score,omitempty"`
	Confidence float32 `json:"Confidence,omitempty"`
	Count      int32   `json:"Count,omitempty"`
}

// Get :
//	Executes a get to the behaviobio javascript endpoint.
// Parameters:
//	[Required] r: request struct to make get easy. should be empty for the use in get operations
//	[Required] c: passing in the client containing authorization and host information.
//	[Required] endpoint: the endpoint for the get request.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) Get(c *sa.Client, endpoint string) (*Response, error) {
	httpRequest, err := c.BuildGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	behaveResponse := new(Response)
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, behaveResponse); err != nil {
		return nil, err
	}
	behaveResponse.RawJSON = string(body)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	behaveResponse.HTTPResponse = httpResponse
	httpResponse.Body.Close()
	return behaveResponse, nil
}

// Post :
//	Executes a post to the behavioBio endpoint.
// Parameters:
//	[Required] r: should have all the required fields for the post type.
//	[Required] c: passing in the client containing authorization and host information.
//	[Required] endpoint: the endpoint for the post request.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) Post(c *sa.Client, endpoint string) (*Response, error) {
	jsonRequest, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	httpRequest, err := c.BuildPostRequest(endpoint, string(jsonRequest))
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	behaveResponse := new(Response)
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, behaveResponse); err != nil {
		return nil, err
	}
	behaveResponse.RawJSON = string(body)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	behaveResponse.HTTPResponse = httpResponse
	httpResponse.Body.Close()
	return behaveResponse, nil
}

// Put :
//	Executes a put to the behavioBio endpoint.
// Parameters:
//	[Required] r: should have all the required fields for the put type.
//	[Required] c: passing in the client containing authorization and host information.
//	[Required] endpoint: the endpoint for the put request.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) Put(c *sa.Client, endpoint string) (*Response, error) {
	jsonRequest, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	httpRequest, err := c.BuildPutRequest(endpoint, string(jsonRequest))
	if err != nil {
		return nil, err
	}
	httpResponse, err := c.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	behaveResponse := new(Response)
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, behaveResponse); err != nil {
		return nil, err
	}
	behaveResponse.RawJSON = string(body)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	behaveResponse.HTTPResponse = httpResponse
	httpResponse.Body.Close()
	return behaveResponse, nil
}

// GetBehaveJs :
//	Helper function for Get request to retrieve the behaviorbiometrics javascript source.
// Parameters:
//	[Required] c: passing in the client containing authorization and host information.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) GetBehaveJs(c *sa.Client) (*Response, error) {
	behaveResponse, err := r.Get(c, jsEndpoint)
	if err != nil {
		return nil, err
	}
	return behaveResponse, nil
}

// PostBehaveProfile :
//	Helper function for posting to the behaviobio endpoint.
// Parameters:
//	[Required] c: passing in the client containing authorization and host information.
//	[Required] userID: the username of the user you wish to post behaviobio to.
//	[Required] behaveProfile: from the behavioBio javascript, a json string.
//	[Required] hostAddress: the ip address of the user's host.
//	[Required] userAgent: the user's userAgent string from their request.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) PostBehaveProfile(c *sa.Client, userID string, behaveProfile string, hostAddress string, userAgent string) (*Response, error) {
	r.UserID = userID
	r.BehaviorProfile = behaveProfile
	r.HostAddress = hostAddress
	r.UserAgent = userAgent
	behaveResponse, err := r.Post(c, behaveEndpoint)
	if err != nil {
		return nil, err
	}
	return behaveResponse, nil
}

// ResetBehaveProfile :
//	Helper function for putting to the behaviobio endpoint.
// Parameters:
//	[Required] c: passing in the client containing authorization and host information.
//	[Required] userID: the username of the user you wish to put behaviobio to.
//	[Required] fieldName: name of field to reset (unique to application); or set to ALL for global reset.
//	[Required] fieldType: Type of field, either regulartext (actual values stored in profile) or anonymoustext (no actual values stored
//                 in profile, e.g. password entries); or set to ALL for global reset.
//	[Required] deviceType: Type of device used by user (Desktop or Mobile); or set to ALL for global reset.
// Returns:
//	Response: Struct marshaled from the Json response from the API endpoints.
//	Error: If an error is encountered, response will be nil and the error must be handled.
func (r *Request) ResetBehaveProfile(c *sa.Client, userID string, fieldName string, fieldType string, deviceType string) (*Response, error) {
	r.UserID = userID
	r.FieldName = fieldName
	r.FieldType = fieldType
	r.DeviceType = deviceType
	behaveResponse, err := r.Put(c, behaveEndpoint)
	if err != nil {
		return nil, err
	}
	return behaveResponse, nil
}

//IsSignatureValid :
//	Helper function to validate the SecureAuth Response signature in X-SA-SIGNATURE
// Parameters:
//	[Required] r: response struct with HTTPResponse
//	[Required] c: passing in the client with application id and key
// Returns:
//	bool: if true, computed signature matches X-SA-SIGNATURE. if false, computed signature does not match.
//	error: If an error is encountered, bool will be false and the error must be handled.
func (r *Response) IsSignatureValid(c *sa.Client) (bool, error) {
	saDate := r.HTTPResponse.Header.Get("X-SA-DATE")
	saSignature := r.HTTPResponse.Header.Get("X-SA-SIGNATURE")
	var buffer bytes.Buffer
	buffer.WriteString(saDate)
	buffer.WriteString("\n")
	buffer.WriteString(c.AppID)
	buffer.WriteString("\n")
	buffer.WriteString(r.RawJSON)
	raw := buffer.String()
	byteKey, _ := hex.DecodeString(c.AppKey)
	byteData := []byte(raw)
	sig := hmac.New(sha256.New, byteKey)
	sig.Write([]byte(byteData))
	computedSig := base64.StdEncoding.EncodeToString(sig.Sum(nil))
	if computedSig != saSignature {
		return false, nil
	}
	return true, nil
}
