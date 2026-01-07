package entity

import (
	"Testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T){
	g := NewGomegaWithT(t)

	books := &Books{
		Title: "My Example Title",
		Price: 100.00,
		Code: "BK123456",
	}

	g.Expect(books).To(Equal(books))

}
