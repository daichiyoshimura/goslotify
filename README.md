# TimeSlots

Find available time slots in Go.

Searching for “available time” is done in various situations.

-	Check the “available time” of your work schedule.
-	Check the “available time” of the restaurant you want to use for a meal.
-	Check the “available time” of a hotel room you want to stay at during your trip.
-	Check the “available time” of a car-sharing service you want to use on your day off.

If your application has information about “scheduled events” that are already booked, this library will provide “available time” easily!

## Quick Start

See `example` directory and type `go run examples/prepared/main.go `

```go
 // This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
 span, err := timeslots.NewSpan(now, now.Add(8*time.Hour))
 ...
 
 // This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
 type ScheduledEvent struct {
  Start time.Time
  End   time.Time
 }
 events := []*ScheduledEvent{
  ...
 }

 // Please define a function for the second argument that maps the values passed to `timeslots.NewBlock` to the fields of your struct.
 blocks, err := timeslots.NewBlocks(events, func(s *ScheduledEvent) (*timeslots.Block, error) {
  return timeslots.NewBlock(s.Start, s.End)
 })
 ...
 
 // Find available time slots!
 slots := timeslots.Find(blocks, span)
 fmt.Println("Available Times (slots):\n" + timeslots.ToString(slots))
```

Result 

```sh
Search Span:
2024-09-26 19:42:44, 2024-09-27 03:42:44

Scheduled Events (blocks):
2024-09-26 18:42:44, 2024-09-26 20:42:44
2024-09-26 21:42:44, 2024-09-26 23:42:44
2024-09-26 22:42:44, 2024-09-27 00:42:44

Available Time Slots(slots):
2024-09-26 20:42:44, 2024-09-26 21:42:44
2024-09-27 00:42:44, 2024-09-27 03:42:44
```

## Note

- Do not mix schedules (Blocks) held by different entities. (You should know a smarter way to handle this.)
- In the implementation of the Period, ensure that the start time is always before the end time. (Be mindful of cases where this may inadvertently happen.)

## Specification

- It is acceptable for different schedules (Blocks) to overlap.
