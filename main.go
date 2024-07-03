package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/veraison/corim/comid"
	"github.com/veraison/corim/encoding"
)

type EntityExtensions struct {
	Email string `cbor:"-1,keyasint,omitempty" json:"email,omitempty"`
}

func (o EntityExtensions) ValidEntity(val *comid.Entity) error {
	_, err := uuid.Parse(val.EntityName.String())
	if err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}

	return nil
}

var sampleText = `
{
      "name": "31fb5abf-023e-4992-aa4e-95f9c1503bfa",
      "regid": "https://acme.example",
      "email": "info@acme.com",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
}
`

func main() {
	var entity comid.Entity
	entity.RegisterExtensions(&EntityExtensions{})

	if err := json.Unmarshal([]byte(sampleText), &entity); err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	if err := entity.Valid(); err != nil {
		log.Fatalf("failed to validate: %s", err.Error())
	} else {
		fmt.Println("validation succeeded")
	}

	email := entity.Extensions.MustGetString("email")
	fmt.Printf("entity email: %s\n", email)

	exts := entity.GetExtensions().(*EntityExtensions)
	fmt.Printf("also entity email: %s\n", exts.Email)

	out, err := encoding.SerializeStructToJSON(entity)
	if err != nil {
		log.Fatalf("could not marshal entity: %s", err.Error())
	}
	fmt.Printf("marshaled: %s", string(out))
}
