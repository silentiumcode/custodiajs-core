package extmodules

import cgowrapper "vnh1/extmodules/cgo_wrapper"

type ExternModuleImport struct {
}

type ExternModuleObject struct {
}

type ExternModuleEvent struct {
}

type ExternalModule struct {
	*cgowrapper.CGOWrappedLibModule
}
