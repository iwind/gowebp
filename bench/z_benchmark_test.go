// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by go run gen.go -output z_benchmark_test.go; DO NOT EDIT

// +build go1.6

package webp_bench

import (
	"bytes"
	"io/ioutil"
	"testing"

	chai2010_webp "github.com/iwind/gowebp"
	x_image_webp "golang.org/x/image/webp"
)

func tbLoadData(tb testing.TB, filename string) []byte {
	data, err := ioutil.ReadFile("../testdata/" + filename)
	if err != nil {
		tb.Fatal(err)
	}
	return data
}

func BenchmarkDecode_1_webp_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_webp_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_webp_a_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "1_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_webp_ll_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_webp_ll_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "1_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_1_webp_ll_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "1_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "2_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "2_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_a_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "2_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_ll_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "2_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_ll_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "2_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_2_webp_ll_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "2_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "3_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "3_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_a_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "3_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_ll_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "3_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_ll_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "3_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_3_webp_ll_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "3_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "4_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "4_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_a_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "4_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_ll_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "4_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_ll_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "4_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_4_webp_ll_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "4_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_a_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "5_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_a_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "5_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_a_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "5_webp_a.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_ll_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "5_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_ll_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "5_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_5_webp_ll_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "5_webp_ll.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.no-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.no-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_no_filter_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.no-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.normal-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.normal-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_normal_filter_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.normal-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.simple-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.simple-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_large_simple_filter_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink-large.simple-filter.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_blue_purple_pink_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "blue-purple-pink.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_1bpp_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.1bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_1bpp_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.1bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_1bpp_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.1bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_2bpp_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.2bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_2bpp_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.2bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_2bpp_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.2bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_4bpp_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.4bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_4bpp_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.4bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_4bpp_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.4bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_8bpp_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.8bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_8bpp_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.8bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_gopher_doc_8bpp_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "gopher-doc.8bpp.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_photo_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "photo.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_photo_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "photo.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_photo_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "photo.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_tux_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "tux.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_tux_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "tux.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_tux_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "tux.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "video-001.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "video-001.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "video-001.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "video-001.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "video-001.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_video_001_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "video-001.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossless_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossless_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossless_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossless.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_with_alpha_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy-with-alpha.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_with_alpha_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy-with-alpha.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_with_alpha_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy-with-alpha.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_chai2010_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_x_image_webp(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := x_image_webp.Decode(bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}

func BenchmarkDecode_yellow_rose_lossy_chai2010_webp_tosize(b *testing.B) {
	data := tbLoadData(b, "yellow_rose.lossy.webp")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := chai2010_webp.DecodeRGBAToSize(data, 256, 256)
		if err != nil {
			b.Fatal(err)
		}
		_ = m
	}
}
