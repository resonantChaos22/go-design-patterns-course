package main

type Document struct {
}

type Machine interface {
	Print(Document)
	Fax(Document)
	Scan(Document)
}

// For Multifunction Printer, everything is fine
type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(doc Document) {

}

func (m *MultiFunctionPrinter) Fax(doc Document) {

}

func (m *MultiFunctionPrinter) Scan(doc Document) {

}

// but the old printer can only implement Print() function, so if we forcefully implement other functions, it could lead to confusions
type OldPrinter struct{}

// To Fix that, just create multiple interfaces and then we can combine them
type Printer interface {
	Print(Document)
}

type Faxer interface {
	Fax(Document)
}

type Scanner interface {
	Scan(Document)
}

type MultiFunctionDevice interface {
	Printer
	Faxer
	Scanner
}

func TestIAP() {

}
