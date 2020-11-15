//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package pbgo

import (
	"fmt"
	"strings"

	google_protobuf "github.com/gogo/protobuf/types"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// GetConfBootstrap implements ves.io/stdlib/confreader.Conf interface
func (c *Conf) GetConfBootstrap(ef db.NewEntryFunc) ([]db.Entry, error) {
	retEnts := []db.Entry{}
	for _, bs := range c.GetBootstrap() {
		// convert schema.ves.io/ves.io.examplesvc.objectone.Object to ves.io.examplesvc.objectone.Object
		sl := strings.Split(bs.TypeUrl, "/")
		ot := sl[len(sl)-1]
		ent, err := ef(ot, db.OpWithSerializedBytes(bs.Value))
		if err != nil {
			return nil, errors.Wrap(err, "NewEntry")
		}
		retEnts = append(retEnts, ent)
	}
	return retEnts, nil
}

// GetConfOverrides implements ves.io/stdlib/confreader.Conf interface
func (c *Conf) GetConfOverrides(ef db.NewEntryFunc) ([]db.Entry, error) {
	retEnts := []db.Entry{}
	for _, bs := range c.GetOverrides() {
		// convert schema.ves.io/ves.io.examplesvc.objectone.Object to ves.io.examplesvc.objectone.Object
		sl := strings.Split(bs.TypeUrl, "/")
		ot := sl[len(sl)-1]
		ent, err := ef(ot, db.OpWithSerializedBytes(bs.Value))
		if err != nil {
			return nil, errors.Wrap(err, "NewEntry")
		}
		retEnts = append(retEnts, ent)
	}
	return retEnts, nil
}

// GetDmnConfBootstrap implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfBootstrap() []*google_protobuf.Any {
	return c.GetBootstrap()
}

// GetDmnConfOverrides implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfOverrides() []*google_protobuf.Any {
	return c.GetOverrides()
}

// GetDmnConfGrpcPort() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfGrpcPort() int32 {

	return c.GetGrpcPort()
}

// GetDmnConfRestPort() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfRestPort() int32 {
	return c.GetRestPort()
}

// GetDmnConfGrpcTLSPort() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfGrpcTLSPort() int32 {
	return c.GetGrpcTLSPort()
}

// GetDmnConfRestTLSPort() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfRestTLSPort() int32 {
	return c.GetRestTLSPort()
}

// GetDmnConfServerTLS() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfServerTLS() *svcfw.TLSConfig {
	svrParams := c.GetTls().GetServerParams()
	if svrParams == nil {
		return nil
	}
	if len(svrParams.TlsCertificates) == 0 {
		return nil
	}
	tlsCfg := &svcfw.TLSConfig{
		Key:    svrParams.TlsCertificates[0].PrivateKeyUrl,
		Cert:   svrParams.TlsCertificates[0].CertificateUrl,
		CACert: svrParams.TrustedCaUrl,
	}
	return tlsCfg
}

// GetDmnConfClientTLS() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfClientTLS() *svcfw.TLSConfig {
	clParams := c.GetTls().GetClientParams()
	if clParams == nil {
		return nil
	}
	if len(clParams.TlsCertificates) == 0 {
		return nil
	}
	tlsCfg := &svcfw.TLSConfig{
		Key:    clParams.TlsCertificates[0].PrivateKeyUrl,
		Cert:   clParams.TlsCertificates[0].CertificateUrl,
		CACert: clParams.TrustedCaUrl,
	}
	return tlsCfg
}

// GetDmnConfEtcdKeyPrefix() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfEtcdKeyPrefix() string {
	return c.GetEtcdKeyPrefix()
}

// GetDmnConfEtcdServerURLs() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfEtcdServerURLs() []string {
	return c.GetEtcdServerURLs()
}

// GetDmnConfStatusEtcdServerURLs() implements ves.io/stdlib/svcfw.DaemonConf interface
func (c *Conf) GetDmnConfStatusEtcdServerURLs() []string {
	return []string{}

}