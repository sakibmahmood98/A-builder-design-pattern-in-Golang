package main

type normalBuilder struct {
    data1 int
    data2 int
    data3 int
}

func newNormalBuilder() *normalBuilder {
    return &normalBuilder{}
}

func (b *normalBuilder) setData1() {
    b.data1 = 2
}

func (b *normalBuilder) add() {
    b.data2 = b.data1++
}

func (b *normalBuilder) remove() {
    b.data3 = b.data1--
}

func (b *normalBuilder) getSomething() something {
    return something{
        data1:   b.data1,
        data2:   b.data2,
        data3:   b.data3,
    }
}
