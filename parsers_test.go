package repomd

import (
	"strings"
	"testing"
)

func TestMetadataParsing(t *testing.T) {
	metadata := `
<?xml version="1.0" encoding="UTF-8"?>
<metadata xmlns="http://linux.duke.edu/metadata/common" xmlns:rpm="http://linux.duke.edu/metadata/rpm" packages="6381">
<package type="rpm">
  <name>migrationtools</name>
  <arch>noarch</arch>
  <version epoch="0" ver="47" rel="7.el6"/>
  <checksum type="sha256" pkgid="YES">ec9d688898b0f6438c2d422229071e9cbb26ba2fe4d451fb62e813f7195dc649</checksum>
  <summary>Migration scripts for LDAP</summary>
  <description>The MigrationTools are a set of Perl scripts for migrating users, groups,
aliases, hosts, netgroups, networks, protocols, RPCs, and services from
existing nameservices (flat files, NIS, and NetInfo) to LDAP.</description>
  <packager>CentOS BuildSystem &lt;http://bugs.centos.org&gt;</packager>
  <url>http://www.padl.com/OSS/MigrationTools.html</url>
  <time file="1309667414" build="1282666158"/>
  <size package="25192" installed="106663" archive="111420"/>
<location href="Packages/migrationtools-47-7.el6.noarch.rpm"/>
  <format>
    <rpm:license>BSD</rpm:license>
    <rpm:vendor>CentOS</rpm:vendor>
    <rpm:group>System Environment/Daemons</rpm:group>
    <rpm:buildhost>c6b2.bsys.dev.centos.org</rpm:buildhost>
    <rpm:sourcerpm>migrationtools-47-7.el6.src.rpm</rpm:sourcerpm>
    <rpm:header-range start="1384" end="8748"/>
    <rpm:provides>
      <rpm:entry name="migrationtools" flags="EQ" epoch="0" ver="47" rel="7.el6"/>
    </rpm:provides>
    <rpm:requires>
      <rpm:entry name="/bin/sh"/>
      <rpm:entry name="/usr/bin/perl"/>
      <rpm:entry name="ldif2ldbm"/>
      <rpm:entry name="openldap-clients"/>
      <rpm:entry name="perl"/>
    </rpm:requires>
  </format>
</package>
</metadata>
`

	r := strings.NewReader(metadata)
	m, err := ParseMetadata(r)
	if err != nil {
		t.Error(err)
	}

	packagesLen := len(m.Packages)
	expectedLen := 1
	if packagesLen != expectedLen {
		t.Errorf(`len(v.Packages) = %q want %q`, packagesLen, expectedLen)
	}
}

func TestRepomdParsing(t *testing.T) {

	repomdData := `
<?xml version="1.0" encoding="UTF-8"?>
<repomd xmlns="http://linux.duke.edu/metadata/repo" xmlns:rpm="http://linux.duke.edu/metadata/rpm">
 <revision>1362488530</revision>
<data type="group">
  <checksum type="sha256">2727fcb43fbe4c1a3588992af8c19e4d97167aee2f6088959221fc285cab6f72</checksum>
  <location href="repodata/2727fcb43fbe4c1a3588992af8c19e4d97167aee2f6088959221fc285cab6f72-c6-x86_64-comps.xml"/>
  <timestamp>1362488701.77</timestamp>
  <size>1189260</size>
</data>
<data type="filelists">
  <checksum type="sha256">3ebb49a547c9d6d96b88ca0befda04c40a083515f33bdfd450640f854f3c87ae</checksum>
  <open-checksum type="sha256">519bc11c1ad67cc2af47ecc61d1c3ff80f775b09f955068c436bd8fd61d71007</open-checksum>
  <location href="repodata/3ebb49a547c9d6d96b88ca0befda04c40a083515f33bdfd450640f854f3c87ae-filelists.xml.gz"/>
  <timestamp>1362488681</timestamp>
  <size>5499724</size>
  <open-size>71465356</open-size>
</data>
</repomd>
`

	r := strings.NewReader(repomdData)
	repomd, err := ParseRepomd(r)
	if err != nil {
		t.Error(err)
	}

	dataLen := len(repomd.Data)
	expectedLen := 2
	if dataLen != expectedLen {
		t.Errorf(`len(v.Data) = %q want %q`, dataLen, expectedLen)
	}
}
