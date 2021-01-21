package main

type iBuilder interface {
    setData1()
    add()
    remove()

    getSomething() something
}

func getBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalBuilder{}
    }

    return nil
}
