// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://kerberos.io",
        "contact": {
            "name": "API Support",
            "url": "https://www.kerberos.io",
            "email": "support@kerberos.io"
        },
        "license": {
            "name": "Apache 2.0 - Commons Clause",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/camera/onvif/capabilities": {
            "post": {
                "description": "Will return the ONVIF capabilities for the specific camera.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will return the ONVIF capabilities for the specific camera.",
                "operationId": "camera-onvif-capabilities",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/gotopreset": {
            "post": {
                "description": "Will activate the desired ONVIF preset.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will activate the desired ONVIF preset.",
                "operationId": "camera-onvif-gotopreset",
                "parameters": [
                    {
                        "description": "OnvifPreset",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifPreset"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/inputs": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will get the digital inputs from the ONVIF device.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will get the digital inputs from the ONVIF device.",
                "operationId": "get-digital-inputs",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/login": {
            "post": {
                "description": "Try to login into ONVIF supported camera.",
                "tags": [
                    "onvif"
                ],
                "summary": "Try to login into ONVIF supported camera.",
                "operationId": "camera-onvif-login",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/outputs": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will get the relay outputs from the ONVIF device.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will get the relay outputs from the ONVIF device.",
                "operationId": "get-relay-outputs",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/outputs/{output}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will trigger the relay output from the ONVIF device.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will trigger the relay output from the ONVIF device.",
                "operationId": "trigger-relay-output",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Output",
                        "name": "output",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/pantilt": {
            "post": {
                "description": "Panning or/and tilting the camera using a direction (x,y).",
                "tags": [
                    "onvif"
                ],
                "summary": "Panning or/and tilting the camera.",
                "operationId": "camera-onvif-pantilt",
                "parameters": [
                    {
                        "description": "OnvifPanTilt",
                        "name": "panTilt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifPanTilt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/presets": {
            "post": {
                "description": "Will return the ONVIF presets for the specific camera.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will return the ONVIF presets for the specific camera.",
                "operationId": "camera-onvif-presets",
                "parameters": [
                    {
                        "description": "OnvifCredentials",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/onvif/zoom": {
            "post": {
                "description": "Zooming in or out the camera.",
                "tags": [
                    "onvif"
                ],
                "summary": "Zooming in or out the camera.",
                "operationId": "camera-onvif-zoom",
                "parameters": [
                    {
                        "description": "OnvifZoom",
                        "name": "zoom",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OnvifZoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/record": {
            "post": {
                "description": "Make a recording.",
                "tags": [
                    "camera"
                ],
                "summary": "Make a recording.",
                "operationId": "camera-record",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/restart": {
            "post": {
                "description": "Restart the agent.",
                "tags": [
                    "camera"
                ],
                "summary": "Restart the agent.",
                "operationId": "camera-restart",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/snapshot/base64": {
            "get": {
                "description": "Get a snapshot from the camera in base64.",
                "tags": [
                    "camera"
                ],
                "summary": "Get a snapshot from the camera in base64.",
                "operationId": "snapshot-base64",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/camera/snapshot/jpeg": {
            "get": {
                "description": "Get a snapshot from the camera in jpeg format.",
                "tags": [
                    "camera"
                ],
                "summary": "Get a snapshot from the camera in jpeg format.",
                "operationId": "snapshot-jpeg",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/camera/stop": {
            "post": {
                "description": "Stop the agent.",
                "tags": [
                    "camera"
                ],
                "summary": "Stop the agent.",
                "operationId": "camera-stop",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/camera/verify/{streamType}": {
            "post": {
                "description": "This method will validate a specific profile connection from an RTSP camera, and try to get the codec.",
                "tags": [
                    "camera"
                ],
                "summary": "Validate a specific RTSP profile camera connection.",
                "operationId": "verify-camera",
                "parameters": [
                    {
                        "enum": [
                            "primary",
                            "secondary"
                        ],
                        "type": "string",
                        "description": "Stream Type",
                        "name": "streamType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Camera Streams",
                        "name": "cameraStreams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CameraStreams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/config": {
            "get": {
                "description": "Get the current configuration.",
                "tags": [
                    "config"
                ],
                "summary": "Get the current configuration.",
                "operationId": "config",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "Update the current configuration.",
                "tags": [
                    "config"
                ],
                "summary": "Update the current configuration.",
                "operationId": "config",
                "parameters": [
                    {
                        "description": "Configuration",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/dashboard": {
            "get": {
                "description": "Get all information showed on the dashboard.",
                "tags": [
                    "general"
                ],
                "summary": "Get all information showed on the dashboard.",
                "operationId": "dashboard",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/days": {
            "get": {
                "description": "Get all days stored in the recordings directory.",
                "tags": [
                    "general"
                ],
                "summary": "Get all days stored in the recordings directory.",
                "operationId": "days",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/hub/verify": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will verify the hub connectivity.",
                "tags": [
                    "persistence"
                ],
                "summary": "Will verify the hub connectivity.",
                "operationId": "verify-hub",
                "parameters": [
                    {
                        "description": "Config",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/latest-events": {
            "post": {
                "description": "Get the latest recordings (events) from the recordings directory.",
                "tags": [
                    "general"
                ],
                "summary": "Get the latest recordings (events) from the recordings directory.",
                "operationId": "latest-events",
                "parameters": [
                    {
                        "description": "Event filter",
                        "name": "eventFilter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EventFilter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Get Authorization token.",
                "tags": [
                    "authentication"
                ],
                "summary": "Get Authorization token.",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Authentication"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Authorization"
                        }
                    }
                }
            }
        },
        "/api/onvif/verify": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will verify the ONVIF connectivity.",
                "tags": [
                    "onvif"
                ],
                "summary": "Will verify the ONVIF connectivity.",
                "operationId": "verify-onvif",
                "parameters": [
                    {
                        "description": "Camera Config",
                        "name": "cameraConfig",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.IPCamera"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/persistence/verify": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Will verify the persistence.",
                "tags": [
                    "persistence"
                ],
                "summary": "Will verify the persistence.",
                "operationId": "verify-persistence",
                "parameters": [
                    {
                        "description": "Config",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.APIResponse": {
            "type": "object",
            "properties": {
                "can_pan_tilt": {
                    "type": "boolean"
                },
                "can_zoom": {
                    "type": "boolean"
                },
                "data": {},
                "message": {},
                "ptz_functions": {}
            }
        },
        "models.Authentication": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Authorization": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.CameraStreams": {
            "type": "object",
            "properties": {
                "rtsp": {
                    "type": "string"
                },
                "sub_rtsp": {
                    "type": "string"
                }
            }
        },
        "models.Capture": {
            "type": "object",
            "properties": {
                "continuous": {
                    "type": "string"
                },
                "forwardwebrtc": {
                    "type": "string"
                },
                "fragmented": {
                    "type": "string"
                },
                "fragmentedduration": {
                    "type": "integer"
                },
                "ipcamera": {
                    "$ref": "#/definitions/models.IPCamera"
                },
                "liveview": {
                    "type": "string"
                },
                "maxlengthrecording": {
                    "type": "integer"
                },
                "motion": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pixelChangeThreshold": {
                    "type": "integer"
                },
                "postrecording": {
                    "type": "integer"
                },
                "prerecording": {
                    "type": "integer"
                },
                "raspicamera": {
                    "$ref": "#/definitions/models.RaspiCamera"
                },
                "recording": {
                    "type": "string"
                },
                "snapshots": {
                    "type": "string"
                },
                "transcodingresolution": {
                    "type": "integer"
                },
                "transcodingwebrtc": {
                    "type": "string"
                },
                "usbcamera": {
                    "$ref": "#/definitions/models.USBCamera"
                }
            }
        },
        "models.Config": {
            "type": "object",
            "properties": {
                "auto_clean": {
                    "type": "string"
                },
                "capture": {
                    "$ref": "#/definitions/models.Capture"
                },
                "cloud": {
                    "type": "string"
                },
                "condition_uri": {
                    "type": "string"
                },
                "dropbox": {
                    "$ref": "#/definitions/models.Dropbox"
                },
                "encryption": {
                    "$ref": "#/definitions/models.Encryption"
                },
                "friendly_name": {
                    "type": "string"
                },
                "heartbeaturi": {
                    "description": "obsolete",
                    "type": "string"
                },
                "hub_key": {
                    "type": "string"
                },
                "hub_private_key": {
                    "type": "string"
                },
                "hub_site": {
                    "type": "string"
                },
                "hub_uri": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "kstorage": {
                    "$ref": "#/definitions/models.KStorage"
                },
                "max_directory_size": {
                    "type": "integer"
                },
                "mqtt_password": {
                    "type": "string"
                },
                "mqtt_username": {
                    "type": "string"
                },
                "mqtturi": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "offline": {
                    "type": "string"
                },
                "region": {
                    "$ref": "#/definitions/models.Region"
                },
                "remove_after_upload": {
                    "type": "string"
                },
                "s3": {
                    "$ref": "#/definitions/models.S3"
                },
                "stunuri": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "timetable": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Timetable"
                    }
                },
                "timezone": {
                    "type": "string"
                },
                "turn_password": {
                    "type": "string"
                },
                "turn_username": {
                    "type": "string"
                },
                "turnuri": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Coordinate": {
            "type": "object",
            "properties": {
                "x": {
                    "type": "number"
                },
                "y": {
                    "type": "number"
                }
            }
        },
        "models.Dropbox": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "directory": {
                    "type": "string"
                }
            }
        },
        "models.Encryption": {
            "type": "object",
            "properties": {
                "enabled": {
                    "type": "string"
                },
                "fingerprint": {
                    "type": "string"
                },
                "private_key": {
                    "type": "string"
                },
                "recordings": {
                    "type": "string"
                },
                "symmetric_key": {
                    "type": "string"
                }
            }
        },
        "models.EventFilter": {
            "type": "object",
            "properties": {
                "number_of_elements": {
                    "type": "integer"
                },
                "timestamp_offset_end": {
                    "type": "integer"
                },
                "timestamp_offset_start": {
                    "type": "integer"
                }
            }
        },
        "models.IPCamera": {
            "type": "object",
            "properties": {
                "fps": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "onvif": {
                    "type": "string"
                },
                "onvif_password": {
                    "type": "string"
                },
                "onvif_username": {
                    "type": "string"
                },
                "onvif_xaddr": {
                    "type": "string"
                },
                "rtsp": {
                    "type": "string"
                },
                "sub_rtsp": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "models.KStorage": {
            "type": "object",
            "properties": {
                "access_key": {
                    "type": "string"
                },
                "cloud_key": {
                    "description": "old way, remove this",
                    "type": "string"
                },
                "directory": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "secret_access_key": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "models.OnvifCredentials": {
            "type": "object",
            "properties": {
                "onvif_password": {
                    "type": "string"
                },
                "onvif_username": {
                    "type": "string"
                },
                "onvif_xaddr": {
                    "type": "string"
                }
            }
        },
        "models.OnvifPanTilt": {
            "type": "object",
            "properties": {
                "onvif_credentials": {
                    "$ref": "#/definitions/models.OnvifCredentials"
                },
                "pan": {
                    "type": "number"
                },
                "tilt": {
                    "type": "number"
                }
            }
        },
        "models.OnvifPreset": {
            "type": "object",
            "properties": {
                "onvif_credentials": {
                    "$ref": "#/definitions/models.OnvifCredentials"
                },
                "preset": {
                    "type": "string"
                }
            }
        },
        "models.OnvifZoom": {
            "type": "object",
            "properties": {
                "onvif_credentials": {
                    "$ref": "#/definitions/models.OnvifCredentials"
                },
                "zoom": {
                    "type": "number"
                }
            }
        },
        "models.Polygon": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Coordinate"
                    }
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.RaspiCamera": {
            "type": "object",
            "properties": {
                "device": {
                    "type": "string"
                }
            }
        },
        "models.Rectangle": {
            "type": "object",
            "properties": {
                "x1": {
                    "type": "integer"
                },
                "x2": {
                    "type": "integer"
                },
                "y1": {
                    "type": "integer"
                },
                "y2": {
                    "type": "integer"
                }
            }
        },
        "models.Region": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "polygon": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Polygon"
                    }
                },
                "rectangle": {
                    "$ref": "#/definitions/models.Rectangle"
                }
            }
        },
        "models.S3": {
            "type": "object",
            "properties": {
                "bucket": {
                    "type": "string"
                },
                "proxy": {
                    "type": "string"
                },
                "proxyuri": {
                    "type": "string"
                },
                "publickey": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "secretkey": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Timetable": {
            "type": "object",
            "properties": {
                "end1": {
                    "type": "integer"
                },
                "end2": {
                    "type": "integer"
                },
                "start1": {
                    "type": "integer"
                },
                "start2": {
                    "type": "integer"
                }
            }
        },
        "models.USBCamera": {
            "type": "object",
            "properties": {
                "device": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Kerberos Agent API",
	Description:      "This is the API for using and configure Kerberos Agent.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
