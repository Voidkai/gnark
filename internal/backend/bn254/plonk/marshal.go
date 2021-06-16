// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gnark DO NOT EDIT

package plonk

import (
	curve "github.com/consensys/gnark-crypto/ecc/bn254"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"

	"io"
)

// WriteTo writes binary encoding of ProofRaw to w
func (proof *ProofRaw) WriteTo(w io.Writer) (int64, error) {
	enc := curve.NewEncoder(w)

	if err := enc.Encode(proof.LROZH[:]); err != nil {
		return enc.BytesWritten(), err
	}
	if err := enc.Encode(proof.ZShift); err != nil {
		return enc.BytesWritten(), err
	}

	var n int64
	for i := 0; i < len(proof.CommitmentsLROZH); i++ {
		n2, err := w.Write(proof.CommitmentsLROZH[i].Marshal())
		n += int64(n2)
		if err != nil {
			return enc.BytesWritten() + n, err
		}
	}

	n2, err := w.Write(proof.BatchOpenings.Marshal())
	n += int64(n2)
	if err != nil {
		return enc.BytesWritten() + n, err
	}

	n2, err = w.Write(proof.OpeningZShift.Marshal())
	n += int64(n2)
	if err != nil {
		return enc.BytesWritten() + n, err
	}

	return n + enc.BytesWritten(), nil
}

// WriteTo writes binary encoding of PublicRaw to w
func (p *PublicRaw) WriteTo(w io.Writer) (n int64, err error) {
	n, err = p.DomainNum.WriteTo(w)
	if err != nil {
		return
	}

	n2, err := p.DomainH.WriteTo(w)
	if err != nil {
		return
	}
	n += n2

	enc := curve.NewEncoder(w)

	// note: p.Ql is of type Polynomial, which is handled by default binary.Write(...) op and doesn't
	// encode the size (nor does it convert from Montgomery to Regular form)

	toEncode := []interface{}{
		([]fr.Element)(p.Ql),
		([]fr.Element)(p.Qr),
		([]fr.Element)(p.Qm),
		([]fr.Element)(p.Qo),
		([]fr.Element)(p.Qk),
		p.Shifter[:],
		([]fr.Element)(p.LS1),
		([]fr.Element)(p.LS2),
		([]fr.Element)(p.LS3),
		([]fr.Element)(p.CS1),
		([]fr.Element)(p.CS2),
		([]fr.Element)(p.CS3),
		uint32(len(p.Permutation)),
		p.Permutation,
	}

	for _, v := range toEncode {
		if err := enc.Encode(v); err != nil {
			return n + enc.BytesWritten(), err
		}
	}

	return n + enc.BytesWritten(), nil
}

// ReadFrom  reads from binary representation in r into PublicRaw
func (p *PublicRaw) ReadFrom(r io.Reader) (int64, error) {
	n, err := p.DomainNum.ReadFrom(r)
	if err != nil {
		return n, err
	}

	n2, err := p.DomainH.ReadFrom(r)
	if err != nil {
		return n, err
	}
	n += n2

	dec := curve.NewDecoder(r)
	var lenPermutations uint32
	pShifter := make([]fr.Element, 2)
	toDecode := []interface{}{
		(*[]fr.Element)(&p.Ql),
		(*[]fr.Element)(&p.Qr),
		(*[]fr.Element)(&p.Qm),
		(*[]fr.Element)(&p.Qo),
		(*[]fr.Element)(&p.Qk),
		&pShifter,
		(*[]fr.Element)(&p.LS1),
		(*[]fr.Element)(&p.LS2),
		(*[]fr.Element)(&p.LS3),
		(*[]fr.Element)(&p.CS1),
		(*[]fr.Element)(&p.CS2),
		(*[]fr.Element)(&p.CS3),
		&lenPermutations,
	}

	for _, v := range toDecode {
		if err := dec.Decode(v); err != nil {
			return n + dec.BytesRead(), err
		}
	}

	p.Permutation = make([]int64, lenPermutations)
	if err := dec.Decode(&p.Permutation); err != nil {
		return n + dec.BytesRead(), err
	}

	p.Shifter[0] = pShifter[0]
	p.Shifter[1] = pShifter[1]

	return n + dec.BytesRead(), nil
}
