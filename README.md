# Slotify

Find available time slots in Go.

## Quick Start

See `example` directory and type `go run cmd/example1/main.go`

```go

 // This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
 span, err := slotify.NewSpan(now, now.Add(8*time.Hour))
 ...
 
 // This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
 type ScheduledEvent struct {
  Start time.Time
  End   time.Time
 }
 events := []*ScheduledEvent{
  ...
 }

 // Please define a function for the second argument that maps the values passed to `slotify.NewBlock` to the fields of your struct.
 blocks, err := slotify.NewBlocks(events, func(s *ScheduledEvent) (*slotify.Block, error) {
  return slotify.NewBlock(s.Start, s.End)
 })
 ...
 
 // Find available time slots!
 slots := slotify.Find(blocks, span)
 fmt.Println("Available Times (slots):\n" + slotify.ToString(slots))
```
