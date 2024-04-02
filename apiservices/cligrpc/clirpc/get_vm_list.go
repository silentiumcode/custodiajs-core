package clirpc

import (
	"context"
	"strings"
	"vnh1/grpc/cligrpc"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CliGrpcServer) ListVMs(ctx context.Context, _ *emptypb.Empty) (*cligrpc.VmListResponse, error) {
	// Die Werte werden abgeabreitet
	entry := []*cligrpc.VmListEntry{}
	for _, item := range s.core.GetAllVMs() {
		// Die geteilten Funktionen werden abgerufen
		sharf := make([]string, 0)
		for _, sharfnc := range item.GetAllSharedFunctions() {
			sharf = append(sharf, sharfnc.GetName())
		}

		// Die Erlaubten Domains werden abgerufen
		allowedDomains := []string{}
		for _, item := range item.GetWhitelist() {
			allowedDomains = append(allowedDomains, item.URL())
		}

		// Der Eintrag wird hinzugefügt
		entry = append(entry, &cligrpc.VmListEntry{
			Name:            item.GetVMName(),
			Id:              strings.ToUpper(item.GetFingerprint()),
			State:           uint32(item.GetState()),
			StartTime:       item.GetStartingTimestamp(),
			NodeJsModules:   item.GetVMModuleNames(),
			DomainWhiteList: allowedDomains,
			UsedHostKeyIds:  item.GetMemberCertKeyIds(),
			SharedFunctions: sharf,
		})
	}

	// Der Rückgabewert wird erzeugt
	returnValue := &cligrpc.VmListResponse{Vms: entry}

	// Die Daten werden zurückgesendet
	return returnValue, nil
}