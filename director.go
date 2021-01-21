package main

type director struct {
    builder iBuilder
}

func newDirector(b iBuilder) *director {
    return &director{
        builder: b,
    }
}

func (d *director) setBuilder(b iBuilder) {
    d.builder = b
}

func (d *director) buildSomething() something {
    d.builder.setData1()
    d.builder.add()
    d.builder.remove()
    return d.builder.getSomething()
}
