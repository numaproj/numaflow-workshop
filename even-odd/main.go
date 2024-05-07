package main

import (
	"context"
	"log"
	"strconv"

	"github.com/numaproj/numaflow-go/pkg/mapper"
)

func handle(ctx context.Context, keys []string, d mapper.Datum) mapper.Messages {
	msg := d.Value()
	// If msg is not an integer, drop it, otherwise return it with "even" or "odd" tag.
	if num, err := strconv.Atoi(string(msg)); err != nil {
		return mapper.MessagesBuilder().Append(mapper.MessageToDrop())
	} else if num%2 == 0 {
		// TODO: Do something different
		return mapper.MessagesBuilder().Append(mapper.NewMessage(msg).WithTags([]string{"even-tag"}))
	} else {
		return mapper.MessagesBuilder().Append(mapper.NewMessage(msg).WithTags([]string{"odd-tag"}))
	}
}

func main() {
	err := mapper.NewServer(mapper.MapperFunc(handle)).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start the function server: ", err)
	}
}
