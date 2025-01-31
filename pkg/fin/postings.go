package fin

import (
	db "github.com/clstb/phi/pkg/db/core"
	"github.com/clstb/phi/pkg/pb"
)

// Postings is a slice of posting.
type Postings []Posting

// PostingsFromPB creates a new posting slice from it's protobuf representation.
func PostingsFromPB(pb *pb.Postings) (Postings, error) {
	var postings Postings
	for _, v := range pb.Data {
		posting, err := PostingFromPB(v)
		if err != nil {
			return Postings{}, err
		}
		postings = append(postings, posting)
	}

	return postings, nil
}

// PostingsFromDB creates a new posting slice from it's database representation.
func PostingsFromDB(db ...db.Posting) (Postings, error) {
	var postings Postings
	for _, p := range db {
		posting, err := PostingFromDB(p)
		if err != nil {
			return Postings{}, err
		}
		postings = append(postings, posting)
	}

	return postings, nil
}

// PB marshalls the postings into their protobuf representation.
func (p Postings) PB() []*pb.Posting {
	var postings []*pb.Posting
	for _, posting := range p {
		postings = append(postings, posting.PB())
	}

	return postings
}
