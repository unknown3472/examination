package models

type UserBookingEmail struct {
    Email string `json:"email"`
}

type PostBookingRequest struct {
    UserId       string `json:"userId"`
    HotelId      string `json:"hotelId"`
    RoomType     string `json:"roomType"`
    CheckInDate  string `json:"checkInDate"`
    CheckOutDate string `json:"checkOutDate"`
}

type BookingResponse struct {
    UserId        string `json:"userId"`
    HotelId       string `json:"hotelId"`
    RoomType      string `json:"roomType"`
    CheckInDate   string `json:"checkInDate"`
    CheckOutDate  string `json:"checkOutDate"`
    TotalAmount   int32  `json:"totalAmount"`
    BookingId     string `json:"bookingId"`
    Status        string `json:"status"`
}

type PutBookingRequest struct {
    CheckInDate  string `json:"checkInDate"`
    CheckOutDate string `json:"checkOutDate"`
    Status       string `json:"status"`
}

type BookingRequest struct {
    BookingId string `json:"bookingId"`
}

type UserBookingRequest struct {
    UserId string `json:"userId"`
}

type DeleteBookingResponse struct {
    Message   string `json:"message"`
    BookingId string `json:"bookingId"`
}



type DescriptionRequest struct {
    HotelId string `json:"hotelId"`
}

type DescriptionResponse struct {
    HotelId   string  `json:"hotelId"`
    Name      string  `json:"name"`
    Location  string  `json:"location"`
    Rating    int32   `json:"rating"`
    Address   string  `json:"address"`
    Rooms     []Room  `json:"rooms"`
}

type Room struct {
    RoomType      string `json:"roomType"`
    PricePerNight int32  `json:"pricePerNight"`
    Availability  bool   `json:"availability"`
}

type GetHotelResponse struct {
    Hotels []HotelInfo `json:"hotels"`
}

type HotelInfo struct {
    HotelId  string `json:"hotelId"`
    Name     string `json:"name"`
    Location string `json:"location"`
    Rating   int32  `json:"rating"`
    Address  string `json:"address"`
}

type RoomAvailability struct {
    Rooms []RoomInfo `json:"rooms"`
}

type RoomInfo struct {
    RoomType       string `json:"roomType"`
    AvailableRooms int32  `json:"availableRooms"`
    PricePerNight  int32  `json:"pricePerNight"`
    HotelId        string `json:"hotelId"`
}

type VoidHotel struct{}


type UserRequest struct {
    UserName string `json:"user_name"`
    Password string `json:"password"`
    Email    string `json:"email"`
}

type UserResponse struct {
    UserID    string `json:"user_id"`
    UserName  string `json:"user_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}

type LoginRequest struct {
    UserName string `json:"user_name"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token     string `json:"token"`
    ExpiredAt string `json:"expired_at"`
}
