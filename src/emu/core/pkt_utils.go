// Copyright (c) 2020 Cisco Systems and/or its affiliates.
// Licensed under the Apache License, Version 2.0 (the "License");
// that can be found in the LICENSE file in the root of the source
// tree.

package core

import (
	"external/google/gopacket"
	"external/google/gopacket/layers"
	"external/google/gopacket/pcapgo"
	"os"
	"time"
)

func PacketUtlBuild(layers ...gopacket.SerializableLayer) []byte {
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	//ip.SerializeTo(buf, opts)
	gopacket.SerializeLayers(buf, opts, layers...)
	data := buf.Bytes()
	return data
}

func PacketUtl(pcap_name string, o []byte) {
	ts := time.Unix(0, 0)
	f, _ := os.Create("/tmp/" + pcap_name + ".pcap")
	w := pcapgo.NewWriterNanos(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)
	w.WritePacket(gopacket.CaptureInfo{
		Timestamp:     ts,
		Length:        len(o),
		CaptureLength: len(o),
	}, o)
	f.Close()
}
