// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package stackconfigurationserver

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/labstack/echo/v4"
	cli "go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/commands"
	"go.thethings.network/lorawan-stack/pkg/component"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/web"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

// Config contains the Stack Configuration Server configuration.
type Config struct {
}

// StackConfigurationServer implements the Stack Configuration Server component.
type StackConfigurationServer struct {
	*component.Component
	config *Config
}

// Roles returns the roles that the Stack Configuration Server fulfills.
func (scs *StackConfigurationServer) Roles() []ttnpb.ClusterRole {
	return []ttnpb.ClusterRole{0}
}

// RegisterServices registers services provided by gcs at s.
func (scs *StackConfigurationServer) RegisterServices(_ *grpc.Server) {}

// RegisterHandlers registers gRPC handlers.
func (scs *StackConfigurationServer) RegisterHandlers(_ *runtime.ServeMux, _ *grpc.ClientConn) {}

// BuildCLIConfig generates CLI configuration from the Stack configuration.
func (scs *StackConfigurationServer) BuildCLIConfig() cli.Config {
	return cli.Config{
		GatewayServerEnabled: true,
	}
}

// RegisterRoutes registers the web frontend routes.
func (scs *StackConfigurationServer) RegisterRoutes(server *web.Server) {
	group := server.Group(ttnpb.HTTPAPIPrefix + "/scs")
	group.GET("/cli", func(c echo.Context) error {
		msg, err := yaml.Marshal(scs.BuildCLIConfig())
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "application/yml", msg)
	})
}

// New returns new *GatewayConfigurationServer.
func New(c *component.Component) (*StackConfigurationServer, error) {
	scs := &StackConfigurationServer{
		Component: c,
	}

	c.RegisterGRPC(scs)
	c.RegisterWeb(scs)
	return scs, nil
}
