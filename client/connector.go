package client

type FeedConnector interface {
	Connect() error
	ReadMessage() ([]byte, error)
	Close() error
}
