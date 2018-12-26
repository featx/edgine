package supply

type BaseSupplier interface {
	Before()
}

type BasicSupplier struct {
}

func (basic *BasicSupplier) HandleError() {

}
