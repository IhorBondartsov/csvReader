package client

import (
	"context"

	"github.com/IhorBondartsov/csvReader/entity"

	"github.com/IhorBondartsov/datasaver/web/myproto/pb"
	"google.golang.org/grpc"
)

// ClientForSaver - clent for data saver
type ClientForSaver struct {
	Conn pb.CSVSenderClient
}

// NewClient create client with connection to server
func NewConnection(uml string) (*ClientForSaver, error) {
	conn, err := grpc.Dial(uml, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewCSVSenderClient(conn)
	return &ClientForSaver{
		Conn: client,
	}, nil
}

// SendPersonData - it is decorator for API.Save
func (c *ClientForSaver) SendPersonData(data entity.PersonData) error {
	ctx := context.Background()
	_, err := c.Conn.Save(ctx, &pb.PersonData{
		Email:        data.Email,
		ID:           int32(data.Id),
		MobileNumber: data.MobileNumber,
		Name:         data.Name,
	})

	if err != nil {
		return err
	}
	return nil
}
