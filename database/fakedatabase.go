package database

type FakeEntry struct {
	Client     string
	Username   string
	PrivateKey []byte
	PublicKey  []byte
}

type FakeDatabase struct {
	Data      []FakeEntry
	Connected bool
}

func (f *FakeDatabase) Connect(db_name string, db_host string,
	db_user string, db_pass string) error {
	f.Connected = true
	return nil
}
