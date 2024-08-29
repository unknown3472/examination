// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/bookings": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new booking in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Add a new booking",
                "parameters": [
                    {
                        "description": "Booking details",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PostBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.BookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/bookings/{bookingId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the details of a specific booking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Retrieve a booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "bookingId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Booking not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update the details of an existing booking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Update an existing booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "bookingId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated booking details",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PutBookingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BookingResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Booking not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Remove a specific booking from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Delete a booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking ID",
                        "name": "bookingId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeleteBookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Booking not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/hotels": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a list of all hotels from the system. This endpoint requires a valid JWT token in the Authorization header.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotel"
                ],
                "summary": "Get all hotels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetHotelResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Authorization header required or token is invalid",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/hotels/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve detailed description of a specific hotel. This endpoint requires a valid JWT token in the Authorization header.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotel"
                ],
                "summary": "Get hotel description",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DescriptionResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid hotel ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Authorization header required or token is invalid",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/hotels/{id}/rooms/availability": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve the available rooms for a specific hotel. This endpoint requires a valid JWT token in the Authorization header.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotel"
                ],
                "summary": "Get rooms of a hotel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RoomAvailability"
                        }
                    },
                    "400": {
                        "description": "Invalid hotel ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - Authorization header required or token is invalid",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "post": {
                "description": "Registers a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/login": {
            "post": {
                "description": "Logs in a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/{userId}/bookings": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all bookings associated with a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booking"
                ],
                "summary": "Retrieve bookings for a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BookingResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BookingResponse": {
            "type": "object",
            "properties": {
                "bookingId": {
                    "type": "string"
                },
                "checkInDate": {
                    "type": "string"
                },
                "checkOutDate": {
                    "type": "string"
                },
                "hotelId": {
                    "type": "string"
                },
                "roomType": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "totalAmount": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.DeleteBookingResponse": {
            "type": "object",
            "properties": {
                "bookingId": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.DescriptionResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "hotelId": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Room"
                    }
                }
            }
        },
        "models.GetHotelResponse": {
            "type": "object",
            "properties": {
                "hotels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.HotelInfo"
                    }
                }
            }
        },
        "models.HotelInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "hotelId": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "expired_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.PostBookingRequest": {
            "type": "object",
            "properties": {
                "checkInDate": {
                    "type": "string"
                },
                "checkOutDate": {
                    "type": "string"
                },
                "hotelId": {
                    "type": "string"
                },
                "roomType": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.PutBookingRequest": {
            "type": "object",
            "properties": {
                "checkInDate": {
                    "type": "string"
                },
                "checkOutDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Room": {
            "type": "object",
            "properties": {
                "availability": {
                    "type": "boolean"
                },
                "pricePerNight": {
                    "type": "integer"
                },
                "roomType": {
                    "type": "string"
                }
            }
        },
        "models.RoomAvailability": {
            "type": "object",
            "properties": {
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RoomInfo"
                    }
                }
            }
        },
        "models.RoomInfo": {
            "type": "object",
            "properties": {
                "availableRooms": {
                    "type": "integer"
                },
                "hotelId": {
                    "type": "string"
                },
                "pricePerNight": {
                    "type": "integer"
                },
                "roomType": {
                    "type": "string"
                }
            }
        },
        "models.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Api-gateway service",
	Description:      "Api-gateway service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
