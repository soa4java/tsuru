// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/globocom/config"
	"github.com/globocom/tsuru/app"
	"github.com/globocom/tsuru/db"
	ttesting "github.com/globocom/tsuru/testing"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
	conn        *db.Storage
	provisioner *ttesting.FakeProvisioner
}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
	err := config.ReadConfigFile("../etc/tsuru.conf")
	c.Assert(err, IsNil)
	config.Set("database:url", "127.0.0.1:27017")
	config.Set("database:name", "tsuru_collector_test")
	s.conn, err = db.Conn()
	c.Assert(err, IsNil)
	s.provisioner = ttesting.NewFakeProvisioner()
	app.Provisioner = s.provisioner
	config.Set("queue-server", "127.0.0.1:0")
}

func (s *S) TearDownSuite(c *C) {
	s.conn.Apps().Database.DropDatabase()
	s.conn.Close()
}

func (s *S) TearDownTest(c *C) {
	_, err := s.conn.Apps().RemoveAll(nil)
	c.Assert(err, IsNil)
	s.provisioner.Reset()
}
