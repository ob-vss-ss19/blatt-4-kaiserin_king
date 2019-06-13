# Protocol

## From Booking
### To Showing
* How many free seats are currently in a show (ID) - AskSeats
* Update amount of free seats in a show (ID) - UpdateSeats
* Ask if show (ID) does exist - Exist
### To User
* User (ID) gets inform that a booking (ID) got deleted - BookingDeleted
* Send User (ID) of new created bookins (ID) - CreatedBooking
* Ask if user (ID) does exist - Exist

## From Cinema
### To Showing
* Hall (ID) got deleted, need to remove all shows with this hall - FromHallDelete

## From Movie
### To Showing
* Movie (ID) got deleted, need to remove all shows with this movie - FromMovieDelete

## From Showing
### To Booking
* Show got deleted, need to delete all bookings whit this show - FromShowDelete
### To Cinema
* Ask if hall (ID) does exist - Exist
* Ask for number of seats in a hall (ID) - AskSeats
### To Movie
* Ask if movie (ID) does exist - Exist

## User
* User only gets informed