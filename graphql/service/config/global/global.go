/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package global

import (
	_ "golang.org/x/tools/imports"
)

//go:generate go run ./generator/resolver generated_resolver.go
//go:generate sed -i "s/func (r \*Resolver) () bool/func (r *Resolver) SoMarkFromDaeSet() bool/" generated_resolver.go
//go:generate go run ./generator/input generated_input.go
//go:generate sed -i "s/i\. != nil/i.SoMarkFromDaeSet != nil/; s/i\. /i.SoMarkFromDaeSet /; s/\\t\\*bool/SoMarkFromDaeSet\\t*bool/" generated_input.go
//go:generate go run -mod=mod golang.org/x/tools/cmd/goimports -w generated_resolver.go generated_input.go
//go:generate go fmt
