package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Mutability string

const (
	Mutability_Immutable Mutability = "Immutable"
	Mutability_ReadOnly  Mutability = "ReadOnly"
	Mutability_ReadWrite Mutability = "ReadWrite"
	Mutability_WriteOnly Mutability = "WriteOnly"
)

func PossibleValuesForMutability() []string {
	return []string{
		string(Mutability_Immutable),
		string(Mutability_ReadOnly),
		string(Mutability_ReadWrite),
		string(Mutability_WriteOnly),
	}
}

func (s *Mutability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMutability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMutability(input string) (*Mutability, error) {
	vals := map[string]Mutability{
		"immutable": Mutability_Immutable,
		"readonly":  Mutability_ReadOnly,
		"readwrite": Mutability_ReadWrite,
		"writeonly": Mutability_WriteOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Mutability(input)
	return &out, nil
}