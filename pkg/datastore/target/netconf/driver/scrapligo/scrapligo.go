// Copyright 2024 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scrapligo

import (
	"fmt"

	"github.com/beevik/etree"
	scraplinetconf "github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/util"

	"github.com/sdcio/data-server/pkg/config"
	"github.com/sdcio/data-server/pkg/datastore/target/netconf/types"
)

type ScrapligoNetconfTarget struct {
	driver *scraplinetconf.Driver
}

// NewScrapligoNetconfTarget inits a new ScrapligoNetconfTarget which is already connected to the target node
func NewScrapligoNetconfTarget(cfg *config.SBI) (*ScrapligoNetconfTarget, error) {
	opts := []util.Option{
		options.WithAuthNoStrictKey(),
		options.WithNetconfForceSelfClosingTags(),
		options.WithTransportType("standard"),
		options.WithPort(int(cfg.Port)),
		options.WithTimeoutOps(cfg.Timeout),
	}

	if cfg.Credentials != nil {
		opts = append(opts,
			options.WithAuthUsername(cfg.Credentials.Username),
			options.WithAuthPassword(cfg.Credentials.Password),
		)
	}
	if cfg.NetconfOptions.PreferredNCVersion != "" {
		opts = append(opts,
			options.WithNetconfPreferredVersion(cfg.NetconfOptions.PreferredNCVersion),
		)
	}
	// init the netconf driver
	d, err := scraplinetconf.NewDriver(cfg.Address, opts...)
	if err != nil {
		return nil, err
	}

	err = d.Open()
	if err != nil {
		return nil, err
	}

	return &ScrapligoNetconfTarget{
		driver: d,
	}, nil
}

func (snt *ScrapligoNetconfTarget) Close() error {
	if snt == nil {
		return nil
	}
	if snt.driver == nil {
		return nil
	}
	return snt.driver.Close()
}

func (snt *ScrapligoNetconfTarget) IsAlive() bool {
	if snt == nil || snt.driver == nil {
		return false
	}
	if snt.driver.Transport == nil {
		return false
	}
	return snt.driver.Transport.IsAlive()
}

// EditConfig transforms the generalized EditConfig into the scrapligo implementation
func (snt *ScrapligoNetconfTarget) EditConfig(target string, config string) (*types.NetconfResponse, error) {
	// add the <config/> tag to the provided config data
	xdoc := fmt.Sprintf("<config>%s</config>", config)

	// send the edit config rpc
	resp, err := snt.driver.EditConfig(target, xdoc)
	if err != nil {
		return nil, err
	}
	if len(resp.ErrorMessages) > 0 {
		return nil, resp.Failed
	}

	// creating a new etree Document and parsing the netconf rpc result
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	// return the rpc result
	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) GetConfig(source string, filter string) (*types.NetconfResponse, error) {
	// prepare the filter to hand it to scrapli
	filterDoc := createFilterOption(filter)

	// execute the GetConfig rpc
	resp, err := snt.driver.GetConfig(source, filterDoc, options.WithNetconfForceSelfClosingTags())
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}

	// creating a new etree Document and parsing the netconf rpc result
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	// the actual config is contained under /rpc-reply/data/ in the result document.
	// so we are extracting that portion
	newRootXpath := "/rpc-reply/data"
	r := x.FindElement(newRootXpath)
	if r == nil {
		return nil, fmt.Errorf("unable to find %q in %s", newRootXpath, resp.Result)
	}

	// making the new retrieved path the new root element of the xml doc
	x.SetRoot(r)

	// return the rpc result
	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) Commit() error {
	// execute the Commit rpc
	resp, err := snt.driver.Commit()
	if err != nil {
		return err
	}
	if resp.Failed != nil {
		return resp.Failed
	}
	return nil
}

func (snt *ScrapligoNetconfTarget) Discard() error {
	resp, err := snt.driver.Discard()
	if err != nil {
		return err
	}
	if resp.Failed != nil {
		return resp.Failed
	}
	return nil
}

func (snt *ScrapligoNetconfTarget) Get(filter string) (*types.NetconfResponse, error) {
	resp, err := snt.driver.Get(filter)
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	newRootXpath := "/rpc-reply/data"
	r := x.FindElement(newRootXpath)
	if r == nil {
		return nil, fmt.Errorf("unable to find %q in %s", newRootXpath, resp.Result)
	}

	x.SetRoot(r)

	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) Lock(target string) (*types.NetconfResponse, error) {
	resp, err := snt.driver.Lock(target)
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) Unlock(target string) (*types.NetconfResponse, error) {
	resp, err := snt.driver.Unlock(target)
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) Validate(source string) (*types.NetconfResponse, error) {
	resp, err := snt.driver.Validate(source)
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	return types.NewNetconfResponse(x), nil
}

func (snt *ScrapligoNetconfTarget) CopyConfig(source, target string) (*types.NetconfResponse, error) {
	resp, err := snt.driver.CopyConfig(source, target)
	if err != nil {
		return nil, err
	}
	if resp.Failed != nil {
		return nil, resp.Failed
	}
	x := etree.NewDocument()
	err = x.ReadFromString(resp.Result)
	if err != nil {
		return nil, err
	}

	return types.NewNetconfResponse(x), nil
}

// createFilterOption is a helper function that populates the Filter field for the internal Scrapligo RPC instantiation
func createFilterOption(filter string) util.Option {
	return func(x interface{}) error {
		oo, ok := x.(*scraplinetconf.OperationOptions)

		if !ok {
			return util.ErrIgnoredOption
		}
		oo.Filter = filter
		return nil
	}
}
